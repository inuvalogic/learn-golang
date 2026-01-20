package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
    "strconv"
)
type Produk struct {
    ID    int    `json:"id"`
    Nama  string `json:"nama"`
    Harga int    `json:"harga"`
    Stok  int    `json:"stok"`
}
var produk = []Produk{
    {ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
    {ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
}

func getProdukById(w http.ResponseWriter, r *http.Request, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid product ID", http.StatusBadRequest)
        return
    }

    // loop & ambil data produk sesuai id
    for _, produk := range produk {
        if produk.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(produk)
            return
        }
    }

    http.Error(w, "product not found", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid product ID", http.StatusBadRequest)
        return
    }

    // get data dari request
    var updateProduk Produk
    err = json.NewDecoder(r.Body).Decode(&updateProduk)
    if err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    // loop & ubah data produk sesuai id
    for _, p := range produk {
        if p.ID == id {
            p.Nama = updateProduk.Nama
            p.Harga = updateProduk.Harga
            p.Stok = updateProduk.Stok

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(p)
            return
        }
    }

    http.Error(w, "product not found", http.StatusNotFound)
}

func deleteProduk(w http.ResponseWriter, r *http.Request, idStr string) {
    // cek apakah id valid
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid product ID", http.StatusBadRequest)
        return
    }

    // loop, get produk sesuai id, dan return 
    for i, p := range produk {
        if p.ID == id {
            produk = append(produk[:i], produk[i+1:]...)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(produk)
            return
        }
    }

    http.Error(w, "product not found", http.StatusNotFound)
}

func main() {

    // GET localhost:8080/api/produk/{id}
    // PUT localhost:8080/api/produk/{id}
    // DELETE localhost:8080/api/produk/{id}
    http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
        idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

        switch r.Method {
            case "GET":
                getProdukById(w, r, idStr)
                return
            case "PUT":
                updateProduk(w, r, idStr)
                return
            case "DELETE":
                deleteProduk(w, r, idStr)
                return
            default:
                http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
                return
        }
    })

    // GET localhost:8080/api/produk
    // POST localhost:8080/api/produk
    http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(produk)
        } else if r.Method == "POST" {
            // baca dari request
            var produkBaru Produk
            err := json.NewDecoder(r.Body).Decode(&produkBaru)
            if err != nil {
                http.Error(w, "invalid request", http.StatusBadRequest)
                return
            }
            // masukkan data ke dalam variable produk
            produkBaru.ID = len(produk) + 1
            produk = append(produk, produkBaru)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusCreated) // 201
            json.NewEncoder(w).Encode(produkBaru)
        }
    })
    // localhost:8080/health
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status":  "OK",
            "message": "API running",
        })
    })
    fmt.Println("server running di localhost:8000")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        fmt.Println("gagal running server")
    }
}