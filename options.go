package king

import (
	"github.com/juju/errors"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	"github.com/altipla-consulting/king/runtime"
)

func WithHttprouter(router *httprouter.Router) ServerOption {
	return func(server *Server) {
		server.router = router
	}
}

func Debug(debug bool) ServerOption {
	return func(server *Server) {
		server.debug = debug
	}
}

func WithLogrus() ServerOption {
	return func(server *Server) {
		server.errorMiddlewares = append(server.errorMiddlewares, func(appErr error) {
			if server.debug {
				log.WithFields(log.Fields{"err": appErr.Error()}).Error("call failed")
				log.Error("Error stack:\n", errors.ErrorStack(appErr))
			} else {
				log.WithFields(log.Fields{"err": appErr.Error(), "stack": errors.ErrorStack(appErr)}).Error("call failed")
			}
		})
	}
}

func HttpClient(client *http.Client) runtime.ClientOption {
	return func(caller *runtime.ClientCaller) {
		caller.Client = client
	}
}
