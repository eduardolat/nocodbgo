package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eduardolat/nocodbgo"
)

// User represents a user in the table
type User struct {
	ID             int    `json:"Id"`
	SingleLineText string `json:"SingleLineText"`
	Email          string `json:"Email"`
	Number         int    `json:"Number"`
}

func main() {

	// Create a new client for NocoDB API v2 using the chain pattern
	client, err := nocodbgo.NewClient().
		WithBaseURL("https://example.com").
		WithAPIToken("your-api-token").
		WithHTTPTimeout(30 * time.Second).
		Build()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Get a table by ID
	table := client.Table("your-table-id")

	// Create a user
	user := map[string]interface{}{
		"SingleLineText": "John Doe",
		"Email":          "john@example.com",
		"Number":         30,
	}

	userID, err := table.Create(context.Background(), user)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	// Read the user using the chain pattern
	readResponse, err := table.Read(context.Background(), userID).
		Fields("SingleLineText", "Email", "Number").
		Execute()
	if err != nil {
		log.Fatalf("Error reading user: %v", err)
	}

	fmt.Printf("Read user: %v\n", readResponse.Data)

	// Decode the user into a struct
	var userStruct User
	err = readResponse.Decode(&userStruct)
	if err != nil {
		log.Fatalf("Error decoding user: %v", err)
	}

	fmt.Printf("Decoded user: %+v\n", userStruct)

	// Update the user
	updateUser := map[string]interface{}{
		"SingleLineText": "John Smith",
	}

	err = table.Update(context.Background(), userID, updateUser)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	// List users with filters using the chain pattern
	result, err := table.List(context.Background()).
		GreaterThan("Number", 18).
		SortAsc("SingleLineText").
		Limit(10).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users: %v", err)
	}

	fmt.Printf("Users: %v\n", result.List)
	fmt.Printf("Page Info: Total Rows: %d, Page: %d, Page Size: %d, Is First Page: %t, Is Last Page: %t\n",
		result.PageInfo.TotalRows, result.PageInfo.Page, result.PageInfo.PageSize, result.PageInfo.IsFirstPage, result.PageInfo.IsLastPage)

	// Decode the list of users into a struct
	var users []User
	err = result.Decode(&users)
	if err != nil {
		log.Fatalf("Error decoding users: %v", err)
	}

	fmt.Printf("Decoded users: %+v\n", users)

	// Count users using the chain pattern
	count, err := table.Count(context.Background()).
		GreaterThan("Number", 18).
		Execute()
	if err != nil {
		log.Fatalf("Error counting users: %v", err)
	}

	fmt.Printf("User count: %d\n", count)

	// Delete the user
	err = table.Delete(context.Background(), userID)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}

	fmt.Println("User deleted")

	// Bulk create users
	bulkUsers := []map[string]interface{}{
		{
			"SingleLineText": "Jane Doe",
			"Email":          "jane@example.com",
			"Number":         25,
		},
		{
			"SingleLineText": "Bob Smith",
			"Email":          "bob@example.com",
			"Number":         40,
		},
	}

	createdUsers, err := table.BulkCreate(context.Background(), bulkUsers)
	if err != nil {
		log.Fatalf("Error bulk creating users: %v", err)
	}

	fmt.Printf("Created users: %v\n", createdUsers)

	// Bulk update users
	bulkUpdateUsers := []map[string]interface{}{
		{
			"Id":             createdUsers[0],
			"SingleLineText": "Jane Smith",
		},
		{
			"Id":             createdUsers[1],
			"SingleLineText": "Robert Smith",
		},
	}

	err = table.BulkUpdate(context.Background(), bulkUpdateUsers)
	if err != nil {
		log.Fatalf("Error bulk updating users: %v", err)
	}

	fmt.Println("Users updated")

	// Bulk delete users
	err = table.BulkDelete(context.Background(), createdUsers)
	if err != nil {
		log.Fatalf("Error bulk deleting users: %v", err)
	}

	fmt.Println("Users deleted")

	// Complex filtering using the chain pattern
	complexResult, err := table.List(context.Background()).
		EqualTo("SingleLineText", "John Smith").
		GreaterThan("Number", 18).
		LessThan("Number", 30).
		SortAsc("SingleLineText").
		Limit(10).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users with complex filter: %v", err)
	}

	fmt.Printf("Users with complex filter: %v\n", complexResult.List)
	fmt.Printf("Page Info: Total Rows: %d, Page: %d, Page Size: %d, Is First Page: %t, Is Last Page: %t\n",
		complexResult.PageInfo.TotalRows, complexResult.PageInfo.Page, complexResult.PageInfo.PageSize, complexResult.PageInfo.IsFirstPage, complexResult.PageInfo.IsLastPage)
}
