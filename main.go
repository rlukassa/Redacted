// File ini adalah entry point backend Little Alchemy 2.
// Berisi setup server Gin, middleware CORS, dan routing endpoint pencarian resep.

package main

import (
	"fmt"              // Untuk format string
	"main/controllers" // Import controller pencarian resep
	"net/http"         // Untuk kebutuhan HTTP
	"os"               // Untuk membaca environment variable

	"github.com/gin-gonic/gin" // Framework web Gin
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Izinkan semua origin (untuk pengembangan)
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Metode yang diizinkan
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Header yang diizinkan
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Izinkan kredensial

        if c.Request.Method == "OPTIONS" { // Tangani preflight request CORS
            c.AbortWithStatus(http.StatusOK)
            return
        }

        c.Next() // Lanjutkan ke handler berikutnya
    }
} // ye intinya ini cuek aja lah 

func main() {
    r := gin.Default() // Inisialisasi Gin
    r.Use(CORSMiddleware()) // Pasang middleware CORS
    r.POST("/api/search", controllers.SearchRecipe) // Endpoint pencarian resep
    // Tambahkan endpoint GET / untuk healthcheck Railway
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }
    addr := fmt.Sprintf(":%s", port)
    r.Run(addr) // Jalankan server di port yang sesuai
}