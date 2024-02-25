package main

import (
	"html/template"
	"os"
	"strings"
)

type Data struct {
	Name string
}

const uuid_tmpl = `package ynab

import "github.com/google/uuid"

func (id {{.Name}}) String() string {
	return (uuid.UUID)(id).String()
}

func (id *{{.Name}}) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id {{.Name}}) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id {{.Name}}) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id {{.Name}}) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func Parse{{.Name}}(s string) ({{.Name}}, error) {
	id, err := uuid.Parse(s)
	return ({{.Name}})(id), err
}

func MustParse{{.Name}}(s string) {{.Name}} {
	return ({{.Name}})(uuid.MustParse(s))
}
`

func main() {
	uuid_names := strings.Split(os.Args[1], ",")

	tmpl, err := template.New("test").Parse(uuid_tmpl)
	if err != nil {
		panic(err)
	}

	for _, name := range uuid_names {
		data := Data{name}

		fp, err := os.Create("zz_uuid_" + strings.ToLower(name) + ".go")
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		err = tmpl.Execute(fp, data)
		if err != nil {
			panic(err)
		}
	}
}
