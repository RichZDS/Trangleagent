# 数据库迁移说明

## 三角机构 TRPG 结构优化

执行 `001_optimize_trpg_role_structure.sql` 以完成以下优化：

1. **用户表**：新增 `active_role_id`，用于记录当前选中的角色卡
2. **用户表**：`red_trace`/`yellow_trace`/`blue_trace` 标记为弃用（轨道数据以 `role_cards` 为准）
3. **角色卡表**：添加 `user_id`、`department_id` 索引
4. **房间成员表**：新建 `trpg_room_member`，记录用户加入房间时选用的角色

### 执行方式

```bash
mysql -u root -p trangleagent < manifest/sql/001_optimize_trpg_role_structure.sql
```

或使用数据库客户端执行该 SQL 文件。

### 执行后

- 若使用 GoFrame 代码生成，可运行 `gf gen dao` 更新 DAO（`hack/config.yaml` 已包含 `trpg_room_member`）
- 前端加入房间时会弹出角色选择框，需选择角色卡后再加入
