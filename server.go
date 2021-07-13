package main

import (
	"GolangEcho/handler"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("api/task", handler.GetAllTask)
	e.POST("api/task", handler.CreateTask)
	e.PUT("api/task/:id", handler.TaskComplete)
	e.PUT("api/undoTask/:id", handler.UndoTask)
	e.DELETE("api/deleteTask/:id", handler.DeleteTask)
	e.DELETE("api/deleteAllTask", handler.DeleteAllTask)

	e.Logger.Fatal(e.Start(":8081"))
}
