build:
	docker build -t userapi:1.0 .

run:
	docker run -p 8080:8080 userapi:1.0
#	docker run -p <host port>:<container port>	


