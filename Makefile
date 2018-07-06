BINARY_NAME := swagger-version-manager
PLATFORMS := linux/amd64 darwin/amd64 windows/amd64
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

all: test build

build: $(PLATFORMS)

$(PLATFORMS):
				GOOS=$(os) GOARCH=$(arch) go build -o '$(BINARY_NAME)-$(os)-$(arch)'

test:
				go test -v ./...

clean:
				go clean

