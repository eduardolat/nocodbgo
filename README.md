# NocoDB Golang Client

A Go client for the [NocoDB API](https://nocodb.com/). This client provides a
simple and intuitive way to interact with the NocoDB API using a fluent chain
pattern.

## Installation

```bash
go get github.com/eduardolat/nocodbgo
```

## Usage

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
    Build()
if err != nil {
    // Handle error
}
```

### Working with tables

```go
// Get a table
table := client.Table("your-table-id")

// Create a record
user := map[string]any{
    "Name": "John Doe",
    "Email": "john@example.com",
    "Age": 30,
}

userID, err := table.Create(context.Background(), user)
if err != nil {
    // Handle error
}

// Read a record
readResponse, err := table.Read(context.Background(), userID).
    Fields("Name", "Email", "Age").
    Execute()
if err != nil {
    // Handle error
}

// Decode into a struct
type User struct {
    ID    int    `json:"Id"`
    Name  string `json:"Name"`
    Email string `json:"Email"`
    Age   int    `json:"Age"`
}

var userStruct User
err = readResponse.Decode(&userStruct)
if err != nil {
    // Handle error
}

// Update a record
updateUser := map[string]any{
    "Name": "John Smith",
}

err = table.Update(context.Background(), userID, updateUser)
if err != nil {
    // Handle error
}

// Delete a record
err = table.Delete(context.Background(), userID)
if err != nil {
    // Handle error
}
```

### Listing and filtering

```go
// List records with options using the chain pattern
result, err := table.List(context.Background()).
    GreaterThan("Age", 18).
    SortAsc("Name").
    Limit(10).
    Execute()
if err != nil {
    // Handle error
}

// Decode the list into a struct
var users []User
err = result.Decode(&users)
if err != nil {
    // Handle error
}

// Count records
count, err := table.Count(context.Background()).
    GreaterThan("Age", 18).
    Execute()
if err != nil {
    // Handle error
}
```

### Complex filters

```go
// Query with complex filters
result, err := table.List(context.Background()).
    EqualTo("Name", "John Smith").
    GreaterThan("Age", 18).
    LessThan("Age", 30).
    SortAsc("Name").
    Limit(10).
    Execute()
if err != nil {
    // Handle error
}
```

### Bulk operations

```go
// Bulk create records
bulkUsers := []map[string]any{
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

createdUsers, err := table.BulkCreate(context.Background(), bulkUsers)
if err != nil {
    // Handle error
}

// Bulk update records
bulkUpdateUsers := []map[string]any{
    {
        "Id": createdUsers[0],
        "Name": "Jane Smith",
    },
    {
        "Id": createdUsers[1],
        "Name": "Robert Smith",
    },
}

err = table.BulkUpdate(context.Background(), bulkUpdateUsers)
if err != nil {
    // Handle error
}

// Bulk delete records
err = table.BulkDelete(context.Background(), createdUsers)
if err != nil {
    // Handle error
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
for details.
