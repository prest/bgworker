package main

/*
#include "main_c.h"
*/
import "C"

func main() {}

//export BackgroundWorkerMain
func BackgroundWorkerMain() C.int {

	C.elog_log(C.CString("Starting GoBackgroundWorker"))

	for C.get_got_sigterm() == 0 {
		var rc C.int

		/* Wait 10s */
		rc = C.wait_latch(C.long(10000))

		C.reset_latch()

		/* Emergency bailout if postmaster has died */
		if C.postmaster_is_dead(rc) != 0 {
			C.elog_log(C.CString("Aborting GoBackgroundWorker"))
			return 1
		}

		C.elog_log(C.CString("Hello World from GoLang!!! Yeahhhhh!!!!"));
	}

	C.elog_log(C.CString("Finishing GoBackgroundWorker"))
	return 0
}

