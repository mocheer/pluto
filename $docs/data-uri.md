# data-uri

Data URI scheme 在前端开发中是个常用的技术，可以在浏览器地址栏打开，也可以通过标签嵌入

```
data:[<mime type>][;charset=<charset>][;base64],<encoded data>

data:image/png;base64,<base64 data>

```

1. 第一部分是 data: 协议头，它标识这个内容为一个 data URI 资源。
2. 第二部分是 MIME 类型，表示这串内容的展现方式，比如：text/plain（默认），则以文本类型展示，image/jpeg，以 jpeg 图片形式展示，同样，客户端也会以这个 MIME 类型来解析数据。
3. 第三部分是编码设置，默认编码是 charset=US-ASCII, 即数据部分的每个字符都会自动编码为 %xx
4. 第四部分是 base64 编码设定，这是一个可选项，base64 编码中仅包含 0-9,a-z,A-Z,+,/,=，其中 = 是用来编码补白的。
5. 最后一部分为这个 Data URI 承载的内容，它可以是纯文本编写的内容，也可以是经过 base64编码 的内容。