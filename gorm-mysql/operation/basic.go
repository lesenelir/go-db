package operation

import (
	"fmt"
	"gorm.io/gorm"
)

// Product struct === products table
type Product struct {
	gorm.Model // id, created_at, updated_at, deleted_at
	Code       string
	Price      uint
}

func BasicOperation(db *gorm.DB) {
	// migrate schema：基于 struct 自定义表结构
	err := db.AutoMigrate(&Product{})
	if err != nil {
		fmt.Println("auto migrate error: ", err)
		return
	}

	// create row data
	db.Create(&Product{Code: "code01", Price: 100})
	db.Create(&Product{Code: "code02", Price: 200})
	db.Create(&Product{Code: "code03", Price: 300})

	// read row data
	var product Product
	db.First(&product, "code = ?", "code01")
	fmt.Println(product.Code, product.Price) // code01 100
	var product2 Product
	db.First(&product2, "code = ?", "code03")  // SELECT * FROM products WHERE code = 'code01';
	fmt.Println(product2.Code, product2.Price) // code03 300

	// update row data
	db.Model(&product).Update("Price", 400)                               // 针对特定的 product 进行更新，只会更新某一条数据
	db.Model(&Product{}).Where("code = ?", "code02").Update("Price", 150) // 针对所有符合条件的 product 进行更新，整表扫描更新多行
	// 更新多列
	db.Model(&product2).Updates(Product{Price: 500, Code: "code04"})                                     // 更新某一行的多列数据
	db.Model(&Product{}).Where("code = ?", "code03").Updates(Product{Price: 600, Code: "code03-update"}) // 更新多行的多列数据

	// delete row data
	// config deletetime data
	db.Where("1 = 1").Delete(&Product{}) //db.Exec("DELETE FROM products")
}
