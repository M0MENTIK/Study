package main

import (
	"fmt"
	"study/user"
)

func main() {
	u := user.User{}
	//u.Name = "Bob"
	//u.Age = 23
	fmt.Println(u)

	u2 := user.NewUser("alex", 45)
	fmt.Println(u2.GetName())
	fmt.Println(u2.GetAge())
	u2.SetNewAge(34)
	fmt.Println(u2.GetAge())

}
