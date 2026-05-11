package ds_xml_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_xml"
)

//go:embed testdata/tmp.xml
var tplXML []byte

func TestReadBufferToDocument(t *testing.T) {
	doc, err := ds_xml.ReadBufferToDocument([]byte(tplXML))
	if err != nil {
		fmt.Printf("解析 XML 失败: %v\n", err)
		return
	}
	// 打印解析后的树结构，验证父子关系
	printTree(doc.Root, 0)
}

// printTree 递归打印树结构，用于展示结果
func printTree(node *ds_xml.Node, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	fmt.Printf("%s<%s", indent, node.Name.Local)
	for _, attr := range node.Attrs {
		fmt.Printf(" %s=\"%s\"", attr.Name.Local, attr.Value)
	}
	fmt.Printf(">")

	if node.CharData != "" {
		fmt.Printf("%s", node.CharData)
	}

	// 递归打印所有子节点
	for _, child := range node.Children {
		printTree(child, depth+1)
	}

	fmt.Printf("%s</%s>\n", indent, node.Name.Local)
}
