package create

import (
	"fmt"
	"github.com/zhouwu4740/tobegopher/gorm/quickstart"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Create() {
	dsn := fmt.Sprintf("root:123456@tcp(localhost:3306)/crawler?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("建立数据库连接失败")
	}
	// 设置表选项标签，指定数据库引擎和字符集
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8")
	students := []quickstart.Student{
		{Name: "xiaohong1"},
		{Name: "xiaohuang1"},
		{Name: "xiaohei1"},
	}
	db.CreateInBatches(students, 1)

	fmt.Println("done")
}
