extern int BackgroundWorkerMain();
void elog_log(char *string);
int wait_latch(long miliseconds);
void reset_latch(void);
int postmaster_is_dead(int rc);
int get_got_sigterm();
