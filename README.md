## Run locally:
### Pre-requisite:
 - Golang installed (<https://golang.org/>)

### Run:
```go run main.go```

### Access:
Just connect to ```localhost:1123```

## Run within Docker container
### Pre-requisite:
 - Docker daemon running on the local machine or on a docker host (<https://www.docker.com/products/docker-toolbox>)

### Build:
```docker build -t slack-tube-service .```

### Run:
```docker run -it --rm --name slack-tube-service -p 1123:1123 slack-tube-service```

### Access:

#### If docker daemon is running on localhost
Just connect to ```localhost:1123```

#### If using docker machine
Find out your IP: ``` `docker-machine ip `docker-machine active` ```
Connect to that IP, port ```1123```.
