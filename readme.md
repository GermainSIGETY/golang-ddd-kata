# Domain Driven Design Kata In Golang
> A To-Do list application to Get Things Done

## Presentation     
A simple Rest API for a Todo list management, developed with Golang.
Project applies DDD concepts :
- hexagonal architecture
- entities and value objects
- domain events (coming soon...)

## Article

Please read https://medium.com/@gsigety/domain-driven-design-golang-kata-1-d76d01459806

## Architecture

project layout applies Golang's Standard project layout
https://github.com/golang-standards/project-layout

### internal/domain

**The hexagon**
- contains all business logic; validation rules of users inputs ; mandatory information for a Todo, format of fields and so on => it validates invariants of value objects entering into the system
- contains domain objects. Business concepts and words have their objects in your code (ubiquitous language). Here just one domain object which is a Todo
- transaction boundaries : manage database transactions ; decide whether actions results should be persisted or rollbacked.

**dependencies :** (almost) Nothing. Your domain is not coupled with any web framework nor persistence/infrastructure Framework.

### internal/infrastructure

**Infrastructure layer**
- contains implementation of means for hexagon to communicate with the outside world
- here GetThingsDone-infra is responsible for one major topic ; manage persistence of Todos with a persistent storage : here a database

But it could be many others responsibilities ; sending email, SMS, push notification, read configuration info etc.

**dependencies** : domain, and some Golang stuff for persistence : GORM, SQL drivers etc.

### internal/ui  

**UI and final runnable artifact**
- contains Go main class : entry point to launch the whole stuff
- user interface : user interface is a HTTP Rest/JSON API, but it could be html pages, CLI, gRPC etc.
- application packaging : an executable file

**dependencies** : domain, infrastructure, and Golang stuff for http/Rest/JSON

### internal/bootstrap

**Code to initiate and launch API**
- create the API, repository, http and and so one

**dependencies** : domain, infrastructure and ui

## Build, run, automated tests

### build

```shell script
make
```

### run
```shell script
make run
```

### automated tests
```shell script
make tests
```
# Usage :

endpoints URL is

http://localhost:8080/todos

### Create 
```shell script
curl -i -X POST \
  http://localhost:8080/todos \
  -H 'Content-Type: application/json' \
  -d '{
   "title":"plant a tree",
   "description" : "because green is pleasing",
   "dueDate": 1557847007
}'
```
### Update 
```shell script
curl -i -X PUT \
  http://localhost:8080/todos/1 \
  -H 'Content-Type: application/json' \
  -d '{
   "title":"cut a tree",
   "description" : "to burn it un my fireplace",
   "dueDate": 155784755
}'
```
### Delete
```shell script
curl -i -X DELETE http://localhost:8080/todos/5
```
### Read todo
```shell script
curl -i -X GET http://localhost:8080/todos/1
```
### list
```shell script
curl -X GET http://localhost:8080/todos
```

## Swagger / OpenAPI
### Install Swaggo / swag
```
go get -u github.com/swaggo/swag/cmd/swag
```
### Generate doc 
```
make swagger
```

Doc is generated/updated into ./doc directtory

See https://github.com/swaggo/swag

## Bibliography

**Hexagonal architecture :**

https://blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example/

**DDD quickly :**

https://www.infoq.com/minibooks/domain-driven-design-quickly

**Tell, Don't ask :**

https://www.martinfowler.com/bliki/TellDontAsk.html

**Implementing DDD, Vaughn Vernon**

https://www.oreilly.com/library/view/implementing-domain-driven-design/9780133039900/



