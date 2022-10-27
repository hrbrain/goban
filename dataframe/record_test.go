package dataframe

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hrbrain/goban/element"
	"github.com/hrbrain/goban/series"
)

func TestRecord_GetSeriesNames(t *testing.T) {
	type fields struct {
		r Record
	}
	tests := []struct {
		name string
		fields
		want []series.Name
	}{
		{
			name: "pass",
			fields: fields{
				r: Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  123,
						IsNull: false,
					},
				},
			},
			want: []series.Name{"series_1", "series_2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.GetSeriesNames()
			opt := cmpopts.SortSlices(func(i, j series.Name) bool {
				return i < j
			})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf(diff)
			}
		})
	}

}

func TestRecord_AddField(t *testing.T) {
	type fields struct {
		r Record
	}
	type args struct {
		name    series.Name
		element element.Element
	}
	tests := []struct {
		name string
		fields
		args
		want      Record
		wantError bool
	}{
		{
			name: "pass",
			fields: fields{
				r: Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  123,
						IsNull: false,
					},
				},
			},
			args: args{
				name: "series_3",
				element: element.StringElement{
					Value:  "いろは",
					IsNull: false,
				},
			},
			want: Record{
				"series_1": element.StringElement{
					Value:  "abc",
					IsNull: false,
				},
				"series_2": element.NumericElement{
					Value:  123,
					IsNull: false,
				},
				"series_3": element.StringElement{
					Value:  "いろは",
					IsNull: false,
				},
			},
			wantError: false,
		},
		{
			name: "fail (duplicated series name)",
			fields: fields{
				r: Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  123,
						IsNull: false,
					},
				},
			},
			args: args{
				name: "series_1",
				element: element.StringElement{
					Value:  "いろは",
					IsNull: false,
				},
			},
			want:      Record{},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.AddField(tt.args.name, tt.args.element)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantError); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})
	}
}

func TestRecord_GetElement(t *testing.T) {
	type fields struct {
		Record
	}
	type args struct {
		name series.Name
	}
	tests := []struct {
		name string
		fields
		args
		want      element.Element
		wantError bool
	}{
		{
			name: "pass",
			fields: fields{
				Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  123,
						IsNull: false,
					},
				},
			},
			args: args{
				name: "series_2",
			},
			want: element.NumericElement{
				Value:  123,
				IsNull: false,
			},
			wantError: false,
		},
		{
			name: "fail(series does not exist)",
			fields: fields{
				Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  123,
						IsNull: false,
					},
				},
			},
			args: args{
				name: "series_3",
			},
			want:      nil,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.Record.GetElement(tt.args.name)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantError); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})
	}
}

func TestRecord_Has(t *testing.T) {
	type fields struct {
		Record
	}
	type args struct {
		name series.Name
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
				Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  123,
						IsNull: false,
					},
				},
			},
			args: args{
				name: "series_2",
			},
			want: true,
		},
		{
			name: "fail(series does not exist)",
			fields: fields{
				Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  123,
						IsNull: false,
					},
				},
			},
			args: args{
				name: "series_3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.Record.Has(tt.args.name)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestRecord_HasNAElement(t *testing.T) {
	type fields struct {
		Record
	}
	type args struct {
		name series.Name
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
				Record{
					"series_1": element.StringElement{
						Value:  "abc",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  math.NaN(),
						IsNull: true,
					},
				},
			},
			args: args{
				name: "series_2",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.Record.HasNAElement()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
