SOURCES=main.go
TARGET_BINARY=SubtitleConverter
GO_CMD=go
GO_FLAGS=-O2 -fstack-protector-all -D_FORTIFY_SOURCE=2 -fPIE -pie
GO_LDFLAGS="-linkmode=external -extldflags=-Wl,-z,now,-z,relro,-z,noexecstack -s"

export GO111MODULE=on
export CGO_ENABLED=1
export CGO_CFLAGS=$(GO_FLAGS)
export CGO_CXXFLAGS=$(GO_FLAGS)

.PHONY: all
all: binary

.PHONY: binary
binary:
	$(GO_CMD) build --buildmode=exe --buildmode=pie -ldflags $(GO_LDFLAGS) -o $(TARGET_BINARY) $(SOURCES)

.PHONY: clean
clean: $(RM) $(TARGET_BINARY)
