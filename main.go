package main

/*
extern void BackgroundWorkerMain();
*/
import "C"
import "github.com/prest/bgworker/pg/log"

func main() {}

//export BackgroundWorkerMain
func BackgroundWorkerMain() {
	log.Info("Hello World from GoLang!!! Yeahhhhh!!!!")
}
