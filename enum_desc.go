package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
	"text/template"
)

//go:embed desc.tpl
var enumDesc string

type EnumDesc struct {
	EnumName string
}

func NewEnumDesc() *EnumDesc {
	return &EnumDesc{}
}

func (desc *EnumDesc) genDescMethod(name string) string {
	desc.EnumName = name
	var err error
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(enumDesc))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, desc); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

// genDescMap
// var (
//
//	TaskStatus_desc = map[TaskStatus]string{
//		CREATED: "CREATED",
//		RUNNING: "RUNNING",
//		STOPPED: "STOPPED",
//		FAILED: "FAILED",
//		SUCCESS: "SUCCESS",
//	}
//
// )
func (desc *EnumDesc) genDescMap(enumName string, values []*protogen.EnumValue) string {
	r := strings.Builder{}
	r.WriteString("var (")
	r.WriteString("\n")
	r.WriteString(fmt.Sprintf("	%s_desc = map[%s]string{", enumName, enumName))
	r.WriteString("\n")
	for _, value := range values {
		r.WriteString("	")
		wrapDesc := "\"" + getComment(value.Comments.Trailing) + "\""
		r.WriteString(value.GoIdent.GoName + ": " + wrapDesc)

		r.WriteString(",")
		r.WriteString("\n")

	}
	r.WriteString("	}\n")
	r.WriteString(")\n")
	return r.String()
}

func (desc *EnumDesc) hasComment(values []*protogen.EnumValue) bool {
	for _, value := range values {
		if getComment(value.Comments.Trailing) == "" {
			return false
		}
	}
	return true
}

func getComment(c protogen.Comments) string {
	if c == "" {
		return ""
	}
	var b []byte
	for _, line := range strings.Split(strings.TrimSuffix(string(c), "\n"), "\n") {
		b = append(b, line...)
	}
	return string(b)
}
