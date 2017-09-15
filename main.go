package main

/*
#cgo CFLAGS: -I/usr/include/postgresql/9.6/server
#cgo LDFLAGS: -L/usr/lib/postgresql/9.6/lib

#include "postgres.h"
#include "pgstat.h"
#include "postmaster/bgworker.h"
#include "storage/ipc.h"
#include "storage/latch.h"
#include "storage/proc.h"
#include "fmgr.h"
#include "utils/elog.h"

#ifdef PG_MODULE_MAGIC
PG_MODULE_MAGIC;
#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var worker C.BackgroundWorker

func main() {
	_PG_init()
}

func _PG_init() {
	worker.bgw_flags = C.BGWORKER_SHMEM_ACCESS
	worker.bgw_start_time = C.BgWorkerStart_RecoveryFinished

	bgw_name := unsafe.Pointer(C.CString("prest"))
	defer C.free(bgw_name)
	bgw_function_name := unsafe.Pointer(C.CString("prest_main"))
	defer C.free(bgw_function_name)

	worker.bgw_name = *(*[C.BGW_MAXLEN]C.char)(bgw_name)
	worker.bgw_library_name = *(*[C.BGW_MAXLEN]C.char)(bgw_name)
	worker.bgw_function_name = *(*[C.BGW_MAXLEN]C.char)(bgw_function_name)

	worker.bgw_restart_time = C.BGW_NEVER_RESTART
	worker.bgw_notify_pid = 0
	worker.bgw_main_arg = 0
	// TODO: undefined reference to 'RegisterBackgroundWorker' (not found)
	C.RegisterBackgroundWorker(&worker)
	fmt.Println("Hi!!", worker)
}
