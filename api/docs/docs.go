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
            "name": "Pejman Hkh",
            "url": "https://www.peji.ir",
            "email": "pejman.hkh@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/movie/{link}": {
            "get": {
                "description": "Get movie details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Details movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Link",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/movies": {
            "get": {
                "description": "Get list of all movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "List movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Year",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Imdb Code",
                        "name": "imdb_code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/movies/detail": {
            "get": {
                "description": "Get movie details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Details movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IMDB Code",
                        "name": "imdb",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/movies/search": {
            "get": {
                "description": "Search in movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Search movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Subtitle API",
	Description:      "Subtitle api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
