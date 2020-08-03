package tasks

import (
	"github.com/zkynet/golang-lessons-for-beginners/src/users"
)

type Task struct {
	Title string
	Done    bool
	D    string
	User  users.User
}
