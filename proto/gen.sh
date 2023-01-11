protoc --go_out=. ./proto/landscape.proto
protoc --go-grpc_out=. ./proto/landscape.proto
protoc --grpc-gateway_opt=logtostderr=true --grpc-gateway_out=. ./proto/landscape.proto
protoc --openapiv2_opt=logtostderr=true,json_names_for_fields=false --openapiv2_out=./ ./proto/landscape.proto
protoc  --validate_out="lang=go:./"  ./proto/landscape.proto
