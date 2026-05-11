package main

import (
	"fmt"
)

const exchangeRate = 320.50

func convertUSDtoLKR(usd float64) float64 {
	return usd * exchangeRate
}

func convertLKRtoUSD(lkr float64) float64 {
	return lkr / exchangeRate
}

func showMenu() {
	fmt.Println("-----------------------------")
	fmt.Println("1. USD to LKR")
	fmt.Println("2. LKR to USD")
	fmt.Println("0. Exit")
	fmt.Println("-----------------------------")
	fmt.Print("Choose an option: ")
}

func main() {
	fmt.Println("USD to LKR Converter")
	fmt.Println("Exchange Rate: 1 USD =", exchangeRate, "LKR")

	for {
		showMenu()

		var option int
		_, err := fmt.Scan(&option)

		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			var discard string
			fmt.Scan(&discard)
			continue
		}

		if option == 0 {
			fmt.Println("Goodbye!")
			break
		}

		var amount float64
		fmt.Print("Enter amount: ")
		_, err = fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var discard string
			fmt.Scan(&discard)
			continue
		}

		switch option {
		case 1:
			result := convertUSDtoLKR(amount)
			fmt.Printf("%.2f USD = %.2f LKR\n", amount, result)
		case 2:
			result := convertLKRtoUSD(amount)
			fmt.Printf("%.2f LKR = %.2f USD\n", amount, result)
		default:
			fmt.Println("Invalid option! Please choose 1, 2 or 0.")
		}
	}
}