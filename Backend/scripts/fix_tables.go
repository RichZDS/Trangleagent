//go:build ignore
// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/trangleagent")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS forum_comments;")
	if err != nil {
		log.Fatal("Drop forum_comments error:", err)
	}

	createCommentsTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论ID',
		user_id BIGINT UNSIGNED NOT NULL COMMENT '评论用户ID',
		target_type VARCHAR(32) NOT NULL COMMENT '被评论对象类型（如 post/image/video 等）',
		target_id BIGINT UNSIGNED NOT NULL COMMENT '被评论对象ID',
		parent_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父评论ID（回复某条评论时填写）',
		root_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '根评论ID（同一楼的顶层评论ID，便于树查询）',
		content TEXT NOT NULL COMMENT '评论内容',
		status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常 0屏蔽 2删除',
		like_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
		reply_count INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复数',
		created_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		updated_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
		deleted_at DATETIME NULL DEFAULT NULL COMMENT '软删除时间',
		PRIMARY KEY (id),
		KEY idx_target (target_type, target_id),
		KEY idx_user_id (user_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='多态通用评论表';
	`
	_, err = db.Exec(createCommentsTable)
	if err != nil {
		log.Fatal("Create comments error:", err)
	}

	// Ensure the user_follows exists, and users/role_cards are altered
	refactorSQL := []string{
		"ALTER TABLE users DROP COLUMN reality_role;",
		"ALTER TABLE users DROP COLUMN abnormal_role;",
		"ALTER TABLE users DROP COLUMN job_title;",
		"ALTER TABLE users ADD COLUMN extra_info JSON NULL COMMENT '扩展配置信息';",
		"ALTER TABLE role_cards ADD COLUMN extra_info JSON NULL COMMENT '扩展配置信息';",
	}
	
	for _, sqlStmt := range refactorSQL {
		_, err := db.Exec(sqlStmt)
		if err != nil {
			log.Printf("Executing %s error: %v", sqlStmt, err)
		}
	}

	fmt.Println("Tables fixed.")
}
