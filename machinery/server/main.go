package main

import (
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gaius-qi/golang-examples/machinery/etask"
)

func main() {
	server, err := startServer()
	if err != nil {
		panic(err)
	}

	send(server, "global_tasks")
	send(server, "worker1_tasks")
}

func startServer() (*machinery.Server, error) {
	var cnf = &config.Config{
		Broker:        "redis://dragonfly2_dev:dragonfly2_dev@127.0.0.1:6379/1",
		DefaultQueue:  "global_tasks",
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

func send(server *machinery.Server, queue string) {
	addAll := tasks.Signature{
		Name:       "Add",
		RoutingKey: queue,
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	addAllAsyncResult, err := server.SendTask(&addAll)
	if err != nil {
		log.Fatal(err)
	}

	res, err := addAllAsyncResult.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("get return %v\n", tasks.HumanReadableResults(res))
}
