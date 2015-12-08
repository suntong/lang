////////////////////////////////////////////////////////////////////////////
// Porgram: Composition.go
// Purpose: Go composition & field promotion demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: Richard Sugg, ArdanLabs (Bill Kennedy)
//          https://gist.github.com/rsperl/01d7a4da7c01706d50f6
//          https://groups.google.com/d/msg/golang-nuts/uj7qH2qzD1E/Xq1Wya0tBQAJ
////////////////////////////////////////////////////////////////////////////

// I thought I needed to do inheritance, but given that golang doesn't
// support inheritance, I knew I wasn't thinking in go idioms. The solution
// is composition, and it's made possible by method and field promotion

// For e.g., say I have a type User with some methods applicable to all Users:

//package user
package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) SendEmail(subject, body string) bool {
	// ... send an email
	fmt.Printf("User %s sent an email with Subject: '%s'\n", u.Name, subject)
	return true
}

// You may also have an Admin, who is a User, but has additional abilities
// and fields not applicable to ordinary users. How do you add the
// additional properties and methods to an Admin while retaining the
// original implementation of a User? The answer is in composition. When a
// type is embedded in another type, the fields and methods of the subtype
// are automatically promoted to the type that is embedding. Think of it as
// having "super" being called automatically.

// Because Admin embeds User, Admin now has all the properties of User as well
// as the additional ones defined specifically for Admin. In other words, given
// "adm" is of type *Admin, you can do
//
// Call the User's Name property and methods explicitly
//   adm.User.Name
//   adm.User.SendEmail("hi", "there")
// Or take advantage of type promotion
//   adm.Name
//   adm.SendEmail("hi", "there")
// Likewise, you also have Admin's own properties that are not
// avaiable at all to type User
//   adm.UnlockGate()

type Admin struct {
	User
	SecretPassword string
}

func (a *Admin) UnlockGate() bool {
	// ... do admin stuff
	fmt.Printf("Admin %s has unlocked the door to mystery with PW: '%s'\n",
		a.Name, a.SecretPassword)
	return true
}

func main() {
	// User
	usr := User{"Sam", "s@abc.ca"}
	usr.SendEmail("hi", "there")

	// Admin
	adm := Admin{User{"Jack", "j@abc.ca"}, "confidential"}

	// Call the User's Name property and methods explicitly
	fmt.Printf("Admin's name: %s\n", adm.User.Name)
	adm.User.SendEmail("Tick", "")

	// Or take advantage of type promotion
	fmt.Printf("Admin's name: %s\n", adm.Name)
	adm.SendEmail("Tock", "")

	// Likewise, you also have Admin's own properties that are not
	// avaiable at all to type User
	adm.UnlockGate()

}
