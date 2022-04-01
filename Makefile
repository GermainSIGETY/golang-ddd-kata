default: build

run: tests
	go run main.go

build: tests
	@echo "building Get things done app..."
	go build -o GetThingsDone
	@echo "built executable 'GetThingsDone'"

# disabling usage of cached results of test because feature files are not considered in cache keys
# => you modify features files, but tests are not executed against modified feature files.
tests :
	go test -v -count=1 ./...

swagger:
	@swag init -d ./GetThingsDone-app
