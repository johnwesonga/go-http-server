 Overview

## go-http-server 

To build the container run:

```
docker build --tag johnwesonga/go-http-server:latest .
```

```
podman build --tag johnwesonga/go-http-server:latest .
```

This will build an arm64 image with the tag `johnwesonga/go-http-server:latest` and load it into your local docker registry.

To run the docker container:
```
docker run -p 8080:8080 johnwesonga/go-http-server:latest
```

To run the docker container in detached mode:
```
docker run -d -p 8080:8080 johnwesonga/go-http-server:latest
```