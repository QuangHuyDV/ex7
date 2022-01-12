package main

import "fmt"

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) GetAge() int64 {
	return u.age
}
func (u *User) GetGender() bool {
	return u.gender
}
func (u *User) GetAddress() string {
	return u.address
}

func main() {
	user1 := User{name: "Trung", age: 12, gender: true, address: "Hue"}
	user2 := User{name: "Hoa", age: 15, gender: true, address: "HN"}
	user3 := User{name: "Ha", age: 13, gender: false, address: "HN"}
	user4 := User{name: "Trang", age: 12, gender: true, address: "HN"}
	user5 := User{name: "Minh", age: 13, gender: false, address: "HN"}
	user6 := User{name: "Lan", age: 11, gender: true, address: "HN"}

	a := map[User]string{user1: user1.name, user2: user2.name, user3: user3.name, user4: user4.name, user5: user5.name, user6: user6.name}

	for str, i := range a {
		fmt.Println(str, i)
	}
}
