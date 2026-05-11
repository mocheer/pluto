# ds_pptx


## PPTX 文件结构

PPTX 文件本质上是一个 ZIP 压缩包，包含以下主要文件：

```
pptx
├── [Content_Types].xml       # 内容类型定义
├── docProps/
│   ├── app.xml               # 应用程序属性
│   └── core.xml              # 核心属性（作者、时间等）
├── _rels/
│   └── .rels                 # 根关系文件
└── ppt/
    ├── presentation.xml      # 主演示文稿文件
    ├── presProps.xml         # 演示文稿属性
    ├── viewProps.xml         # 视图属性
    ├── tableStyles.xml       # 表格样式
    ├── theme/                # 主题文件夹
    ├── slideMasters/         # 幻灯片母版
    ├── slideLayouts/         # 幻灯片版式
    ├── slides/               # 幻灯片内容
    ├── notesSlides/          # 备注页
    ├── media/                # 媒体资源（图片、视频等）
    └── charts/               # 图表定义
```

## 参考
- https://github.com/pipipi-pikachu/pptxtojson
- https://github.com/g21589/PPTX2HTML
- https://github.com/meshesha/PPTXjs
- https://github.com/gitbrent/PptxGenJS/
- https://github.com/pipipi-pikachu/PPTist