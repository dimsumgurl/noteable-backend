package app

import (
	"github.com/dimsumgurl/noteable-backend/pkg/api/restapi"
	"github.com/dimsumgurl/noteable-backend/pkg/api/restapi/operations"
	"github.com/go-openapi/loads"
)

type App struct {
	Server *restapi.Server
}

func NewApp() (*App, error) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewNoteableBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	server.Host = "0.0.0.0"
	server.Port = 8080
	server.ConfigureAPI()
	return &App{Server: server}, nil
}

func (a *App) Run() error {
	err := a.Server.Serve()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	return nil
}
