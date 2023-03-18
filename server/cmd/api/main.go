package main

import (
	"flag"
	_ "github.com/lib/pq"
	"kristianpilegaard.dk/todo-app/pkg/api"
)

func main() {
	var cfg api.TodoAppConfig

	flag.IntVar(&cfg.ListenPort, "listenPort", 4000, "listen port")
	flag.IntVar(&cfg.PostgresqlPort, "postgresqlPort", 5432, "postgresql port")
	flag.Parse()

	api.StartTodoApp(cfg)

}
