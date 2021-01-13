package version

import (
	"github.com/dimsumgurl/noteable-backend/pkg/api/models"
	"github.com/dimsumgurl/noteable-backend/pkg/api/restapi/operations/version"
	"github.com/go-openapi/runtime/middleware"
)

func GetVersion(params version.GetVersionParams) middleware.Responder {
	return version.NewGetVersionOK().WithPayload(&models.Version{
		CommitHash: "poe release",
	})
}
