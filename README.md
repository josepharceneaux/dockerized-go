# Stelligent Mini Project - Go

Provision a system running a Go web server.

## Building the Docker Image

On a docker-enabled system, clone this repo, then build the container:

```console
$ git clone git@github.com:josepharceneaux/stelligent-go.git
$ make
```

## Running the Docker Image

The image has not been pushed to the docker hub, but once pushed it can be run like this:

```console
$ docker run -it --rm -P -p 80:8080 josepharceneaux/go-server:latest
```

This starts the container and maps port 8080 on it to port 80 on the host. This can also be done via ssh, salt, et al.
