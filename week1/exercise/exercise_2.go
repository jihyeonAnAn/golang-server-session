package exercise

import "fmt"

func Exec2() {
	var number int

	fmt.Println("숫자를 입력하세요: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println(number, "는 짝수입니다.")
	} else {
		fmt.Println(number, "는 홀수입니다.")
	}
}
