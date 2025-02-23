package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
	"strconv"
)

func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role //slice, similar to an array
	database.DB.Find(&roles)

	return c.JSON(roles)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	role := models.Role{
		Id:uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	
	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id,_ := strconv.Atoi(permissionId.(string)) // cast as string
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name: roleDto["name"].(string),
		Permissions: permissions,

	}

	database.DB.Create(&role); // create role by reference

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	
	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id,_ := strconv.Atoi(permissionId.(string)) // cast as string
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result interface{}

	database.DB.Table("role_permissions").Where("role_id", 1).Delete(&result)

	role := models.Role{
		Id: uint(id),
		Name: roleDto["name"].(string),
		Permissions: permissions,
	}


	database.DB.Preload("Permissions").Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	role := models.Role{
		Id:uint(id),
	}
	
	database.DB.Delete(&role)

	return nil
}