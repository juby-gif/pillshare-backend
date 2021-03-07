package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	cache "github.com/go-redis/cache/v8"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/internal/repositories"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

type Controller struct {
	db       *sql.DB
	UserRepo models.UserRepo
	cache    *cache.Cache
}

func New(db *sql.DB) *Controller {
	userRepo := repositories.NewUserRepo(db)
	cache := utils.RedisCache()
	return &Controller{
		db:       db,
		UserRepo: userRepo,
		cache:    cache,
	}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get URL,n(length),authStatus and sessionUUID from the context
	URL := ctx.Value("url_split").([]string)
	n := ctx.Value("length").(int)
	authStatus := ctx.Value("is_authorized").(bool)
	sessionUUID := ctx.Value("session_uuid").(string)

	// Get User modal from cache
	var user models.User
	if err := c.cache.Get(ctx, sessionUUID, &user); err == nil {
		fmt.Println(user)
	}

	switch {
	case n == 3 && URL[2] == "version" && r.Method == "GET":
		c.getVersion(w, r)
	case n == 3 && URL[2] == "login" && r.Method == "POST":
		c.postLogin(w, r)
	case n == 3 && URL[2] == "register" && r.Method == "POST":
		c.postRegister(w, r)
	case n == 3 && URL[2] == "hello" && r.Method == "GET":
		if authStatus != true {
			http.Error(w, "You are not authorized", http.StatusUnauthorized)
		} else {
			c.getHello(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}
