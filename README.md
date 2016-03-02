### Build:
```docker build -t slack-tube-service .```

### Run:
```docker run -it --rm --name slack-tube-service -p 1123:1123 slack-tube-service```

### Access:
#### If using docker machine
Find out your IP: ``` `docker-machine ip `docker-machine active` ```
Connect to that IP, port ```1123```.
