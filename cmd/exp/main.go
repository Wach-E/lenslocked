package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// user := struct {
	// 	Name string
	// }{
	// 	Name: "Wachukwu Emmanuel",
	// }

	user := User{
		Name: "Wachukwu Emmanuel",
		Age:  23,
		Meta: UserMeta{
			Visits: 4,
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
