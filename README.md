# NocoDB Golang Client

<p>
  <a href="https://pkg.go.dev/github.com/eduardolat/nocodbgo">
    <img src="https://pkg.go.dev/badge/github.com/eduardolat/nocodbgo" alt="Go Reference"/>
  </a>
  <a href="https://goreportcard.com/report/eduardolat/nocodbgo">
    <img src="https://goreportcard.com/badge/eduardolat/nocodbgo" alt="Go Report Card"/>
  </a>
  <a href="https://github.com/eduardolat/nocodbgo/releases/latest">
    <img src="https://img.shields.io/github/release/eduardolat/nocodbgo.svg" alt="Release Version"/>
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/eduardolat/nocodbgo.svg" alt="License"/>
  </a>
  <a href="https://github.com/eduardolat/nocodbgo">
    <img src="https://img.shields.io/github/stars/eduardolat/nocodbgo?style=flat&label=github+stars"/>
  </a>
</p>

A Zero-Dependency Go client for the
[NocoDB API](https://docs.nocodb.com/developer-resources/rest-APIs/overview).
This client provides a simple and intuitive way to interact with the NocoDB API
using a fluent chain pattern.

> [!WARNING]
> This client is not yet stable and the API signature may change in the future
> until it reaches version 1.0.0, so be careful when upgrading. However, the API
> signature should not change too much.

## Installation

Go version 1.22 or higher is required.

```bash
go get github.com/eduardolat/nocodbgo
```

## Usage

Most modern code editors will provide autocompletion to guide you through the
available options as you type. This makes using the SDK intuitive and easy to
learn.

You can find examples of how to use the client in the [examples](examples)
directory.

You can also hover over the methods provided by the client to read more about
them.

### Creating a client

```go
import (
    "github.com/eduardolat/nocodbgo"
    "time"
)

// Create a new client using the chain pattern
client, err := nocodbgo.NewClient().
    WithBaseURL("https://example.com").
    WithAPIToken("your-api-token").
    WithHTTPTimeout(30*time.Second).
    Create()
if err != nil {
    // Handle error
}
```

### Basic CRUD Operations

```go
// Get a table
table := client.Table("your-table-id")

// Create a record (can be a map[string]any or a struct with JSON tags)
user := map[string]any{
    "Name": "John Doe",
    "Email": "john@example.com",
    "Age": 30,
}

// Create a record
userID, err := table.CreateRecord(user).Execute()

// Read a record
readResponse, err := table.ReadRecord(userID).
    // Optional, if not provided a context.Background() will be used
    WithContext(context.Background()). 
    ReturnFields("Name", "Email", "Age").
    Execute()

// Decode into a struct
type User struct {
    ID    int    `json:"Id"`
    Name  string `json:"Name"`
    Email string `json:"Email"`
    Age   int    `json:"Age"`
}

var userStruct User
err = readResponse.DecodeInto(&userStruct)

// Update a record
updateUser := map[string]any{
    "Id": userID,  // ID must be included
    "Name": "John Smith",
}

err = table.UpdateRecord(updateUser).Execute()

// Delete a record
err = table.DeleteRecord(userID).Execute()
```

### Listing and Filtering Records

```go
// List records with options using the chain pattern
result, err := table.ListRecords().
    Where("(Age,gt,18)").
    SortAscBy("Name").
    Limit(10).
    Execute()

// Decode the list into a struct
var users []User
err = result.DecodeInto(&users)

// Count records
count, err := table.CountRecords().
    Where("(Age,gt,18)").
    Execute()
```

### Complex Filters

```go
// Query with complex filters using specific methods
result, err := table.ListRecords().
    WhereIsEqualTo("Name", "John Smith").
    WhereIsGreaterThan("Age", "18").
    WhereIsLessThan("Age", "30").
    SortAscBy("Name").
    Limit(10).
    Execute()

// Query with custom filters
result, err := table.ListRecords().
    Where("(Age,gt,20)~or(Email,like,%@example.com)").
    SortDescBy("Age").
    Limit(5).
    Execute()
```

### Operations with Multiple Records

```go
// Create multiple records
users := []map[string]any{
    {
        "Name": "Jane Doe",
        "Email": "jane@example.com",
        "Age": 25,
    },
    {
        "Name": "Bob Smith",
        "Email": "bob@example.com",
        "Age": 40,
    },
}

createdIDs, err := table.CreateRecords(users).Execute()

// Update multiple records
updateUsers := []map[string]any{
    {
        "Id": createdIDs[0],
        "Name": "Jane Smith",
    },
    {
        "Id": createdIDs[1],
        "Name": "Robert Smith",
    },
}

err = table.UpdateRecords(updateUsers).Execute()

// Delete multiple records
err = table.DeleteRecords(createdIDs).Execute()
```

### Working with Linked Records

```go
// List linked records
linkedRecords, err := table.ListLinks("link-field-id", recordID).
    ReturnFields("Name", "Email").
    SortAscBy("Name").
    Limit(10).
    Where("(Age,gt,18)").
    Execute()

// Decode the linked records
var linkedUsers []User
err = linkedRecords.DecodeInto(&linkedUsers)

// Create a link
err = table.CreateLink("link-field-id", recordID, targetID).Execute()

// Create multiple links
err = table.CreateLinks("link-field-id", recordID, []int{1, 2, 3}).Execute()

// Delete a link
err = table.DeleteLink("link-field-id", recordID, targetID).Execute()

// Delete multiple links
err = table.DeleteLinks("link-field-id", recordID, []int{1, 2}).Execute()
```

### Additional Options

```go
// Use a specific view
result, err := table.ListRecords().
    WithViewId("view-id").
    Execute()

// Shuffle results randomly
result, err := table.ListRecords().
    Shuffle().
    Limit(5).
    Execute()

// Pagination
result, err := table.ListRecords().
    Page(2, 10). // Page 2, 10 records per page
    Execute()
```

## Context Control

All operations support the use of `context.Context` for cancellation and timeout
control:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

result, err := table.ListRecords().
    WithContext(ctx).
    Where("(Age,gt,18)").
    Execute()
```

If you don't provide a context, a context.Background() will be used.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
for details.
