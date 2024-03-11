package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		ID: 1, Name: "Mr.A", Age: 18,
	}

	customers := []Customer{
		{ID: 1, Name: "Mr.X", Age: 20},
		{ID: 2, Name: "Mr.Y", Age: 30},
	}

	byte_for_user, err := json.Marshal(user)
	fmt.Printf("type : %T \n", byte_for_user)
	fmt.Printf("byte : %v  \n", byte_for_user)
	fmt.Printf("string: %s \n", byte_for_user)
	fmt.Println(err)

	byte_for_customer, err := json.Marshal(customers)
	fmt.Printf("type : %T \n", byte_for_customer)
	fmt.Printf("byte : %v  \n", byte_for_customer)
	fmt.Printf("string: %s \n", byte_for_customer)
	fmt.Println(err)
}
