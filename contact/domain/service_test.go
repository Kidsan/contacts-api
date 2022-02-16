package domain

import (
	"context"
	"reflect"
	"testing"

	contactsapi "github.com/kidsan/contacts-api"
	"github.com/kidsan/contacts-api/mock"
)

func TestContactService_Get(t *testing.T) {
	type fields struct {
		repository ContactRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []contactsapi.Contact
		wantErr bool
	}{
		{
			name: "Two Elements",
			fields: fields{
				repository: &mock.MockContactHandler{
					Data: []contactsapi.Contact{
						{Name: "test-one"},
						{Name: "test-two"},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: []contactsapi.Contact{
				{Name: "test-one"},
				{Name: "test-two"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &ContactService{
				repository: tt.fields.repository,
			}
			got, err := cs.Get(tt.args.ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("ContactService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContactService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactService_Save(t *testing.T) {
	type fields struct {
		repository ContactRepository
	}
	type args struct {
		ctx        context.Context
		newContact contactsapi.Contact
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    contactsapi.Contact
		wantErr bool
	}{
		{
			name: "New Contact",
			fields: fields{
				repository: &mock.MockContactHandler{
					Data: []contactsapi.Contact{
						{Name: "test-one"},
						{Name: "test-two"},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				newContact: contactsapi.Contact{
					Name: "test-three",
				},
			},
			want: contactsapi.Contact{
				Name: "test-three",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &ContactService{
				repository: tt.fields.repository,
			}

			got, err := cs.Save(tt.args.ctx, tt.args.newContact)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContactService.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Name != tt.want.Name {
				t.Errorf("ContactService.Save() = %v, want %v", got, tt.want)
				return
			}
			if got.ID.String() == tt.want.ID.String() {
				t.Errorf("ContactService.Save() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
