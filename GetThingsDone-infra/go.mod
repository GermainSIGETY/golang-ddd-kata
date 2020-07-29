module github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-infra

go 1.13

replace github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain v0.0.1 => ../GetThingsDone-todo-domain

require (
	github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain v0.0.1
	github.com/jinzhu/gorm v1.9.14
	github.com/stretchr/testify v1.6.1
	gorm.io/gorm v0.2.4 // indirect
)
