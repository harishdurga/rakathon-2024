package handlers

import (
	"adapt-ecomm/repos"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RootHandler struct {
	ProductGetter repos.ProductGetter
}

func NewRootHandler(pg repos.ProductGetter) *RootHandler {
	return &RootHandler{ProductGetter: pg}
}

func (rh *RootHandler) GetProducts(c echo.Context) error {
	products, err := rh.ProductGetter.GetProducts(getProductFiltersFromContext(c))
	if err != nil {
		return err
	}
	return c.JSON(200, products)
}

func getProductFiltersFromContext(c echo.Context) repos.ProductFilters {
	filters := repos.ProductFilters{}
	//product filters are part of the query string
	categoryID := c.QueryParam("category_id")
	if categoryID != "" {
		catIDInt, _ := strconv.Atoi(categoryID)
		filters.CategoryID = catIDInt
	}
	limit := c.QueryParam("limit")
	if limit != "" {
		limitInt, _ := strconv.Atoi(limit)
		filters.Limit = limitInt
	}
	sort := c.QueryParam("sort")
	if sort != "" {
		filters.Sort = sort
	}
	minPrice := c.QueryParam("min_price")
	if minPrice != "" {
		minPriceInt, _ := strconv.Atoi(minPrice)
		filters.MinPrice = minPriceInt
	}
	maxPrice := c.QueryParam("max_price")
	if maxPrice != "" {
		maxPriceInt, _ := strconv.Atoi(maxPrice)
		filters.MaxPrice = maxPriceInt
	}
	keywords := c.QueryParam("keywords")
	if keywords != "" {
		filters.Keywords = []string{keywords}
	}
	offset := c.QueryParam("offset")
	if offset != "" {
		offsetInt, _ := strconv.Atoi(offset)
		filters.Offset = offsetInt
	}
	return filters
}
