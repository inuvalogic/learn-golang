package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
    "strconv"
)

type Category struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Product struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Price    int    `json:"price"`
    Stock    int    `json:"stock"`
    Category int    `json:"category"` // FIXME: belum tau kalau bikin relation gimana
}

var categories = []Category{
    {ID: 1, Name: "Main Course", Description: "Burgers, steaks, pasta, seafood, pizza"},
    {ID: 2, Name: "Beverages", Description: "Soft drinks, coffee, tea"},
    {ID: 3, Name: "Sides", Description: "Fries, vegetables, rice, bread"},
    {ID: 4, Name: "Desserts", Description: "Cakes, ice cream, pastries"},
    {ID: 5, Name: "Appetizers", Description: "Wings, nachos, soups, salads"},
}

var products = []Product{
    {ID: 1, Name: "Ribs Eye Steak", Price: 135000, Stock: 100, Category: 1},
    {ID: 2, Name: "Ice Lemon", Price: 25000, Stock: 50, Category: 2},
    {ID: 3, Name: "Spicy Lobster", Price: 180800, Stock: 10, Category: 1},
}

/*
 * handle API Response & Error
 * agar response API lebih rapi dan mudah dikembangkan
 *
 */
func handleResponse(httpCode int, w http.ResponseWriter, data interface{}, message ...string) {
    msg := ""
    if len(message) > 0 {
        msg = message[0]
    }

    var output = map[string]interface{}{
        "status": httpCode,
        "message": msg,
        "data": data,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(httpCode)
    json.NewEncoder(w).Encode(output)
}

func handleError(httpCode int, w http.ResponseWriter, message string) {
    var output = map[string]interface{}{
        "status": httpCode,
        "message": message,
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(httpCode)
    json.NewEncoder(w).Encode(output)
}

/**
 * CRUD Category
 * 
 * createCategory = membuat data kategori baru
 * getCategories = mengambil semua data kategori
 * getCategory = mengambil data sebuah kategori sesuai id
 * updateCategory = mengubah data kategori sesuai id
 * deleteCategory = menghapus data kategori sesuai id
 * 
 */
func createCategory(w http.ResponseWriter, r *http.Request) {
    // get data dari request
    var newCategory Category
    err := json.NewDecoder(r.Body).Decode(&newCategory)

    // jika request tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid request")
        return
    }

    // masukkan data ke dalam variable categories
    newCategory.ID = len(categories) + 1
    categories = append(categories, newCategory)

    handleResponse(http.StatusCreated, w, newCategory)
}

func getCategories(w http.ResponseWriter) {
    // jika ingin custom response, bisa diubah disini
    handleResponse(http.StatusOK, w, categories)
}

func getCategory(w http.ResponseWriter, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid category ID")
        return
    }

    // loop & ambil data kategori sesuai id
    for _, category := range categories {
        if category.ID == id {
            handleResponse(http.StatusOK, w, category)
            return
        }
    }

    // jika id kategori tidak ditemukan muncul error not found
    handleError(http.StatusNotFound, w, "category not found")
}

func updateCategory(w http.ResponseWriter, r *http.Request, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)

    // jika id tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid category ID")
        return
    }

    // get data dari request
    var updateCategory Category
    err = json.NewDecoder(r.Body).Decode(&updateCategory)

    // jika request tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid request")
        return
    }

    // loop & ubah data kategori sesuai id
    for i := range categories {
        if categories[i].ID == id {
            updateCategory.ID = id
            categories[i] = updateCategory

            handleResponse(http.StatusOK, w, updateCategory)
            return
        }
    }

    // jika id kategori tidak ditemukan muncul error not found
    handleError(http.StatusNotFound, w, "category not found")
}

func deleteCategory(w http.ResponseWriter, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)
    
    // jika id tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid request")
        return
    }

    // loop, get kategori sesuai id, dan return 
    for i, c := range categories {
        if c.ID == id {
            categories = append(categories[:i], categories[i+1:]...)
            handleResponse(http.StatusOK, w, nil, "category deleted")
            return
        }
    }

    // jika id kategori tidak ditemukan muncul error not found
    handleError(http.StatusNotFound, w, "category not found")
}

