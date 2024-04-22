package users

// concrete User bernama Admin, implement User interface (Admin)
type Admin struct {
    User
}

// encapsulate User creation
func NewAdmin() User {
    return &Admin{
        User: User{
            name: "Admin",
        },
    }
}


