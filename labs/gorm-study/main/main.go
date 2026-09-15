package main

import (
	gorm_study "gorm-study"
)

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	ID       uint
	Name     string
	Articles []Article `gorm:"many2many:article_tags;"` // 用于反向引用
}

type ArticleTag struct {
	ArticleID uint
	TagID     uint
}

func main() {
	DB := gorm_study.Initialization()
	
	DB.AutoMigrate(&Article{}, &Tag{}, &ArticleTag{})
	//DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
	DB.SetupJoinTable(&Tag{}, "Articles", &ArticleTag{})
	DB.Create(&Article{
		Title: "flask零基础入门",
		Tags: []Tag{
			{Name: "python"},
			{Name: "后端"},
			{Name: "web"},
		},
	})
}