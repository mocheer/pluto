# ds_docx

入口文件为`word/document.xml`

```text
docx 根目录
├── [Content_Types].xml       # 内容类型定义（必须，定义各部分的 MIME 类型）
├── _rels/
│   └── .rels                 # 根关系文件（定义文档与各部分之间的起始关联）
├── docProps/                 # 文档属性（核心元数据）
│   ├── core.xml              # 核心属性（标题、作者、创建时间等，对应 Dublin Core）
│   └── app.xml               # 应用属性（页数、字数、程序版本等）
├── word/                     # 主文档内容所在目录
│   ├── document.xml          # 主文档内容（所有文本、段落、表格、书签等）
│   ├── styles.xml            # 样式定义（段落样式、字符样式、表格样式等）
│   ├── numbering.xml         # 编号定义（多级列表、项目符号、编号格式）
│   ├── fontTable.xml         # 字体表（文档中使用的字体映射）
│   ├── settings.xml          # 文档设置（页面设置、安全设置、显示选项等）
│   ├── webSettings.xml       # Web 视图设置（针对网页浏览的优化）
│   ├── comments.xml          # 批注（所有审阅批注）
│   ├── commentsExtended.xml  # 批注扩展信息（如提及、回复关系）
│   ├── people.xml            # 人员信息（批注者、审阅者身份）
│   ├── footnotes.xml         # 脚注
│   ├── endnotes.xml          # 尾注
│   ├── header1.xml           # 页眉（可能有多个 header2.xml 等）
│   ├── footer1.xml           # 页脚（可能有多个 footer2.xml 等）
│   ├── theme/
│   │   └── theme1.xml        # 主题定义（颜色、字体、效果方案）
│   ├── media/                # 嵌入资源（图片、视频、音频、OLE 对象等）
│   │   ├── image1.png
│   │   ├── image2.jpg
│   │   └── ...
│   ├── embeddings/           # 嵌入的 OLE 对象或外部文件（可选）
│   │   ├── oleObject1.bin
│   │   └── ...
│   ├── customXml/            # 自定义 XML 数据（用于内容控件绑定）
│   │   ├── item1.xml
│   │   ├── itemProps1.xml
│   │   └── ...
│   └── _rels/                # word 目录下的关系文件
│       ├── document.xml.rels # 主文档中嵌入资源、页眉、脚注等的关系映射
│       ├── header1.xml.rels  # 页眉中资源的关系映射
│       ├── footer1.xml.rels  # 页脚中资源的关系映射
│       └── ...
├── customXml/                # 根级自定义 XML（若存在全局自定义项）
│   ├── item1.xml
│   ├── itemProps1.xml
│   └── _rels/
│       └── item1.xml.rels
└── ...
```

## 主要标签和含义

|元素|描述|对应 HTML
|---|---|---|
|<w:p>|段落|<p>|
|<w:r>|文本运行（run），一段具有相同格式的文本|	<span> 或直接放在 <p> 内|
|<w:t>|文本内容|文本节点|
|<w:rPr>|运行属性（字体、大小、粗体等）|CSS 样式|
|<w:pPr>|段落属性（对齐、缩进、行距等）|CSS 样式|
|<w:tbl>|表格|<table>|
|<w:tr>|表格行|<tr>|
|<w:tc>|表格单元格|<td>|
|<w:drawing>|绘图对象（图片、形状等）|<img> 或 <svg>|
|<w:hyperlink>|超链接|<a>|
|<w:bookmarkStart/End>|书签|可忽略或转为锚点|


## 参考
- https://github.com/fumiama/go-docx: 推荐
- https://github.com/sajari/docconv：太多依赖，连windows的依赖安装都没有说明
- https://github.com/gomutex/godocx：推荐
- https://learn.microsoft.com/zh-cn/previous-versions/office/gg607163(v=office.14)：docx 格式规范
- https://ecma-international.org/publications-and-standards/standards/ecma-376/

