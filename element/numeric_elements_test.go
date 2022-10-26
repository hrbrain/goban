package element

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumericElements_Floats(t *testing.T) {
	type field struct {
		NumericElements
	}
	tests := []struct {
		name string
		field
		want      []float64
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				NumericElements{
					NumericElement{
						Value:  3,
						IsNull: false,
					},
					NumericElement{
						Value:  3,
						IsNull: false,
					},
					NumericElement{
						Value:  6,
						IsNull: false,
					},
				},
			},
			want:      []float64{3, 3, 6},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.Floats()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantError); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})
	}
}

func TestNumericElements_GetElement(t *testing.T) {
	type fields struct {
		numericElements NumericElements
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args
		want    Element
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				numericElements: NumericElements{
					NumericElement{
						Value:  11,
						IsNull: false,
					},
					NumericElement{
						Value:  22,
						IsNull: false,
					},
					NumericElement{
						Value:  33,
						IsNull: false,
					},
				},
			},
			args: args{
				index: 2,
			},
			want: NumericElement{
				Value:  33,
				IsNull: false,
			},
			wantErr: false,
		},
		{
			name: "fail (index is equal to length)",
			fields: fields{
				numericElements: NumericElements{
					NumericElement{
						Value:  11,
						IsNull: false,
					},
					NumericElement{
						Value:  22,
						IsNull: false,
					},
					NumericElement{
						Value:  33,
						IsNull: false,
					},
				},
			},
			args: args{
				index: 3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail (index is less than 0)",
			fields: fields{
				numericElements: NumericElements{
					NumericElement{
						Value:  11,
						IsNull: false,
					},
					NumericElement{
						Value:  22,
						IsNull: false,
					},
					NumericElement{
						Value:  33,
						IsNull: false,
					},
				},
			},
			args: args{
				index: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.numericElements.GetElement(tt.args.index)
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

func TestNumericElements_GetGroupedElement(t *testing.T) {
	type fields struct {
		numericElements NumericElements
	}
	tests := []struct {
		name    string
		fields  fields
		want    Element
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				numericElements: NumericElements{
					NumericElement{
						Value:  11,
						IsNull: false,
					},
					NumericElement{
						Value:  11,
						IsNull: false,
					},
				},
			},
			want: NumericElement{
				Value:  11,
				IsNull: false,
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				numericElements: NumericElements{
					NumericElement{
						Value:  11,
						IsNull: false,
					},
					NumericElement{
						Value:  12,
						IsNull: false,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.numericElements.GetGroupedElement()
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
