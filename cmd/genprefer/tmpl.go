package main

import (
	_ "embed"
	"strings"
	"text/template"
)

func args(s []string) string {
	return strings.Join(s, ", ")
}

var (
	//go:embed gentmpl.tmpl
	rawTemplate string

	funcs = template.FuncMap{
		"join": strings.Join,
		"args": args,
	}
	tmpl = template.Must(template.New("").Funcs(funcs).Parse(rawTemplate))
)
