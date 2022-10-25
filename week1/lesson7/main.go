package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

// NewUser функція конструктор
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		weight: weight,
		height: height,
	}
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{m: make(map[string]string)}
}

func main() {

	user3 := NewUser("Sasha", "Male", 40, 81, 186)

	user1 := User{"Vasya", 23, "Male", 75, 185}
	//user2 := User{"Petya", 23, "Male", 84, 197}

	fmt.Println(user1)
	fmt.Printf("%+v\n", user3)

	user3.printUserInfo()
}
