package users

type Customer struct {
    User
}

// encapsulate User creation
func NewCustomer() User {
    return &Customer{
        User: User{
            name: "Customer",
        },
    }
}


