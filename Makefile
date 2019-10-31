
BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

build: clean go.sum

ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/ens-go.exe .
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/ens-go
endif

build-linux:
	@GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	rm -fr ./build
