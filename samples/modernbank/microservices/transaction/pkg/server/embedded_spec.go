// Code generated by go-swagger; DO NOT EDIT.

package server

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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Transaction Processing Microservice, ensuring transactions are processed in Modern Bank.",
    "title": "Transaction",
    "version": "1.0.0"
  },
  "host": "transaction",
  "basePath": "/v1",
  "paths": {
    "/transactions": {
      "post": {
        "description": "Sends money from one account to another",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Sends money from one account to another",
        "operationId": "CreateTransaction",
        "parameters": [
          {
            "description": "Transaction to create",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newtransaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/transaction"
            }
          },
          "400": {
            "description": "Nice try! You can't send negative amounts..."
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "newtransaction": {
      "type": "object",
      "required": [
        "sender",
        "receiver",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "number",
          "format": "currency"
        },
        "receiver": {
          "type": "integer",
          "format": "int64"
        },
        "sender": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "transaction": {
      "allOf": [
        {
          "type": "object",
          "required": [
            "sender",
            "receiver",
            "amount"
          ],
          "properties": {
            "amount": {
              "type": "number",
              "format": "currency"
            },
            "receiver": {
              "type": "integer",
              "format": "int64"
            },
            "sender": {
              "type": "integer",
              "format": "int64"
            }
          }
        },
        {
          "type": "object",
          "required": [
            "id"
          ],
          "properties": {
            "id": {
              "type": "string"
            }
          }
        }
      ]
    }
  },
  "tags": [
    {
      "description": "Operations about handling transactions",
      "name": "transaction"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Transaction Processing Microservice, ensuring transactions are processed in Modern Bank.",
    "title": "Transaction",
    "version": "1.0.0"
  },
  "host": "transaction",
  "basePath": "/v1",
  "paths": {
    "/transactions": {
      "post": {
        "description": "Sends money from one account to another",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "transactions"
        ],
        "summary": "Sends money from one account to another",
        "operationId": "CreateTransaction",
        "parameters": [
          {
            "description": "Transaction to create",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newtransaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/transaction"
            }
          },
          "400": {
            "description": "Nice try! You can't send negative amounts..."
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "newtransaction": {
      "type": "object",
      "required": [
        "sender",
        "receiver",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "number",
          "format": "currency"
        },
        "receiver": {
          "type": "integer",
          "format": "int64"
        },
        "sender": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "transaction": {
      "allOf": [
        {
          "type": "object",
          "required": [
            "sender",
            "receiver",
            "amount"
          ],
          "properties": {
            "amount": {
              "type": "number",
              "format": "currency"
            },
            "receiver": {
              "type": "integer",
              "format": "int64"
            },
            "sender": {
              "type": "integer",
              "format": "int64"
            }
          }
        },
        {
          "type": "object",
          "required": [
            "id"
          ],
          "properties": {
            "id": {
              "type": "string"
            }
          }
        }
      ]
    }
  },
  "tags": [
    {
      "description": "Operations about handling transactions",
      "name": "transaction"
    }
  ]
}`))
}