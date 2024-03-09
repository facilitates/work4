package models

func migration(){
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Video{}).
		AutoMigrate(&Comment{}).
		AutoMigrate(&ChatList{}).
		AutoMigrate(&ChatHistory{})
	DB.Model(&Video{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")//User是父表
	DB.Model(&Comment{}).AddForeignKey("uid", "Video(id)", "CASCADE","CASCADE")
	DB.Model(&ChatHistory{}).AddForeignKey("uid", "ChatList(id)", "CASCADE", "CASCADE")
}