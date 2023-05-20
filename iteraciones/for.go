package iteraciones

import (
	"fmt"
)

func Iterar() {

	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := 10; i > 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}

}
