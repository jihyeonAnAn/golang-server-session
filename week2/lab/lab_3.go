package lab

import "fmt"

type BankAccount struct {
	Balance int
}

func (bankAccount *BankAccount) Deposit(inputPrice int) error {
	if inputPrice <= 0 {
		return fmt.Errorf("입금 금액은 양수여야 합니다: %d", inputPrice)
	}
	bankAccount.Balance += inputPrice
	return nil
}

func (bankAccount *BankAccount) Withdraw(inputPrice int) error {
	if inputPrice <= 0 {
		return fmt.Errorf("출금 금액은 양수여야 합니다: %d", inputPrice)
	}
	if bankAccount.Balance < inputPrice {
		return fmt.Errorf("잔고 부족: 요청한 금액: %d, 현재 잔고: %d", inputPrice, bankAccount.Balance)
	}
	bankAccount.Balance -= inputPrice
	return nil
}

func Lab_3() {
	account := &BankAccount{100}
	var inputNumber int

	for true {
		fmt.Println("입금 (1), 출금 (2), 종료 (Others) : ")
		fmt.Scan(&inputNumber)
		inputPrice := 0

		switch inputNumber {
		case 1:
			fmt.Println("입금할 금액을 입력하세요: ")
			fmt.Scan(&inputPrice)
			err := account.Deposit(inputPrice)
			if err != nil {
				fmt.Println("입금 오류:", err)
			} else {
				fmt.Printf("입금 성공! 현재 잔액: %d원\n\n", account.Balance)
			}
		case 2:
			fmt.Println("출금할 금액을 입력하세요: ")
			fmt.Scan(&inputPrice)
			err := account.Withdraw(inputPrice)
			if err != nil {
				fmt.Println("출금 오류:", err)
			} else {
				fmt.Printf("출금 성공! 현재 잔액: %d원\n\n", account.Balance)
			}
		case 3:
			return
		}
	}
}
