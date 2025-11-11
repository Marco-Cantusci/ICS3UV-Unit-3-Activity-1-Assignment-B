/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-10
 * @fileoverview Calculate the restaurant bill
 */

package main

import "fmt"

func main() {

	// store subtotal as 35.85
	const subtotal float64 = 35.85

	// store tip amount as 20%
	const tipAmount float64 = 20.0

	// store people as 3
	const people float64 = 3.0

	// set the tax as 15%
	const tax float64 = 15.0

	// calculate tip value
	const tipValue float64 = subtotal * tipAmount / 100

	// calculate total before tax
	const totalBeforeTax float64 = subtotal + tipValue

	// calculate tax value
	const taxValue float64 = totalBeforeTax * tax / 100

	// calculate total value
	const total float64 = totalBeforeTax + taxValue

	// calculate split amongst people
	const split float64 = total / people

	// display the subtotal
	fmt.Println("Subtotal($):", subtotal)

	// display the tax rate
	fmt.Println("Tax(%):", tax)

	// display the tip amount
	fmt.Println("Tip amount(%):", tipAmount)
	
	// display the total
	fmt.Println("Total($):", total)

	// display split
	fmt.Println("split($):", split, "per person")

	fmt.Println("\nDone.")

}
