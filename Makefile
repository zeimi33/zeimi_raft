#GO_SRC_PATH := $(GOPATH)/src/
#GOGO_PROTOBUF_PATH := $(GOPATH)/src/github.com/gogo/protobuf/protobuf/
#GO_PROTOS := $(*.proto)
#PROTOC := protoc
#.makeProto: $(PROTOC) $(GO_PROTOS)
#    build/werror.sh protoc -I=raftpb: \
#    $(GOGO_PROTOBUF_PATH): \
#    $(GO_SRC_PATH): \
#    --gofast_out=. $(GO_PROTOS)
