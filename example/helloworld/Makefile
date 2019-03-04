# Variables
GOCMD=GOOS=linux go
SERVICE=helloworld
IMG_HUB?=registry.cn-beijing.aliyuncs.com/weihua/test
TAG?=latest
# Version information
VERSION=1.0.0
REVISION=${shell git rev-parse --short HEAD}
RELEASE=production
BUILD_HASH=${shell git rev-parse HEAD}
BUILD_TIME=${shell date "+%Y-%m-%d@%H:%M:%SZ%z"}
LD_FLAGS:=-X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Release=$(RELEASE) -X main.BuildHash=$(BUILD_HASH) -X main.BuildTime=$(BUILD_TIME)

ifeq (${shell uname -s}, Darwin)
	SED=gsed
else
	SED=sed
endif

prepare:
	@echo prepare
	$(GOCMD) install learn/api-gateway/plugin/...

generate-go:
	@echo generate-go
	@protoc -I${GOPATH}/src -I./service \
	-I${GOPATH}/src/learn/api-gateway/third_party \
	--gogogo_out=\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
	plugins=grpc:./service service/*.proto
	@$(SED) -i '/google\/api/d' service/*.pb.go
	@echo Generate successfully.

build:prepare generate-go
	$(GOCMD) build -ldflags="$(LD_FLAGS)" -o bundles/$(SERVICE) internal/cmd/main.go

image:build
	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .

run:image
	@-docker service rm $(SERVICE) > /dev/null 2>&1  || true	
	@docker service create --name $(SERVICE) --network devel $(IMG_HUB)/$(SERVICE):$(TAG)

test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
