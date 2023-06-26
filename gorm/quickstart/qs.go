package quickstart

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	Height      float32 `json:"height"`
	Nearsighted bool    `json:"nearsighted"`
	gorm.Model
}

// QuickStart 快速入门实例
func QuickStart() {
	// root:123456@tcp(localhost:3306)/crawler?charset=utf8&parseTime=True&loc=Local
	// root:123456@tcp(localhost:3306)/crawler?charset=utf8&parseTime=True&loc=Local&timeout=10s
	// root:123456@tcp(localhost:3306)/crawler?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s
	dsn := fmt.Sprintf("root:123456@tcp(localhost:3306)/crawler?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("建立数据库连接失败")
	}
	// 设置表选项标签，指定数据库引擎和字符集
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8")
	// 同步schema
	//db.AutoMigrate(&Student{})

	// 插入数据
	//s := Student{Name: "张三", Age: 41, Height: 167.5, Nearsighted: true}
	//db.Create(&s)

	// 查询
	//var students []Student
	//db.Where("id >= ?", 1).Find(&students)
	//
	//for _, student := range students {
	//	fmt.Println(student)
	//}

	// 更新
	//var s Student
	//db.Model(&s).Where("id=?", 1).Updates(&Student{Name: "lisi", Age: 37})

	// 开启事务
	//txn := db.Begin()
	//txn.Create(&Student{Name: "tianqi", Age: 20, Height: 188.5, Nearsighted: true})
	//txn.Create(&Student{Name: "wangwu", Age: 10, Height: 145.5, Nearsighted: true})
	//txn.Create(&Student{Name: "zhaoliu", Age: 6, Height: 98.5, Nearsighted: true})
	////txn.Rollback()
	//txn.Commit()

	s := Student{Name: "select", Age: 20, Height: 188.5, Nearsighted: true}
	db.Select("Name", "Age").Create(&s)
	fmt.Println("done")
}
