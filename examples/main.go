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
	ID    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Age   int    `json:"Age"`
}

func main() {
	// Create a new client for NocoDB v2 API using the chain pattern
	client, err := nocodbgo.NewClient().
		WithBaseURL("https://example.com").
		WithAPIToken("your-api-token").
		WithHTTPTimeout(30 * time.Second).
		Create()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Reference a table instance by ID
	table := client.Table("your-table-id")

	// ==========================================
	// Basic CRUD Operations
	// ==========================================

	// Create a record, this can be a map or a struct with json tags matching the table fields
	user := map[string]any{
		"Name":  "John Doe",
		"Email": "john@example.com",
		"Age":   30,
	}

	userID, err := table.CreateRecord(user).
		WithContext(context.Background()). // Optional, if not provided a context.Background() will be used
		Execute()
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Printf("User created with ID: %d\n", userID)

	// Read a record
	readResponse, err := table.ReadRecord(userID).
		ReturnFields("Name", "Email", "Age").
		Execute()
	if err != nil {
		log.Fatalf("Error reading user: %v", err)
	}
	fmt.Printf("User read: %v\n", readResponse.Data)

	// Decode into a struct
	var userStruct User
	err = readResponse.DecodeInto(&userStruct)
	if err != nil {
		log.Fatalf("Error decoding user: %v", err)
	}
	fmt.Printf("Decoded user: %+v\n", userStruct)

	// Update a record
	updateUser := map[string]any{
		"Id":   userID, // ID must be included
		"Name": "John Smith",
	}

	err = table.UpdateRecord(updateUser).Execute()
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Println("User updated")

	// ==========================================
	// Listing Records with Filters and Sorting
	// ==========================================

	// List records with basic filters
	listResponse, err := table.ListRecords().
		WithContext(context.Background()).
		WhereIsGreaterThan("Age", "18"). // Same as Where("(Age,gt,18)")
		SortAscBy("Name").
		Limit(10).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users: %v", err)
	}

	fmt.Printf("Users: %v\n", listResponse.List)
	fmt.Printf(
		"Pagination info: Total: %d, Page: %d, Size: %d, IsFirstPage: %t, IsLastPage: %t\n",
		listResponse.PageInfo.TotalRows, listResponse.PageInfo.Page, listResponse.PageInfo.PageSize,
		listResponse.PageInfo.IsFirstPage, listResponse.PageInfo.IsLastPage,
	)

	// Decode the list into a struct
	var users []User
	err = listResponse.DecodeInto(&users)
	if err != nil {
		log.Fatalf("Error decoding users: %v", err)
	}
	fmt.Printf("Decoded users: %+v\n", users)

	// Count records
	count, err := table.CountRecords().
		WhereIsGreaterThan("Age", "18"). // Same as Where("(Age,gt,18)")
		Execute()
	if err != nil {
		log.Fatalf("Error counting users: %v", err)
	}
	fmt.Printf("User count: %d\n", count)

	// Delete a record
	err = table.DeleteRecord(userID).Execute()
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	fmt.Println("User deleted")

	// ==========================================
	// Operations with Multiple Records
	// ==========================================

	// Create multiple records
	multipleUsers := []map[string]any{
		{
			"Name":  "Jane Doe",
			"Email": "jane@example.com",
			"Age":   25,
		},
		{
			"Name":  "Bob Smith",
			"Email": "bob@example.com",
			"Age":   40,
		},
	}

	createdUserIDs, err := table.CreateRecords(multipleUsers).Execute()
	if err != nil {
		log.Fatalf("Error creating multiple users: %v", err)
	}
	fmt.Printf("Created user IDs: %v\n", createdUserIDs)

	// Update multiple records
	updateUsers := []map[string]any{
		{
			"Id":   createdUserIDs[0],
			"Name": "Jane Smith",
		},
		{
			"Id":   createdUserIDs[1],
			"Name": "Robert Smith",
		},
	}

	err = table.UpdateRecords(updateUsers).Execute()
	if err != nil {
		log.Fatalf("Error updating multiple users: %v", err)
	}
	fmt.Println("Multiple users updated")

	// Delete multiple records
	err = table.DeleteRecords(createdUserIDs).Execute()
	if err != nil {
		log.Fatalf("Error deleting multiple users: %v", err)
	}
	fmt.Println("Multiple users deleted")

	// ==========================================
	// Complex Filters
	// ==========================================

	// Filters with specific operators
	complexResult, err := table.ListRecords().
		WhereIsEqualTo("Name", "John Smith").
		WhereIsGreaterThan("Age", "18").
		WhereIsLessThan("Age", "30").
		SortAscBy("Name").
		Limit(10).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users with complex filter: %v", err)
	}
	fmt.Printf("Users with complex filter: %v\n", complexResult.List)

	// Custom filters with Where
	customFilterResult, err := table.ListRecords().
		Where("(Age,gt,20)~or(Email,like,%@example.com)").
		SortDescBy("Age").
		Limit(5).
		Execute()
	if err != nil {
		log.Fatalf("Error listing users with custom filter: %v", err)
	}
	fmt.Printf("Users with custom filter: %v\n", customFilterResult.List)

	// Count with custom filters
	customFilterCount, err := table.CountRecords().
		Where("(Name,like,%Smith)~and(Age,gt,20)").
		Execute()
	if err != nil {
		log.Fatalf("Error counting users with custom filter: %v", err)
	}
	fmt.Printf("User count with custom filter: %d\n", customFilterCount)

	// ==========================================
	// Working with Linked Records
	// ==========================================

	// Create a link between records
	err = table.CreateLink("link-field-id", userID, 123).Execute()
	if err != nil {
		log.Fatalf("Error creating link: %v", err)
	}
	fmt.Println("Link created")

	// Create multiple links
	err = table.CreateLinks("link-field-id", userID, []int{124, 125, 126}).Execute()
	if err != nil {
		log.Fatalf("Error creating multiple links: %v", err)
	}
	fmt.Println("Multiple links created")

	// List linked records
	linkedRecords, err := table.ListLinks("link-field-id", userID).
		ReturnFields("Name", "Email").
		SortAscBy("Name").
		Limit(10).
		Where("(Age,gt,18)").
		Execute()
	if err != nil {
		log.Fatalf("Error listing linked records: %v", err)
	}

	// Decode linked records
	var linkedUsers []User
	err = linkedRecords.DecodeInto(&linkedUsers)
	if err != nil {
		log.Fatalf("Error decoding linked records: %v", err)
	}
	fmt.Printf("Linked records: %+v\n", linkedUsers)

	// Delete a link
	err = table.DeleteLink("link-field-id", userID, 123).Execute()
	if err != nil {
		log.Fatalf("Error deleting link: %v", err)
	}
	fmt.Println("Link deleted")

	// Delete multiple links
	err = table.DeleteLinks("link-field-id", userID, []int{124, 125}).Execute()
	if err != nil {
		log.Fatalf("Error deleting multiple links: %v", err)
	}
	fmt.Println("Multiple links deleted")

	// ==========================================
	// Additional Query Options
	// ==========================================

	// Use specific views
	viewResult, err := table.ListRecords().
		WithViewId("view-id").
		Execute()
	if err != nil {
		log.Fatalf("Error listing records with specific view: %v", err)
	}
	fmt.Printf("Records with specific view: %v\n", viewResult.List)

	// Shuffle results randomly
	shuffleResult, err := table.ListRecords().
		Shuffle().
		Limit(5).
		Execute()
	if err != nil {
		log.Fatalf("Error listing random records: %v", err)
	}
	fmt.Printf("Random records: %v\n", shuffleResult.List)

	// Pagination
	pageResult, err := table.ListRecords().
		Page(2, 10). // Page 2, 10 records per page
		Execute()
	if err != nil {
		log.Fatalf("Error listing paginated records: %v", err)
	}
	fmt.Printf("Paginated records: %v\n", pageResult.List)
}
