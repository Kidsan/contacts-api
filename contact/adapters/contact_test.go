package adapters

import (
	"reflect"
	"testing"
)

func TestContactHandler_Get(t *testing.T) {
	type fields struct {
		data []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "basic test",
			fields: fields{data: []string{"a", "b", "c"}},
			want:   []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContactHandler{
				data: tt.fields.data,
			}
			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContactHandler.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
