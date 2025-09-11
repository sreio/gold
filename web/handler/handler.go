package handler

import (
	"github.com/sreio/gold/web/handler/Auth"
	"github.com/sreio/gold/web/handler/gold"
	"github.com/sreio/gold/web/handler/notification"
	"github.com/sreio/gold/web/handler/source"
	"github.com/sreio/gold/web/handler/task"
	"github.com/sreio/gold/web/handler/user"
)

type Handler struct {
	Auth         Auth.Auth
	Source       source.Source
	User         user.User
	Gold         gold.Gold
	Task         task.Task
	Notification notification.Notification
}

func NewHandler() *Handler {
	return &Handler{
		Auth:         Auth.Auth{},
		Source:       source.Source{},
		User:         user.User{},
		Gold:         gold.Gold{},
		Task:         task.Task{},
		Notification: notification.Notification{},
	}
}
