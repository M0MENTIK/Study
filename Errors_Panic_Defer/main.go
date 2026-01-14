package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) error {
	if user.Ballance-usd < 0 {
		return errors.New("Don't have enoght money")
	}
	if user.Ballance-usd >= 0 {
		user.Ballance -= usd
	}
	return nil
}

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor-10 <= 0 {
		return 0, errors.New("Don't gas. Because car breaking")
	}
	kmmh := rand.Intn(150)

	c.Armor -= 10

	return kmmh, nil
}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Was panic:", p)
		}
	}()

	slice := []int{1, 2, 3}
	fmt.Println(slice[4])

	car := Car{
		Armor: 25,
	}

	for {
		pp.Println("Car before:", car)
		kmch, err := car.Gas()
		pp.Println("Car after:", car)
		if err != nil {
			fmt.Println("Problem with gas:", err)
			break
		}
		fmt.Println("Pazgon:", kmch)
	}

	//

	//

	//

	//

	// user := User{
	// 	Name:     "Oleg",
	// 	Ballance: 300,
	// }
	// pp.Println("User before:", user)
	// err := Pay(&user, 300)
	// pp.Println("User after:", user)

	// if err != nil {
	// 	fmt.Println("Payment doesn't have. Problem:", err.Error())
	// } else {
	// 	fmt.Println("Was payment")

	// }
}
