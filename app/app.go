// Package app contains the front end application.
package app

import (
	"fmt"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/norunners/hellovecty/dao"
	"strconv"
)

// App is a trivial addition calculator app.
type App struct {
	vecty.Core
	dao dao.Dao

	a, b, sum int
	err       error
}

// New creates an app with a given dao.
func New(dao dao.Dao) *App {
	return &App{dao: dao}
}

// Render returns the component's html.
func (app *App) Render() *vecty.HTML {
	return elem.Body(
		elem.Div(
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypeText),
					event.Change(app.onChangeA),
				),
			),
		),
		elem.Div(
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypeText),
					event.Change(app.onChangeB),
				),
			),
		),
		elem.Div(
			elem.Button(
				vecty.Text("Add"),
				vecty.Markup(
					event.Click(app.onClick).PreventDefault(),
				),
			),
		),
		elem.Div(
			vecty.Text(fmt.Sprintf("%v + %v = %v", app.a, app.b, app.sum)),
		),
		elem.Div(
			vecty.Text(fmt.Sprintf("Error: %v", app.err)),
		),
	)
}

// onChangeA parses the input text and stores the result.
func (app *App) onChangeA(e *vecty.Event) {
	value := e.Target.Get("value").String()
	val, err := strconv.Atoi(value)
	app.err = err
	if err != nil {
		return
	}
	app.a = val
}

// onChangeB parses the input text and stores the result.
func (app *App) onChangeB(e *vecty.Event) {
	value := e.Target.Get("value").String()
	val, err := strconv.Atoi(value)
	app.err = err
	if err != nil {
		return
	}
	app.b = val
}

// onClick adds the two input value and renders the result.
// NOTE: The extra go-routine ensures the call is non-blocking.
func (app *App) onClick(e *vecty.Event) {
	go func() {
		defer vecty.Rerender(app)
		if app.err != nil {
			return
		}
		sum, err := app.dao.Add(app.a, app.b)
		app.err = err
		if err != nil {
			return
		}
		app.sum = sum
	}()
}
