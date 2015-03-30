default: docker

docker: Dockerfile
	docker build -t josepharceneaux/go-server:latest .
