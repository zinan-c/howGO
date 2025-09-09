// cmd/api/main.go
package main

/*
 * not sure the import items could change to from local rather than github.com/zinan-c/howGO
 * because the go.mod module name is changed to howgo
 */
import (
    "log"
    "net/http"
    
    "github.com/zinan-c/howGO/internal/handler"
    "github.com/zinan-c/howGO/internal/database"
)

func main() {
    db := database.New()
    handler := handler.New(db)
    
    log.Println("Server starting on :8080")
    http.ListenAndServe(":8080", handler)
}