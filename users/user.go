package users

import (
	"fmt"
	"time"

	"github.com/AngelTravieso/go/modelos"
)

func AltaUsuario() {

	user := new(modelos.User)

	user.AddUser(10, "angel", time.Now(), true)

	fmt.Println(user)

}
