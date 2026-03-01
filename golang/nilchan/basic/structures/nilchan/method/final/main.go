package main

import (
	"fmt"
	"time"
)

type User struct {
	Name    string  // must be not empty
	Phone   string  // must be not empty
	Age     int     // Age must be > 0 and < 150
	IsClose bool    // Is profile closed?
	Rating  float64 // rating must be > 0 and <= 10
}

func main() {
	user1 := &User{
		Name:    "Jora",
		Age:     15,
		Phone:   "+525000000",
		IsClose: false,
		Rating:  15.0,
	}
	user2 := &User{
		Name:    "Volodya",
		Age:     4,
		Phone:   "+91203343200",
		IsClose: true,
		Rating:  15.0,
	}

	// Before changes
	fmt.Println("Before changes:")
	fmt.Println(user2)
	fmt.Println(user1)

	user1.CloseProfile()
	user2.OpenProfile()

	user1.ChangeAge(590)
	user2.ChangeAge(12)

	user1.ChangeName("Gosha")
	user1.ChangeName("Pavel")

	user1.RatingUp(12.9)
	user2.RatingDown(4.0)

	user1.ChangePhone("+88005553535")
	user2.ChangePhone("+49328423434")

	time.Sleep(1 * time.Millisecond)

	// After changes
	fmt.Println("After changes:")
	fmt.Println(user2)
	fmt.Println(user1)

	// Creating new user
	fmt.Println("Creting new user Jopa")
	time.Sleep(1 * time.Millisecond)
	fmt.Println(NewUser("Jopa", "+8940234299", 12, true, 129))

}

func NewUser(
	name string,
	phone string,
	age int,
	isClose bool, // Not must have
	rating float64, // Not must have
) *User {
	fmt.Println("Name Validation!")
	time.Sleep(500 * time.Millisecond)
	if name == "" {
		fmt.Println("Not allowed to use empty name")
		return &User{}
	}
	fmt.Println("Name validation completed")

	fmt.Println("Phone Validation!")
	time.Sleep(500 * time.Millisecond)
	if phone == "" {
		fmt.Println("Not allowed to use empty phone")
		return &User{}
	}
	fmt.Println("Phone validation completed")

	fmt.Println("Age Validation!")
	time.Sleep(500 * time.Millisecond)
	if age == 0 || age > 150 {
		fmt.Println("Not allowed age (if it's your real age write in the support)")
		return &User{}
	}
	fmt.Println("Age validation completed")

	fmt.Println("Profile Validation!")
	time.Sleep(500 * time.Millisecond)
	if isClose {
		fmt.Println("Your account now is closed")
	} else {
		fmt.Println("Set your profile as opened")
	}
	fmt.Println("Profile validation completed")

	fmt.Println("Rating Validation!")
	time.Sleep(500 * time.Millisecond)
	if rating > 100 {
		fmt.Println("Maximum rating is 100, so your rating is 100")
		rating = 100
	}
	fmt.Println("Rating validation completed")

	return &User{
		Name:    name,
		Phone:   phone,
		Age:     age,
		IsClose: isClose,
		Rating:  rating,
	}
}

func (u *User) CloseProfile() {
	if u.IsClose == true {
		fmt.Println("Your account is already closed.")
	} else {
		u.IsClose = true
	}
}
func (u *User) OpenProfile() {
	if u.IsClose == false {
		fmt.Println("Your account is already opened.")
	} else {
		u.IsClose = false
	}
}
func (u *User) ChangeAge(y int) {
	if y <= 0 {
		fmt.Println("Not allowed age (lower or equals 0)")
	} else if y > 150 {
		fmt.Println("Not allowed age (more than 150), if your age is over 150, text to the support")
	} else {
		u.Age = y
	}
}
func (u *User) ChangePhone(y string) {
	if y == "" {
		fmt.Println("Not supported phone number")
	} else {
		u.Phone = y
	}
}
func (u *User) ChangeName(y string) {
	if y == "" {
		fmt.Println("Name cannot be empty")
	} else {
		u.Name = y
	}
}
func (u *User) RatingUp(y float64) {
	if u.Rating+y > 100.0 {
		fmt.Println("Rating cannot be over 100, the maximal rating is 100")
		u.Rating = 100.0
	} else {
		u.Rating += y
	}
}
func (u *User) RatingDown(y float64) {
	u.Rating -= y
}
