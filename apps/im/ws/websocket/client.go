package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/url"
)

type Client interface {
	Close() error
	Send(v any) error
	Read(v any) error
}

type client struct {
	*websocket.Conn
	host string

	opt dialOption
}

func NewClient(host string, opts ...DialOptions) *client {
	opt := newDialOptions(opts...)
	c := &client{
		Conn: nil,
		host: host,
		opt:  opt,
	}

	conn, err := c.dial()
	if err != nil {
		panic(err)
	}
	c.Conn = conn

	return c
}

func (c *client) dial() (*websocket.Conn, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   c.host,
		Path:   c.opt.pattern,
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	return conn, err
}

func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = c.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}

	// todo 有错误重连在发送一次
	conn, err := c.dial()
	if err != nil {
		return err
	}
	c.Conn = conn

	return c.WriteMessage(websocket.TextMessage, data)
}

func (c *client) Read(v any) error {
	_, data, err := c.ReadMessage()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
