package router

import (
	"github.com/gin-gonic/gin"
	"storage/database"
	"storage/utils"
)

func Initialized() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("template/*.html")
	router.Static("assets", "assets")

	router.GET("/", index)
	router.GET("/addProduct", addProductGet)
	router.PUT("/api/motionProduct/:motion", motionProduct)
	router.GET("/delProduct", delProductGet)
	router.GET("/journal", getJournal)

	return router
}

func getJournal(c *gin.Context) {
	journal := database.GetJournal()

	c.HTML(200, "journal", gin.H{
		"Title":   "Журнал",
		"Journal": journal,
	})
}

func delProductGet(c *gin.Context) {
	products := database.GetProducts()

	c.HTML(200, "product", gin.H{
		"Title":    "Использование товара",
		"JS":       "del_product.js",
		"Products": products,
		"Btn":      "Использовать",
	})
}

func motionProduct(c *gin.Context) {
	motion := c.Param("motion")
	var product database.Product
	e := c.BindJSON(&product)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	alert := product.MotionProduct(motion)
	if motion == "add" {
		if alert != "Ошибка" {
			c.JSON(200, true)
		} else {
			c.JSON(200, false)
		}
	} else {
		if alert != "Ошибка" {
			c.JSON(200, alert)
		}
	}
}

func addProductGet(c *gin.Context) {
	products := database.GetProducts()

	c.HTML(200, "product", gin.H{
		"Title":    "Поступление товара",
		"JS":       "add_product.js",
		"Products": products,
		"Btn":      "Добавить",
	})
}

func index(c *gin.Context) {
	categories := database.GetCategory()
	storage := database.GetStorage()

	c.HTML(200, "index", gin.H{
		"Categories": categories,
		"Storage":    storage,
		"Title":      "Склад",
	})
}
