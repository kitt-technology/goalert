package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

func main() {
	out := flag.String("out", "", "Output file.")
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	var src []*ast.Source
	for _, file := range flag.Args() {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("ERROR:", err)
		}
		src = append(src, &ast.Source{
			Name:  file,
			Input: string(data),
		})
	}

	doc, err := parser.ParseSchemas(src...)
	if err != nil {
		log.Fatal("ERROR:", err)
	}

	w := os.Stdout
	if *out != "" {
		fd, err := os.Create(*out)
		if err != nil {
			log.Fatal("ERROR:", err)
		}
		defer fd.Close()
		w = fd
	}

	typeName := func(n *ast.Type) string {
		result := ""

		switch n.Name() {
		case "String", "ID":
			result = "string"
		case "Int":
			result = "number"
		case "Boolean":
			result = "boolean"
		default:
			result = n.Name()
		}

		isArrayType := (n.String())[:1] == "["
		if isArrayType {
			result += "[]"
		}

		return result
	}

	fmt.Fprintf(w, "// Code generated by devtools/gqltsgen DO NOT EDIT.\n\n")

	for _, def := range doc.Definitions {
		switch def.Kind {
		case ast.Enum:
			fmt.Fprintf(w, "export type %s = ", def.Name)
			for _, e := range def.EnumValues {
				fmt.Fprintf(w, " | '%s'", e.Name)
			}
			fmt.Fprintf(w, "\n\n")
		case ast.InputObject, ast.Object:
			fmt.Fprintf(w, "export interface %s {\n", def.Name)
			for _, e := range def.Fields {
				mod := "?"
				if e.Type.NonNull {
					mod = ""
				}
				fmt.Fprintf(w, "  %s: %s\n", e.Name+mod, typeName(e.Type))
			}
			fmt.Fprintf(w, "}\n\n")
		case ast.Scalar:
			fmt.Fprintf(w, "export type %s = string\n\n", def.Name)
		default:
			log.Fatal("Unsupported kind:", def.Name, def.Kind)
		}
	}
}
