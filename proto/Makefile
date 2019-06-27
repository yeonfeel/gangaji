space := $(empty) $(empty)
comma := ,
empty :=

import_path = github.com/yeonfeel/gangaji/pkg/pb/v1beta1
dst = ../pkg/pb

gateway_repo = github.com/grpc-ecosystem/grpc-gateway
gogo_repo = github.com/gogo/protobuf
go_src = $(GOPATH)/src
gogo_path = $(go_src)/$(gogo_repo)
apps_path = $(go_src)/$(apps_repo)
gateway_path = $(go_src)/$(gateway_repo)
gateway_gapis_path = $(gateway_path)/third_party/googleapis
pb_v1beta1_path = $(apps_path)/pb/v1beta1

gogo_out = gogofaster_out
grpc_plugins = plugins=grpc
logtoerr = logtostderr=true
swagger_out = swagger_out
swagger_api_dst = ../third-party/swagger-ui
gateway_out = grpc-gateway_out

gogo_protos = $(gogo_path)/gogoproto/*.proto

v1beta1_common_protos = v1beta1/common/*.proto
v1beta1_model_path = v1beta1/model

model_ias = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(model_pkg),$(addprefix M,$(model_pbs))))
model_pbs = $(sort $(wildcard $(v1beta1_model_protos)))
model_pkg = model

google_deps = Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any

api_imports = -I . -I $(go_src) -I $(gateway_path) -I $(gateway_gapis_path)

.PHONY: all
all: common

.PHONY: model
model:
	protoc -I . -I $(go_src) --$(gogo_out)=$(grpc_plugins):$(dst) $(v1beta1_model_protos)

.PHONY: common
common:
	protoc $(api_imports) --$(gogo_out)=$(grpc_plugins):$(dst) $(v1beta1_common_protos)
	protoc $(api_imports) --$(swagger_out)=$(logtoerr):$(swagger_api_dst) $(v1beta1_common_protos)
	protoc $(api_imports) --$(gateway_out)=$(logtoerr):$(dst) $(v1beta1_common_protos)

.PHONY: clean
clean:
	@rm -rf $(dst)/v1beta1/* 2>/dev/null
	# @find $(dst)/v1beta1/* -type d -not -path "$(dst)/v1beta1/utils" -exec rm -r {} +
