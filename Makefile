default: build

run: tests
	@cd GetThingsDone-app && go run main.go

build: tests
	@echo "building Get things done app..."
	@cd GetThingsDone-app && go build -o ../GetThingsDone
	@echo "built executable 'GetThingsDone'"

tests : test-app

# disabling usage of cached results of test because feature files are not considered in cache keys
# => you modify features files, but tests are not executed against modified feature files.
test-app: test-infra
	@cd GetThingsDone-app && go test -v -count=1 ./...

test-infra: test-domain
	@cd GetThingsDone-infra && go test ./...

test-domain:
	@cd GetThingsDone-todo-domain && go test ./...