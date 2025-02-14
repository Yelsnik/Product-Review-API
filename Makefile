proto-nlp:
	python3 -m grpc_tools.protoc -I./protos/nlp \
	--python_out=./nlp-service/nlp \
	--pyi_out=./nlp-service/nlp \
	--grpc_python_out=./nlp-service/nlp \
	./protos/nlp/*.proto

proto-review:
	protoc --proto_path=protos/review --go_out=$(name)-service/review --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/review --go-grpc_opt=paths=source_relative \
    protos/review/*.proto

proto-go-nlp:
	protoc --proto_path=protos/nlp --go_out=$(name)-service/nlp --go_opt=paths=source_relative \
    --go-grpc_out=$(name)-service/nlp --go-grpc_opt=paths=source_relative \
    protos/nlp/*.proto

protot-review:
	protoc --plugin=./$(service)/node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_out=$(service)/pb/ \
    --ts_proto_opt=nestJs=true \
    --proto_path=protos/review \
    --protos_path=/usr/local/include \
    protos/review/*.proto

.PHONY: proto-review protot-review proto-nlp-py