# loading env if file exists
ifeq ($(shell test -s .env && echo -n yes),yes)
	include .env
	export $(shell sed 's/=.*//' .env)
endif

default: build

run: tests tests-e2e
	go run main.go

build: tests tests-e2e
	@echo "building Get things done app..."
	go build -o GetThingsDone
	@echo "built executable 'GetThingsDone'"

# disabling usage of cached results of test because feature files are not considered in cache keys
# => you modify features files, but tests are not executed against modified feature files.
# don't execute e2e tests
tests :
	go test $(shell go list ./... | grep -v e2e) -count=1 -p 1

tests-e2e :
	go test -v -count=1 ./internal/e2e -count=1

generate_swagger:
	# 1. Génération des fichiers swagger
	@swag init -o ./deployments/swagger
	# 2. On créer le swagger de prod avec le host de prod
	sed "s/$(STAGING_HOST)/$(PRODUCTION_HOST)/" ./deployments/swagger/swagger.json > ./deployments/swagger/production.swagger.json
	# 3. On vire la clé d'API pour les environnements de staging puis on créer le staging avec les hosts correspondant
	sed -i '/^\s*\"security\": \[\s*/{N;N;N;N;d}' ./deployments/swagger/swagger.json
	sed "s/$(STAGING_HOST)/$(STAGING_HOST)/" ./deployments/swagger/swagger.json > ./deployments/swagger/staging.swagger.json
	# 4. On supprime les fichiers dont on n'a pas besoin
	rm ./deployments/swagger/swagger.json
	rm ./deployments/swagger/swagger.yaml
	rm ./deployments/swagger/docs.go

# TODO: écrire une target qui check si les packages internal/domain n'utilisent pas de packages externe à internal/doamin
