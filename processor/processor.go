package processor

import (
	"fmt"
	"gitlab.tech.mvideo.ru/mvideoru/debug/generator"
	"gitlab.tech.mvideo.ru/mvideoru/debug/parser"
	"os"
	"path/filepath"
)

func ProcessDir(dir string) error {
	fset, files, err := parser.ParseDir(dir)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}

	pkgName := files[0].Name.Name

	fileData := parser.ExtractFileData(files)

	allImports := parser.CollectImports(fileData.Imports)
	structs := parser.CollectStructs(fset, fileData.Decls)
	if len(structs) == 0 {
		return nil
	}

	needed := parser.NeededImports(structs, allImports)
	src, err := generator.GenerateSource(pkgName, structs, needed)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, fmt.Sprintf("%s_getters.go", files[0].Name)), src, 0644)
}
