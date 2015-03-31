default: docker

docker: Dockerfile
	docker build -t josepharceneaux/go-server:latest .

test: tests
tests: server_test.go
	go test

server: server.go
	go build server.go

clean:
	rm -rf server