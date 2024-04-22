package factories

import "design-pattern/users"

// CreateUser creates a user based on the given name
func CreateUser(name string) users.User {
    if name == "Admin" {
        return users.NewAdmin()
    } else if name == "Customer" {
        return users.NewCustomer()
    }
    return nil
}
