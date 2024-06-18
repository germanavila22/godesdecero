package user

import (
	"fmt"
	"time"

	"github.com/godesdecero/godesdecero/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "pablo", time.Now(), true)
	fmt.Println(u)
}
