package app

import (
	"context"
	"net/http"

	"github.com/dimsumgurl/noteable-backend/pkg/api/restapi"
	"github.com/dimsumgurl/noteable-backend/pkg/api/restapi/operations"
	"github.com/go-openapi/loads"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Server  *restapi.Server
	DB      *mongo.Database
	handler *http.Handler
}

func NewApp() (*App, error) {
	return &App{}, nil
}

func (a *App) Initialize() error {
	var err error
	a.DB, err = a.initDB()
	if err != nil {
		return err
	}
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return err
	}
	api := operations.NewNoteableBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	server.Host = "0.0.0.0"
	server.Port = 8080
	server.ConfigureAPI()
	a.Server = server
	return nil
}

func (a *App) Run() error {
	err := a.Server.Serve()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	err := a.DB.Client().Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initDB() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017/?connect=direct"))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client.Database("mongo"), nil
}
