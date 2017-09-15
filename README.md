# PostgreSQL BGWorker

Background Worker Processes for PostgreSQL written in Go


## Build

```
go build -buildmode=c-shared -o prest.so
cp prest.so /usr/lib/postgresql/9.6/lib/prest
```
