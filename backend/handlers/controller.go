package handlers

import (
	"adapt-ecomm/repos"
	"log"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type RootHandler struct {
	ProductGetter repos.ProductGetter
	BannerGetter  repos.BannerGetter
}

func NewRootHandler(pg repos.ProductGetter, bg repos.BannerGetter) *RootHandler {
	return &RootHandler{ProductGetter: pg, BannerGetter: bg}
}

func (rh *RootHandler) GetProducts(c echo.Context) error {
	products, err := rh.ProductGetter.GetProducts(getProductFiltersFromContext(c))
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(200, echo.Map{"products": products})
}

func (rh *RootHandler) HomePageLayout(c echo.Context) error {
	return c.String(200, `{
    "layout": [
        {
            "type": "main-banner-slider",
            "filters": {
                "categories": "1,2"
            }
        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 1,
                "limit": 5,
                "sort": "price_asc",
                "min_price": 15000
            },
            "display_title": "Electronics you need"
        },
        {
            "type": "category-products",
            "filters": {
                "category_id": 2,
                "limit": 5,
                "sort": "price_asc",
                "min_price": 15000
            },
            "display_title": "Fashion you love"
        }
    ]
}
`)
}

func (rh *RootHandler) GetBanners(c echo.Context) error {
	banners, err := rh.BannerGetter.GetBanners(getBannerFiltersFromContext(c))
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(200, echo.Map{"banners": banners})
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
		sortParts := strings.Split(sort, "_")
		filters.Sort = strings.Join(sortParts, " ")
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

func getBannerFiltersFromContext(c echo.Context) repos.BannerFilters {
	filters := repos.BannerFilters{}
	//banner filters are part of the query string
	categoryIDs := c.QueryParam("category_ids")
	if categoryIDs != "nil" {
		// convert categoryIds string to slice of ints
		categoryIDsSlice := strings.Split(categoryIDs, ",")
		for _, categoryID := range categoryIDsSlice {
			catIDInt, _ := strconv.Atoi(categoryID)
			filters.CategoryIDs = append(filters.CategoryIDs, catIDInt)
		}
	}
	limit := c.QueryParam("limit")
	if limit != "" {
		limitInt, _ := strconv.Atoi(limit)
		filters.Limit = limitInt
	}
	return filters
}
