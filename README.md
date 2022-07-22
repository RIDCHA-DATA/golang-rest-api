# golang-rest-api

## get all deps
```bash
go get
```
## Start the api
```bash
go run main.go
```

# Todo list

- [ ] Check & fix pb with go 1.18
- [ ] Check migration on sqllite
- [ ] Prepare dev env with psql
- [ ] Add CI process
    - [ ] go test & build
    - [ ] Dcoker image build
    - [ ] Helm chart package/build


commands:

## Start the api
```bash
GO111MODULE=on go run main.go
```

# Generate Swagger docs (version 1.7.9) ref: https://github.com/swaggo/swag/issues/1126
```bash
swag init --parseDependency --parseInternal -g main.go
```