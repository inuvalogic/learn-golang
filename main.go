package main
import (
    "encoding/json"
    "fmt"
    "net/http"
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
func main() {
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