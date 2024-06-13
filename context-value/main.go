package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strings"
)

type Logic interface {
	BusinessLogic(ctx context.Context, user string, data string) (string, error)
}
type Controller struct {
	Logic Logic
}

// Login implements the worst authentication system known.
func (c Controller) Login(rw http.ResponseWriter, req *http.Request) {
	userName := req.URL.Query().Get("user")
	if len(strings.TrimSpace(userName)) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := rw.Write([]byte("No user specified"))
		if err != nil {
			return
		}
		return
	}
	SetUser(userName, rw)
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("user logged in"))
	if err != nil {
		return
	}
}

func (c Controller) DoLogic(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user, ok := UserFromContext(ctx)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := req.URL.Query().Get("data")
	result, err := c.Logic.BusinessLogic(ctx, user, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, err := rw.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	_, err = rw.Write([]byte(result))
	if err != nil {
		return
	}
}

func (c Controller) Logout(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := UserFromContext(ctx)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	DeleteUser(rw)
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("user logged out"))
	if err != nil {
		return
	}
}

type LogicImpl struct{}

func (l LogicImpl) BusinessLogic(ctx context.Context, user string, data string) (string, error) {
	return fmt.Sprintf("Hello %s, thank you for sending me %s", user, data), nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	controller := Controller{
		Logic: LogicImpl{},
	}
	r.Get("/login", controller.Login)
	r.Route("/business", func(r chi.Router) {
		r = r.With(Middleware)
		r.Get("/", controller.DoLogic)
		r.Get("/logout", controller.Logout)
	})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
