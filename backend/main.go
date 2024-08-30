package main

import (
	"adapt-ecomm/handlers"
	"adapt-ecomm/repos"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//init db and mysql
	dsn := "root:root@tcp(127.0.0.1:3306)/adapt_ecomm?parseTime=true"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	// Verify the connection is valid
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully connected to the database!")
	defer db.Close()
	//init product getter
	productGetter := repos.NewProductGetterImpl(db)
	bannerGetter := repos.NewBannerGetterImpl(db)
	//init root handler
	rootHandler := handlers.NewRootHandler(productGetter, bannerGetter)
	e := echo.New()
	e.Use(middleware.CORS())
	e.Static("/static", "assets")
	e.GET("/products", rootHandler.GetProducts)
	e.GET("/banners", rootHandler.GetBanners)
	e.GET("/home-page-layout", rootHandler.HomePageLayout)
	e.Logger.Fatal(e.Start(":1323"))
}
