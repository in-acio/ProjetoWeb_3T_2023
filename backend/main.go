package main

import (
	"backend/internal/database"
	"backend/internal/entity"
	"backend/internal/webserver/handlers"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
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

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Item{}, &entity.Vote{});

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExpiresIn)

	itemDB := database.NewItemDB(db)
	itemHandler := handlers.NewItemHandler(itemDB)

	voteDB := database.NewVoteDB(db)
	voteHandler := handlers.NewVoteHandler(voteDB)

	FileServer(r, "/images", http.Dir("./internal/webserver/images"))

	r.Post("/users", userHandler.Create)
	r.Post("/users/login", userHandler.GetJWT)

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