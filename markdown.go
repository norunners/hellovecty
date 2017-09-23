// Copyright (c) 2016 The Vecty Authors. All rights reserved.
// See source: https://github.com/gopherjs/vecty/tree/master/example/markdown

package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/microcosm-cc/bluemonday"
	"github.com/norunners/pubsub"
	"github.com/russross/blackfriday"
)

// Markdown is a simple component which renders the Input markdown as sanitized
// HTML into a div.
type markdown struct {
	vecty.Core
	ps pubsub.PubSub
	us pubsub.UnSub

	input string
}

// NewMd creates a new markdown component.
func NewMd(ps pubsub.PubSub, input string) vecty.Component {
	m := &markdown{ps: ps, input: input}
	m.us = ps.Sub(m, "onType")
	return m
}

// Render implements the vecty.Component interface.
func (m *markdown) Render() *vecty.HTML {
	// Render the markdown input into HTML using Blackfriday.
	unsafeHTML := blackfriday.MarkdownCommon([]byte(m.input))

	// Sanitize the HTML.
	safeHTML := string(bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML))

	// Return the HTML, which we guarantee to be safe / sanitized.
	return elem.Div(
		vecty.Markup(
			vecty.UnsafeHTML(safeHTML),
		),
	)
}

// Receive stores the input and re-renders.
func (m *markdown) Receive(msg interface{}) {
	if input, ok := msg.(string); ok {
		m.input = input
		vecty.Rerender(m)
	}
}

// Unmount unsubscribes the marketdown.
func (m *markdown) Unmount() {
	m.us()
}
