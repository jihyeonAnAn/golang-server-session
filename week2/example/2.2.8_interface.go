package example

import "fmt"

// 인터페이스가 없는 구조체
type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name  string
	Owner *Person
	Age   int
}

func (person *Person) incrementAge() {
	person.Age++
}

func (person *Person) getAge() int {
	return person.Age
}

func (dog *Dog) incrementAge() {
	dog.Age++
}

func (dog *Dog) getAge() int {
	return dog.Age
}

// 두 개의 구조체를 표준화하는 인터페이스
type Living interface {
	incrementAge()
	getAge() int
}

func incrementAndPrintAge(being Living) {
	being.incrementAge()
	fmt.Println(being.getAge())
}

func Code227() {
	harley := Person{"Harley", 21}
	snowy := Dog{"Snowy", &harley, 6}
	incrementAndPrintAge(&harley)
	incrementAndPrintAge(&snowy)
}
