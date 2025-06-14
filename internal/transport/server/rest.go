package server

import (
	"cart-api/config"
	v1 "cart-api/internal/transport/rest/api/v1"
	ctrl "cart-api/internal/transport/rest/controllers/v1"
	"cart-api/pkg/log"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Transporter interface {
	StartServer()
	StopServer()
}

type Rest struct {
	server *http.Server
	log    *log.Logger
}

func (r *Rest) StartServer() {
	const op = "transport.server.rest.StartServer"
	r.log.InfoLog.Printf("%s: Starting server\n", op)

	if err := r.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		r.log.ErrorLog.Fatalf("%s: Failed to listen server: %s\n", op, err)
	}
}

func (r *Rest) StopServer() {
	const op = "transport.server.rest.StopServer"
	r.log.InfoLog.Printf("%s: Stopping server\n", op)

	if err := r.server.Shutdown(context.Background()); err != nil {
		r.log.ErrorLog.Fatalf("%s: Error stop server: %s\n", op, err)
	}
}

func New(conf *config.AppConf, log *log.Logger, productCtrl *ctrl.Product) *Rest {
	engine := gin.Default()

	gin.SetMode(conf.ServerMode)

	v1.RegisterRoutes(engine, productCtrl)

	server := &http.Server{
		Addr:     conf.ServerPort,
		Handler:  engine,
		ErrorLog: log.ErrorLog,
	}

	return &Rest{
		server: server,
		log:    log,
	}
}
