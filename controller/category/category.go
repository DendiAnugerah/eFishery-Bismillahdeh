package categorycontroller

import (
	"Fish/model/category"
	categoryservice "Fish/service/category"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	service categoryservice.CategoryService
}

func NewCatController(repository category.Repository) *CategoryController {
	return &CategoryController{
		service: *categoryservice.NewCatService(repository),
	}
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
	req := new(categoryservice.CreateCategory)
	if err := c.Bind(req); err != nil {
		log.Printf("Result err")
		return err
	}

	err := cc.service.Create(req)
	if err != nil {
		log.Printf("Controller err")
		return err
	}

	return c.JSON(200, req)
}

func (cc *CategoryController) GetAllCategory(c echo.Context) error {
	cat, err := cc.service.FindAll()
	if err != nil {
		log.Printf("Srvice err")
		return err
	}
	return c.JSON(200, cat)
}

func (cc *CategoryController) GetCategory(c echo.Context) error {
	req := new(categoryservice.FindById)
	id := c.Param("id")

	var err error

	req.Id_category, err = strconv.Atoi(id)
	if err != nil {
		log.Printf("err")
		return err
	}
	category, err := cc.service.FindById(req)
	if err != nil {
		log.Printf("err")
		return err
	}
	return c.JSON(200, category)
}

func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	type M map[string]interface{}
	req := new(categoryservice.UpdateCategory)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := cc.service.Update(req)
	if err != nil {
		return err
	}
	return c.JSON(200, M{"message": "Category updated"})
}

func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	type M map[string]interface{}
	req := new(categoryservice.DeleteCategory)
	id := c.Param("id")
	var err error
	req.Id_category, err = strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = cc.service.Delete(req)
	if err != nil {
		return err
	}
	return c.JSON(200, M{"message": "Category deleted"})
}
