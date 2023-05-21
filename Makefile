all: protoc python-proto

protoc:
	protoc \
		--go_out=./api/proto \
		--go-grpc_out=./api/proto \
		./api/proto/via.proto

python-proto:
	python -m grpc_tools.protoc \
		--proto_path=./api/proto \
		--python_out=./python \
		--grpc_python_out=./python \
		./api/proto/via.proto

docker-build:
	docker build -t app .
	docker build -t python-grpc-server ./python

docker-run:
	docker run --publish 32314:32314 via

go-server-run:
	go run ./cmd/via.go

py-server-run:
	python ./python/server.py

py-client-run:
	python ./python/client.py

clean:
	rm -rf .pytest_cache .coverage .pytest_cache coverage.xml

docker-clean:
	@docker system prune -f --filter "label=name=$(MODULE)"