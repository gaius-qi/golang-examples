package main

import (
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/gaius-qi/golang-examples/machinery/etask"
)

func main() {
	globalServer, err := startServer("global_tasks")
	if err != nil {
		panic(err)
	}
	globalWorker1 := globalServer.NewWorker("worker1", 10)
	go func() {
		if err := globalWorker1.Launch(); err != nil {
			panic(err)
		}
		globalWorker2 := globalServer.NewWorker("worker2", 10)
		if err := globalWorker2.Launch(); err != nil {
			panic(err)
		}
	}()

	worker1Server, err := startServer("worker1_tasks")
	if err != nil {
		panic(err)
	}
	worker1 := worker1Server.NewWorker("worker1", 10)
	go func() {
		if err := worker1.Launch(); err != nil {
			panic(err)
		}
	}()

	worker2Server, err := startServer("worker2_tasks")
	if err != nil {
		panic(err)
	}
	worker2 := worker2Server.NewWorker("worker2", 10)
	if err := worker2.Launch(); err != nil {
		panic(err)
	}
}

func startServer(queue string) (*machinery.Server, error) {
	var cnf = &config.Config{
		Broker:        "redis://dragonfly2_dev:dragonfly2_dev@127.0.0.1:6379/1",
		DefaultQueue:  queue,
		ResultBackend: "redis://dragonfly2_dev:dragonfly2_dev@127.0.0.1:6379/2",
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Fatal(err)
	}

	etasks := map[string]interface{}{
		"Add": etask.Add,
		"Sub": etask.Sub,
	}

	return server, server.RegisterTasks(etasks)
}
