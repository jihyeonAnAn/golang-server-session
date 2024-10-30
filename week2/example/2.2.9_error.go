package example

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

// 사용자 정의 에러 타입
type errorString struct {
	s string // 에러 메시지를 저장할 필드
}

// 에러 메서드 구현
func (e *errorString) Error() string {
	return e.s
}

//Division By Zero 오류 처리 예시

// errors.New를 사용한 오류 생성
var DivisionByZero = errors.New("division by zero")

func Divide(number, d float32) (float32, error) {
	if d == 0 {
		return 0, DivisionByZero
	}
	return number / d, nil
}

func Code228() {
	n1, e1 := Divide(1, 1)
	fmt.Println(n1)
	if e1 != nil {
		fmt.Println(e1.Error())
	}
	n2, e2 := Divide(1, 0)
	fmt.Println(n2)
	if e2 != nil {
		fmt.Println(e2.Error())
	}
}

var SampleError = errors.New("This is a test error")

func testPanic() {
	panic(SampleError)
	fmt.Println("Hello from testPanic!")
}

func testRecover() {
	defer func() {
		if recover() != nil {
			fmt.Println("got an error!")
		} else {
			fmt.Println("no error")
		}
	}()
	testPanic()
	fmt.Println("Hello from testRecover!")
}

func Code229() {
	testRecover()
}
