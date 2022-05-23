package main

import(
	"fmt" 
)

func main() {
	slice := map[string]int{"foo":20,"poo":30,"doo":40}

	for _, v := range slice {
		if v == 20 {
			continue
		}
		fmt.Println(v)
	}
}
