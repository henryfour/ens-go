
BUILD_FLAGS := -tags "$(build_tags)" -ldflags "$(ldflags)"

build: clean go.sum

ifeq ($(OS),Windows_NT)
	CGO_ENABLED=0 go build -mod=readonly $(BUILD_FLAGS) -o build/ens-go.exe ./cmd
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/ens-go ./cmd
endif

build-linux:
	@GOOS=linux GOARCH=amd64 $(MAKE) build

build-windows:
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -mod=readonly $(BUILD_FLAGS) -o build/ens-go.exe ./cmd

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	rm -fr ./build
