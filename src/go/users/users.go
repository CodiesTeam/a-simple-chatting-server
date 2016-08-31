package users

type User struct {
    BasicInfo
}

type BasicInfo struct {
    ID string
    Name, Email string
    PhoneNumber int
}

func getUserById(id string) User {
    // TODO
}

func InsertUser(user User) error {
    // TODO
}