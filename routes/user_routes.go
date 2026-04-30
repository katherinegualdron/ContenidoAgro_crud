// package routes

// import (
// 	"API_GO_CRUD/controllers"
// 	"github.com/gorilla/mux"
// )

// func RegisterUserRoutes(r *mux.Router){
// 	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
// 	r.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
// 	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
// 	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
// 	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
// }