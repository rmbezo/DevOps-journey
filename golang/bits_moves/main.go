package main

import "fmt"

func main() {
	const (
		IsAdmin   = 1 << 0 // 0001 (в десятичной 1)
		IsEditor  = 1 << 1 // 0010 (в десятичной это 2)
		CanSub    = 1 << 2 // 0100 (В десятичной это 4)
		IsPremium = 1 << 3 // 1000 (в десятичной это 8)
	)

	// create a user with permissions
	var userPermissions = IsEditor | IsPremium // 1010

	fmt.Printf("Права юзера в двоичной системе: %b\n", userPermissions)

	// Check permission

	fmt.Println("--- Check permissions ---")

	if userPermissions&IsAdmin != 0 {
		fmt.Println("Hello, Boss!")
	} else {
		fmt.Println("You dont have admin permissions")
	}

	if userPermissions&IsEditor != 0 {
		fmt.Println("You can edit blogs")
	} else {
		fmt.Println("You dont have access to edit blogs")
	}

	if userPermissions&IsPremium != 0 {
		fmt.Println("Thanks for subscription!")
	} else {
		fmt.Println("You dont have Premium status")
	}
}
