// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for the Noteable",
    "title": "Noteable",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/version": {
      "get": {
        "tags": [
          "Version"
        ],
        "summary": "Returns the Noteable version",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "version": {
      "type": "object",
      "properties": {
        "buildDate": {
          "type": "string"
        },
        "commitHash": {
          "type": "string"
        },
        "releaseVersion": {
          "type": "string"
        },
        "runtime": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for the Noteable",
    "title": "Noteable",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/version": {
      "get": {
        "tags": [
          "Version"
        ],
        "summary": "Returns the Noteable version",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "version": {
      "type": "object",
      "properties": {
        "buildDate": {
          "type": "string"
        },
        "commitHash": {
          "type": "string"
        },
        "releaseVersion": {
          "type": "string"
        },
        "runtime": {
          "type": "string"
        }
      }
    }
  }
}`))
}