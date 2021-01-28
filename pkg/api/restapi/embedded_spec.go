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
    "/login": {
      "post": {
        "tags": [
          "Authentication"
        ],
        "summary": "Authenticate a user",
        "parameters": [
          {
            "name": "user login body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userAuth"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/userToken"
            }
          },
          "401": {
            "description": "unauthorized",
            "schema": {
              "$ref": "#/responses/unauthorized"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/responses/internalServerError"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "Authentication"
        ],
        "summary": "Create a user",
        "parameters": [
          {
            "name": "user creation",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userAuth"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "bad user request error",
            "schema": {
              "$ref": "#/responses/badRequest"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/responses/internalServerError"
            }
          }
        }
      }
    },
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
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "userAuth": {
      "type": "object",
      "required": [
        "emailaddress",
        "password"
      ],
      "properties": {
        "emailaddress": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "userToken": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
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
  },
  "responses": {
    "200": {
      "description": "OK"
    },
    "badRequest": {
      "description": "Bad request from user",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "internalServerError": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "unauthorized": {
      "description": "Authorization information is missing or invalid.",
      "schema": {
        "$ref": "#/definitions/error"
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
    "/login": {
      "post": {
        "tags": [
          "Authentication"
        ],
        "summary": "Authenticate a user",
        "parameters": [
          {
            "name": "user login body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userAuth"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/userToken"
            }
          },
          "401": {
            "description": "unauthorized",
            "schema": {
              "description": "Authorization information is missing or invalid.",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "description": "Internal Server Error",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "Authentication"
        ],
        "summary": "Create a user",
        "parameters": [
          {
            "name": "user creation",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userAuth"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "bad user request error",
            "schema": {
              "description": "Bad request from user",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "description": "Internal Server Error",
              "schema": {
                "$ref": "#/definitions/error"
              }
            }
          }
        }
      }
    },
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
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "userAuth": {
      "type": "object",
      "required": [
        "emailaddress",
        "password"
      ],
      "properties": {
        "emailaddress": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "userToken": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
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
  },
  "responses": {
    "200": {
      "description": "OK"
    },
    "badRequest": {
      "description": "Bad request from user",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "internalServerError": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "unauthorized": {
      "description": "Authorization information is missing or invalid.",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  }
}`))
}
