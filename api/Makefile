build:
	echo "Building lambda binaries"

build-CoreFunction:
	env GOOS=linux GOARCH=arm64 go build -o build/handler/core/bootstrap cmd/functions/core/core.go
	cp build/handler/core/bootstrap $(ARTIFACTS_DIR)/bootstrap
