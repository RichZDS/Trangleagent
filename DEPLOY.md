# Docker Compose 部署

## 快速启动

```bash
# 构建并启动所有服务
docker compose up -d --build

# 查看日志
docker compose logs -f backend
```

首次构建可能需要几分钟（Go 编译 + npm 构建）。

## 服务说明

| 服务 | 端口 | 说明 |
|------|------|------|
| frontend | 80 | 前端（Nginx 静态 + API 代理） |
| backend | 8888 | 后端 API |
| mysql | 3306 | 数据库 |
| minio | 9000, 9001 | 对象存储（API / 控制台） |
| rabbitmq | 5672, 15672 | 消息队列（AMQP / 管理界面） |

## 访问

- **前端**: http://localhost
- **MinIO 控制台**: http://localhost:9001（minioadmin / minioadmin）
- **RabbitMQ 管理**: http://localhost:15672（guest / guest）

## 首次部署

1. MySQL 会自动执行 `Backend/manifest/sql/` 下的 SQL 初始化数据库
2. MinIO 需在控制台创建 bucket `trangleagent`（或修改 config 中的 bucket 名）
3. 注册首个用户后即可登录

## 环境变量（可选）

可在 `docker-compose.yml` 中为 backend 添加：

- `GF_GCFG_FILE`: 配置文件路径（默认 `manifest/config/config.docker.yaml`）

## 停止

```bash
docker compose down

# 同时删除数据卷
docker compose down -v
```
