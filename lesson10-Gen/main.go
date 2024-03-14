package main

import (
	"Gen/dal/model"
	"Gen/dal/query"
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const MySQLDSN = "root:abc123@tcp(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True"

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	fmt.Println("gen demo start-------")
	db := ConnectDB(MySQLDSN)

	//重要 调用dal层生成的代码。并设置默认的db对象
	query.SetDefault(db)

	//CRUD
	b1 := model.Post{
		Title:    "你好",
		AuthorID: 2,
		Content:  "这个是gen框架的学习 ",
		Status:   1,
	}

	//调用生成的增删改查代码
	err := query.Post.WithContext(context.Background()).Create(&b1)
	if err != nil {
		fmt.Printf("create post fail err :%v", err)
	}
	fmt.Printf("b1:%#v\n", b1)

	//查询方法1 ：调用查询的操作
	b, err := query.Post.WithContext(context.Background()).Where(query.Post.PostID.Eq(1)).First()
	fmt.Printf("b :%#v\n  err:%#v", b, err)

	//查询方法2 ：
	post := query.Post
	post.WithContext(context.Background()).Where(post.ID.Eq(2)).First()

	//更新操作
	ret, err := query.Post.WithContext(context.Background()).Where(query.Post.ID.Eq(4)).Update(query.Post.Content, "你好")
	if err != nil {
		fmt.Printf("update post fail,err :%#v\n", err)
	}
	fmt.Println(ret.RowsAffected)

	//删除操作
	//方法1
	b3 := model.Post{ID: 4}
	ret, err = query.Post.WithContext(context.Background()).Delete(&b3)
	if err != nil {
		fmt.Printf("detete failed err:%#v", err)
	}
	fmt.Println(ret.RowsAffected)

	//方法2
	ret, err = query.Post.WithContext(context.Background()).Where(query.Post.ID.Eq(3)).Delete()
	if err != nil {
		fmt.Printf("detete failed err:%#v", err)
	}
	fmt.Printf("RowsAfected :%v", ret)

}
