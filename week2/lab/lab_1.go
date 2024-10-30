package lab

import "fmt"

type User struct {
	Name string
	Age  int
}

func change(a *User, b *User) {
	temp := a.Name
	a.Name = b.Name
	b.Name = temp
}

func sorting(list []User) {

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j].Name > list[j+1].Name {
				change(&list[j], &list[j+1])
			}
		}
	}
}

func Lab_1() {
	list := []User{
		{"Paul", 19},
		{"John", 21},
		{"Jane", 35},
		{"Abraham", 25},
	}

	sorting(list)

	for _, user := range list {
		fmt.Println(user.Name)
	}
}
