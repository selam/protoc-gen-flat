package main

import (
	"flag"
	"fmt"
	"time"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var suffix string

func main() {
	var flags flag.FlagSet
	flags.StringVar(&suffix, "suffix", "_flat", "The suffix to use for generated files")
	protogen.Options{ParamFunc: flags.Set}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			// Generate code for each file
			generateFlatCode(file, plugin)
		}
		return nil
	})
}

func generateFlatCode(file *protogen.File, plugin *protogen.Plugin) {
	// Create a new file for generated structs and enums
	filename := file.GeneratedFilenamePrefix + suffix + ".go"
	g := plugin.NewGeneratedFile(filename, file.GoImportPath)

	// Write package declaration
	g.P("package ", file.GoPackageName)
	g.P()
	g.P("// Code generated by protoc-gen-flat. DO NOT EDIT.")
	g.P("// source: ", file.Desc.Path(), " (", file.Desc.Package(), ") ", time.Now().Format(time.RFC3339))
	g.P()

	// Generate enums
	for _, enum := range file.Enums {
		generateEnum(g, enum)
	}

	// Generate structs
	for _, message := range file.Messages {
		generateStruct(g, message)
	}
}

func generateEnum(g *protogen.GeneratedFile, enum *protogen.Enum) {
	enumName := enum.GoIdent.GoName
	g.P(fmt.Sprintf("// %s represents the enum values for %s.", enumName, enumName))
	g.P(fmt.Sprintf("type %s int32", enumName))
	g.P("const (")

	for _, value := range enum.Values {
		g.P(fmt.Sprintf("    %s = %d", value.GoIdent.GoName, value.Desc.Number()))
	}

	g.P(")")
	g.P()
}

func generateStruct(g *protogen.GeneratedFile, message *protogen.Message) {
	structName := message.GoIdent.GoName
	g.P("type ", structName, " struct {")

	// Generate fields for each message field
	for _, field := range message.Fields {
		fieldName := field.GoName
		fieldType := mapProtoTypeToGo(field.Desc.Kind(), field.Message, field.Desc.IsList())
		jsonTag := fmt.Sprintf("`json:\"%s\"`", field.Desc.JSONName())
		g.P(fmt.Sprintf("    %s %s %s", fieldName, fieldType, jsonTag))
	}

	g.P("}")
	g.P()
}

func mapProtoTypeToGo(protoType protoreflect.Kind, message *protogen.Message, isRepeated bool) string {
	goType := ""
	switch protoType {
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.Int32Kind:
		goType = "int32"
	case protoreflect.Int64Kind:
		goType = "int64"
	case protoreflect.Uint32Kind:
		goType = "uint32"
	case protoreflect.Uint64Kind:
		goType = "uint64"
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.BytesKind:
		goType = "[]byte"
	case protoreflect.EnumKind:
		goType = "int32"
	case protoreflect.MessageKind:
		goType = message.GoIdent.GoName
	case protoreflect.GroupKind:
		goType = "struct{}"
	case protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Sint32Kind:
		goType = "int32"
	case protoreflect.Sint64Kind:
		goType = "int64"
	case protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Fixed64Kind:
		goType = "uint64"
	default:
		goType = "interface{}" // Generic type for unhandled cases
	}
	if isRepeated {
		goType = "[]" + goType
	}
	return goType
}
