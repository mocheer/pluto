package ds_xml

import (
	"encoding/xml"
	"io"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts"
)

// ReadFile 读取xml文件
func ReadFile(fileName string, e any) error {
	return xml.Unmarshal(ds.MustReadFile(fileName), &e)
}

// ReadFileToMap
func ReadFileToMap(fileName string) (data ts.Map[any], err error) {
	err = ReadFile(fileName, data)
	return
}

// CollectNamespacePrefixes 遍历 XML 文档，收集所有命名空间声明。
// 返回 map[uri]prefix，其中 prefix 可能为空字符串（表示默认命名空间）。
// 如果同一个 URI 被多次声明，后出现的会覆盖先出现的。
// Go解析的xml的属性key时，会将命名空间前缀去掉，只保留本地名称和命名空间URI。
// 所以这里做一些处理，可用于测试对照。
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
				if attr.Name.Space == "xmlns" {
					nsMap[attr.Value] = attr.Name.Local
				}
			}
		}
	}
	return nsMap, nil
}
