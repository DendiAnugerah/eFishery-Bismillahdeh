package productcontroller

import (
	"Fish/model/category"
	"Fish/model/product"
	categoryservice "Fish/service/category"
	productservice "Fish/service/product"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service  productservice.ProductService
	category categoryservice.CategoryService
}

func NewProdController(prodrep product.Repository, catrep category.Repository) *ProductController {
	return &ProductController{service: *productservice.NewProdService(prodrep, catrep)}
}

func (p *ProductController) CreateProduct(c echo.Context) error {
	req := new(productservice.CreateProduct)
	if err := c.Bind(req); err != nil {
		log.Print("result error")
		return err
	}

	err := p.service.Create(req)
	if err != nil {
		log.Print("service error")
		return err
	}
	return c.JSON(200, req)
}

func (p *ProductController) GetProduct(c echo.Context) error {
	type P map[string]interface{}
	req := new(productservice.FindById)
	id := c.Param("id")
	var err error

	req.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(400, http.StatusBadRequest)
		return err
	}

	product, err := p.service.FindById(req)

	if err != nil {
		c.JSON(400, http.StatusBadRequest)
		return err
	}
	return c.JSON(200, P{"product": product})
}

func (p *ProductController) GetProducts(c echo.Context) error {
	type P map[string]interface{}
	fmt.Println(c.QueryParams())

	products, err := p.service.FindAll()
	if err != nil {
		c.JSON(400, http.StatusBadRequest)
		return err
	}

	return c.JSON(200, P{"products": products})
}

func (p *ProductController) UpdateProduct(c echo.Context) error {
	type P map[string]interface{}
	req := new(productservice.UpdateProduct)
	if err := c.Bind(req); err != nil {
		c.JSON(400, http.StatusBadRequest)
		return err
	}

	err := p.service.Update(req)
	if err != nil {
		c.JSON(400, http.StatusBadRequest)
		return err
	}
	return c.JSON(200, P{"message": "Product telah diupdate"})
}

func (p *ProductController) DeleteProduct(c echo.Context) error {
	type P map[string]interface{}
	req := new(productservice.DeleteProduct)

	id := c.Param("id")
	var err error

	req.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(400, http.StatusBadRequest)
		return err
	}

	err = p.service.Delete(req)
	if err != nil {
		c.JSON(400, http.StatusBadGateway)
		return err
	}
	return c.JSON(200, P{"message": "Product deleted successfully"})
}
