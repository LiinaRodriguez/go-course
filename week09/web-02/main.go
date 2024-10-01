package main

import (
	"fmt"
	"web-02/web-02/controllers"
)

func main() {
	per1 := controllers.Person{Id: 1, FirstName: "Gabriela", LastName: "Torres", Age: 26}
	per2 := controllers.Person{
		Id:        2,
		FirstName: "Juan",
		LastName:  "Rojas",
		Age:       24,
	}
	per1.Greet()
	per1.AddFriend(per2)
	fmt.Println(per1)
	per1.RemoveFriend(2)
	fmt.Println(per1)
	per1.AddFriend(per2)
	per1.ListFriends()
}
