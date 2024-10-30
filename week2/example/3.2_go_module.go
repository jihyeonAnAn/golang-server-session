package example

import (
	"fmt"
	"github.com/otiai10/primes"
	"os"
	"strconv"
)

func Code320() {
	args := os.Args[1:] // 명령줄 인자에서 첫 번째 인자를 가져옴
	if len(args) != 1 {
		fmt.Println("Usage:", os.Args[0], "<number>")
		os.Exit(1)
	}

	number, err := strconv.Atoi(args[0]) //문자열로 전달된 명령줄 인자를 정수로 변환
	if err != nil {
		panic(err) // 변환에 실패하면 프로그램을 종료하고 오류 메시지 출력
	}

	f := primes.Factorize(int64(number)) //number를 소인수분해하고 그 결과는 Factorization 구조체로 반환된다.
	fmt.Println("primes:", len(f.Powers()) == 1)
}
