package main

import "fmt"

func main() {

	phonebook := make(map[string]int)
	var count int
	fmt.Scanf("%d", &count)

	for i := 0; i < count; i++ {
		var name string
		var phoneNumber int
		fmt.Scanf("%s", &name)
		fmt.Scanf("%d", &phoneNumber)
		phonebook[name] = phoneNumber
	}

	for {
		var contactName string
		_, err := fmt.Scanf("%s", &contactName)
		if err != nil {
			break
		}
		phoneNumber, exist := phonebook[contactName]
		if exist {
			fmt.Printf("%s=%d\n", contactName, phoneNumber)
		} else {
			fmt.Println("Not Found")
		}
	}

}
