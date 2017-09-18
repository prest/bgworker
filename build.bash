#!/bin/bash

if [ ! $(which pg_config) ]; then
	echo "ERROR: pg_config not found"
	exit 1
fi

BUILDDIR=".build"

[ -d ${BUILDDIR} ] || mkdir ${BUILDDIR}

INCLUDEDIR=$(pg_config --includedir-server)
LIBDIR=$(pg_config --pkglibdir)
LIBNAME="go_background_worker.so"
LIBOUTPUT="${BUILDDIR}/${LIBNAME}"

export CGO_CFLAGS="-I ${INCLUDEDIR}"
export CGO_LDFLAGS="-shared"

echo "Building ${LIBOUTPUT}"
go build -buildmode=c-shared -o ${LIBOUTPUT}

if [ $? -eq 0 ]; then
	echo "Sending ${LIBOUTPUT} to ${LIBDIR}"
	cp -p ${LIBOUTPUT} ${LIBDIR}
fi
