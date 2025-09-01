package main

import (
	"gorm.io/gorm"
)

// （模型定义同上，此处省略，实际使用时需包含）

// 查询某个用户发布的所有文章及其对应的评论信息
func getUserArticlesWithComments(db *gorm.DB, userID uint) (*User, error) {
	var user User
	// 使用Preload嵌套加载关联数据  当需要查询主模型时同时加载其关联数据，可以使用 Preload 指定关联字段。
	result := db.Preload("Posts.Comments").First(&user, userID)
	return &user, result.Error
}

// 查询评论数量最多的文章信息
func getMostCommentedPost(db *gorm.DB) (*Post, int64, error) {
	// 先统计各文章的评论数量
	type CommentCount struct {
		PostID uint
		Count  int64
	}
	var commentCounts []CommentCount

	// 分组统计评论数  select count(*) as count,post_id from comments group by post_id order by count desc
	err := db.Model(&Comment{}).
		Select("post_id, count(*) as count"). //查询的字段
		Group("post_id").
		Order("count desc").
		Limit(1).
		Scan(&commentCounts).Error

	if err != nil || len(commentCounts) == 0 {
		return nil, 0, err
	}

	// 获取评论最多的文章详情
	var post Post
	err = db.Preload("Comments").First(&post, commentCounts[0].PostID).Error
	return &post, commentCounts[0].Count, err
}

//func main() {
//	// 连接数据库
//	dsn := "root:123456@tcp(127.0.0.1:3306)/web3?parseTime=true&charset=utf8mb4&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		println("failed to connect database: " + err.Error())
//	}
//
//	// 示例：查询ID为1的用户的所有文章及评论
//	user, err := getUserArticlesWithComments(db, 1)
//	if err != nil {
//		fmt.Println("查询用户文章失败:", err)
//	} else {
//		fmt.Printf("用户 %s 的文章列表:\n", user.Name)
//		for _, post := range user.Posts {
//			fmt.Printf("- 文章: %s (评论数: %d)\n", post.Title, len(post.Comments))
//		}
//	}
//
//	// 示例：查询评论最多的文章
//	post, count, err := getMostCommentedPost(db)
//	if err != nil {
//		fmt.Println("查询评论最多的文章失败:", err)
//	} else {
//		fmt.Printf("\n评论最多的文章: %s (评论数: %d)\n", post.Title, count)
//	}
//}
