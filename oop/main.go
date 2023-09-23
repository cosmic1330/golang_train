package main

import "fmt"

type Person struct {
	Name string
	password string
}

func (p *Person) SetPassword(password string) {
	p.password = password
}

// 工廠模式
func NewPersion(name string) *Person {
	return &Person{
		Name: name,
	}
}

// 繼承看struct
// 多型看interface

func main() {
	p1 := NewPersion("Mary")
	fmt.Println(p1.Name)
}