package log

/*
#include "postgres.h"
#include "utils/elog.h"

void elog_start(const char *filename, int lineno, const char *funcname);
void elog_finish(int elevel, const char *fmt,...) pg_attribute_printf(2, 3);

void elog_log(char *string) {
	elog(LOG, string, "");
}
*/
import "C"

// Info : Throws a message in the log file
func Info(message String) error {

	if err := C.elog_log(C.CString(message)); err != nil {
		return err
	}

	return nil
}
