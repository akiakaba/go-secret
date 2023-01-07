package secret

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	sec := New("aaa")
	tests := []struct {
		format string
		want   string
	}{
		{
			format: "%v",
			want:   "******",
		},
		{
			format: "%+v",
			want:   "******",
		},
		{
			format: "%#v",
			want:   "******",
		},
		{
			format: "%s",
			want:   "******",
		},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			got := fmt.Sprintf(tt.format, sec)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fmt.Sprintf(%v, New(\"aaa\")) = %v, want %v", tt.format, got, tt.want)
			}
		})
	}
}

func TestSecret_String(t *testing.T) {
	type fields struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "secret is empty",
			fields: fields{v: ""},
			want:   "*secret is empty*",
		},
		{
			name:   "secret is non-empty",
			fields: fields{v: "aaa"},
			want:   "******",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Secret{
				v: tt.fields.v,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret_Value(t *testing.T) {
	type fields struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "secret is empty",
			fields: fields{v: ""},
			want:   "",
		},
		{
			name:   "secret is non-empty",
			fields: fields{v: "aaa"},
			want:   "aaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Secret{
				v: tt.fields.v,
			}
			if got := s.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
