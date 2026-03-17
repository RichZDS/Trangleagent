# 项目启动说明

## 前置条件

1. **MySQL**：需创建数据库 `trangleagent`，并执行迁移脚本
2. **Node.js**：用于运行前端
3. **RabbitMQ**（可选）：用于聊天室多实例广播，不启动则自动降级为本地模式

## 数据库

```bash
# 创建数据库
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS trangleagent CHARACTER SET utf8mb4;"

# 方式一：全新建表（推荐，直接可用的完整建表脚本）
mysql -u root -p trangleagent < Backend/manifest/sql/000_init_schema.sql

# 方式二：已有基础表时，执行增量迁移
# mysql -u root -p trangleagent < Backend/manifest/sql/001_optimize_trpg_role_structure.sql
# mysql -u root -p trangleagent < Backend/manifest/sql/002_user_exp_cleanup.sql
```

> `000_init_schema.sql` 会 DROP 并重建所有表，仅适用于全新环境。

## 配置

- **后端**：`Backend/manifest/config/config.yaml`
  - 数据库：`database.default.link`（默认 `root:123456@127.0.0.1:3306`）
  - 端口：`server.address`（默认 `:8888`）
- **前端**：`Frontend/vite.config.js` 代理目标为 `localhost:8888`

## 启动

### 1. 后端

```bash
cd Backend
go run main.go
```

访问：http://localhost:8888/swagger/

### 2. 前端

```bash
cd Frontend
npm install   # 首次需安装依赖
npm run dev
```

访问：http://localhost:5173/

## 验证

- 后端编译：`cd Backend && go build ./...`
- 前端构建：`cd Frontend && npm run build`
