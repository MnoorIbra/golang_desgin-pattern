package main

import (
    "fmt"
    "design-pattern/users"
	"design-pattern/factories"
)

func main() {
    admin := CreateUser("Admin")
    fmt.Println(admin.doStuff())
    customer := CreateUser("Customer")
    fmt.Println(customer.doStuff())
}
