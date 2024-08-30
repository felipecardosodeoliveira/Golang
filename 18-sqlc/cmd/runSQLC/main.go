package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/felipecardosodeoliveira/Golang/18-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	// Open a connection to the database
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/course")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	// Create a new queries instance
	queries := db.New(dbConn)

	// Print a test message
	fmt.Println("Running CreateCategory...")

	// Execute the CreateCategory query
	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// })
	// if err != nil {
	// 	log.Fatalf("Error creating category: %v", err)
	// }

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		log.Fatalf("Error listing  category: %v", err)
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description)
	}

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "c84a362c-bb65-4cc9-b872-3d4efe6750f3",
	// 	Name:        "Java",
	// 	Description: sql.NullString{String: "Spring", Valid: true},
	// })
	// if err != nil {
	// 	log.Fatalf("Error update category %v", err)
	// }

}
