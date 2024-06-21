package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BankAccount struct {
	Balance float64
	Owner   string
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
}
func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.Balance < amount {
		return fmt.Errorf("недостаточно средств на счету")
	}
	ba.Balance -= amount
	return nil
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.Balance
}

func main() {
	account := BankAccount{Balance: 1000, Owner: "Иван Иванов"}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать в банковское приложение!")

	for {
		fmt.Println("\nВыберите операцию:")
		fmt.Println("1. Пополнить счет")
		fmt.Println("2. Снять средства")
		fmt.Println("3. Проверить баланс")
		fmt.Println("4. Выход")

		fmt.Print("Введите номер операции: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Введите сумму для пополнения: ")
			scanner.Scan()
			amountStr := scanner.Text()
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Некорректный ввод суммы.")
				continue
			}
			account.Deposit(amount)
			fmt.Println("Баланс после пополнения:", account.GetBalance())

		case "2":
			fmt.Print("Введите сумму для снятия: ")
			scanner.Scan()
			amountStr := scanner.Text()
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Некорректный ввод суммы.")
				continue
			}
			err = account.Withdraw(amount)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Баланс после снятия:", account.GetBalance())
			}

		case "3":
			fmt.Println("Текущий баланс:", account.GetBalance())

		case "4":
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Неверный выбор операции.")
		}
	}
}
