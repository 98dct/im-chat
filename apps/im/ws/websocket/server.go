package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
	"time"
)

type AckType int

const (
	NoAck AckType = iota
	OnlyAck
	RigorAck
)

func (a AckType) ToString() string {
	switch a {
	case OnlyAck:
		return "OnlyAck"
	case RigorAck:
		return "RigorAck"
	}
	return "NoAck"
}

type Server struct {
	sync.RWMutex
	connToUser map[*Conn]string
	userToConn map[string]*Conn

	opt            *serverOption
	authentication Authentication

	routes   map[string]HandlerFunc
	addr     string
	patten   string
	upgrader websocket.Upgrader
	logx.Logger
}

func NewServer(addr string, opts ...ServerOption) *Server {

	opt := newServerOptions(opts...)

	return &Server{
		connToUser:     make(map[*Conn]string),
		userToConn:     make(map[string]*Conn),
		opt:            &opt,
		authentication: opt.Authentication,
		routes:         make(map[string]HandlerFunc),
		addr:           addr,
		patten:         opt.patten,
		upgrader:       websocket.Upgrader{},
		Logger:         logx.WithContext(context.Background()),
	}
}

// 接受请求并处理请求
func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			s.Errorf("server handler ws recover err %v", r)
		}
	}()

	conn := NewConn(s, w, r)
	if conn == nil {
		return
	}
	//conn, err := s.upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	s.Errorf("upgrade err %v", err)
	//	return
	//}

	// 对连接进行鉴权
	if !s.authentication.Auth(w, r) {
		s.Send(Message{FrameType: FrameData, Data: fmt.Sprint("不具备访问权限")}, conn)
		// conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint("不具备访问权限")))
		conn.Close()
		return
	}

	// 记录连接
	s.addConn(conn, r)

	go s.handlerConn(conn)
}

func (s *Server) addConn(conn *Conn, req *http.Request) {
	userId := s.authentication.UserId(req)

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	// 验证用户是否之前登陆过
	if c := s.userToConn[userId]; c != nil {
		// 关闭之前的连接
		c.Close()
	}

	s.connToUser[conn] = userId
	s.userToConn[userId] = conn
}

func (s *Server) GetConn(uid string) *Conn {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	return s.userToConn[uid]
}

func (s *Server) GetConns(uids ...string) []*Conn {
	if len(uids) == 0 {
		return nil
	}
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	res := make([]*Conn, 0, len(uids))
	for _, uid := range uids {
		res = append(res, s.userToConn[uid])
	}
	return res
}

func (s *Server) GetUsers(conns ...*Conn) []string {

	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}

func (s *Server) Close(conn *Conn) {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	uid := s.connToUser[conn]
	if uid == "" {
		// 已经被关闭
		return
	}

	delete(s.connToUser, conn)
	delete(s.userToConn, uid)

	conn.Close()
}

func (s *Server) SendByIds(msg interface{}, sendIds ...string) error {

	if len(sendIds) == 0 {
		return nil
	}

	return s.Send(msg, s.GetConns(sendIds...)...)
}

func (s *Server) Send(msg interface{}, conns ...*Conn) error {
	if len(conns) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}

	return nil

}

// 根据连接对象执行任务处理
func (s *Server) handlerConn(conn *Conn) {

	uids := s.GetUsers(conn)
	conn.Uid = uids[0]

	// 处理任务
	go s.handlerWrite(conn)

	if s.isAck(nil) {
		go s.readAck(conn)
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn read message err %v", err)
			// todo 关闭连接
			s.Close(conn)
			return
		}

		var message Message
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal err %v", err)
			s.Close(conn)
			return
		}

		// todo 给客户端回复ack

		if s.isAck(&message) {
			s.Infof("conn message read ack message %v:", message)
			conn.appendMsgMq(&message)
		} else {
			conn.message <- &message
		}

	}
}

func (s *Server) isAck(message *Message) bool {

	if message == nil {
		return s.opt.ack != NoAck
	}
	return s.opt.ack != NoAck && message.FrameType != FrameNoAck
}

// 读取消息的ack
func (s *Server) readAck(conn *Conn) {
	for {
		select {
		case <-conn.done:
			s.Infof("close message ack uid %v:", conn.Uid)
			return
		default:

		}

		// 从队列中获取新的消息
		conn.messageMu.Lock()
		if len(conn.readMessage) == 0 {
			conn.messageMu.Unlock()
			time.Sleep(100 * time.Microsecond)
			continue
		}

		// 读取第一条
		message := conn.readMessage[0]

		// 判断ack的方式
		switch s.opt.ack {
		case OnlyAck:
			// 直接给客户端回复
			s.Send(Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq + 1,
			}, conn)

			// 把消息从队列中删除
			conn.readMessage = conn.readMessage[1:]
			conn.messageMu.Unlock()

			conn.message <- message
		case RigorAck:

			// 先回
			if message.AckSeq == 0 {
				conn.readMessage[0].AckSeq++
				conn.readMessage[0].AckTime = time.Now()
				s.Send(&Message{
					FrameType: FrameAck,
					Id:        message.Id,
					AckSeq:    message.AckSeq,
				})
				s.Infof("message ack RigorAck send mid:%v ackSeq:%v acktime:%v",
					message.Id, message.AckSeq, message.AckTime)
				conn.messageMu.Unlock()
				continue
			}

			// 在验证

			// 1.客户端返回结果，再一次确认
			msgSeq := conn.readMessageSeq[message.Id]
			if msgSeq.AckSeq > message.AckSeq {
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				conn.message <- message
				s.Infof("message ack rigorAck success mid: %v", message.Id)
				continue
			}

			// 2.客户端没有返回结果，考虑是否超过ack确认时间
			v := s.opt.ackTimeout - time.Since(message.AckTime)
			if !message.AckTime.IsZero() && v <= 0 {
				// 2.2 超过了就结束
				delete(conn.readMessageSeq, message.Id)
				conn.readMessage = conn.readMessage[1:]
				conn.messageMu.Unlock()
				continue
			}
			// 2.1 未超过，重新发送
			conn.messageMu.Unlock()
			s.Send(&Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq,
			}, conn)

			// 测试时，时间长一点
			time.Sleep(3 * time.Second)

		}

	}
}

// 任务的处理
func (s *Server) handlerWrite(conn *Conn) {

	for {
		select {
		case <-conn.done:
			// 连接关闭
			return
		case message := <-conn.message:
			// 根据消息进行处理
			switch message.FrameType {
			case FramePing:
				s.Send(Message{
					FrameType: FramePing,
				}, conn)
			case FrameData:
				// 根据请求的方法分发路由，并执行
				if handler, ok := s.routes[message.Method]; ok {
					handler(s, conn, message)
				} else {
					s.Send(Message{FrameType: FrameData, Data: fmt.Sprintf("不存在方法%v", message.Method)}, conn)
					//conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("不存在方法%v", message.Method)))
				}
			}

			if s.isAck(message) {
				conn.messageMu.Lock()
				delete(conn.readMessageSeq, message.Id)
				conn.messageMu.Unlock()
			}

		}
	}
}

func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) Start() {
	http.HandleFunc(s.patten, s.ServerWs)
	s.Info(http.ListenAndServe(s.addr, nil))

}

func (s *Server) Stop() {
	fmt.Println("停止服务！")
}
