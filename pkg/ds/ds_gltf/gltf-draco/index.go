package gltfdraco

import (
	"log"
	"os"

	"github.com/qmuntal/draco-go/draco"
)

func ReadFile() error {
	data, err := os.ReadFile("../../testdata/test_nm.obj.edgebreaker.cl4.2.2.drc")
	if err != nil {
		return err
	}
	m := draco.NewMesh()
	d := draco.NewDecoder()
	if err := d.DecodeMesh(m, data); err != nil {
		log.Fatalf("failed to decode mesh: %v", err)
	}
	log.Println("point count:", m.NumPoints())
	log.Println("face count:", m.NumFaces())
	log.Println("faces:", m.Faces(nil))
	return nil
}

// func ReadGltf() error {
// 	doc, err := gltf.Open("testdata/box/Box.gltf")
// 	if err != nil {
// 		return err
// 	}

// 	pd, _ := gltfDraco.UnmarshalMesh(doc, doc.BufferViews[0])
// 	p := doc.Meshes[0].Primitives[0]
// 	fmt.Println(pd.ReadIndices(nil))
// 	fmt.Println(pd.ReadAttr(p, "POSITION", nil))
// 	fmt.Println(pd.ReadAttr(p, "NORMAL", nil))
// 	return nil
// }
