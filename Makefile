kill-all-servers:
	@for i in {0..4}; do\
		kill -9 $$(lsof -t -i :500$$i);\
	done

run-python-servers:
	python loadbalancer/server.py server-0 5000 & python loadbalancer/server.py server-1 5001 & \
	python loadbalancer/server.py server-2 5002 & python loadbalancer/server.py server-3 5003 & \
	python loadbalancer/server.py server-4 5004