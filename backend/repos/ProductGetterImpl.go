package repos

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ProductGetterImpl struct {
	db *sqlx.DB
}

type productSqlRow struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	Price          int    `db:"price"`
	CategoryID     int    `db:"category_id"`
	Thumbnail      string `db:"thumbnail"`
	CustomerRating int    `db:"customer_rating"`
	Attributes     string `db:"attributes"`
	Keywords       string `db:"keywords"`
}

func NewProductGetterImpl(db *sqlx.DB) *ProductGetterImpl {
	return &ProductGetterImpl{db: db}
}

func (p *ProductGetterImpl) GetProducts(filters ProductFilters) ([]Product, error) {
	var productRows []productSqlRow
	var products []Product
	// Generate SQL query using all the filters
	query := "SELECT * FROM products WHERE 1=1"
	if filters.CategoryID != 0 {
		query += fmt.Sprintf(" AND category_id = %d", filters.CategoryID)
	}
	// Min and max price conditions
	if filters.MinPrice != 0 {
		query += fmt.Sprintf(" AND price >= %d", filters.MinPrice)
	}
	if filters.MaxPrice != 0 {
		query += fmt.Sprintf(" AND price <= %d", filters.MaxPrice)
	}
	// Keywords
	if len(filters.Keywords) > 0 {
		query += " AND ("
		for i, keyword := range filters.Keywords {
			if i == 0 {
				query += fmt.Sprintf("keywords LIKE '%%%s%%'", keyword)
			} else {
				query += fmt.Sprintf(" OR keywords LIKE '%%%s%%'", keyword)
			}
		}
		query += ")"
	}
	// Sort
	if filters.Sort != "" {
		query += " ORDER BY " + filters.Sort
	}
	// Limit and offset
	if filters.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", filters.Limit)
	}
	if filters.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", filters.Offset)
	}
	err := p.db.Select(&productRows, query)
	if err != nil {
		return nil, err
	}
	for _, row := range productRows {
		products = append(products, Product{
			ID:             row.ID,
			Name:           row.Name,
			Description:    row.Description,
			Price:          row.Price,
			CategoryID:     row.CategoryID,
			Thumbnail:      row.Thumbnail,
			CustomerRating: row.CustomerRating,
			Keywords:       keyWordsStringToSlice(row.Keywords),
		})
	}
	return products, nil
}

func keyWordsStringToSlice(keywords string) []string {
	// Convert the comma separated string to a slice
	return strings.Split(keywords, ",")
}
