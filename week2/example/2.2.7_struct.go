package example

import "fmt"

type User struct {
	Name string
	Age  int
}

// 구조체를 이용하여 여러 개의 반환 값 변경하기
func nameAndAge(uid int) User {
	switch uid {
	case 0:
		return User{"Baheer Kamal", 24}
	case 1:
		return User{"Tanmay Bakshi", 16}
	default:
		return User{"", -1}
	}
}

func Code219() {
	user := nameAndAge(1)
	fmt.Println(user)
}

// 함수에 인자로 전달된 구조체 변화
func incrementAge(user User) {
	user.Age++
	fmt.Println(user.Age)
}

func Code220() {
	kathy := User{"Kathy", 19}
	incrementAge(kathy)
	fmt.Println(kathy.Age)
}

// 함수에 포인터로 전달된 구조체의 변화
func incrementAgePointer(user *User) {
	user.Age++
	fmt.Println(user.Age)
}

func Code221() {
	kathy := User{"Kathy", 19}
	incrementAgePointer(&kathy)
	fmt.Println(kathy.Age)
}

// 값 리시버를 이용한 구조체 변화
func (user User) incrementAge() {
	user.Age++
	fmt.Println(user.Age)
}

func Code222() {
	kathy := User{"Kathy", 19}
	kathy.incrementAge()
	fmt.Println(kathy.Age)
}

// 포인터 리시버를 이용한 구조체 변화
func (user *User) incrementAgePointerReceiver() {
	user.Age++
	fmt.Println(user.Age)
}

func Code223() {
	kathy := User{"Kathy", 19}
	kathy.incrementAgePointerReceiver()
	fmt.Println(kathy.Age)
}
