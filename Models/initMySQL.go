package Models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"wPan/v1/Config"
)

var DB *gorm.DB

func InitMySQL() error {
	var err error
	DB, err = gorm.Open(
		Config.DatabaseSetting.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			Config.DatabaseSetting.User,
			Config.DatabaseSetting.Password,
			Config.DatabaseSetting.Host,
			Config.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
		return err
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return Config.DatabaseSetting.TablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	return DB.DB().Ping()
}
