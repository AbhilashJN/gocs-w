package helpers

import "fmt"

func RenderMapping(m map[string]int, title string) {
	fmt.Println("==================================")
	fmt.Println(title, ":")
	for k, v := range m {
		fmt.Printf("=  %s : %d\n", k, v)
	}
	fmt.Println("==================================")

	fmt.Println("")

}
