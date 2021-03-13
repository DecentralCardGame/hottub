# Hottub
Hottub - The glue between experience and management

## Generating Documentation
Before you can compile hottub, you first have to run the following commands
```shell script
go get -u github.com/swaggo/swag/cmd/swag
swag init
```

## Dump PostgreSQL
```shell script
docker-compose exec database pg_dumpall -c -U crowdcontrol > dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql
```
