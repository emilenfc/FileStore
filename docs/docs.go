// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "iyaemile4@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/folders": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Get list of folders for the current user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "folders"
                ],
                "summary": "Get user folders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.FolderResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/folders/{folder}": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Get list of files in a specific folder",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "folders"
                ],
                "summary": "Get folder contents",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Folder name",
                        "name": "folder",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.FileResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/regenerate-secret": {
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Generate a new API secret for the current user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Regenerate API secret",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.SecretResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "Get current user's profile information",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate user and receive JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Create a new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "Upload a file to a specific folder",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "files"
                ],
                "summary": "Upload file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key for authentication",
                        "name": "X-API-Key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "API Secret for authentication",
                        "name": "X-API-Secret",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Folder name",
                        "name": "folder",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.UploadResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/uploads/{path}": {
            "get": {
                "description": "Download a file",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "files"
                ],
                "summary": "Get file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File path",
                        "name": "path",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "File contents",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Invalid credentials"
                }
            }
        },
        "handlers.FileResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2025-01-08T12:00:00Z"
                },
                "name": {
                    "type": "string",
                    "example": "document.pdf"
                },
                "size": {
                    "type": "integer",
                    "example": 1024
                },
                "url": {
                    "type": "string",
                    "example": "http://localhost:8085/uploads/ak_123/DOCUMENTS/document.pdf"
                }
            }
        },
        "handlers.FolderResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2025-01-08T12:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "DOCUMENTS"
                }
            }
        },
        "handlers.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "handlers.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIs..."
                }
            }
        },
        "handlers.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "firstname": {
                    "type": "string",
                    "example": "John"
                },
                "lastname": {
                    "type": "string",
                    "example": "Doe"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "phone": {
                    "type": "string",
                    "example": "+250783544364"
                }
            }
        },
        "handlers.RegisterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "User created successfully"
                },
                "user": {
                    "$ref": "#/definitions/handlers.UserResponse"
                }
            }
        },
        "handlers.SecretResponse": {
            "type": "object",
            "properties": {
                "api_secret": {
                    "type": "string",
                    "example": "as_987654321"
                }
            }
        },
        "handlers.UserResponse": {
            "type": "object",
            "properties": {
                "api_key": {
                    "type": "string",
                    "example": "ak_123456789"
                },
                "api_secret": {
                    "type": "string",
                    "example": "as_987654321"
                },
                "created_at": {
                    "type": "string",
                    "example": "2025-01-08T12:00:00Z"
                },
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "firstname": {
                    "type": "string",
                    "example": "John"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lastname": {
                    "type": "string",
                    "example": "Doe"
                },
                "phone": {
                    "type": "string",
                    "example": "+250783544364"
                }
            }
        },
        "main.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Invalid credentials"
                }
            }
        },
        "main.UploadResponse": {
            "type": "object",
            "properties": {
                "dir": {
                    "type": "string"
                },
                "file_name": {
                    "type": "string"
                },
                "file_url": {
                    "type": "string"
                },
                "full_path": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "TokenAuth": {
            "description": "JWT token for authentication",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8085",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "File Store API",
	Description:      "A secure file storage API with user authentication and folder management",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
