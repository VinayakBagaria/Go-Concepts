kill-all-servers:
	@for i in {0..4}; do\
		kill -9 $$(lsof -t -i :500$$i);\
	done

run-python-servers:
	python loadbalancer/server.py server-0 5000 & python loadbalancer/server.py server-1 5001 & \
	python loadbalancer/server.py server-2 5002 & python loadbalancer/server.py server-3 5003 & \
	python loadbalancer/server.py server-4 5004

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative todo.proto

run-consul:
	docker run -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0