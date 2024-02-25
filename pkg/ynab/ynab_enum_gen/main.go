package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type EnumValue struct {
	Name     string
	JSONName string
	Idx      int
}

type Data struct {
	Name string
	Enum []EnumValue
}

const enum_tmpl = `package ynab

import (
	"encoding/json"
	"fmt"
	"strings"
)

// #region {{.Name}}Type

type {{.Name}}Type uint8

const (
	No{{.Name}} {{.Name}}Type = iota{{block "enumlist" .Enum}}{{"\n"}}{{range .}}	{{println .Name}}{{end}}{{end}})

var (
	{{.Name}}_jsonname = map[uint8]string{
{{block "jsonnamelist" .Enum}}{{range .}}		{{.Idx}}: "{{.JSONName}}",{{"\n"}}{{end}}{{end}}	}
)

var (
	{{.Name}}_name = map[uint8]string{
		0: "No{{.Name}}",
{{block "namelist" .Enum}}{{range .}}		{{.Idx}}: "{{.Name}}",{{"\n"}}{{end}}{{end}}	}
)

var (
	{{.Name}}_value = map[string]uint8{
{{block "valuelist" .Enum}}{{range .}}		"{{.JSONName}}": {{.Idx}},{{"\n"}}{{end}}{{end}}	}
)

func (s {{.Name}}Type) String() string {
	return {{.Name}}_name[uint8(s)]
}

func (s {{.Name}}Type) MarshalJSON() ([]byte, error) {
	if s == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal({{.Name}}_jsonname[uint8(s)])
}

func (s *{{.Name}}Type) UnmarshalJSON(data []byte) (err error) {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		*s = {{.Name}}Type(0)
	} else if *s, err = Parse{{.Name}}Type(*value); err != nil {
		return err
	}
	return nil
}

func Parse{{.Name}}Type(s string) ({{.Name}}Type, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return {{.Name}}Type(0), nil
	}
	value, ok := {{.Name}}_value[s]
	if !ok {
		return {{.Name}}Type(0), fmt.Errorf("%q is not a valid {{.Name}}Type", s)
	}
	return {{.Name}}Type(value), nil
}

// #endregion
`

func main() {
	enum_name := os.Args[1]

	data := Data{enum_name, []EnumValue{}}
	for i, enumvar := range os.Args[2:] {
		enum := strings.Split(enumvar, ":")
		data.Enum = append(data.Enum, EnumValue{enum[0], enum[1], i + 1})
	}
	fmt.Println(data)

	tmpl, err := template.New("test").Parse(enum_tmpl)
	if err != nil {
		panic(err)
	}

	fp, err := os.Create("zz_enum_" + strings.ToLower(enum_name) + ".go")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	err = tmpl.Execute(fp, data)
	if err != nil {
		panic(err)
	}

}
