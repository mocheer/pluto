package pkg

import (
	"fmt"
	"go/importer"
	"go/types"
)

func GetObjects(packageName string) []*types.Named {
	pkg, err := importer.Default().Import(packageName)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return nil
	}
	scope := pkg.Scope()
	names := scope.Names()
	objs := make([]*types.Named, len(names))
	for _, name := range names {
		obj := scope.Lookup(name)
		if tn, ok := obj.Type().(*types.Named); ok {
			objs = append(objs, tn)
		}
	}
	return objs
}
