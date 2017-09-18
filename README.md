# PostgreSQL BGWorker

Background Worker Processes for PostgreSQL written in Go


## Build

1) Make sure the 'pg\_config' tool is in your $PATH, if not then config like the example below:

```
export PATH=/usr/lib/postgresql/9.6/bin:$PATH
```

2) Build the library

```
./build.bash
```

## PostgreSQL Config (postgresql.conf)

```
shared_preload_libraries = 'go_brackground_worker'
```
