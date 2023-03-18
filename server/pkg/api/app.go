package api

import (
	"fmt"
	"kristianpilegaard.dk/todo-app/pkg/db"
	"kristianpilegaard.dk/todo-app/pkg/todo"
	"net/http"
	"time"
)

type TodoAppConfig struct {
	ListenPort     int
	PostgresqlPort int
}

type todoApp struct {
	config        TodoAppConfig
	todoComponent todo.Component
}

func newTodoApp(conf TodoAppConfig, component todo.Component) *todoApp {
	return &todoApp{config: conf,
		todoComponent: component}
}

func StartTodoApp(config TodoAppConfig) {
	dbConn, err := db.ConnectAndMigrate(config.PostgresqlPort)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	component := todo.NewTodoComponent(dbConn)
	err = component.PrepareDB()
	if err != nil {
		panic(fmt.Sprintf("failed when preparing DB: %s", err))
	}

	app := newTodoApp(config, component)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", config.ListenPort),
		Handler:      app.routes(),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}
	fmt.Printf("starting server on listenPort %d\n", config.ListenPort)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
