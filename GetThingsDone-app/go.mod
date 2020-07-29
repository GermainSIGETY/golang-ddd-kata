module github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-app

go 1.13

replace github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain v0.0.1 => ../GetThingsDone-todo-domain

replace github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-infra v0.0.1 => ../GetThingsDone-infra

require (
	github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-infra v0.0.1
	github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain v0.0.1
	github.com/cucumber/godog v0.10.0
	github.com/gin-gonic/gin v1.6.3
)
