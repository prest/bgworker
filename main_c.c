/* PostgreSQL related headers */
#include "postgres.h"
#include "pgstat.h"
#include "postmaster/bgworker.h"
#include "storage/ipc.h"
#include "storage/latch.h"
#include "storage/proc.h"
#include "utils/elog.h"
#include "fmgr.h"

/* Go related headers */
#include "_cgo_export.h"

PG_MODULE_MAGIC;

/* Entry point of library loading */
void _PG_init(void);

/* Main loop of process */
void background_main(Datum main_arg) pg_attribute_noreturn();

/* Wrappers to using in CGO */
void elog_log(char *string) {
	elog(LOG, string, "");
}

/* Signal handling */
static volatile sig_atomic_t got_sigterm = false;

/*
 * background_sigterm
 *
 * SIGTERM handler.
 */
static void
background_sigterm(SIGNAL_ARGS)
{
	int save_errno = errno;
	got_sigterm = true;
	if (MyProc)
		SetLatch(&MyProc->procLatch);
	errno = save_errno;
}

/*
 * background_main
 *
 * Main loop processing.
 */
void
background_main(Datum main_arg)
{
	/* Set up the sigterm signal before unblocking them */
	pqsignal(SIGTERM, background_sigterm);

	/* We're now ready to receive signals */
	BackgroundWorkerUnblockSignals();
	while (!got_sigterm)
	{
		int rc;

		/* Wait 10s */
		rc = WaitLatch(&MyProc->procLatch,
					   WL_LATCH_SET | WL_TIMEOUT | WL_POSTMASTER_DEATH,
					   10000L,
					   PG_WAIT_EXTENSION);
		ResetLatch(&MyProc->procLatch);

		/* Emergency bailout if postmaster has died */
		if (rc & WL_POSTMASTER_DEATH)
			proc_exit(1);

		BackgroundWorkerMain();

	}
	proc_exit(0);
}

void _PG_init(void) {
	BackgroundWorker worker;

	/* Register the worker processes */
	MemSet(&worker, 0, sizeof(BackgroundWorker));
	worker.bgw_flags = BGWORKER_SHMEM_ACCESS;
	worker.bgw_start_time = BgWorkerStart_RecoveryFinished;
	snprintf(worker.bgw_library_name, BGW_MAXLEN, "go_background_worker");
	snprintf(worker.bgw_function_name, BGW_MAXLEN, "background_main");
	snprintf(worker.bgw_name, BGW_MAXLEN, "GoBackgroundWorker");
	worker.bgw_restart_time = BGW_NEVER_RESTART;
	worker.bgw_main_arg = (Datum) 0;
	worker.bgw_notify_pid = 0;
	RegisterBackgroundWorker(&worker);
}
