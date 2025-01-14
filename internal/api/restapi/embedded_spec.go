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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "bomond-tenis",
    "version": "1.0.0"
  },
  "host": "bomond.vn",
  "paths": {
    "/v1/bomond.vn/auth/logout": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authentication"
        ],
        "responses": {
          "200": {
            "description": "User successfully logged out",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/auth/sign-in": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authentication"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/UserSignin"
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully signed in",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/auth/sign-up": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authentication"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/UserSignup"
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully signed up",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/courts": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "responses": {
          "200": {
            "description": "Courts list retrieved",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/users/{user_id}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/UserId"
          }
        ],
        "responses": {
          "200": {
            "description": "User details retrieved",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/UserUpdate"
          },
          {
            "$ref": "#/parameters/UserId"
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully updated",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/UserId"
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully deleted",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/{court_id}/book": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/CourtId"
          }
        ],
        "responses": {
          "200": {
            "description": "Booking details retrieved",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/BookingRequest"
          },
          {
            "$ref": "#/parameters/CourtId"
          }
        ],
        "responses": {
          "200": {
            "description": "Court successfully booked",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/{court_id}/book/{book_id}": {
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "parameters": [
          {
            "$ref": "#/parameters/CourtId"
          },
          {
            "$ref": "#/parameters/BookingId"
          }
        ],
        "responses": {
          "200": {
            "description": "Booking successfully deleted",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResult": {
      "description": "Error response",
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string",
          "maxLength": 10
        },
        "debugInfo": {
          "description": "Detailed error information",
          "type": "string",
          "maxLength": 512
        },
        "message": {
          "description": "Error message",
          "type": "string",
          "maxLength": 255
        },
        "status": {
          "type": "integer",
          "maxLength": 3
        },
        "timestamp": {
          "type": "string",
          "format": "date-time",
          "readOnly": true
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "data": {
          "type": "object"
        },
        "err": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "status": {
          "type": "integer"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "parameters": {
    "Authorization": {
      "type": "string",
      "description": "JWT token for authorization",
      "name": "Authorization",
      "in": "header",
      "required": true
    },
    "BookingId": {
      "type": "string",
      "name": "book_id",
      "in": "path",
      "required": true
    },
    "BookingRequest": {
      "name": "bookingRequest",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "date": {
            "type": "string",
            "format": "datetime"
          },
          "duration": {
            "type": "integer"
          },
          "userId": {
            "type": "string"
          }
        }
      }
    },
    "CourtId": {
      "type": "string",
      "name": "court_id",
      "in": "path",
      "required": true
    },
    "UserId": {
      "type": "string",
      "name": "user_id",
      "in": "path",
      "required": true
    },
    "UserSignin": {
      "name": "userSignin",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "email": {
            "type": "string"
          },
          "password": {
            "type": "string"
          }
        }
      }
    },
    "UserSignup": {
      "name": "userSignup",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "email": {
            "type": "string"
          },
          "password": {
            "type": "string"
          },
          "username": {
            "type": "string"
          }
        }
      }
    },
    "UserUpdate": {
      "name": "userUpdate",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "email": {
            "type": "string"
          },
          "password": {
            "type": "string"
          },
          "username": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "bomond-tenis",
    "version": "1.0.0"
  },
  "host": "bomond.vn",
  "paths": {
    "/v1/bomond.vn/auth/logout": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authentication"
        ],
        "responses": {
          "200": {
            "description": "User successfully logged out",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/auth/sign-in": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authentication"
        ],
        "parameters": [
          {
            "name": "userSignin",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "email": {
                  "type": "string"
                },
                "password": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully signed in",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/auth/sign-up": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authentication"
        ],
        "parameters": [
          {
            "name": "userSignup",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "email": {
                  "type": "string"
                },
                "password": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully signed up",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/courts": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "responses": {
          "200": {
            "description": "Courts list retrieved",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/users/{user_id}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "User details retrieved",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "parameters": [
          {
            "name": "userUpdate",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "email": {
                  "type": "string"
                },
                "password": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          },
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully updated",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "User successfully deleted",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/{court_id}/book": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "court_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Booking details retrieved",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "parameters": [
          {
            "name": "bookingRequest",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "date": {
                  "type": "string",
                  "format": "datetime"
                },
                "duration": {
                  "type": "integer"
                },
                "userId": {
                  "type": "string"
                }
              }
            }
          },
          {
            "type": "string",
            "name": "court_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Court successfully booked",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    },
    "/v1/bomond.vn/{court_id}/book/{book_id}": {
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Courts"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "court_id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "book_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Booking successfully deleted",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResult"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResult": {
      "description": "Error response",
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string",
          "maxLength": 10
        },
        "debugInfo": {
          "description": "Detailed error information",
          "type": "string",
          "maxLength": 512
        },
        "message": {
          "description": "Error message",
          "type": "string",
          "maxLength": 255
        },
        "status": {
          "type": "integer",
          "maxLength": 3
        },
        "timestamp": {
          "type": "string",
          "format": "date-time",
          "readOnly": true
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "data": {
          "type": "object"
        },
        "err": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "status": {
          "type": "integer"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "parameters": {
    "Authorization": {
      "type": "string",
      "description": "JWT token for authorization",
      "name": "Authorization",
      "in": "header",
      "required": true
    },
    "BookingId": {
      "type": "string",
      "name": "book_id",
      "in": "path",
      "required": true
    },
    "BookingRequest": {
      "name": "bookingRequest",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "date": {
            "type": "string",
            "format": "datetime"
          },
          "duration": {
            "type": "integer"
          },
          "userId": {
            "type": "string"
          }
        }
      }
    },
    "CourtId": {
      "type": "string",
      "name": "court_id",
      "in": "path",
      "required": true
    },
    "UserId": {
      "type": "string",
      "name": "user_id",
      "in": "path",
      "required": true
    },
    "UserSignin": {
      "name": "userSignin",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "email": {
            "type": "string"
          },
          "password": {
            "type": "string"
          }
        }
      }
    },
    "UserSignup": {
      "name": "userSignup",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "email": {
            "type": "string"
          },
          "password": {
            "type": "string"
          },
          "username": {
            "type": "string"
          }
        }
      }
    },
    "UserUpdate": {
      "name": "userUpdate",
      "in": "body",
      "schema": {
        "type": "object",
        "properties": {
          "email": {
            "type": "string"
          },
          "password": {
            "type": "string"
          },
          "username": {
            "type": "string"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
