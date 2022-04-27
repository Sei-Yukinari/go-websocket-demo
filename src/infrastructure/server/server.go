package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-websocket-demo/src/config"
	"net/http"
)

func Run(handler *gin.Engine) {
	srv := &http.Server{
		Addr:    ":" + config.Conf.App.Port,
		Handler: handler,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		_ = fmt.Errorf("error:%v", err)
	}
}
