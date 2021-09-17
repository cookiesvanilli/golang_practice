//tax calculator
package main

import "fmt"

func main() {
var price float64
var tax float64

fmt.Println("What is the price? ($)")
_, err := fmt.Scanf("%f", &price)
if err != nil {
	fmt.Println("Error: ", err)
	return
}

fmt.Println("What is the tax? (%)")
_,err = fmt.Scanf("%f", &tax)
if err != nil {
	fmt.Println("Error: ", err)
	return
}

fmt.Printf("A purchase of $%.2f with a tax rate of %.2f%%\n", price, tax)
	totalPrice, tax := CalculateTax(price, tax/100)
	fmt.Printf("Total Tax: $%.2f\n", tax)
	fmt.Printf("Total Purchase Price: $%.2f\n", totalPrice)
}

func CalculateTax(prices float64, taxRate float64) (totalPrice float64, taxes float64) {
	taxes = prices * taxRate
	totalPrice = prices + taxes

	return totalPrice, taxes
}

