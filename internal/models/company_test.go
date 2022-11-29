package models

import (
	"database/sql"
	"reflect"
	"strings"
	"testing"
)

func TestCompany_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       *Company
		wantErr bool
	}{
		{
			name:    "name too short",
			c:       &Company{Name: "A"},
			wantErr: true,
		},
		{
			name:    "name too long",
			c:       &Company{Name: "ABCDEFGHIJKLMNOPR"},
			wantErr: true,
		},
		{
			name:    "description too long",
			c:       &Company{Name: strings.Repeat("D", 4000)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nullStringToJson(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		wantNil bool
		wantStr string
	}{
		{
			name:    "invalid string",
			ns:      sql.NullString{Valid: false},
			wantNil: true,
		},
		{
			name:    "valid string",
			ns:      sql.NullString{Valid: true, String: "test"},
			wantStr: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nullStringToJson(tt.ns)
			if tt.wantNil && got != nil {
				t.Errorf("nullStringToJson() = %v, want %v", got, nil)
			}
			if tt.wantStr != "" && *got != tt.wantStr {
				t.Errorf("nullStringToJson() = %v, want %v", *got, tt.wantStr)
			}
		})
	}
}

func Test_jsonStringToNullString(t *testing.T) {
	str := "test"
	tests := []struct {
		name string
		s    *string
		want sql.NullString
	}{
		{
			name: "not nil string",
			s:    &str,
			want: sql.NullString{Valid: true, String: str},
		},
		{
			name: "nil string",
			s:    nil,
			want: sql.NullString{Valid: false, String: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsonStringToNullString(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonStringToNullString() = %v, want %v", got, tt.want)
			}
		})
	}
}