/**
 * CRUD Product
 * 
 * createProduct = membuat data produk baru
 * getProducts = mengambil semua data produk
 * getProduct = mengambil data sebuah produk sesuai id
 * updateProduct = mengubah data produk sesuai id
 * deleteProduct = menghapus data produk sesuai id
 * 
 */
func createProduct(w http.ResponseWriter, r *http.Request) {
    // get data dari request
    var newProduct Product
    err := json.NewDecoder(r.Body).Decode(&newProduct)

    // jika request tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid request")
        return
    }

    // masukkan data ke dalam variable categories
    newProduct.ID = len(products) + 1
    products = append(products, newProduct)

    handleResponse(http.StatusCreated, w, newProduct)
}

func getProducts(w http.ResponseWriter) {
    // jika ingin custom response, bisa diubah disini
    handleResponse(http.StatusOK, w, products)
}

func getProduct(w http.ResponseWriter, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid product ID")
        return
    }

    // loop & ambil data produk sesuai id
    for _, product := range products {
        if product.ID == id {
            handleResponse(http.StatusOK, w, product)
            return
        }
    }

    // jika id produk tidak ditemukan muncul error not found
    handleError(http.StatusNotFound, w, "product not found")
}

func updateProduct(w http.ResponseWriter, r *http.Request, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)

    // jika id tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid product ID")
        return
    }

    // get data dari request
    var updateProduct Product
    err = json.NewDecoder(r.Body).Decode(&updateProduct)

    // jika request tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid request")
        return
    }

    // loop & ubah data produk sesuai id
    for i := range products {
        if products[i].ID == id {
            updateProduct.ID = id
            products[i] = updateProduct

            handleResponse(http.StatusOK, w, updateProduct)
            return
        }
    }

    // jika id kategori tidak ditemukan muncul error not found
    handleError(http.StatusNotFound, w, "product not found")
}

func deleteProduct(w http.ResponseWriter, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)
    
    // jika id tidak valid, muncul error
    if err != nil {
        handleError(http.StatusBadRequest, w, "invalid request")
        return
    }

    // loop, get produk sesuai id, dan return 
    for i, p := range products {
        if p.ID == id {
            products = append(products[:i], products[i+1:]...)
            handleResponse(http.StatusOK, w, nil, "product deleted")
            return
        }
    }

    // jika id produk tidak ditemukan muncul error not found
    handleError(http.StatusNotFound, w, "product not found")
}


/*
 * Main Function
 */
func main() {

    // GET localhost:8080/api/categories/{id}
    // PUT localhost:8080/api/categories/{id}
    // DELETE localhost:8080/api/categories/{id}
    http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
        idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

        switch r.Method {
            case "GET":
                getCategory(w, idStr)
                return
            case "PUT":
                updateCategory(w, r, idStr)
                return
            case "DELETE":
                deleteCategory(w, idStr)
                return
            default:
                handleError(http.StatusMethodNotAllowed, w, "method not allowed")
                return
        }
    })

    // GET localhost:8080/api/categories
    // POST localhost:8080/api/categories
    http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case "GET":
                getCategories(w)
                return
            case "POST":
                createCategory(w, r)
                return
            default:
                handleError(http.StatusMethodNotAllowed, w, "method not allowed")
                return
        }
    })

    // GET localhost:8080/api/products/{id}
    // PUT localhost:8080/api/products/{id}
    // DELETE localhost:8080/api/products/{id}
    http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
        idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

        switch r.Method {
            case "GET":
                getProduct(w, idStr)
                return
            case "PUT":
                updateProduct(w, r, idStr)
                return
            case "DELETE":
                deleteProduct(w, idStr)
                return
            default:
                handleError(http.StatusMethodNotAllowed, w, "method not allowed")
                return
        }
    })

    // GET localhost:8080/api/products
    // POST localhost:8080/api/products
    http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case "GET":
                getProducts(w)
                return
            case "POST":
                createProduct(w, r)
                return
            default:
                handleError(http.StatusMethodNotAllowed, w, "method not allowed")
                return
        }
    })

    // GET localhost:8080/health
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        handleResponse(http.StatusOK, w, nil, "API running")
    })

    // Running Server di port 8080
    fmt.Println("server running di localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("gagal running server")
    }
}