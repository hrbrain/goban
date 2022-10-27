package dataframe

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hrbrain/goban/element"
	"github.com/hrbrain/goban/series"
)

func TestNewColumns(t *testing.T) {
	type args struct {
		series []series.Series
	}
	tests := []struct {
		name string
		args
		want    Columns
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				[]series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "abc",
								IsNull: false,
							},
							element.StringElement{
								Value:  "def",
								IsNull: false,
							},
						},
					},
					{
						Name: "series_2",
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
			},
			want: Columns{
				{
					Name: "series_1",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "abc",
							IsNull: false,
						},
						element.StringElement{
							Value:  "def",
							IsNull: false,
						},
					},
				},
				{
					Name: "series_2",
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
			wantErr: false,
		},
		{
			name: "pass (empty columns)",
			args: args{
				[]series.Series{},
			},
			want:    Columns{},
			wantErr: false,
		},
		{
			name: "fail (series length mismatch)",
			args: args{
				[]series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "abc",
								IsNull: false,
							},
						},
					},
					{
						Name: "series_2",
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
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewColumns(tt.args.series)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestColumns_RecordCount(t *testing.T) {
	type fields struct {
		columns Columns
	}
	tests := []struct {
		name string
		fields
		want int
	}{
		{
			name: "pass",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  3,
								IsNull: false,
							},
							element.NumericElement{
								Value:  4,
								IsNull: false,
							},
							element.NumericElement{
								Value:  5,
								IsNull: false,
							},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "pass (empty columns)",
			fields: fields{
				columns: Columns{},
			},
			want: 0,
		},
		{
			name: "pass (nil columns)",
			fields: fields{
				columns: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.columns.RecordCount()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestColumns_Append(t *testing.T) {
	type fields struct {
		columns Columns
	}
	type args struct {
		series series.Series
	}
	tests := []struct {
		name string
		fields
		args
		want    Columns
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
				},
			},
			args: args{
				series.Series{
					Name: "series_2",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  4,
							IsNull: false,
						},
						element.NumericElement{
							Value:  5,
							IsNull: false,
						},
					},
				},
			},
			want: Columns{
				series.Series{
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
				},
				series.Series{
					Name: "series_2",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  4,
							IsNull: false,
						},
						element.NumericElement{
							Value:  5,
							IsNull: false,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "pass (same column name but different aggregation method)",
			fields: fields{
				columns: Columns{
					series.Series{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "sales",
								IsNull: false,
							},
							element.StringElement{
								Value:  "dev",
								IsNull: false,
							},
						},
						AggregatedMethod: series.None,
					},
				},
			},
			args: args{
				series.Series{
					Name: "series_1",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  4,
							IsNull: false,
						},
					},
					AggregatedMethod: series.Mean,
				},
			},
			want: Columns{
				series.Series{
					Name: "series_1",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "sales",
							IsNull: false,
						},
						element.StringElement{
							Value:  "dev",
							IsNull: false,
						},
					},
					AggregatedMethod: series.None,
				},
				series.Series{
					Name: "series_1",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  3,
							IsNull: false,
						},
						element.NumericElement{
							Value:  4,
							IsNull: false,
						},
					},
					AggregatedMethod: series.Mean,
				},
			},
			wantErr: false,
		},
		{
			name: "pass (empty columns and series)",
			fields: fields{
				columns: Columns{
					series.Series{
						Name:     "series_1",
						Elements: element.NumericElements{},
					},
				},
			},
			args: args{
				series.Series{
					Name:     "series_2",
					Elements: element.NumericElements{},
				},
			},
			want: Columns{
				series.Series{
					Name:     "series_1",
					Elements: element.NumericElements{},
				},
				series.Series{
					Name:     "series_2",
					Elements: element.NumericElements{},
				},
			},
			wantErr: false,
		},
		{
			name: "pass (allow record count mismatch if columns are empty)",
			fields: fields{
				columns: Columns{},
			},
			args: args{
				series.Series{
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
				},
			},
			want: Columns{
				series.Series{
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
				},
			},
			wantErr: false,
		},
		{
			name: "fail (record count mismatch)",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
				},
			},
			args: args{
				series.Series{
					Name: "series_2",
					Elements: element.NumericElements{
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
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail (duplicated series name)",
			fields: fields{
				columns: Columns{
					series.Series{
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
			},
			args: args{
				series.Series{
					Name: "series_1",
					Elements: element.NumericElements{
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
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.columns.Append(tt.args.series)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
		})
	}
}

func TestColumns_GetSeriesByNameAndMethod(t *testing.T) {
	type fields struct {
		Columns
	}
	type args struct {
		name   series.Name
		method series.AggregationMethod
	}
	tests := []struct {
		name string
		fields
		args
		want    series.Series
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				Columns{
					series.Series{
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
						AggregatedMethod: series.None,
					},
					series.Series{
						Name: "series_2",
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
						AggregatedMethod: series.None,
					},
				},
			},
			args: args{
				name: "series_2",
			},
			want: series.Series{
				Name: "series_2",
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
			wantErr: false,
		},
		{
			name: "pass (same column name but different aggregation method)",
			fields: fields{
				Columns{
					series.Series{
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
						AggregatedMethod: series.Count,
					},
					series.Series{
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
							element.StringElement{
								Value:  "c",
								IsNull: false,
							},
						},
						AggregatedMethod: series.None,
					},
				},
			},
			args: args{
				name:   "series_1",
				method: series.Count,
			},
			want: series.Series{
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
				AggregatedMethod: series.Count,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.FindSeriesBy(tt.args.name, tt.args.method)
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

func TestColumns_Replace(t *testing.T) {
	type fields struct {
		columns Columns
	}
	type args struct {
		index  int
		series series.Series
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Columns
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				index: 0,
				series: series.Series{
					Name: "series_3",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
					},
				},
			},
			want: Columns{
				series.Series{
					Name: "series_3",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
					},
				},
				series.Series{
					Name: "series_2",
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
			wantErr: false,
		},
		{
			name: "fail (index is equal to length)",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				index: 2,
				series: series.Series{
					Name: "series_3",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail (index is less than 0)",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				index: -1,
				series: series.Series{
					Name: "series_3",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.columns.Replace(tt.args.index, tt.args.series)
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

func TestColumns_ReplaceByName(t *testing.T) {
	type fields struct {
		columns Columns
	}
	type args struct {
		series series.Series
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Columns
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				series.Series{
					Name: "series_2",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
					},
				},
			},
			want: Columns{
				series.Series{
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
				},
				series.Series{
					Name: "series_2",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail (series not found)",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				series.Series{
					Name: "series_3",
					Elements: element.StringElements{
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
						element.StringElement{
							Value:  "d",
							IsNull: false,
						},
					},
				},
			},
			want:    Columns{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.columns.ReplaceByName(tt.args.series)
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

func TestColumns_GetSeries(t *testing.T) {
	type fields struct {
		columns Columns
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    series.Series
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				index: 1,
			},
			want: series.Series{
				Name: "series_2",
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
			wantErr: false,
		},
		{
			name: "fail (index is equal to length)",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				index: 2,
			},
			want:    series.Series{},
			wantErr: true,
		},
		{
			name: "fail (index is less than 0)",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			args: args{
				index: -1,
			},
			want:    series.Series{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.columns.GetSeries(tt.args.index)
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

func TestColumns_Delete(t *testing.T) {
	type fields struct {
		columns Columns
	}
	tests := []struct {
		name    string
		fields  fields
		want    Columns
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				columns: Columns{
					series.Series{
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
					},
					series.Series{
						Name: "series_2",
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
			},
			want: Columns{
				series.Series{
					Name:     "series_1",
					Elements: element.NumericElements{},
				},
				series.Series{
					Name:     "series_2",
					Elements: element.NumericElements{},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.columns.Delete()
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
