package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jszwec/csvutil"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Student struct {
	Name      string `csv:"name"`
	Age       int    `csv:"age,omitempty"`
	CreatedAt time.Time
}

type StudentCompo struct {
	app.Compo
	Students []Student
}

func (sc *StudentCompo) OnMount(ctx app.Context) {
	app.Log("StudentCompo OnMount called")
	sc.Students = createStudents()
}

func (sc *StudentCompo) Render() app.UI {
	app.Log("StudentCompo Render called")
	l := len(sc.Students)
	app.Log("len(sc.Students)", l)
	for i := 0; i < len(sc.Students); i++ {
		app.Log("Name", sc.Students[i].Name)
		app.Log("Age", sc.Students[i].Age)
		app.Log("CreatedAt", sc.Students[i].CreatedAt)
	}
	return app.Div().Body(
		app.Table().Body(
			app.Tr().Body(
				app.Th().Text("Name"),
				app.Th().Text("Age"),
				app.Th().Text("CreatedAt"),
			),
			/*
				app.Tr().Body(
					app.Td().Text("jacek"),
					app.Td().Text("23"),
					app.Td().Text("2012-04-01T15:00:00Z"),
				),
				app.Tr().Body(
					app.Td().Text("john"),
					app.Td().Text("21"),
					app.Td().Text("2001-05-21T16:57:00Z"),
				),
			*/
			app.Range(sc.Students).Slice(func(i int) app.UI {
				return app.Tr().Body(app.Td().Text(sc.Students[i].Name),
					app.Td().Text(fmt.Sprintf("%d", sc.Students[i].Age)),
					app.Td().Text(sc.Students[i].CreatedAt.Format(time.RFC3339)),
				)
			}),
		),
	)
}

func createStudents() []Student {
	app.Log("createStudents called")
	var csvInput = []byte(`
name,age,CreatedAt
jacek,23,2012-04-01T15:00:00Z
john,21,2001-05-21T16:57:00Z`,
	)

	var students []Student
	if err := csvutil.Unmarshal(csvInput, &students); err != nil {
		fmt.Println("error:", err)
	}

	return students
}

func main() {
	app.Log("Hello, World!")
	app.Route("/", &StudentCompo{})

	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name: "Student",
	})

	log.Println("Listening on http://:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
