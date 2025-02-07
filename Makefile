proto-review:
	rm -f $(name)-service/review/*.go
	protoc --proto_path=proto/review --go_out=$(name)-service/review --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/review --go-grpc_opt=paths=source_relative \
    proto/review/*.proto


protot-review:
	protoc --plugin=./$(service)/node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=$(service)/pb/ \
    --ts_proto_opt=nestJs=true \
    --proto_path=proto/review \
    --proto_path=/usr/local/include \
    proto/review/*.proto

.PHONY: proto-review protot-review

