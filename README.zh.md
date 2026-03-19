# docx-go

> **重要说明**  
> 本仓库由 [python-openxml/python-docx](https://github.com/python-openxml/python-docx) 翻译而来，是其在 Go 语言中的实现版本。  
> [English](README.md)

---

### 项目概述

**docx-go** 是一个用于读取、创建和修改 Microsoft Word 2007+ (.docx) 文件的 Go 语言库。

本仓库从 [python-openxml/python-docx](https://github.com/python-openxml/python-docx) 翻译而来，将原 Python 库的核心功能与 API 设计移植到 Go 语言，为 Go 开发者提供与 python-docx 相似的文档操作体验。

### 主要功能

- 创建新文档或打开已有 .docx 文件
- 添加段落、标题、表格
- 支持样式、分页、分节
- 处理图片、形状、绘图
- 支持批注（Comments）
- 基于 OPC (Open Packaging Conventions) 和 OOXML 的底层实现

### 安装

```bash
go get github.com/docx-go
```

### 快速示例

```go
package main

import (
    "github.com/docx-go/docx"
)

func main() {
    // 创建新文档
    doc := docx.New()
    doc.AddParagraph("这是一个风雨交加的夜晚。", "")
    doc.Save("output.docx")

    // 打开已有文档
    doc2, _ := docx.Open("output.docx")
    paras := doc2.Paragraphs()
    // paras[0].Text() 返回 "这是一个风雨交加的夜晚。"
}
```

### 项目结构

- `docx/` - 文档、段落、表格、节等核心 API
- `text/` - 文本、段落、Run 等
- `opc/` - OPC 包与部件
- `oxml/` - OOXML 解析与生成
- `styles/` - 样式
- `comments/` - 批注
- `image/`、`drawing/`、`dml/` - 图片与绘图

---

## License

MIT（与 [python-docx](https://github.com/python-openxml/python-docx) 一致）
