package ds_xml

import (
	"bytes"
	"encoding/xml"
	"io"
)

// CollectNamespacePrefixes 遍历 XML 文档，收集所有命名空间声明。
// 返回 map[uri]prefix，其中 prefix 可能为空字符串（表示默认命名空间）。
// 如果同一个 URI 被多次声明，后出现的会覆盖先出现的。
// Go解析的xml的属性key时，会将命名空间前缀去掉，只保留本地名称和命名空间URI，这是直接解析到struct时的行为。
// 所以这里做一些处理，可用于测试对照。
//
// 解析后，nsMap 将包含：
func CollectNamespacePrefixes(reader io.Reader) (map[string]string, error) {
	decoder := xml.NewDecoder(reader)
	nsMap := make(map[string]string) // URI -> 前缀
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		switch se := token.(type) {
		case xml.StartElement:
			for _, attr := range se.Attr {
				// 判断是否为命名空间声明
				// attr.Name.Space 为空字符串表示默认命名空间
				// attr.Name.Local 为本地名称，例如
				// attr.Value 为命名空间URI
				if attr.Name.Space == "xmlns" {
					nsMap[attr.Value] = attr.Name.Local
				}
			}
		}
	}
	return nsMap, nil
}

// ParseToTree 将 XML 字节流解析为一棵完整的树，保留所有父子关系
func ParseToTree(reader io.Reader) (*Document, error) {
	decoder := xml.NewDecoder(reader)
	// 使用栈来跟踪当前的解析路径，遇到关闭标签时弹出栈顶节点，最后一个栈节点就是当前解析节点
	var stack []*Node
	doc := &Document{
		NsMap: make(map[string]string), // URI -> 前缀
	}
	for {
		// 读取下一个 Token
		token, err := decoder.Token()
		if err != nil {
			// 解析完成或遇到错误
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			// 遇到开始标签，创建一个新节点
			node := &Node{
				Document: doc,
				Name:     t.Name,
				Attrs:    t.Attr,
				Children: make([]*Node, 0),
			}
			// 更新命名空间URI映射表
			for _, attr := range t.Attr {
				if attr.Name.Space == "xmlns" {
					doc.NsMap[attr.Value] = attr.Name.Local
				}
			}
			// 如果是根节点
			if doc.Root == nil {
				doc.Root = node
			} else {
				// 否则，将新节点添加到栈顶节点（即其父节点）的 Children 中
				parent := stack[len(stack)-1]
				parent.AddChild(node)
			}
			// 将新节点压入栈，它将成为后续子元素的父节点
			stack = append(stack, node)

		case xml.EndElement:
			// 遇到结束标签，表示当前节点解析完成，从栈中弹出
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			// 遇到文本内容，将其添加到当前栈顶节点
			if len(stack) > 0 {
				// 去除首尾空白字符
				text := string(bytes.TrimSpace(t))
				if text != "" {
					stack[len(stack)-1].CharData += text
				}
			}
		case xml.Comment:
			// 遇到注释
		case xml.Directive:
			// 遇到指令
		case xml.ProcInst:
			// 遇到处理指令
		}
	}
	return doc, nil
}
