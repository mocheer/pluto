package ds_xml_test

import (
	"fmt"
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_xml"
)

var tplXML = `<w:document xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas"
	xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
	xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"
	xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math"
	xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"
	xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
	xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
	xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
	xmlns:w10="urn:schemas-microsoft-com:office:word"
	xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml"
	xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup"
	xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk"
	xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml"
	mc:Ignorable="w14 w15 wp14">
	<w:body>
		<w:p>
		<w:rPr>
			<w:rFonts w:hint="eastAsia" w:ascii="仿宋_GB2312" w:hAnsi="仿宋_GB2312"
				w:eastAsia="仿宋_GB2312" w:cs="仿宋_GB2312" />
			<w:sz w:val="32" />
			<w:szCs w:val="32" />
			<w:highlight w:val="none" />
			<w:lang w:val="en-US" w:eastAsia="zh-CN" />
		</w:rPr>
			<w:r>
				<w:t>Go 语言实战</w:t>
			</w:r>
		</w:p>
	</w:body>
</w:document>`

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
	fmt.Println()

	// 递归打印所有子节点
	for _, child := range node.Children {
		printTree(child, depth+1)
	}

	fmt.Printf("%s</%s>\n", indent, node.Name.Local)
}
