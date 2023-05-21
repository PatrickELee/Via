all: protoc python-proto

protoc:
	protoc \
		--go_out=./viapb \
		--go-grpc_out=./viapb \
		./viapb/via.proto

python-proto:
	python -m grpc_tools.protoc \
		--proto_path=./viapb \
		--python_out=./python \
		--grpc_python_out=./python \
		./viapb/via.proto

docker-build:
	docker build -t via .

docker-run:
	docker run --publish 32314:32314 via

client-run:
	python ./python/client.py

clean:
	rm -rf .pytest_cache .coverage .pytest_cache coverage.xml

docker-clean:
	@docker system prune -f --filter "label=name=$(MODULE)"