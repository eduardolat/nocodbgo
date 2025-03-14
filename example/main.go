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
		Create()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Get a table by ID
	table := client.Table("your-table-id")

	// Create a user
	user := map[string]any{
		"SingleLineText": "John Doe",
		"Email":          "john@example.com",
		"Number":         30,
	}

	userID, err := table.CreateRecord(user).
		WithContext(context.Background()).
		Execute()
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	// Read the user using the chain pattern
	readResponse, err := table.ReadRecord(userID).
		WithContext(context.Background()).
		Fields("SingleLineText", "Email", "Number").
		Execute()
	if err != nil {
		log.Fatalf("Error reading user: %v", err)
	}

	fmt.Printf("Read user: %v\n", readResponse.Data)

	// Decode the user into a struct
	var userStruct User
	err = readResponse.DecodeInto(&userStruct)
	if err != nil {
		log.Fatalf("Error decoding user: %v", err)
	}

	fmt.Printf("Decoded user: %+v\n", userStruct)

	// Update the user
	updateUser := map[string]any{
		"SingleLineText": "John Smith",
	}

	err = table.UpdateRecord(userID, updateUser).
		WithContext(context.Background()).
		Execute()
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	// List users with filters using the chain pattern
	listResponse, err := table.ListRecords().
		WithContext(context.Background()).
		GreaterThan("Number", "18").
		SortAsc("SingleLineText").
		Limit(10).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users: %v", err)
	}

	fmt.Printf("Users: %v\n", listResponse.List)
	fmt.Printf(
		"Page Info: Total Rows: %d, Page: %d, Page Size: %d, Is First Page: %t, Is Last Page: %t\n",
		listResponse.PageInfo.TotalRows, listResponse.PageInfo.Page, listResponse.PageInfo.PageSize, listResponse.PageInfo.IsFirstPage, listResponse.PageInfo.IsLastPage,
	)

	// Decode the list of users into a struct
	var users []User
	err = listResponse.DecodeInto(&users)
	if err != nil {
		log.Fatalf("Error decoding users: %v", err)
	}

	fmt.Printf("Decoded users: %+v\n", users)

	// Count users using the chain pattern
	count, err := table.Count().
		WithContext(context.Background()).
		GreaterThan("Number", "18").
		Execute()
	if err != nil {
		log.Fatalf("Error counting users: %v", err)
	}

	fmt.Printf("User count: %d\n", count)

	// Delete the user
	err = table.DeleteRecord(userID).
		WithContext(context.Background()).
		Execute()
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}

	fmt.Println("User deleted")

	// Bulk create users
	bulkUsers := []map[string]any{
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

	createdUsers, err := table.BulkCreateRecords(bulkUsers).
		WithContext(context.Background()).
		Execute()
	if err != nil {
		log.Fatalf("Error bulk creating users: %v", err)
	}

	fmt.Printf("Created users: %v\n", createdUsers)

	// Bulk update users
	bulkUpdateUsers := []map[string]any{
		{
			"Id":             createdUsers[0],
			"SingleLineText": "Jane Smith",
		},
		{
			"Id":             createdUsers[1],
			"SingleLineText": "Robert Smith",
		},
	}

	err = table.BulkUpdateRecords(bulkUpdateUsers).
		WithContext(context.Background()).
		Execute()
	if err != nil {
		log.Fatalf("Error bulk updating users: %v", err)
	}

	fmt.Println("Users updated")

	// Bulk delete users
	err = table.BulkDeleteRecords(createdUsers).
		WithContext(context.Background()).
		Execute()
	if err != nil {
		log.Fatalf("Error bulk deleting users: %v", err)
	}

	fmt.Println("Users deleted")

	// Complex filtering using the chain pattern
	complexResult, err := table.ListRecords().
		WithContext(context.Background()).
		EqualTo("SingleLineText", "John Smith").
		GreaterThan("Number", "18").
		LessThan("Number", "30").
		SortAsc("SingleLineText").
		Limit(10).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users with complex filter: %v", err)
	}

	fmt.Printf("Users with complex filter: %v\n", complexResult.List)
	fmt.Printf(
		"Page Info: Total Rows: %d, Page: %d, Page Size: %d, Is First Page: %t, Is Last Page: %t\n",
		complexResult.PageInfo.TotalRows, complexResult.PageInfo.Page, complexResult.PageInfo.PageSize, complexResult.PageInfo.IsFirstPage, complexResult.PageInfo.IsLastPage,
	)

	// Using the Where method for custom filter expressions
	customFilterResult, err := table.ListRecords().
		WithContext(context.Background()).
		Where("(Number,gt,20)~or(Email,like,%@example.com)").
		SortDesc("Number").
		Limit(5).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users with custom filter: %v", err)
	}

	fmt.Printf("Users with custom filter: %v\n", customFilterResult.List)

	// Using Where with Count
	customFilterCount, err := table.Count().
		WithContext(context.Background()).
		Where("(SingleLineText,like,%Smith)~and(Number,gt,20)").
		Execute()
	if err != nil {
		log.Fatalf("Error counting users with custom filter: %v", err)
	}

	fmt.Printf("User count with custom filter: %d\n", customFilterCount)
}
