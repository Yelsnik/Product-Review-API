server:
	python3 main.py

test1:
	python3 -m unittest $(file)

test:
	python3 -m unittest discover

proto:
	python3 -m grpc_tools.protoc -I. \
	--python_out=. \
	--pyi_out=. \
	--grpc_python_out=. \
	./*.proto

.PHONY: run