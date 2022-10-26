package series

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hrbrain/goban/element"
)

func TestNewSeries(t *testing.T) {
	type args struct {
		name             Name
		elements         element.Elements
		aggregatedMethod AggregationMethod
	}
	tests := []struct {
		name string
		args
		want    Series
		wantErr bool
	}{
		{
			name: "pass (numeric elements)",
			args: args{
				name: "series_1",
				elements: element.NumericElements{
					element.NumericElement{
						Value:  1,
						IsNull: false,
					},
					element.NumericElement{
						Value:  2,
						IsNull: false,
					},
					element.NumericElement{
						Value:  3,
						IsNull: false,
					},
				},
				aggregatedMethod: Count,
			},
			want: Series{
				Name: "series_1",
				Elements: element.NumericElements{
					element.NumericElement{
						Value:  1,
						IsNull: false,
					},
					element.NumericElement{
						Value:  2,
						IsNull: false,
					},
					element.NumericElement{
						Value:  3,
						IsNull: false,
					},
				},
				AggregatedMethod: Count,
			},
			wantErr: false,
		},
		{
			name: "pass (string elements)",
			args: args{
				name: "series_1",
				elements: element.StringElements{
					element.StringElement{
						Value:  "a",
						IsNull: false,
					},
					element.StringElement{
						Value:  "b",
						IsNull: false,
					},
				},
				aggregatedMethod: None,
			},
			want: Series{
				Name: "series_1",
				Elements: element.StringElements{
					element.StringElement{
						Value:  "a",
						IsNull: false,
					},
					element.StringElement{
						Value:  "b",
						IsNull: false,
					},
				},
				AggregatedMethod: None,
			},
			wantErr: false,
		},
		{
			name: "fail (nil elements)",
			args: args{
				name:             "series_1",
				elements:         nil,
				aggregatedMethod: None,
			},
			want:    Series{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSeries(tt.args.name, tt.args.elements, tt.args.aggregatedMethod)
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

func TestSeries_Mean(t *testing.T) {
	type field struct {
		Series
	}
	tests := []struct {
		name string
		field
		want      float64
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  6,
							IsNull: false,
						},
					},
				},
			},
			want:      4,
			wantError: false,
		},
		{
			name: "fail (string type)",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
						element.StringElement{
							Value:  "b",
							IsNull: false,
						},
						element.StringElement{
							Value:  "c",
							IsNull: false,
						},
					},
				},
			},
			want:      0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.Mean()
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

func TestSeries_Sum(t *testing.T) {
	type field struct {
		Series
	}
	tests := []struct {
		name string
		field
		want      float64
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  6,
							IsNull: false,
						},
					},
				},
			},
			want:      12,
			wantError: false,
		},
		{
			name: "fail (string type)",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
						element.StringElement{
							Value:  "b",
							IsNull: false,
						},
						element.StringElement{
							Value:  "c",
							IsNull: false,
						},
					},
				},
			},
			want:      0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.Sum()
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

func TestSeries_Split(t *testing.T) {
	type field struct {
		Series
	}
	type args struct {
		separator string
		limit     int
	}
	tests := []struct {
		name string
		field
		args
		want      Series
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a,b,c",
							IsNull: false,
						},
					},
				},
			},
			args: args{
				separator: ",",
				limit:     10,
			},
			want: Series{
				Name: "test",
				Elements: element.StringListElements{
					element.StringListElement{"a", "b", "c"},
				},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.Split(tt.args.separator, tt.args.limit)
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

func TestSeries_Slice(t *testing.T) {
	type field struct {
		Series
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		field
		args
		want      Series
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringListElements{
						element.StringListElement{
							"a", "b", "c",
						},
					},
				},
			},
			args: args{
				start: 0,
				end:   2,
			},
			want: Series{
				Name: "test",
				Elements: element.StringListElements{
					element.StringListElement{"a", "b"},
				},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.Slice(tt.args.start, tt.args.end)
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

func TestSeries_GetGroupedElement(t *testing.T) {
	type field struct {
		Series
	}
	tests := []struct {
		name string
		field
		want      element.Element
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
					},
				},
			},
			want: element.NumericElement{
				Value:  3,
				IsNull: false,
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.GetGroupedElement()
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

func TestSeries_Aggregate(t *testing.T) {
	type field struct {
		Series
	}
	type args struct {
		method AggregationMethod
	}
	tests := []struct {
		name string
		field
		args
		want      element.Element
		wantError bool
	}{
		{
			name: "pass (count)",
			field: field{
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
					},
				},
			},
			args: args{
				method: Count,
			},
			want: element.NumericElement{
				Value:  2,
				IsNull: false,
			},
			wantError: false,
		},
		{
			name: "pass (mean)",
			field: field{
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
					},
				},
			},
			args: args{
				method: Mean,
			},
			want: element.NumericElement{
				Value:  3,
				IsNull: false,
			},
			wantError: false,
		},
		{
			name: "pass (sum)",
			field: field{
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
					},
				},
			},
			args: args{
				method: Sum,
			},
			want: element.NumericElement{
				Value:  6,
				IsNull: false,
			},
			wantError: false,
		},
		{
			name: "pass (none)",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
					},
				},
			},
			args: args{
				method: None,
			},
			want: element.StringElement{
				Value:  "a",
				IsNull: false,
			},
			wantError: false,
		},
		{
			name: "fail (not grouped)",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
						element.StringElement{
							Value:  "b",
							IsNull: false,
						},
					},
				},
			},
			args: args{
				method: None,
			},
			want:      nil,
			wantError: true,
		},
		{
			name: "fail (invalid aggregation method)",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
						element.StringElement{
							Value:  "b",
							IsNull: false,
						},
					},
				},
			},
			args: args{
				method: Mean,
			},
			want:      nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.Aggregate(tt.args.method)
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

