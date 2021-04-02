package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAuthJWT_Create(t *testing.T) {
	type fields struct {
		Secret      string
		ExpDuration time.Duration
		Issuer      string
	}
	type args struct {
		userID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "测试正常产生 token",
			fields: fields{
				Secret:      "123456",
				ExpDuration: time.Minute * 15,
				Issuer:      "test",
			},
			args: args{
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := JWT{
				Secret:      tt.fields.Secret,
				ExpDuration: tt.fields.ExpDuration,
				Issuer:      tt.fields.Issuer,
			}
			got, err := A.Create(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("token len must than 0")
				return
			}
			// 进行校验
			claims, err := A.Parse(got)
			assert.Nil(t, err)
			if !claims.Authorized {
				t.Errorf("claims.Authorized must be true")
				return
			}
			if claims.UserID != tt.args.userID {
				t.Errorf("claims.UserID need %d, got %d", tt.args.userID, claims.UserID)
				return
			}
		})
	}
}
