
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
  - 1 profile, listening on 10001

* to simulate the create and update event, the `profile` provide the restful apis below
  - `PUT` localhost:10001/profile/add  
    create one patient profile
  - `POST` localhost:10001/profile/update/one
    update one patient profile
  - `POST` localhost:10001/profile/update/multiple
    update multiple patient profiles
  - `GET` localhost:10001/profile/query?id=1
    query one patient profile by patient id

### stop and remove all containers
```shell
docker compose -f docker/deploy/dev.yml stop && docker compose -f docker/deploy/dev.yml rm -y
```
