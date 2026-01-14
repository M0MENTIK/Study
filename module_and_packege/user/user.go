package user

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) User {
	if name == "" {
		return User{}
	}
	if age <= 0 || age >= 150 {
		return User{}

	}
	return User{
		name: name,
		age:  age,
	}
}

func (u *User) SetNewName(name string) {
	if name != "" {
		u.name = name
	}
}
func (u *User) SetNewAge(age int) {
	if age > 0 && age < 150 {
		u.age = age
	}
}
func (u *User) GetName() string {
	return u.name
}
func (u *User) GetAge() int {
	return u.age
}

func helper1() {

}
func helper2() {

}
