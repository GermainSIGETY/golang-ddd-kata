.PHONY: default run build tests tests-e2e generate_swagger use-external-package

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

use-external-package: # Check if packages in internal/domain use external package (they must not)
	@# 1. List all imports in the `./internal/domain/` package
	@# 2. Remove all imports that contain "internal/domain" because we want to allow import from this package inside it self
	@# 3. Keep only imports that have a period in it (this will remore standard lib import that have a / in them ex: net/http)
	@! go list -f '{{join .Imports "\n"}}' ./internal/domain/... | grep -v internal/domain | grep "\."
