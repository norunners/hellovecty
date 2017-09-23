// Copyright (c) 2016 The Vecty Authors. All rights reserved.
// See source: https://github.com/gopherjs/vecty/tree/master/example/markdown

package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/norunners/pubsub"
)

// PageView is our main page component.
type PageView struct {
	vecty.Core
	ps pubsub.PubSub

	input string
}

// NewPv creates a new page view component.
func NewPv(ps pubsub.PubSub, input string) vecty.Component {
	return &PageView{ps: ps, input: input}
}

// onType publishes the input.
func (p *PageView) onType(e *vecty.Event) {
	p.input = e.Target.Get("value").String()
	p.ps.Pub(p.input, "onType")
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() *vecty.HTML {
	// Display a textarea on the right-hand side of the page.
	return elem.Div(
		vecty.Markup(
			vecty.Style("float", "right"),
		),
		elem.TextArea(
			vecty.Markup(
				vecty.Style("font-family", "monospace"),
				vecty.Property("rows", 14),
				vecty.Property("cols", 70),

				// When input is typed into the textarea, update the local
				// component state and rerender.
				event.Input(p.onType),
			),
			vecty.Text(p.input), // initial textarea text.
		),
	)
}
