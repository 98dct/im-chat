package websocket

import "time"

type serverOption struct {
	Authentication
	patten            string
	maxConnectionIdle time.Duration

	ack        AckType
	ackTimeout time.Duration
}

type ServerOption func(opt *serverOption)

func newServerOptions(opts ...ServerOption) serverOption {
	o := serverOption{
		Authentication:    new(authentication),
		maxConnectionIdle: defaultMaxConnectionIdle,
		ackTimeout:        defaultAckTimeout,
		patten:            "/ws",
	}

	for _, opt := range opts {
		opt(&o)
	}

	return o
}

func WithServerAuthentication(auth Authentication) ServerOption {
	return func(opt *serverOption) {
		opt.Authentication = auth
	}
}

func WithServerPatten(patten string) ServerOption {
	return func(opt *serverOption) {
		opt.patten = patten
	}
}

func WithServerMaxConnectionIdle(maxConnectionIdle time.Duration) ServerOption {
	return func(opt *serverOption) {
		if maxConnectionIdle > 0 {
			opt.maxConnectionIdle = maxConnectionIdle
		}
	}
}

func WithServerAck(ack AckType) ServerOption {
	return func(opt *serverOption) {
		opt.ack = ack
	}
}
