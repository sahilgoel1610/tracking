package users
import "fmt"
type User struct {
	Id int
	Friends []User
}

var currentId = 0

func (u *User) AddFriend(subUser *User) {
	u.Friends = append(u.Friends, *subUser)
	fmt.Println(u.Friends)
	// return u
}

func New() (u User) {
	currentId = currentId + 1
	return User {currentId, []User{}}
}
