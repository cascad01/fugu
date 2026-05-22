package models

import "go/ast"

type FieldInfo struct {
	Name     string
	TypeName string
}

type StructInfo struct {
	Name     string
	RecvName string
	Fields   []FieldInfo
}

type ExtractedFileData struct {
	Imports []*ast.ImportSpec
	Decls   []ast.Decl
}
