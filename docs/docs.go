// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bookmarks": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Add a new bookmark",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookmarks"
                ],
                "summary": "Create a bookmark",
                "parameters": [
                    {
                        "description": "Add bookmark",
                        "name": "bookmark",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkCreateModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkCreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkCreateResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkCreateResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkCreateResponse"
                        }
                    }
                }
            }
        },
        "/bookmarks/{id}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get bookmark by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookmarks"
                ],
                "summary": "Show a bookmark",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bookmark ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkShowResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkShowResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkShowResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/bookmarks.BookmarkShowResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "bookmarks.BookmarkCreateModel": {
            "type": "object",
            "required": [
                "version"
            ],
            "properties": {
                "version": {
                    "type": "string"
                }
            }
        },
        "bookmarks.BookmarkCreateResponse": {
            "type": "object",
            "properties": {
                "bookmark": {
                    "$ref": "#/definitions/bookmarks.BookmarkShowModel"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "bookmarks.BookmarkShowModel": {
            "type": "object",
            "properties": {
                "bookmarks": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastUpdated": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "bookmarks.BookmarkShowResponse": {
            "type": "object",
            "properties": {
                "bookmark": {
                    "$ref": "#/definitions/bookmarks.BookmarkShowModel"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}