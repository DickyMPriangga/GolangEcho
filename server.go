package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, data)
}

func getUsers(c echo.Context) error {
	for _, d := range data {
		if d.ID == c.Param("id") {
			//return c.JSON(http.StatusOK, d)
			return c.Render(http.StatusOK, "template.html", d)
		}
	}

	return c.String(http.StatusOK, "Data not found")
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("template.html")),
	}
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/users", getAllUsers)
	e.GET("/users/:id", getUsers)

	e.Logger.Fatal(e.Start(":8081"))
}
