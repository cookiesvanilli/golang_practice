package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age uint16
	Money                int16
	AvgGrades, Happiness float64
	Hobbies []string
}

func (u User) getAllInfo() string { //Когда структура небольшая, проще передавать объект не как ссылку (т.е. * здесь необязательна)
	return fmt.Sprintf("User name is: %s. He is %d and " +
		"he has money equal: %d",
		u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) { // Не будет происходить корректное изменение данных без ссылки
	u.Name = newName
}

func homePage(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	bob := User{
		Name:      "Bob",
		Age:       25,
		Money:     -50,
		AvgGrades: 4.2,
		Happiness: 0.8,
		Hobbies: []string{"Football", "Music"},
	}
	/* bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())*/
/*	fmt.Println("Work!")*/
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	handleRequest()
}


