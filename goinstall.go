package goinstall

import (
	"embed"
)

//go:embed scripts/*
var scriptsFS embed.FS

const DefaultVersion = "1.25.2"

type Goinstall struct {
	version      string
	logger       func(string)
	afterInstall func(goPath string) error
}

type Option func(*Goinstall)

func New(opts ...Option) *Goinstall {
	g := &Goinstall{
		version: DefaultVersion,
		logger:  func(msg string) {},
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func WithVersion(v string) Option {
	return func(g *Goinstall) {
		g.version = v
	}
}

func WithLogger(f func(string)) Option {
	return func(g *Goinstall) {
		g.logger = f
	}
}

func WithAfterInstall(f func(goPath string) error) Option {
	return func(g *Goinstall) {
		g.afterInstall = f
	}
}

func (g *Goinstall) log(msg string) {
	if g.logger != nil {
		g.logger(msg)
	}
}
