package main

import (
	"fmt"
	"learngo.com/crawler_distributed/config"
	"learngo.com/crawler_distributed/rpcsupport"
	"learngo.com/crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