func TestSeries_Floats(t *testing.T) {
	type field struct {
		Series
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
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  6,
							IsNull: false,
						},
					},
				},
			},
			want:      []float64{3, 3, 6},
			wantError: false,
		},
		{
			name: "fail (string type)",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
						element.StringElement{
							Value:  "b",
							IsNull: false,
						},
						element.StringElement{
							Value:  "c",
							IsNull: false,
						},
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

func TestSeries_AddElement(t *testing.T) {
	type field struct {
		series Series
	}
	type args struct {
		element element.Element
	}
	tests := []struct {
		name string
		field
		args
		want      Series
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				series: Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  1,
							IsNull: false,
						},
					},
				},
			},
			args: args{
				element: element.NumericElement{
					Value:  2,
					IsNull: false,
				},
			},
			want: Series{
				Name: "test",
				Elements: element.NumericElements{
					element.NumericElement{
						Value:  1,
						IsNull: false,
					},
					element.NumericElement{
						Value:  2,
						IsNull: false,
					},
				},
			},
			wantError: false,
		},
		{
			name: "fail (invalid element)",
			field: field{
				series: Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  1,
							IsNull: false,
						},
					},
				},
			},
			args: args{
				element: element.StringElement{
					Value:  "2",
					IsNull: false,
				},
			},
			want:      Series{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.series.AddElement(tt.args.element)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantError); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
		})
	}
}

func TestSeries_UpdateElements(t *testing.T) {
	type field struct {
		series Series
	}
	type args struct {
		element.Elements
	}
	tests := []struct {
		name string
		field
		args
		want      Series
		wantError bool
	}{
		{
			name: "pass",
			field: field{
				series: Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  1,
							IsNull: false,
						},
						element.NumericElement{
							Value:  2,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
					},
				},
			},
			args: args{
				element.NumericElements{
					element.NumericElement{
						Value:  2,
						IsNull: false,
					},
					element.NumericElement{
						Value:  3,
						IsNull: false,
					},
					element.NumericElement{
						Value:  4,
						IsNull: false,
					},
				},
			},
			want: Series{
				Name: "test",
				Elements: element.NumericElements{
					element.NumericElement{
						Value:  2,
						IsNull: false,
					},
					element.NumericElement{
						Value:  3,
						IsNull: false,
					},
					element.NumericElement{
						Value:  4,
						IsNull: false,
					},
				},
			},
			wantError: false,
		},
		{
			name: "fail (elements are nil)",
			field: field{
				series: Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  1,
							IsNull: false,
						},
						element.NumericElement{
							Value:  2,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
					},
				},
			},
			args: args{
				nil,
			},
			want:      Series{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.series.UpdateElements(tt.args.Elements)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantError); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
		})
	}
}

func TestSeries_Len(t *testing.T) {
	type fields struct {
		Series
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "pass",
			fields: fields{
				Series{
					Name: "series_1",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  1,
							IsNull: false,
						},
						element.NumericElement{
							Value:  2,
							IsNull: false,
						},
					},
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Series.Len()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestSeries_Delete(t *testing.T) {
	type field struct {
		Series
	}
	tests := []struct {
		name string
		field
		want      Series
		wantError bool
	}{
		{
			name: "pass(numeric)",
			field: field{
				Series{
					Name: "test",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  6,
							IsNull: false,
						},
					},
				},
			},
			want: Series{
				Name:     "test",
				Elements: element.NumericElements{},
			},
			wantError: false,
		},
		{
			name: "pass (string)",
			field: field{
				Series{
					Name: "test",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "a",
							IsNull: false,
						},
						element.StringElement{
							Value:  "b",
							IsNull: false,
						},
					},
				},
			},
			want: Series{
				Name:     "test",
				Elements: element.StringElements{},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Delete()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
