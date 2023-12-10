proto:
	rm -f doc/swagger/*.swagger.json && \
	rm -f pb/*.go && \
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=brood \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	./proto/*.proto

.PHONY: proto