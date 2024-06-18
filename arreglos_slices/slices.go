package arreglos_slices

import "fmt"

var tabla_slice []int = []int{1, 2, 3}
var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MuestroSlice() {
	fmt.Println(tabla_slice)
	porsion := arreglo[3:]
	porsion2 := arreglo[:5]
	porsion3 := arreglo[6:7]

	fmt.Println(porsion, porsion2, porsion3)
}
func Capacidad() {
	elemento := make([]int, 5, 20)

	fmt.Printf("Largo %d , capacidad %d", len(elemento), cap(elemento))
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d , capacidad %d", len(nums), cap(nums))
}
