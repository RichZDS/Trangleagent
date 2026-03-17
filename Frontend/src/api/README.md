# Swagger 文档自动生成配置说明

本文档介绍了如何使用 Swagger/OpenAPI 文档自动生成 TypeScript 接口和服务代码。

## 安装全局 CLI 工具

```bash
npm install -g @zeronejs/cli
```

## 配置文件

在项目中创建 `swagger.config.json` 文件：

```json
{
  "docsUrl": "./docs.json",
  "axiosInstanceUrl": "../../http.js"
}
```

其中：
- `docsUrl`: 指向 Swagger 的 JSON 文档地址，可以是远程 URL 或本地文件路径
- `axiosInstanceUrl`: 指向项目中 Axios 实例的位置

## 使用方法

### 生成 TypeScript 代码

```bash
# 在项目根目录执行
zerone api

# 或者指定特定路径
zerone api -p ./src/api
```

### 生成 JavaScript 代码

```bash
# 生成 JS 代码（基于 axios）
zerone api -js

# 指定路径生成 JS 代码
zerone api -js -p ./src/api
```

## 项目结构

生成的代码将包含以下内容：

1. `interface/` - 包含所有从 Swagger 模式生成的 TypeScript 接口定义
2. `controller/` - 包含基于 API 路径生成的服务函数

这些生成的代码具有完整的 TypeScript 类型支持和 IDE 代码提示功能。