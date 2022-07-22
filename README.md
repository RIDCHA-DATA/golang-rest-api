# golang-rest-api

## get all deps
```bash
GO111MODULE=on go get
```
## Start the api
```bash
GO111MODULE=on go run main.go
```

>> swagger docs url : http://localhost:8081/v1/api-docs/index.html


![alt text](https://lh3.googleusercontent.com/chat_attachment/ADMKSccW-tG4ZNLj8491E84QuZ4ouCvdCLQrXV6EeMir12AP4g0cLN9dAIMbdbUMbgTwuBmD350A5wYqGTRMvQJmwLB0XkzDe3-e-cxq8g3ks8-L8Gl6qN-zJi1iWNFZRbIxLW30oH3cRY3EewBT838g9-sdhlwYmuovFI8vjC-4k_cTOxh2qJkr55EtTAq_pwFbLqUaeWvLl1DyQrU34YU=w512 "Title")
# Generate Swagger docs (version 1.7.9) ref: https://github.com/swaggo/swag/issues/1126
```bash
swag init --parseDependency --parseInternal -g main.go
```

# E2E test with curl

## Get all actions

```bash
curl -X 'GET'   'http://localhost:8081/v1/actions'   -H 'accept: application/json'
```

## Get all actions

```bash
curl -X 'POST'   'http://localhost:8081/v1/actions'   -H 'accept: application/json'  -d '{"action": "TestAddAction"}'
```


# Todo list

- [ x ] Check & fix pb with go 1.18
- [ x ] Check migration on sqllite
- [ x ] Prepare dev env with psql
- [ x ] Add CI process
    - [ x ] go test & build
    - [ x ] Dcoker image build
    - [ ] Helm chart package/build