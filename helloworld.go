package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

type Address struct {
	City, Province, Country string
}

func (address *Address) getAddress() {
	address.City = "City " + address.City
}

func main() {
	result, err := Pembagi(10, 4)

	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", err)
	}

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(months)

	newSlice := make([]string, 3, 5)
	newSlice[0] = "Anggi"
	newSlice[1] = "Fergian"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	person := map[string]string{
		"name":    "Anggi",
		"address": "Pasar Minggu",
	}
	fmt.Println(person)
	fmt.Println(person["name"])

	if person["name"] == "Anggi" {
		fmt.Println("Nama saya anggi")
	}
	if len(copySlice) == 3 {
		fmt.Println("Length", len(copySlice))
	}

	counter := 1

	for counter <= 10 {
		fmt.Println("Looping", counter)
		counter++
	}

	firstName, _ := getFullName()
	fmt.Println(firstName)

	sum := sumAll(1, 2, 3)
	fmt.Println(sum)

	myNumbers := []int{1, 2, 4, 123, 123}
	total := sumAll(myNumbers...)
	fmt.Println(total)

	saySayonara := getGoodBye

	fmt.Println(saySayonara("anggi"))

	showFilter("anggi", getGoodBye)

	fmt.Println(factorialRecursive(5))

	runApp(true)

	var Anggi Customer
	Anggi.Name = "Anggi"
	Anggi.Age = 20
	Anggi.Address = "Kebagusan"
	fmt.Println(Anggi)

	person2nd := Customer{
		Name:    "Ben",
		Address: "Pasar Minggu",
		Age:     30,
	}
	fmt.Println(person2nd)

	person2nd.sayHello()

	// 36. Interface
	myPayment := RequestPayment{
		BankAccount: "BNI",
		BankName:    "Anggi",
		Amount:      20_000,
	}
	fmt.Println(myPayment)
	printBankAccount(myPayment)

	// 41. Pointer
	address1 := Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	address2 := &address1
	address3 := &address1
	address2.City = "Cirebon"
	*address2 = Address{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	address4 := new(Address)
	fmt.Println(address4)
	address1.getAddress()
	fmt.Println(address1)
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, my name is ", customer.Name)
}

func getFullName() (string, string) {
	return "Anggi", "Fergian"
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func getGoodBye(name string) string {
	return "Sayonara " + name
}

func showFilter(name string, sayHello Filter) {
	name2 := sayHello(name)
	fmt.Println("func as parameter", name2)
}

type Filter func(string) string

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error", message)
	}

	fmt.Println("End app")
}

func runApp(err bool) {
	defer endApp()

	if err {
		panic("Application error")
	}

	fmt.Println("Application running successfully")
}

// struct
type Customer struct {
	Name, Address string
	Age           int
}

type RequestPayment struct {
	BankAccount, BankName string
	Amount                int
}

func (payment RequestPayment) getBankAccount() string {
	return payment.BankName
}

type Payment interface {
	getBankAccount() string
}

func printBankAccount(payment Payment) {
	fmt.Println(payment.getBankAccount())
}
