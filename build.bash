#!/bin/bash

PGROOT="/home/fabrizio/pgsql"

LIBNAME="go_background_worker.so"

export CGO_CFLAGS="-I ${PGROOT}/include/server"
export CGO_LDFLAGS="-shared"

go build -buildmode=c-shared -v -o ${PGROOT}/lib/${LIBNAME}
