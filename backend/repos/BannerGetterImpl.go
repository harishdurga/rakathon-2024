package repos

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BannerGetterImpl struct {
	db *sqlx.DB
}

type bannerSqlRow struct {
	ID         int    `db:"id"`
	CategoryID int    `db:"category_id"`
	Url        string `db:"url"`
	Keywords   string `db:"keywords"`
}

func NewBannerGetterImpl(db *sqlx.DB) *BannerGetterImpl {
	return &BannerGetterImpl{db: db}
}

func (b *BannerGetterImpl) GetBanners(filters BannerFilters) ([]Banner, error) {
	var bannerRows []bannerSqlRow
	var banners []Banner
	// Generate SQL query using all the filters
	query := "SELECT * FROM banners WHERE 1=1"
	if len(filters.CategoryIDs) > 0 {
		query += " AND category_id IN ("
		for i, categoryID := range filters.CategoryIDs {
			if i == 0 {
				query += fmt.Sprintf("%d", categoryID)
			} else {
				query += "," + fmt.Sprintf("%d", categoryID)
			}
		}
		query += ")"
	}
	// Limit
	if filters.Limit != 0 {
		query += " LIMIT " + string(filters.Limit)
	}
	err := b.db.Select(&bannerRows, query)
	if err != nil {
		return nil, err
	}
	for _, bannerRow := range bannerRows {
		banner := Banner{
			ID:         bannerRow.ID,
			CategoryID: bannerRow.CategoryID,
			Url:        bannerRow.Url,
			Keywords:   keyWordsStringToSlice(bannerRow.Keywords),
		}
		banners = append(banners, banner)
	}
	return banners, nil
}
