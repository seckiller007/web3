package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null"         json:"name"`
	Email     string    `gorm:"size:128;uniqueIndex;not null" json:"email"`
	PostCount int64     `gorm:"not null;default:0"       json:"post_count"` // ← 统计字段
	Posts     []Post    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"posts"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint64    `gorm:"index"                    json:"user_id"`
	User          *User     `json:"user"`
	Title         string    `gorm:"size:200;not null"        json:"title"`
	Content       string    `gorm:"type:longtext;not null"   json:"content"`
	CommentStatus string    `gorm:"size:16;not null;default:'无评论'" json:"comment_status"` // ← 评论状态
	Comments      []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    uint64    `gorm:"index;not null"           json:"post_id"`
	Post      Post      `json:"post"`
	UserID    uint64    `gorm:"index"                    json:"user_id"`
	User      *User     `json:"user"`
	Content   string    `gorm:"type:text;not null"       json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// （模型定义同上，此处添加钩子方法）

// Post模型的创建后钩子：更新用户的文章数量统计
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 找到对应的用户并将文章数量+1
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1")).Error
}

// Comment模型的删除后钩子：检查文章评论数量并更新状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 统计该文章剩余的评论数量
	var count int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	// 如果评论数量为0，更新文章的评论状态
	if count == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3?parseTime=true&charset=utf8mb4&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("failed to connect database: " + err.Error())
	}

	// 示例：创建文章（会自动触发Post的AfterCreate钩子）
	newPost := Post{Title: "测试文章", Content: "这是一篇测试文章", UserID: 1}
	db.Create(&newPost)

	// 示例：删除评论（会自动触发Comment的AfterDelete钩子）
	db.Delete(&Comment{}, 1)
}
