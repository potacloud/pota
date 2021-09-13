# Run CMD
run-potad:
	go run $(shell ls cmd/potad/*.go)

run-potac:
	go run $(shell ls cmd/potac/*.go)


# Proto Generate
gen-images-v1:
	protoc -I=api/images/v1 --go_out=api/images/v1 --go_opt=module=github.com/potacloud/pota/api/images/v1 \
		--go-grpc_out=api/images/v1 --go-grpc_opt module=github.com/potacloud/pota/api/images/v1 \
		$(shell ls api/images/v1/*.proto)
