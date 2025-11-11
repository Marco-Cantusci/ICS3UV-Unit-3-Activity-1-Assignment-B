// store subtotal as 35.85
const subtotal: number = 35.85

// store tip amount as 20%
const tipAmount: number = 20

// store people as 3
const people: number = 3

// set the tax as 15%
const tax: number = 15

// calculate tip value
const tipValue: number = subtotal * tipAmount / 100

// calculate total before tax
const totalBeforeTax: number = subtotal + tipValue

// calculate tax value
const taxValue: number = totalBeforeTax * tax / 100

// calculate total value
const total: number = totalBeforeTax + taxValue

// calculate split amongst people
const split: number = total / people

// display the subtotal
console.log("Subtotal($):", subtotal)

// display the tax rate
console.log("Tax(%):", tax)

// display the tip amount
console.log("Tip amount(%):", tipAmount)

// display the total
console.log("Total($):", total)

// display split
console.log("split($):", split, "per person")

console.log("\nDone.")
