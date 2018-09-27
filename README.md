# minerserver

REST API server dedicated to serve [minerclient](https://bitbucket.org/boomstarternetwork/minerclient/src).

For now it has only one method `/projects/list` which returns all projects
list registered in `miningcore` postges database.

This project intended to be run in docker container like it done in 
[miningpool](https://bitbucket.org/boomstarternetwork/miningpool/src/master/).

## How to install

```bash
dep ensure -v
go install -v .
```

## How to run

It has required environment variable `MINERSERVER_POSTGRES_CONNECTION_STRING`
like:
```
postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full
```
or
```
user=pqgotest password=password dbname=pqgotest sslmode=verify-full
```

To run locally:
```
MINERSERVER_POSTGRES_CONNECTION_STRING=<your-string> minerserver
```