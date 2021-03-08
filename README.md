# Myservice :  a sample REST API in go

This repository contains the code of dummy service, written in Golang. This service is used is different demo I run. 

## Build

Import the module dependency 
```
go mod tidy
```

Build the app
```
go build -o <app-name>
```

## Run 

Run the app without build
```
go run
```

Once you build your application, you could use the binary to execute it

## test
```
curl http://localhost:8080 
```
the answer should be :
```
{"message":"Hello from go api server"}
``` 

run the go test
```
go test
```

## Container
The repo contains also a `Dockerfile` to build a container with the compiled application.
``` 
Docker build -t <your_repo>/myservice:<version> .
``` 
