package users


type User interface {
    doStuff() string
}

type User struct {
    name string
}

func (u *User) doStuff() string {
    return "hello, i am " + u.name
}
