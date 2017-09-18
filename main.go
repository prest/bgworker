package main

/*
extern void BackgroundWorkerMain();
void elog_log(char *string);
*/
import "C"

func main() {}

//export BackgroundWorkerMain
func BackgroundWorkerMain(){
	C.elog_log(C.CString("Hello World from GoLang!!! Yeahhhhh!!!!"));
}

