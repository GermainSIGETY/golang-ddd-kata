module github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-app

go 1.13

replace github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain v0.0.1 => ../GetThingsDone-todo-domain

replace github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-infra v0.0.1 => ../GetThingsDone-infra

require (
	github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-infra v0.0.1
	github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain v0.0.1
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/cucumber/godog v0.10.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/swaggo/swag v1.6.7 // indirect
	github.com/urfave/cli/v2 v2.2.0 // indirect
	golang.org/x/sys v0.0.0-20200831180312-196b9ba8737a // indirect
	golang.org/x/tools v0.0.0-20200903005429-2364a5e8fdcf // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
