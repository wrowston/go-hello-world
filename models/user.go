package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	// &u takes the address of the pointer because
	// users collection stores pointers to user objects
	users = append(users, &u)
	return u, nil
}
