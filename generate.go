package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Enums) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_desc.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-enum-desc. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc             ", protocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)

	enumGen := NewEnumDesc()
	for _, e := range file.Enums {
		if !enumGen.hasComment(e.Values) {
			continue
		}
		g.P(enumGen.genDescMap(e.GoIdent.GoName, e.Values))
		g.P(enumGen.genDescMethod(e.GoIdent.GoName))
	}

	g.P()
	return g
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}