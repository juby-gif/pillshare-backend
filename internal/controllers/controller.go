package controllers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/internal/repositories"
)

type Controller struct {
	db       *sql.DB
	UserRepo models.UserRepo
}

func New(db *sql.DB) *Controller {
	userRepo := repositories.NewUserRepo(db)
	return &Controller{
		db:       db,
		UserRepo: userRepo,
	}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	URL := strings.Split(r.URL.Path, "/")[1:]
	n := len(URL)

	switch {
	case n == 3 && URL[2] == "version" && r.Method == "GET":
		c.getVersion(w, r)
	case n == 3 && URL[2] == "login" && r.Method == "POST":
		c.postLogin(w, r)
	case n == 3 && URL[2] == "register" && r.Method == "POST":
		c.postRegister(w, r)
	}
}
