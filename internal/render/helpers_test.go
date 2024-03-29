package render

import (
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func Test_defaultValueForType(t *testing.T) {
	type args struct {
		f *models.Field
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing some default value texts",
			args: args{f: &models.Field{
				Name: "myvar",
				Type: &models.Expression{
					Value:      "value",
					Underlying: "string",
					IsStar:     false,
					IsVariadic: false,
					IsWriter:   false,
				},
			}},
			want: `"hello"`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := defaultValueForType(tt.args.f); got != tt.want {
				t.Errorf("defaultValueForType() = %v, want %v", got, tt.want)
			}
		})
	}
}
