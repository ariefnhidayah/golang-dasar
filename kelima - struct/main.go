package main

import "fmt"

// struct

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// method
func (user User) display() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	user1 := User{
		Id:        1,
		FirstName: "Arief",
		LastName:  "Hidayah",
		Email:     "ariefnhidayah@gmail.com",
		IsActive:  true,
	}

	result := user1.display()
	fmt.Println(result)

	// user2 := User{
	// 	Id:        2,
	// 	FirstName: "Uki",
	// 	LastName:  "Popa",
	// 	Email:     "ukipopa@mail.com",
	// 	IsActive:  true,
	// }

	// users := []User{
	// 	user1,
	// 	user2,
	// }

	// group := Group{
	// 	Name:  "Gamer",
	// 	Users: users,
	// 	Admin: User{
	// 		Id:        3,
	// 		FirstName: "admin",
	// 		Email:     "admin@admin.com",
	// 		IsActive:  true,
	// 	},
	// 	IsAvailable: true,
	// }

	// displayGroup(group)

	// method

}

// func displayGroup(group Group) {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member Count: %d", len(group.Users))
// 	fmt.Println("")
// 	fmt.Println("Users Name: ")

// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}
// }

// struct menjadi parameter
// func displayUser(user User) string {
// 	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
// }
