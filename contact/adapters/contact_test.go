package adapters

import (
	"testing"

	uuid "github.com/google/uuid"
)

func TestContactHandler_Get(t *testing.T) {
	type fields struct {
		data []Contact
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic test",
			fields: fields{data: []Contact{
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
			got := c.Get()

			for i, v := range got {
				if v.Name != tt.want[i] {
					t.Errorf("ContactHandler.Get() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
