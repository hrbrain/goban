package element

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumericElement_String(t *testing.T) {
	type fields struct {
		numericElement NumericElement
	}
	tests := []struct {
		name string
		fields
		want    string
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				NumericElement{
					Value:  99,
					IsNull: false,
				},
			},
			want:    "99.000000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.numericElement.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})
	}
}

func TestNumericElement_Equal(t *testing.T) {
	type fields struct {
		numericElement NumericElement
	}
	type args struct {
		element Element
	}
	tests := []struct {
		name string
		fields
		args
		want bool
	}{
		{
			name: "pass",
			fields: fields{
				numericElement: NumericElement{
					Value:  99,
					IsNull: false,
				},
			},
			args: args{
				element: NumericElement{
					Value:  99,
					IsNull: false,
				},
			},
			want: true,
		},
		{
			name: "fail (type mismatch)",
			fields: fields{
				numericElement: NumericElement{
					Value:  99,
					IsNull: false,
				},
			},
			args: args{
				element: StringElement{
					Value:  "99",
					IsNull: false,
				},
			},
			want: false,
		},
		{
			name: "fail (both elements are null)",
			fields: fields{
				numericElement: NumericElement{
					Value:  99,
					IsNull: true,
				},
			},
			args: args{
				element: NumericElement{
					Value:  0,
					IsNull: true,
				},
			},
			want: true,
		},
		{
			name: "fail (different value)",
			fields: fields{
				numericElement: NumericElement{
					Value:  99,
					IsNull: false,
				},
			},
			args: args{
				element: NumericElement{
					Value:  88,
					IsNull: false,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.numericElement.Equal(tt.args.element)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
