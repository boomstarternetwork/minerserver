# minerserver

REST API server dedicated to serve [minerclient](https://github.com/boomstarternetwork/minerclient).

For now it has only one method `/projects` which returns all projects
list registered in `miningcore` postges database.

## How to install

```bash
dep ensure -v
go install -v .
```

## How to run

Read help:

```bash
minerserver --help
```