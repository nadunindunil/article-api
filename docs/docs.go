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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Nadun Indunil",
            "url": "https://www.linkedin.com/in/nadunindunil/",
            "email": "nadun1indunil@gmail.com"
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
        "/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "get all articles",
                "operationId": "get-all-articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/article.ArticleResponseDto"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "add a new article",
                "operationId": "create-article",
                "parameters": [
                    {
                        "description": "article data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.ArticleCreateDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "get a article",
                "operationId": "get-article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/article.ArticleResponseDto"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "get all tags",
                "operationId": "get-all-tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tag.Tag"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "add a new tag",
                "operationId": "create-tag",
                "parameters": [
                    {
                        "description": "tag data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tag.TagCreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tag.Tag"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tags/{name}/{date}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "get tags by name and date",
                "operationId": "get-tags-by-name-and-date",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tag name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date",
                        "name": "date",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tag.Tag"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "article.ArticleCreateDto": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string",
                    "example": "Article Body"
                },
                "date": {
                    "type": "string",
                    "example": "2022-04-23T00:00:00Z"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/article.TagDto"
                    }
                },
                "title": {
                    "type": "string",
                    "example": "Article Title"
                }
            }
        },
        "article.ArticleResponseDto": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tag.Tag"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "article.TagDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "tag.Tag": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "tag.TagCreateDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Article API",
	Description:      "This is a sample api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
