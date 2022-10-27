package element

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringListElements_GetElement(t *testing.T) {
	type fields struct {
		StringListElements StringListElements
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
				StringListElements: StringListElements{
					StringListElement{
						"a", "b",
					},
					StringListElement{
						"c", "d",
					},
					StringListElement{
						"e", "f",
					},
				},
			},
			args: args{
				index: 1,
			},
			want: StringListElement{
				"c", "d",
			},
			wantErr: false,
		},
		{
			name: "fail (index is equal to length)",
			fields: fields{
				StringListElements: StringListElements{
					StringListElement{
						"a", "b",
					},
					StringListElement{
						"c", "d",
					},
					StringListElement{
						"e", "f",
					},
				},
			},
			args: args{
				index: 3,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.StringListElements.GetElement(tt.args.index)
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

func TestStringListElements_Floats(t *testing.T) {
	type field struct {
		StringListElements
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
				StringListElements{
					StringListElement{
						"a", "b",
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

func TestStringListElements_GetGroupedElement(t *testing.T) {
	type fields struct {
		StringListElements StringListElements
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
				StringListElements: StringListElements{
					StringListElement{
						"a", "b",
					},
					StringListElement{
						"a", "b",
					},
				},
			},
			want: StringListElement{
				"a", "b",
			},
			wantErr: false,
		},
		{
			name: "fail (not grouped)",
			fields: fields{
				StringListElements: StringListElements{
					StringListElement{
						"a", "b",
					},
					StringListElement{
						"a", "c",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.StringListElements.GetGroupedElement()
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
