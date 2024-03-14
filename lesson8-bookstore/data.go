package main

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	defaluShelfSize = 10
)

// 使用GORM  --创建mysql连接使用了 gorm.Open() 函数来打开一个数据库连接。根据传入的 DSN 字符串，GORM 会自动推断出要连接的数据库类型，然后选择对应的驱动程序进行连接。
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//迁移schema 调用 db.AutoMigrate() 方法来执行数据库迁移。这个方法会根据我们定义的模型结构自动创建表格，并将数据库 schema 更新到最新版本。
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// 定义模型 models
// shelf 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

// / Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct {
	db *gorm.DB
}

// 数据库操作
// CreateShelf 创建书架入库操作
func (b *bookstore) CreateShelf(ctx context.Context, data Shelf) (*Shelf, error) {
	//参数校验
	if len(data.Theme) <= 0 {
		return nil, errors.New("invalid theme")
	}
	size := data.Size
	if size <= 0 {
		size = defaluShelfSize
	}
	//构建模型
	v := Shelf{Theme: data.Theme, Size: size, CreateAt: time.Now(), UpdateAt: time.Now()}

	//数据入库
	err := b.db.WithContext(ctx).Create(&v).Error
	return &v, err
}

// GetShelf 获取书架
func (b *bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	//构建模型
	v := Shelf{}
	err := b.db.WithContext(ctx).First(&v, id).Error
	return &v, err
}

// ListShelves 书架列表
func (b *bookstore) ListShelves(ctx context.Context) ([]*Shelf, error) {
	var vl []*Shelf
	err := b.db.WithContext(ctx).Find(&vl).Error
	return vl, err
}

// DeleShelf 删除书架
func (b *bookstore) DeleShelf(ctx context.Context, id int64) error {
	return b.db.WithContext(ctx).Delete(&Shelf{}, id).Error
}
