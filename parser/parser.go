package parser

import (
	"bytes"
	"fmt"
	"gitlab.tech.mvideo.ru/mvideoru/debug/models"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"
)

func ExtractFileData(files []*ast.File) models.ExtractedFileData {
	res := models.ExtractedFileData{
		Imports: make([]*ast.ImportSpec, 0),
		Decls:   make([]ast.Decl, 0),
	}

	for _, f := range files {
		res.Decls = append(res.Decls, f.Decls...)
		res.Imports = append(res.Imports, f.Imports...)
	}

	return res
}

func ParseDir(dir string) (*token.FileSet, []*ast.File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil, err
	}

	fset := token.NewFileSet()
	var files []*ast.File

	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".go") ||
			name == "getters.go" ||
			strings.HasSuffix(name, "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(dir, name), nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, nil, fmt.Errorf("parse %s: %w", name, err)
		}
		files = append(files, f)
	}
	return fset, files, nil
}

func CollectImports(imports []*ast.ImportSpec) map[string]string {
	result := make(map[string]string)
	for _, imp := range imports {
		importPath := strings.Trim(imp.Path.Value, `"`)
		var alias string
		if imp.Name != nil && imp.Name.Name != "_" && imp.Name.Name != "." {
			alias = imp.Name.Name
		} else {
			alias = path.Base(importPath)
		}
		result[alias] = imp.Path.Value
	}

	return result
}

func CollectStructs(fset *token.FileSet, decls []ast.Decl) []models.StructInfo {
	var structs []models.StructInfo

	for _, decl := range decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			si := models.StructInfo{
				Name:     typeSpec.Name.Name,
				RecvName: string(unicode.ToLower([]rune(typeSpec.Name.Name)[0])),
			}

			for _, field := range structType.Fields.List {
				for _, ident := range field.Names {
					if ast.IsExported(ident.Name) {
						continue
					}
					si.Fields = append(si.Fields, models.FieldInfo{
						Name:     ident.Name,
						TypeName: exprToString(fset, field.Type),
					})
				}
			}

			if len(si.Fields) > 0 {
				structs = append(structs, si)
			}
		}
	}

	return structs
}

func exprToString(fset *token.FileSet, expr ast.Expr) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, expr)
	return buf.String()
}

func NeededImports(structs []models.StructInfo, allImports map[string]string) map[string]string {
	used := make(map[string]bool)
	for _, s := range structs {
		for _, f := range s.Fields {
			parts := strings.FieldsFunc(f.TypeName, func(r rune) bool {
				return r == '*' || r == '[' || r == ']' || r == ' ' || r == '(' || r == ')'
			})
			for _, p := range parts {
				if dot := strings.Index(p, "."); dot > 0 {
					used[p[:dot]] = true
				}
			}
		}
	}

	result := make(map[string]string)
	for alias, importPath := range allImports {
		if used[alias] {
			result[alias] = importPath
		}
	}
	return result
}
