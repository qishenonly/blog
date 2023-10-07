package cmd

import (
	"blog/global"
	"blog/models"
)

// MakeMigrations 数据库迁移
func MakeMigrations() {
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectArticles", &models.UserCollectModel{})
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.ArticleModel{},
		&models.BannerModel{},
		&models.CommentModel{},
		&models.MessageModel{},
		&models.TagModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Logger.Fatalf("数据库迁移失败：%v", err)
		return
	}
	global.Logger.Infof("数据库迁移成功!")

}
