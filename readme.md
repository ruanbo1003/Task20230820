
## notes
* the project depends on `docker` and `docker compose`. 
* tested on `go version go1.21.0 darwin/amd64` and `Docker Compose version v2.0.0-rc.3`


## build and run 
### build docker images
```shell
docker build -f docker/Dockerfile -t heidi_task:v1 .
```

### use docker compose to start the services  
```shell
docker compose -f docker/deploy/dev.yml up
```  
* by doing this, it will start the services below
  - 1 zookeeper
  - 1 kafka, listening on 9092
  - 1 mysql, listening on 3036
  - 2 consultations  
* for some reason, the `profile` service can not be started by `docker cmopose`, so we need to open another terminal 
and follow the steps bellow to start the `profile` service
  - run `make` if you have installed `make` in your computer, otherwise run `go build -o ./bin/ ./cmd/profile`
  - run `./bin/profile` to start the `profile` service

* to simulate the create and update event, the `profile` provide the restful apis below
  - `PUT` localhost:10001/profile/add  
    create one patient profile
  - `POST` localhost:10001/profile/update/one
    update one patient profile
  - `POST` localhost:10001/profile/update/multiple
    update multiple patient profiles
  - `GET` localhost:10001/profile/query?id=1
    query one patient profile by patient id
