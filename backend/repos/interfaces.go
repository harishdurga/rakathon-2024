package repos

type ProductFilters struct {
	CategoryID int
	Limit      int
	Sort       string
	MinPrice   int
	MaxPrice   int
	Keywords   []string
	Offset     int
}

type BannerFilters struct {
	CategoryIDs []int
	Limit       int
}

type Product struct {
	ID             int                    `json:"id"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	Price          int                    `json:"price"`
	CategoryID     int                    `json:"category_id"`
	Thumbnail      string                 `json:"thumbnail"`
	CustomerRating int                    `json:"customer_rating"`
	Attributes     map[string]interface{} `json:"attributes"`
	Keywords       []string               `json:"keywords"`
}

type Banner struct {
	ID         int      `json:"id"`
	CategoryID int      `json:"category_id"`
	Url        string   `json:"url"`
	Keywords   []string `json:"keywords"`
}

type ProductGetter interface {
	GetProducts(filters ProductFilters) ([]Product, error)
}

type BannerGetter interface {
	GetBanners(filters BannerFilters) ([]Banner, error)
}
