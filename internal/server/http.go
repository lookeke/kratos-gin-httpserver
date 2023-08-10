package server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"tiktok/internal/conf"
	"tiktok/internal/interfaces"
	"tiktok/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userRouter *interfaces.UserUseCase, user *service.UserService, logger log.Logger) *http.Server {
	// http server

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/v1/user", interfaces.RegisterHTTPServer(userRouter))
	// user.RegisterUserHTTPServer(srv, user)
	return srv
}
