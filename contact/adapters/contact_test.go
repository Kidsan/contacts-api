package adapters

import (
	"context"
	"testing"

	uuid "github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
)

func TestContactHandler_Get(t *testing.T) {
	type fields struct {
		data []contactsapi.Contact
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic test",
			fields: fields{data: []contactsapi.Contact{
				{
					Name: "a",
					ID:   uuid.New(),
				},
				{
					Name: "b",
					ID:   uuid.New(),
				},
				{
					Name: "c",
					ID:   uuid.New(),
				},
			}},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContactHandler{
				data: tt.fields.data,
			}
			ctx := context.Background()
			got, _ := c.Get(ctx)

			for i, v := range got {
				if v.Name != tt.want[i] {
					t.Errorf("contactsapi.ContactHandler.Get() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
