package main

import (
	"backend/configs"
	"backend/internal/database"
	"backend/internal/entity"
	"backend/internal/webserver/handlers"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName),
		DefaultStringSize: 256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex: true,
		DontSupportRenameColumn: true,
		SkipInitializeWithVersion: false,
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
	itemHandler := handlers.NewItemHandler(itemDB, cfg.ImagesFolder)

	voteDB := database.NewVoteDB(db)
	voteHandler := handlers.NewVoteHandler(voteDB)

	FileServer(r, "/images", http.Dir(cfg.ImagesFolder))

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
		r.Get("/ranking", itemHandler.Ranking)

		r.Group(func(r chi.Router){
			r.Use(jwtauth.Verifier(cfg.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Use(isAdmin)
			r.Post("/", itemHandler.Create)
			r.Put("/{id}", itemHandler.Update)
			r.Delete("/{id}", itemHandler.Delete)
		})

		r.Group(func (r chi.Router){
			r.Use(jwtauth.Verifier(cfg.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/", itemHandler.List)
			r.Get("/{id}", itemHandler.Show)
		})
	})

	r.Route("/vote", func(r chi.Router){
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", voteHandler.Create)
		r.Get("/", voteHandler.List)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), r)
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

func isAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())

		if isAdminClaim, ok := claims["isAdmin"].(bool); !ok || !isAdminClaim {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}