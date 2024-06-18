package modelos

import (
	"time"
)

type User struct {
	id        int
	Name      string
	Createdat time.Time
	status    bool
}

func (usuario *User) AddUser(id int, name string, createdat time.Time, status bool) {
	usuario.id = id
	usuario.Name = name
	usuario.Createdat = createdat
	usuario.status = status
}
