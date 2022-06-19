# service-template

## Setup
#### Dependencies
- [buf cli](https://docs.buf.build/installation/) - compiles proto files
- [docker](https://www.docker.com/) - containerizes application
- [grpc-go](https://github.com/grpc/grpc-go) - golang gRPC framework
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) - gRPC to JSON reverse proxy generator
- [go-gorm](https://github.com/go-gorm/gorm) - golang SQL ORM library
- [postgres](https://github.com/docker-library/postgres) - SQL backend
- [golang-migrate](https://github.com/golang-migrate/migrate) - golang library to handle SQL migrations
- [jwt-go](https://github.com/dgrijalva/jwt-go) - golang library to interact with json web tokens
- [gqlgen](https://github.com/99designs/gqlgen) - golang library for building GraphQL servers

### Installation
```shell script
# install docker
$ brew install docker

# install buf CLI
$ brew tap bufbuild/buf
$ brew install buf

# one time command to download proto dependencies
$ buf mod update

# generate client/swagger/server code from proto files
$ buf generate

# install the tools listed in tools/tools.go
$ go install \ 
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
  # ... etc  
  
# generate graphql code from schema
$ gqlgen generate
```

#### Spin up the API
```shell script
docker-compose up -d
```
###### Now you can test out curling your gRPC server via [grpcurl](https://github.com/fullstorydev/grpcurl)
##### List all gRPC services
```shell script
$ grpcurl -plaintext localhost:8080 list
grpc.health.v1.Health
grpc.reflection.v1alpha.ServerReflection
template.AuthService
template.HealthService
template.TodoService
```
##### Describe a gRPC service
```shell script
$ grpcurl -plaintext localhost:8080 describe template.AuthService
template.AuthService is a service:
service AuthService {
  rpc Me ( .google.protobuf.Empty ) returns ( .template.User ) {
    option (.google.api.http) = { get:"/auth/me"  };
  }
  rpc Signup ( .template.SignupRequest ) returns ( .template.TokenResponse ) {
    option (.google.api.http) = { post:"/auth/signup" body:"*"  };
    option (.grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = { security:<> };
  }
  rpc Token ( .template.TokenRequest ) returns ( .template.TokenResponse ) {
    option (.google.api.http) = { post:"/auth/token" body:"*"  };
    option (.grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = { security:<> };
  }
}
```
#### Auth
##### Create Account
```shell script
$ grpcurl -plaintext \
    -d '{ "email": "pep@test.com", "password": "mypass", "username": "pepsmooth", "given_name": "Pep", "family_name": "Smooth", "nickname": "Pep" }' \
    localhost:8080 template.AuthService/Signup
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzE3MjgsImp0aSI6IjA1NDcyOTU2LTI5YTctNGI4OS1hOGFjLTcyYjFmN2FhM2U5MiIsImlhdCI6MTYyMjgyODEyOCwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.Zh35EwTtHMK0j0rLe_ZSt1eJZpy2lL11Ig0riPLUSfs",
  "refreshToken": "NzQzNDE2NzUtMjU1Ni00M2VlLWE0NzItZTMxMjM4M2NjNzdi",
  "tokenType": "bearer",
  "expiresIn": 3600,
  "refreshExpiresIn": 604800
}
```
##### Logged In User Details
```shell script
$ grpcurl -plaintext \
    -rpc-header "authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzE3MjgsImp0aSI6IjA1NDcyOTU2LTI5YTctNGI4OS1hOGFjLTcyYjFmN2FhM2U5MiIsImlhdCI6MTYyMjgyODEyOCwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.Zh35EwTtHMK0j0rLe_ZSt1eJZpy2lL11Ig0riPLUSfs" \
    localhost:8080 template.AuthService/Me
{
  "id": "05472956-29a7-4b89-a8ac-72b1f7aa3e92",
  "email": "pep@test.com",
  "username": "pepsmooth",
  "givenName": "Pep",
  "familyName": "Smooth",
  "nickname": "Pep"
}
```
##### Login via Refresh Token
```shell script
$ grpcurl -plaintext \
    -d '{ "grant_type": "refresh_token", "username": "pepsmooth", "refresh_token": "NzQzNDE2NzUtMjU1Ni00M2VlLWE0NzItZTMxMjM4M2NjNzdi"}' \
    localhost:8080 template.AuthService/Token
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzE4ODEsImp0aSI6IjM1ODUyYjZlLTdlNmMtNDI4MS1hMzAxLTI3ZWY2YTk5NGY2YyIsImlhdCI6MTYyMjgyODI4MSwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.VtdCVFwd2S9DJGTjmE0sYSWOYl9eZs1qg-5F9444m6M",
  "refreshToken": "ZmZhNGU3MTMtY2YxOS00ODQ5LWE0YWUtOGU2NTBjMjliYWZm",
  "tokenType": "bearer",
  "expiresIn": 3600,
  "refreshExpiresIn": 604800
}
```
##### Login via Password
```shell script
$ grpcurl -plaintext \
    -d '{ "grant_type": "password", "username": "pepsmooth", "password": "mypass"}' \
    localhost:8080 template.AuthService/Token
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzI1MTYsImp0aSI6IjFhYjk0YWEwLTJjOWMtNDY2NC1hODM5LWM2ZWY1ODYzM2RkNiIsImlhdCI6MTYyMjgyODkxNiwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.esN_ipmbbITUlKSwQIo2rgbJHIe-1MTOsYNDcL1-K5o",
  "refreshToken": "MWZmN2E3Y2EtNTE0Ni00Y2E3LTg2M2QtZWU5OTI4NGQzMTMy",
  "tokenType": "bearer",
  "expiresIn": 3600,
  "refreshExpiresIn": 604800
}
```
#### Todos
##### Add a todo entry
```shell script
$ grpcurl -plaintext \
    -rpc-header "authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzE3MjgsImp0aSI6IjA1NDcyOTU2LTI5YTctNGI4OS1hOGFjLTcyYjFmN2FhM2U5MiIsImlhdCI6MTYyMjgyODEyOCwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.Zh35EwTtHMK0j0rLe_ZSt1eJZpy2lL11Ig0riPLUSfs" \
    -d '{"text": "finish everything", "author": "pep"}' \
    localhost:8080 template.TodoService/Create
{
  "id": "6978690c-5a6d-4771-9182-81aaf6fc333e",
  "text": "finish everything",
  "author": "pep",
  "timestamp": "2021-05-28 17:19:29.5131457 +0000 UTC m=+160.747979501"
}
```
##### List my todos
```shell script
$ grpcurl -plaintext \
    -rpc-header "authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzE3MjgsImp0aSI6IjA1NDcyOTU2LTI5YTctNGI4OS1hOGFjLTcyYjFmN2FhM2U5MiIsImlhdCI6MTYyMjgyODEyOCwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.Zh35EwTtHMK0j0rLe_ZSt1eJZpy2lL11Ig0riPLUSfs" \
    localhost:8080 template.TodoService/ListAll
{
  "todos": [
    {
      "id": "6978690c-5a6d-4771-9182-81aaf6fc333e",
      "text": "finish everything",
      "author": "pep",
      "timestamp": "2021-05-28 17:19:29.5131457 +0000 UTC m=+160.747979501"
    }
  ]
}
```
#### REST Gateway
###### [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/) reverse proxy will setup HTTP/1 endpoints for each gRPC method
##### Logged In User Details
```shell script
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzI1MTYsImp0aSI6IjFhYjk0YWEwLTJjOWMtNDY2NC1hODM5LWM2ZWY1ODYzM2RkNiIsImlhdCI6MTYyMjgyODkxNiwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.esN_ipmbbITUlKSwQIo2rgbJHIe-1MTOsYNDcL1-K5o" \
    -X GET localhost:8080/auth/me | json_pp
{
   "username" : "pepsmooth",
   "nickname" : "Pep",
   "id" : "1ab94aa0-2c9c-4664-a839-c6ef58633dd6",
   "family_name" : "Smooth",
   "email" : "pep@test.com",
   "given_name" : "Pep"
}
```
##### List my todos
```shell script
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI4MzI1MTYsImp0aSI6IjFhYjk0YWEwLTJjOWMtNDY2NC1hODM5LWM2ZWY1ODYzM2RkNiIsImlhdCI6MTYyMjgyODkxNiwiaXNzIjoiYXBpIiwic3ViIjoiMmQ5OTYyMDctYTkxMS00MDVlLWI1OTMtMTI1NjQzZmRiYzc4IiwiZW1haWwiOiJwZXBAdGVzdC5jb20iLCJ1c2VybmFtZSI6InBlcHNtb290aCIsIm5hbWUiOiJQZXAgU21vb3RoIiwiZ2l2ZW5fbmFtZSI6IlBlcCIsImZhbWlseV9uYW1lIjoiU21vb3RoIiwibmlja25hbWUiOiJQZXAiLCJwaWN0dXJlIjoiIn0.esN_ipmbbITUlKSwQIo2rgbJHIe-1MTOsYNDcL1-K5o" \
    -X GET localhost:8080/todos | json_pp
{
  "todos": [
    {
      "id": "6978690c-5a6d-4771-9182-81aaf6fc333e",
      "text": "finish everything",
      "author": "pep",
      "timestamp": "2021-05-28 17:19:29.5131457 +0000 UTC m=+160.747979501"
    }
  ]
}
```

#### GraphQL
##### Navigate to http://localhost:8080/graphiql to interact with the GraphQL playground
```graphql
mutation Signup {
  signup(
    input: {
        email: "pep@test.com", 
        username: "pepsmooth", 
        password: "mypass", 
        givenName: "Pep", 
        familyName: "Smooth"
    }
  ) {
    accessToken
    refreshToken
    tokenType
    expires
  }
}
```

#### Swagger
##### Navigate to http://localhost:8080/swagger-ui/ to interact with the REST gateway via swagger
