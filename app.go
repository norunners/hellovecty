// Copyright (c) 2016 The Vecty Authors. All rights reserved.
// See source: https://github.com/gopherjs/vecty/tree/master/example/markdown

package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/norunners/pubsub"
)

const defaultInput = `# Markdown Example

This is a live editor, try editing the Markdown on the right of the page.
`

type App struct {
	vecty.Core
	ps pubsub.PubSub
}

// NewApp creates a new app component.
func NewApp(ps pubsub.PubSub) vecty.Component {
	return &App{ps: ps}
}

func (app *App) Render() *vecty.HTML {
	return elem.Body(
		NewPv(app.ps, defaultInput),
		NewMd(app.ps, defaultInput),
	)
}

func main() {
	ps := pubsub.New()
	app := NewApp(ps)
	vecty.SetTitle("Markdown Demo")
	vecty.RenderBody(app)
}
