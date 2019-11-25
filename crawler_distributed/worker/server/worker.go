package main

import (
	"flag"
	"fmt"
	"learngo.com/crawler_distributed/rpcsupport"
	"learngo.com/crawler_distributed/worker"
	"log"
)

// 命令行参数
var port = flag.Int("port", 0, "The port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
