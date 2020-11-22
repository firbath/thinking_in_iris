/*
@Time : 2020/11/19 19:39
@Author : Firbath
@File : demo_controller.go
@Software: GoLand
*/
package service

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

// 交易表
type Transaction struct {
	gorm.Model
	Entry       int64
	Exit        int64
	ContractId  uint
	Remark      string
	ActivatedAt time.Time
	Status      int8
}

// 合约表
type Contract struct {
	gorm.Model
	Entry       int64
	TotalNum    int32
	Complete    int32
	Payment     int64
	Remark      string
	ActivatedAt time.Time
	Status      int8
}

// 请求&响应 报文
type Content struct {
	Action string
	Data   Data
}

type Data struct {
}

type Sub struct {
	Data
	Name string
}

func (transaction Transaction) TableName() string {
	return "y_transaction"
}
func (contract Contract) TableName() string {
	return "y_contract"
}

func getAll() []Product {
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		println()
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: fmt.Sprintf("D_%d", rand.Intn(100)), Price: rand.Intn(100)})
	var products []Product
	db.Find(&products)
	return products
}

func demo() {
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}

type DemoController struct {
	/* dependencies */
}

// GET: http://localhost:8080/demo
func (c *DemoController) Get() Content {
	return Content{"haha", Data{}}
}

// POST: http://localhost:8080/demo
func (c *DemoController) Post(b Product) int {
	//println("Received Book: " + b.Title)

	return iris.StatusCreated
}
