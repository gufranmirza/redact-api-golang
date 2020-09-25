# redact-api-golang

A microserver exposes an API /redact/ which will be used to redact a JSON based on a JSON path specified.
- Only redaction of leaf nodes of a JSON are allowed.
- No limitation on the level of nesting that can be expressed using the JSON path configured.


# Running Locally

Clone repository 
``` 
git clone https://github.com/gufranmirza/redact-api-golang
```
go to `src`  folder
```
cd redact-api-golang/src
```
Run the service
```
make run
```
Building the service
```
make build
```
It will create executable binary into the `/bin` folder


# Testing
go to `src`  folder

Run tests
```
make test
```

# Generating Mocks 
go to `src`  folder

Generate mocks for packages using the command

```
make mock
```
# API Documentation
API documentation is maintained with swagger.[ Here is the link to swagger file](https://github.com/gufranmirza/redact-api-golang/blob/master/api/swagger.yaml)

# Docker
A docker files is present in the `src` folder which can be used to create the docker image of this server

### Create Docker Image
To create the docker image run the following command
```
docker build . -t redact-api
```

### Running Docker Image
To run the docker image you have created using the following command
```
docker run -p 8001:8000 redact-api
```
It will run the service on the port:8001
