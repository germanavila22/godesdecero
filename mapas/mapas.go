package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)
	paises["mexico"] = "df"
	paises["argentina"] = "buenos aires"

	fmt.Println(paises)
	fmt.Println(paises["argentina"])
	campeonato := map[string]int{
		"barcelona":   39,
		"otro":        50,
		"otro mas":    55,
		"boca junion": 553,
	}
	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s , tiene un puntahe de %d \n", equipo, puntaje)

	}
	fmt.Println("borro barcelona")
	delete(campeonato, "barcelona")

	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s , tiene un puntahe de %d \n", equipo, puntaje)
	}
	puntaje, existe := campeonato["otro"]
	fmt.Printf(" el puntaje es de %d  y el equipo existe %t ", puntaje, existe)
}
