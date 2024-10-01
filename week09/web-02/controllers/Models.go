package controllers

import "fmt"

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	Friends   []Person
}

func (pe *Person) Greet() {
	fmt.Println("Hi I'm, ", pe.FirstName, " ", pe.LastName, " and I'm ", pe.Age, " years old.")
}

func (pe *Person) ListFriends() {
	fmt.Println("Listing friends")
	for _, friend := range pe.Friends {
		fmt.Println(friend.FirstName, ", ")
	}
	fmt.Println()
}

func (pe *Person) AddFriend(f Person) {
	pe.Friends = append(pe.Friends, f)
}

func (pe *Person) RemoveFriend(Id int) {
	for i, f := range pe.Friends {
		if f.Id == Id {
			pe.Friends = append(pe.Friends[:i], pe.Friends[i+1:]...)
			return
		}
	}
}

func (pe *Person) RemoveAllFriends() {
	pe.Friends = pe.Friends[:0]
}
