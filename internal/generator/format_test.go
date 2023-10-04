package generator

import (
	"errors"
	"testing"
)

func TestFormatFromStr(t *testing.T) {
	tests := []struct {
		name         string
		expression   string
		expectErr    error
		expectFormat format
	}{
		{
			name:       "empty format",
			expression: "",
			expectFormat: format{
				typeStr: "string",
				verb:    "%s",
			},
		}, {
			name:       "[]byte w/o verb",
			expression: "[]byte",
			expectFormat: format{
				typeStr: "[]byte",
				verb:    "%s",
			},
		}, {
			name:       "[]byte w/ verb",
			expression: "[]byte(%q)",
			expectFormat: format{
				typeStr: "[]byte",
				verb:    "%q",
			},
		}, {
			name:       "float32 w/o verb",
			expression: "float32",
			expectFormat: format{
				typeStr: "float32",
				verb:    "%f",
			},
		}, {
			name:       "string type w/ verb",
			expression: "float32(%1.2f)",
			expectFormat: format{
				typeStr: "float32",
				verb:    "%1.2f",
			},
		}, {
			name:       "int32 type w/o verb",
			expression: "int32",
			expectFormat: format{
				typeStr: "int32",
				verb:    "%d",
			},
		}, {
			name:       "int32 type w/ verb",
			expression: "int(%b)",
			expectFormat: format{
				typeStr: "int",
				verb:    "%b",
			},
		}, {
			name:       "custom type w/o verb",
			expression: "potato",
			expectFormat: format{
				typeStr: "potato",
				verb:    "%v",
			},
		}, {
			name:       "custom type w/ verb",
			expression: "potato(%+v)",
			expectFormat: format{
				typeStr: "potato",
				verb:    "%+v",
			},
		}, {
			name:       "invalid format type",
			expression: "]potato(%+v)",
			expectErr:  errInvalidFormat,
		}, {
			name:       "invalid format verb",
			expression: "potato(%(v)",
			expectErr:  errInvalidFormat,
		}, {
			name:       "invalid verb too many letters",
			expression: "potato(%ev)",
			expectErr:  errInvalidVerb,
		}, {
			name:       "invalid verb non existant verb",
			expression: "potato(%a)",
			expectErr:  errInvalidVerb,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFormat, gotErr := formatFromStr(tt.expression)
			if !errors.Is(gotErr, tt.expectErr) {
				t.Fatalf("expected err %v, got %v", tt.expectErr, gotErr)
			}
			if gotFormat != tt.expectFormat {
				t.Fatalf("expected format %+v, got %+v", tt.expectFormat, gotFormat)
			}
		})
	}

}
