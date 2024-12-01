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
            "name": "DEVils"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/employee": {
            "post": {
                "tags": [
                    "Employee"
                ],
                "summary": "Полуение конкретного сотрудника",
                "parameters": [
                    {
                        "description": "json body",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.GetEmployeeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.GetEmployeeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/employees": {
            "post": {
                "tags": [
                    "Employee"
                ],
                "summary": "Полуение сотрудников по фильтрам",
                "parameters": [
                    {
                        "description": "json body",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.GetBaseEmployeesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.GetBaseEmployeesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/hint": {
            "post": {
                "tags": [
                    "Hint"
                ],
                "summary": "Получение подсказок по полям",
                "parameters": [
                    {
                        "description": "json body",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.GetHintRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.GetBaseEmployeesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/unit": {
            "post": {
                "tags": [
                    "Unit"
                ],
                "summary": "Получение данных о юните",
                "parameters": [
                    {
                        "description": "json body",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.GetUnitRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.Unit"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "github_com_cutlery47_employee-service_internal_model.BaseEmployee": {
            "type": "object",
            "properties": {
                "family_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_general": {
                    "type": "boolean"
                },
                "middle_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                },
                "units": {
                    "description": "название текущего юнита -\u003e название высшестоящего юнита",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "github_com_cutlery47_employee-service_internal_model.GetBaseEmployeesRequest": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "position": {
                    "type": "string"
                },
                "project": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "github_com_cutlery47_employee-service_internal_model.GetBaseEmployeesResponse": {
            "type": "object",
            "properties": {
                "employees": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.BaseEmployee"
                    }
                }
            }
        },
        "github_com_cutlery47_employee-service_internal_model.GetEmployeeRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "github_com_cutlery47_employee-service_internal_model.GetEmployeeResponse": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "family_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "middle_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "office_address": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "project": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "teammates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.BaseEmployee"
                    }
                },
                "unit_id": {
                    "type": "integer"
                }
            }
        },
        "github_com_cutlery47_employee-service_internal_model.GetHintRequest": {
            "type": "object",
            "properties": {
                "city_search_term": {
                    "type": "string"
                },
                "name_search_term": {
                    "type": "string"
                },
                "position_search_term": {
                    "type": "string"
                },
                "project_search_term": {
                    "type": "string"
                },
                "role_search_term": {
                    "type": "string"
                },
                "unit_search_term": {
                    "type": "string"
                }
            }
        },
        "github_com_cutlery47_employee-service_internal_model.GetUnitRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "github_com_cutlery47_employee-service_internal_model.Unit": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "leader_full_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "partisipants": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.BaseEmployee"
                    }
                },
                "unit_parent_id": {
                    "type": "integer"
                },
                "units": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_cutlery47_employee-service_internal_model.Unit"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Employee Service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
