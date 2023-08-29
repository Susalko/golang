package main

import (
	"fmt"
	cashe "golang/HW/HW/Cashe"
)

func main() {
	cashe := cashe.New()

	cashe.Set("userId", 42)
	userId, _ := cashe.Get("userId")

	fmt.Println(userId)

	cashe.Delete("userId")
	userId, _ = cashe.Get("userId")

	fmt.Println(userId)

}
