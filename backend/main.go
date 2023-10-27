package main

import (
	"backend/internal/database"
	"backend/internal/entity"
	"backend/internal/webserver/handlers"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type cfgs struct {
	JWTExpiresIn  int 
	TokenAuth *jwtauth.JWTAuth
};

func main() {
	cfg := cfgs{
		TokenAuth: jwtauth.New("HS256", []byte("teste123"), nil),
		JWTExpiresIn: 3600,
	};

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:@tcp(127.0.0.1:3306)/prog3?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	  }), &gorm.Config{})
	  if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Item{}, &entity.Vote{});

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	  }))
	
	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExpiresIn)

	itemDB := database.NewItemDB(db)
	itemHandler := handlers.NewItemHandler(itemDB)

	voteDB := database.NewVoteDB(db)
	voteHandler := handlers.NewVoteHandler(voteDB)

	FileServer(r, "/images", http.Dir("./internal/webserver/images"))

	r.Route("/users", func(r chi.Router){
		r.Post("/", userHandler.Create)
		r.Post("/login", userHandler.GetJWT)

		r.Group(func (r chi.Router){
			r.Use(jwtauth.Verifier(cfg.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/session", userHandler.Show)
			r.Put("/", userHandler.Update)
		})
	})
	

	r.Route("/itens", func(r chi.Router){
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", itemHandler.List)
		r.Post("/", itemHandler.Create)
	})

	r.Route("/vote", func(r chi.Router){
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", voteHandler.Create)
	})

	http.ListenAndServe(":8080", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}