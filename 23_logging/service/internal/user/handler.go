package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"service/internal/handlers"
	"service/pkg/logger"
)

type handler struct {
	//log   *zap.Logger
	//logMW func(next httprouter.Handle) httprouter.Handle
	log logger.AllLogger
}

func New(logger logger.AllLogger) handlers.Handler {
	return &handler{
		//log:   logger.NewLogger(),
		//logMW: logger.NewLoggerMiddleware(logger.NewLogger()),
		log: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {

	router.GET("/users", h.log.LogMV(h.GetList))
	//router.GET("/users", h.GetList)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.log.Log.Info("aaa")
	h.log.Log.Info("aaa")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("list of users"))
}
