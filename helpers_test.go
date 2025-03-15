package nocodbgo

import (
	"testing"
)

func TestDecodeInto(t *testing.T) {
	type User struct {
		ID    int    `json:"Id"`
		Name  string `json:"Name"`
		Email string `json:"Email"`
	}

	tests := []struct {
		name    string
		data    any
		dest    any
		wantErr bool
	}{
		{
			name: "decode map into struct",
			data: map[string]any{
				"Id":    1,
				"Name":  "John",
				"Email": "john@example.com",
			},
			dest:    &User{},
			wantErr: false,
		},
		{
			name: "decode slice of maps into slice of structs",
			data: []map[string]any{
				{
					"Id":    1,
					"Name":  "John",
					"Email": "john@example.com",
				},
				{
					"Id":    2,
					"Name":  "Jane",
					"Email": "jane@example.com",
				},
			},
			dest:    &[]User{},
			wantErr: false,
		},
		{
			name:    "invalid destination",
			data:    map[string]any{"foo": "bar"},
			dest:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := decodeInto(tt.data, tt.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeInto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListResponseDecodeInto(t *testing.T) {
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
	err := response.DecodeInto(&users)
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

func TestReadResponseDecodeInto(t *testing.T) {
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
	err := response.DecodeInto(&user)
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

func TestStructToMap(t *testing.T) {
	type User struct {
		ID    int    `json:"Id"`
		Name  string `json:"Name"`
		Email string `json:"Email"`
	}

	tests := []struct {
		name    string
		data    any
		want    map[string]any
		wantErr bool
	}{
		{
			name: "convert struct to map",
			data: User{
				ID:    1,
				Name:  "John",
				Email: "john@example.com",
			},
			want: map[string]any{
				"Id":    float64(1), // JSON numbers are float64
				"Name":  "John",
				"Email": "john@example.com",
			},
			wantErr: false,
		},
		{
			name:    "invalid input",
			data:    make(chan int),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := structToMap(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("structToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				for k, v := range tt.want {
					if got[k] != v {
						t.Errorf("structToMap() = %v, want %v", got[k], v)
					}
				}
			}
		})
	}
}

func TestStructsToMaps(t *testing.T) {
	type User struct {
		ID    int    `json:"Id"`
		Name  string `json:"Name"`
		Email string `json:"Email"`
	}

	tests := []struct {
		name    string
		data    any
		want    []map[string]any
		wantErr bool
	}{
		{
			name: "convert slice of structs to maps",
			data: []User{
				{
					ID:    1,
					Name:  "John",
					Email: "john@example.com",
				},
				{
					ID:    2,
					Name:  "Jane",
					Email: "jane@example.com",
				},
			},
			want: []map[string]any{
				{
					"Id":    float64(1),
					"Name":  "John",
					"Email": "john@example.com",
				},
				{
					"Id":    float64(2),
					"Name":  "Jane",
					"Email": "jane@example.com",
				},
			},
			wantErr: false,
		},
		{
			name:    "invalid input",
			data:    make(chan int),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := structsToMaps(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("structsToMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("structsToMaps() len = %v, want %v", len(got), len(tt.want))
					return
				}
				for i := range tt.want {
					for k, v := range tt.want[i] {
						if got[i][k] != v {
							t.Errorf("structsToMaps()[%d][%s] = %v, want %v", i, k, got[i][k], v)
						}
					}
				}
			}
		})
	}
}
