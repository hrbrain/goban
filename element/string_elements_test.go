package element

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringElements_Floats(t *testing.T) {
	type field struct {
		StringElements
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
				StringElements{
					StringElement{
						Value:  "a",
						IsNull: false,
					},
					StringElement{
						Value:  "b",
						IsNull: false,
					},
					StringElement{
						Value:  "c",
						IsNull: false,
					},
				},
			},
			want:      nil,
			wantError: true,
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

func TestStringElements_GetElement(t *testing.T) {
	type fields struct {
		stringElements StringElements
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
				stringElements: StringElements{
					StringElement{
						Value:  "a",
						IsNull: false,
					},
					StringElement{
						Value:  "b",
						IsNull: false,
					},
					StringElement{
						Value:  "c",
						IsNull: false,
					},
				},
			},
			args: args{
				index: 1,
			},
			want: StringElement{
				Value:  "b",
				IsNull: false,
			},
			wantErr: false,
		},
		{
			name: "fail (index is equal to length)",
			fields: fields{
				stringElements: StringElements{
					StringElement{
						Value:  "a",
						IsNull: false,
					},
					StringElement{
						Value:  "b",
						IsNull: false,
					},
					StringElement{
						Value:  "c",
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
				stringElements: StringElements{
					StringElement{
						Value:  "a",
						IsNull: false,
					},
					StringElement{
						Value:  "b",
						IsNull: false,
					},
					StringElement{
						Value:  "c",
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
			got, err := tt.fields.stringElements.GetElement(tt.args.index)
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

func TestStringElements_GetGroupedElement(t *testing.T) {
	type fields struct {
		stringElements StringElements
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
				stringElements: StringElements{
					StringElement{
						Value:  "a",
						IsNull: false,
					},
					StringElement{
						Value:  "a",
						IsNull: false,
					},
				},
			},
			want: StringElement{
				Value:  "a",
				IsNull: false,
			},
			wantErr: false,
		},
		{
			name: "fail (not grouped)",
			fields: fields{
				stringElements: StringElements{
					StringElement{
						Value:  "a",
						IsNull: false,
					},
					StringElement{
						Value:  "b",
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
			got, err := tt.fields.stringElements.GetGroupedElement()
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
