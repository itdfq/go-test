package main

import "fmt"

type User struct {
	id   int
	name string
}
type Manager struct {
	User
	address string
	title   string
}

func (self *User) ToString() string {
	return fmt.Sprintf("User:%p,%v", self, self)
}

func (self *Manager) ToString() string { //receiver = &(Manager.User)
	return fmt.Sprintf("Manager: %p, %v", self, self)
}
func main() {
	m := Manager{User{1, "Tom"}, "上海市", "题目"}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
	id := m.id
	fmt.Println(id)
	fmt.Println(m.address)
}
