package nocodbgo

import (
	"testing"
)

func TestDecode(t *testing.T) {
	// Test decoding a map into a struct
	type User struct {
		ID    int    `json:"Id"`
		Name  string `json:"Name"`
		Email string `json:"Email"`
		Age   int    `json:"Age"`
	}

	data := map[string]any{
		"Id":    1,
		"Name":  "John Doe",
		"Email": "john@example.com",
		"Age":   30,
	}

	var user User
	err := decode(data, &user)
	if err != nil {
		t.Errorf("Decode() error = %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Decode() ID = %v, want %v", user.ID, 1)
	}
	if user.Name != "John Doe" {
		t.Errorf("Decode() Name = %v, want %v", user.Name, "John Doe")
	}
	if user.Email != "john@example.com" {
		t.Errorf("Decode() Email = %v, want %v", user.Email, "john@example.com")
	}
	if user.Age != 30 {
		t.Errorf("Decode() Age = %v, want %v", user.Age, 30)
	}

	// Test error case: invalid destination
	err = decode(data, data)
	if err == nil {
		t.Errorf("Decode() error = nil, want error")
	}
}

func TestListResponseDecode(t *testing.T) {
	// Test decoding a slice into a slice of structs
	type User struct {
		ID    int    `json:"Id"`
		Name  string `json:"Name"`
		Email string `json:"Email"`
		Age   int    `json:"Age"`
	}

	dataSlice := []map[string]any{
		{
			"Id":    1,
			"Name":  "John Doe",
			"Email": "john@example.com",
			"Age":   30,
		},
		{
			"Id":    2,
			"Name":  "Jane Doe",
			"Email": "jane@example.com",
			"Age":   25,
		},
	}

	response := ListResponse{
		List: dataSlice,
	}

	var users []User
	err := response.Decode(&users)
	if err != nil {
		t.Errorf("Decode() error = %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Decode() len = %v, want %v", len(users), 2)
	}

	if users[0].ID != 1 {
		t.Errorf("Decode() ID = %v, want %v", users[0].ID, 1)
	}
	if users[0].Name != "John Doe" {
		t.Errorf("Decode() Name = %v, want %v", users[0].Name, "John Doe")
	}
	if users[0].Email != "john@example.com" {
		t.Errorf("Decode() Email = %v, want %v", users[0].Email, "john@example.com")
	}
	if users[0].Age != 30 {
		t.Errorf("Decode() Age = %v, want %v", users[0].Age, 30)
	}

	if users[1].ID != 2 {
		t.Errorf("Decode() ID = %v, want %v", users[1].ID, 2)
	}
	if users[1].Name != "Jane Doe" {
		t.Errorf("Decode() Name = %v, want %v", users[1].Name, "Jane Doe")
	}
	if users[1].Email != "jane@example.com" {
		t.Errorf("Decode() Email = %v, want %v", users[1].Email, "jane@example.com")
	}
	if users[1].Age != 25 {
		t.Errorf("Decode() Age = %v, want %v", users[1].Age, 25)
	}
}

func TestReadResponseDecode(t *testing.T) {
	// Test decoding a map into a struct
	type User struct {
		ID    int    `json:"Id"`
		Name  string `json:"Name"`
		Email string `json:"Email"`
		Age   int    `json:"Age"`
	}

	data := map[string]any{
		"Id":    1,
		"Name":  "John Doe",
		"Email": "john@example.com",
		"Age":   30,
	}

	response := ReadResponse{
		Data: data,
	}

	var user User
	err := response.Decode(&user)
	if err != nil {
		t.Errorf("Decode() error = %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Decode() ID = %v, want %v", user.ID, 1)
	}
	if user.Name != "John Doe" {
		t.Errorf("Decode() Name = %v, want %v", user.Name, "John Doe")
	}
	if user.Email != "john@example.com" {
		t.Errorf("Decode() Email = %v, want %v", user.Email, "john@example.com")
	}
	if user.Age != 30 {
		t.Errorf("Decode() Age = %v, want %v", user.Age, 30)
	}
}
