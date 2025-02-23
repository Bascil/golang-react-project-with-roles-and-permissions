package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
	"strconv"
)

func GetProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page","1"))
	
	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))

}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	product := models.Product{
		Id:uint(id),
	}

	database.DB.Find(&product)

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var data map[string]string //array with key string and value string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	price, _ := strconv.Atoi(data["price"]) //id and error

	product := models.Product{
		Title:data["title"],
		Description:data["description"],
		Image:data["image"],
		Price: float64(price),
	}

	database.DB.Create(&product);

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	product := models.Product{
		Id:uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	product := models.Product{
		Id:uint(id),
	}

	database.DB.Delete(&product)

	return nil
}