package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"someApp/internal/controller"
	router "someApp/internal/http/mux-router"
	"someApp/internal/repository/postgres"
	"someApp/internal/service"
	"someApp/pkg/postgresql"

	"someApp/schema/migrations"
)

func main() {
	mg, err := migrations.NewMigration("localhost", "5432",
		"postgres", "postgres", "postgres")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := mg.Up(); err != nil {
		fmt.Println(err)
	}

	conn, err := postgresql.NewConnection("localhost", "5432",
		"postgres", "postgres", "postgres")
	if err != nil {
		fmt.Println(errors.Wrap(err, fmt.Sprintf("error in /create: %v\n", err)))
		return
	}
	postRepo := postgres.NewPostgresRepo(conn)
	defer postRepo.CloseConnection(context.Background())
	postService := service.NewPostService(postRepo)
	postController := controller.NewController(postService)
	postRouter := router.NewMuxRouter()

	postRouter.POST("/create", postController.CreatePerson)
	postRouter.GET("/get/{id}", postController.GetPersonByID)
	postRouter.PUT("/update/{id}", postController.UpdatePersonByID)
	postRouter.DELETE("/delete/{id}", postController.DeletePersonByID)
	postRouter.SERVE(":8181")
}
