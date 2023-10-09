package main

import (
	"employee-service/internal/app/database"
	"employee-service/internal/app/handler"
	"employee-service/internal/app/middleware"
	"employee-service/internal/app/repository"
	"employee-service/internal/app/usecase"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	repo := repository.NewEmployeeRepository()
	usecase := usecase.NewEmployeeUsecase(repo)
	handler := handler.NewEmployeeHandler(usecase)

	mux := http.NewServeMux()

	mux.Handle("/hire", middleware.ContentTypeMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handler.HireEmployee))))
	mux.Handle("/fire", middleware.ContentTypeMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handler.FireEmployee))))
	mux.Handle("/vacation", middleware.ContentTypeMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handler.GetVacationDays))))
	mux.Handle("/find", middleware.ContentTypeMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handler.FindEmployee))))

	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
