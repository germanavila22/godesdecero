package variables

import (
	"fmt"
)

func Muestroenteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)
	fmt.Println("int comun=", intcomun)
	fmt.Println("int comun=", intde32)
	fmt.Println("int comun=", intde64)
}
