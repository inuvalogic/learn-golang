package main

import (
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/libs"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Port    string `mapstructure:"PORT"`
    DBConn string `mapstructure:"DB_CONN"`
}

/*
 * Main Function
 */
func main() {

    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

    if _, err := os.Stat(".env"); err == nil {
        viper.SetConfigFile(".env")
        _ = viper.ReadInConfig()
    }

    config := Config{
        Port: viper.GetString("PORT"),
        DBConn: viper.GetString("DB_CONN"),
    }

    // Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	defer db.Close()

    // Category
    categoryRepo := repositories.NewCategoryRepository(db)
    categoryService := services.NewCategoryService(categoryRepo)
    categoryHandler := handlers.NewCategoryHandler(categoryService)

    http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
    http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

    // Product
    productRepo := repositories.NewProductRepository(db)
    productService := services.NewProductService(productRepo)
    productHandler := handlers.NewProductHandler(productService)

    http.HandleFunc("/api/products", productHandler.HandleProducts)
    http.HandleFunc("/api/products/", productHandler.HandleProductByID)

    // Transaction
    transactionRepo := repositories.NewTransactionRepository(db)
    transactionService := services.NewTransactionService(transactionRepo)
    transactionHandler := handlers.NewTransactionHandler(transactionService)

    http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST

    // GET localhost:8080/health
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        libs.HandleResponse(http.StatusOK, w, nil, "API running")
    })

    // Running Server di port 8080
    addr := "0.0.0.0:" + config.Port
    fmt.Println("server running di", addr)
    err = http.ListenAndServe(addr, nil)
    if err != nil {
        fmt.Println("gagal running server")
    }
}