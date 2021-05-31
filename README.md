# service-template

## Setup
#### Dependencies
- [buf cli](https://docs.buf.build/installation/) to compile proto files
- [docker](https://www.docker.com/) to run local containers

### Installation
```shell script
# one time command to download proto dependencies
$ buf beta mod update

# generate client/swagger/server code from proto files
$ buf generate
```

##### Spin up the API
```shell script
docker-compose up -d
```
##### Now you can test out curling your gRPC server via [grpcurl](https://github.com/fullstorydev/grpcurl):
```shell script
$ grpcurl -plaintext localhost:8080 list
grpc.reflection.v1alpha.ServerReflection
template.TodoService
```
```shell script
$ grpcurl -plaintext localhost:8080 describe template.TodoService
template.TodoService is a service:
service TodoService {
  rpc Create ( .template.CreateRequest ) returns ( .template.Todo ) {
    option (.google.api.http) = { post:"/todos" body:"*"  };
  }
  rpc Get ( .template.GetRequest ) returns ( .template.Todo ) {
    option (.google.api.http) = { get:"/todos/{id}"  };
  }
  rpc ListAll ( .google.protobuf.Empty ) returns ( .template.ListAllResponse ) {
    option (.google.api.http) = { get:"/todos"  };
  }
}
```
```shell script
$ grpcurl -plaintext -d '{"text": "this is a test", "author": "pep"}' localhost:8080 template.TodoService/Create
{
  "id": "6978690c-5a6d-4771-9182-81aaf6fc333e",
  "text": "this is a test",
  "author": "pep",
  "timestamp": "2021-05-28 17:19:29.5131457 +0000 UTC m=+160.747979501"
}
```
```shell script
$ grpcurl -plaintext localhost:8080 template.TodoService/ListAll
{
  "todos": [
    {
      "id": "f1c4128c-421a-4f47-955d-30ccaeec9492",
      "text": "this is a test",
      "author": "pep",
      "timestamp": "2021-05-28 17:26:34.7979285 +0000 UTC m=+22.439283901"
    }
  ]
}
```
##### You can also test out the REST gateway:
```shell script
$ curl -X GET localhost:8080/todos | json_pp
{
   "todos" : [
      {
         "author" : "pep",
         "text" : "this is a test",
         "id" : "f1c4128c-421a-4f47-955d-30ccaeec9492",
         "timestamp" : "2021-05-28 17:26:34.7979285 +0000 UTC m=+22.439283901"
      }
   ]
}
```

##### Navigate to http://localhost:8080/swagger-ui/ to interact with the REST gateway via swagger

## Demo
##### Swagger
https://service-template-oxm27jqbha-uc.a.run.app/swagger-ui/
##### gRPC Health Endpoint
```shell script
$ grpcurl service-template-oxm27jqbha-uc.a.run.app:443 template.HealthService/Readiness
{
  "ok": true,
  "ready": {
    "datastore": true
  }
}
```
##### REST Health Endpoint
```shell script
$ curl -X GET https://service-template-oxm27jqbha-uc.a.run.app/readyz | json_pp
{
   "ok" : true,
   "ready" : {
      "datastore" : true
   }
}
```