package ds_xml

import (
	"encoding/xml"
	"slices"
	"strconv"
)

// Document 代表 XML 文档
type Document struct {
	Root  *Node
	NsMap map[string]string // 命名空间URI映射表
}

// Node 代表 XML 树中的一个节点
type Node struct {
	Name     xml.Name   // 元素的标签名，包括命名空间
	Attrs    []xml.Attr // 元素的所有属性
	CharData string     // 元素的纯文本内容
	Parent   *Node      // 指向父节点的指针，形成链式关系
	Children []*Node    // 子节点切片
	Document *Document  // 指向文档的指针
}

// NodeAttrValue 代表 XML 元素的属性值
type NodeAttrValue string

// GetTagName 获取元素的标签名，包括命名空间前缀
// tagName 标签名，包含命名空间前缀，例如 "w:p" 表示 "w" 映射的命名空间中的 "p" 元素
func (m *Node) GetTagName(xName xml.Name) string {
	name := xName.Local
	// 命名空间前缀可能不为空
	if xName.Space != "" {
		// 获取命名空间URI映射表+本地名称
		prefix, ok := m.Document.NsMap[xName.Space]
		if ok {
			name = prefix + ":" + name
		}
	}
	return name
}

// IsTagName 是否是指定的标签名
func (n *Node) IsTagName(tagName string) bool {
	return n.GetTagName(n.Name) == tagName
}

// HIndex 查找索引
func (n *Node) HIndex() int {
	if n.Parent == nil {
		return -1
	}
	return slices.Index(n.Parent.Children, n)
}

// Depth
func (n *Node) Depth() int {
	i := 0
	parent := n.Parent
	for parent != nil {
		parent = parent.Parent
		i++
	}
	return i
}

// AddChild 添加子节点
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
	child.Parent = n
}

// First
func (n *Node) First(tagNames ...string) *Node {
	node := n
	for _, tagName := range tagNames {
		var cNode *Node
		for _, child := range node.Children {
			if child.IsTagName(tagName) {
				cNode = child
				break
			}
		}
		if cNode == nil {
			return nil
		}
		node = cNode
	}
	return node
}

// Find 查找所有子节点
func (n *Node) Find(tagName string) []*Node {
	var nodes []*Node
	for _, child := range n.Children {
		if child.IsTagName(tagName) {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// Attr 查找属性值
func (n *Node) Attr(name string) *NodeAttrValue {
	for _, attr := range n.Attrs {
		if n.GetTagName(attr.Name) == name {
			return new(NodeAttrValue(attr.Value))
		}
	}
	return nil
}

// AttrsMap 获取所有属性值
func (n *Node) AttrsMap() map[string]*NodeAttrValue {
	m := make(map[string]*NodeAttrValue, len(n.Attrs))
	for _, attr := range n.Attrs {
		m[n.GetTagName(attr.Name)] = new(NodeAttrValue(attr.Value))
	}
	return m
}

// ToInt
func (n *NodeAttrValue) ToInt() int {
	i, _ := strconv.Atoi(string(*n))
	return i
}

// ToFloat
func (n *NodeAttrValue) ToFloat64() float64 {
	f, _ := strconv.ParseFloat(string(*n), 64)
	return f
}

// ToBool
func (n *NodeAttrValue) ToBool() bool {
	b, _ := strconv.ParseBool(string(*n))
	return b
}
