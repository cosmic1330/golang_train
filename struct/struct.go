package main

import "fmt"

/* 定義結構 */
type person struct {
	name string
	age  int
}

/* 添加結構方法 */
func (u person) printName() {
	fmt.Println("Name:", u.name)
}
func (u *person) setName(name string) {
	(*u).name = name
}

/* 執行 */
func main() {

	/* 宣告結構 */
	// case1
	var User1 person
	User1.name = "John"
	User1.age = 10
	// case2
	var User2 person = person{name: "John", age: 20}
	// case3
	var User3 person = person{"John", 30}
	// case4
	User4 := person{name: "John", age: 40}
	fmt.Println(User1, User2, User3, User4)
	fmt.Println()

	/* 結構指針 */
	var User5 *person = &User1
	User5.name = "Mary"
	(*User5).age = 50
	fmt.Println(User1, User5)
	fmt.Println()

	/* 繼承 */
	type student struct {
		person
		grade int
	}
	var student1 student = student{User2, 3}
	fmt.Println("student1:", student1)
	fmt.Println("student1.name:", student1.name) // same as student1.person.name
	fmt.Println()

	/* 指針繼承 */
	type studentpointer struct {
		*person
		grade int
	}

	var student2 *studentpointer = &studentpointer{
		person: &person{
			name: "John",
			age:  10,
		},
		grade: 3,
	}
	fmt.Println("student2:", student2)
	fmt.Print("student2.name:", student2.name)
	fmt.Println()

	/* 結構標籤 */
	type studentlabel struct {
		name string `json:"name"`
		age  int    `json:"age"`
	}
	var student3 studentlabel = studentlabel{"John", 10}
	fmt.Println("student3:", student3)
	fmt.Println()

	/* 執行結構方法 */
	User1.printName()
	User1.setName("Jeff")
	fmt.Println(User1)
}
