package main

import (
	"log"
	"net/http"

	"github.com/felipecardosodeoliveira/Golang/12-apis/internal/entity"
	"github.com/felipecardosodeoliveira/Golang/12-apis/internal/infra/database"
	"github.com/felipecardosodeoliveira/Golang/12-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/felipecardosodeoliveira/Golang/12-apis/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/felipecardosodeoliveira/Golang/12-apis/configs"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Felipe
// @contact.url    https://github.com/felipecardosodeoliveira
// @contact.email  felipecardosodeoliveira02@gmail.com

// @license.name   Felipe Cardoso
// @license.url    https://github.com/felipecardosodeoliveira

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB /*conf.TokenAuth, conf.JwtExpiresIn*/)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwt", conf.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", conf.JwtExpiresIn))
	// r.Use(LogRequest)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/login", userHandler.GetJWT)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(conf.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	// r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	// ))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
