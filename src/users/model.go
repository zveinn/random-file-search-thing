package users

import "log"

// User ...
// This is a documentation thingy for the user struct
type User struct {
	Name string
	Email string
}

// DoSomething ..
func (u *User)Something(){
	log.Println("Users can view restricted materials..")
}
