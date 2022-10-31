package dataframe

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hrbrain/goban/element"
	"github.com/hrbrain/goban/series"
)

func TestNewDataFrame(t *testing.T) {
	type args struct {
		columns Columns
	}
	tests := []struct {
		name string
		args
		want DataFrame
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
			want: DataFrame{
				Columns: []series.Series{
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
				RecordCount: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDataFrame(tt.args.columns)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestDataFrame_GetColumnNames(t *testing.T) {
	type fields struct {
		DataFrame
	}
	tests := []struct {
		name string
		fields
		want []series.Name
	}{
		{
			name: "pass",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
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
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  4,
									IsNull: false,
								},
								element.NumericElement{
									Value:  5,
									IsNull: false,
								},
								element.NumericElement{
									Value:  6,
									IsNull: false,
								},
							},
						},
					},
				},
			},
			want: []series.Name{
				"series_1",
				"series_2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.GetColumnNames()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestDataFrame_UpdateColumn(t *testing.T) {
	type fields struct {
		DataFrame
	}
	type args struct {
		series series.Series
	}
	tests := []struct {
		name string
		fields
		args
		want    DataFrame
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
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
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  4,
									IsNull: false,
								},
								element.NumericElement{
									Value:  5,
									IsNull: false,
								},
								element.NumericElement{
									Value:  6,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				series: series.Series{
					Name: "series_2",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  7,
							IsNull: false,
						},
						element.NumericElement{
							Value:  8,
							IsNull: false,
						},
						element.NumericElement{
							Value:  9,
							IsNull: false,
						},
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
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
					{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  7,
								IsNull: false,
							},
							element.NumericElement{
								Value:  8,
								IsNull: false,
							},
							element.NumericElement{
								Value:  9,
								IsNull: false,
							},
						},
					},
				},
				RecordCount: 3,
			},
			wantErr: false,
		},
		{
			name: "fail (invalid series name)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
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
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  4,
									IsNull: false,
								},
								element.NumericElement{
									Value:  5,
									IsNull: false,
								},
								element.NumericElement{
									Value:  6,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				series: series.Series{
					Name: "series_3",
					Elements: element.NumericElements{
						element.NumericElement{
							Value:  7,
							IsNull: false,
						},
						element.NumericElement{
							Value:  8,
							IsNull: false,
						},
						element.NumericElement{
							Value:  9,
							IsNull: false,
						},
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.UpdateColumn(tt.args.series)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})
	}
}

func TestDataFrame_GroupBy(t *testing.T) {
	type fields struct {
		DataFrame
	}
	type args struct {
		columnName series.Name
	}
	tests := []struct {
		name string
		fields
		args
		want    Groups
		wantErr bool
	}{
		{
			name: "pass (single string elements)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				columnName: "series_1",
			},
			want: []Group{
				{
					SeriesName: "series_1",
					Element:    element.StringElement{Value: "Apple"},
					DataFrame: DataFrame{
						Columns: []series.Series{
							{
								Name: "series_1",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "Apple",
										IsNull: false,
									},
									element.StringElement{
										Value:  "Apple",
										IsNull: false,
									},
								},
							},
						},
						RecordCount: 2,
					},
				},
				{
					SeriesName: "series_1",
					Element:    element.StringElement{Value: "Orange"},
					DataFrame: DataFrame{
						Columns: []series.Series{
							{
								Name: "series_1",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "Orange",
										IsNull: false,
									},
								},
							},
						},
						RecordCount: 1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "pass (single numeric elements)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
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
									Value:  1,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				columnName: "series_1",
			},
			want: []Group{
				{
					SeriesName: "series_1",
					Element:    element.NumericElement{Value: 1},
					DataFrame: DataFrame{
						Columns: []series.Series{
							{
								Name: "series_1",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  1,
										IsNull: false,
									},
									element.NumericElement{
										Value:  1,
										IsNull: false,
									},
								},
							},
						},
						RecordCount: 2,
					},
				},
				{
					SeriesName: "series_1",
					Element:    element.NumericElement{Value: 2},
					DataFrame: DataFrame{
						Columns: []series.Series{
							{
								Name: "series_1",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  2,
										IsNull: false,
									},
								},
							},
						},
						RecordCount: 1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "pass (multiple columns)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
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
									Value:  1,
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "aaa",
									IsNull: false,
								},
								element.StringElement{
									Value:  "bbb",
									IsNull: false,
								},
								element.StringElement{
									Value:  "aaa",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_3",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  99,
									IsNull: false,
								},
								element.NumericElement{
									Value:  88,
									IsNull: false,
								},
								element.NumericElement{
									Value:  77,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				columnName: "series_2",
			},
			want: []Group{
				{
					SeriesName: "series_2",
					Element:    element.StringElement{Value: "aaa"},
					DataFrame: DataFrame{
						Columns: []series.Series{
							{
								Name: "series_1",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  1,
										IsNull: false,
									},
									element.NumericElement{
										Value:  1,
										IsNull: false,
									},
								},
							},
							{
								Name: "series_2",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "aaa",
										IsNull: false,
									},
									element.StringElement{
										Value:  "aaa",
										IsNull: false,
									},
								},
							},
							{
								Name: "series_3",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  99,
										IsNull: false,
									},
									element.NumericElement{
										Value:  77,
										IsNull: false,
									},
								},
							},
						},
						RecordCount: 2,
					},
				},
				{
					SeriesName: "series_2",
					Element:    element.StringElement{Value: "bbb"},
					DataFrame: DataFrame{
						Columns: []series.Series{
							{
								Name: "series_1",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  2,
										IsNull: false,
									},
								},
							},
							{
								Name: "series_2",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "bbb",
										IsNull: false,
									},
								},
							},
							{
								Name: "series_3",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  88,
										IsNull: false,
									},
								},
							},
						},
						RecordCount: 1,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.DataFrame.GroupBy(tt.args.columnName)
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

func TestDataFrame_Records(t *testing.T) {
	type fields struct {
		DataFrame
	}
	tests := []struct {
		name string
		fields
		want    Records
		wantErr bool
	}{
		{
			name: "pass (single series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			want: Records{
				{
					"series_1": element.StringElement{
						Value:  "Apple",
						IsNull: false,
					},
				},
				{
					"series_1": element.StringElement{
						Value:  "Orange",
						IsNull: false,
					},
				},
				{
					"series_1": element.StringElement{
						Value:  "Apple",
						IsNull: false,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "pass (multiple series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
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
								element.NumericElement{
									Value:  3,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			want: Records{
				{
					"series_1": element.StringElement{
						Value:  "Apple",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  1,
						IsNull: false,
					},
				},
				{
					"series_1": element.StringElement{
						Value:  "Orange",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  2,
						IsNull: false,
					},
				},
				{
					"series_1": element.StringElement{
						Value:  "Apple",
						IsNull: false,
					},
					"series_2": element.NumericElement{
						Value:  3,
						IsNull: false,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.DataFrame.Records()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})

	}
}

func TestDataFrame_LoadRecord(t *testing.T) {
	type fields struct {
		DataFrame
	}
	type args struct {
		Record
	}
	tests := []struct {
		name string
		fields
		args
		want    DataFrame
		wantErr bool
	}{
		{
			name: "pass (single series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				Record{
					"series_1": element.StringElement{
						Value:  "Apple",
						IsNull: false,
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Orange",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
						},
					},
				},
				RecordCount: 3,
			},
			wantErr: false,
		},
		{
			name: "pass (multiple series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				Record{
					"series_2": element.NumericElement{
						Value:  25,
						IsNull: false,
					},
					"series_1": element.StringElement{
						Value:  "Apple",
						IsNull: false,
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Orange",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
						},
					},
					{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  20,
								IsNull: false,
							},
							element.NumericElement{
								Value:  30,
								IsNull: false,
							},
							element.NumericElement{
								Value:  25,
								IsNull: false,
							},
						},
					},
				},
				RecordCount: 3,
			},
			wantErr: false,
		},
		{
			name: "fail (aggregated series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
							AggregatedMethod: series.Count,
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				Record{
					"series_2": element.NumericElement{
						Value:  25,
						IsNull: false,
					},
					"series_1": element.StringElement{
						Value:  "Apple",
						IsNull: false,
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
		{
			name: "fail (series type mismatch)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				Record{
					"series_2": element.StringElement{
						Value:  "20",
						IsNull: false,
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
		{
			name: "fail (series not found)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_3",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				Record{
					"series_2": element.NumericElement{
						Value:  40,
						IsNull: false,
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.DataFrame.LoadRecord(tt.args.Record)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})

	}
}

func TestDataFrame_LoadRecords(t *testing.T) {
	type fields struct {
		DataFrame
	}
	type args struct {
		Records
	}
	tests := []struct {
		name string
		fields
		args
		want    DataFrame
		wantErr bool
	}{
		{
			name: "pass (multiple series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				Records{
					{
						"series_2": element.NumericElement{
							Value:  25,
							IsNull: false,
						},
						"series_1": element.StringElement{
							Value:  "Apple",
							IsNull: false,
						},
					},
					{
						"series_2": element.NumericElement{
							Value:  40,
							IsNull: false,
						},
						"series_1": element.StringElement{
							Value:  "Banana",
							IsNull: false,
						},
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Orange",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Banana",
								IsNull: false,
							},
						},
					},
					{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  20,
								IsNull: false,
							},
							element.NumericElement{
								Value:  30,
								IsNull: false,
							},
							element.NumericElement{
								Value:  25,
								IsNull: false,
							},
							element.NumericElement{
								Value:  40,
								IsNull: false,
							},
						},
					},
				},
				RecordCount: 4,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.DataFrame.LoadRecords(tt.args.Records)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})

	}
}

func TestDataFrame_Append(t *testing.T) {
	type fields struct {
		DataFrame
	}
	type args struct {
		DataFrame
	}
	tests := []struct {
		name string
		fields
		args
		want    DataFrame
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				DataFrame: DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Banana",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  25,
									IsNull: false,
								},
								element.NumericElement{
									Value:  40,
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Orange",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Banana",
								IsNull: false,
							},
						},
						AggregatedMethod: series.None,
					},
					{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  20,
								IsNull: false,
							},
							element.NumericElement{
								Value:  30,
								IsNull: false,
							},
							element.NumericElement{
								Value:  25,
								IsNull: false,
							},
							element.NumericElement{
								Value:  40,
								IsNull: false,
							},
						},
						AggregatedMethod: series.None,
					},
				},
				RecordCount: 4,
			},
			wantErr: false,
		},
		{
			name: "pass (same series name but different aggregation method)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
						{
							Name: "series_1",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
							AggregatedMethod: series.Mean,
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				DataFrame: DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Banana",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
						{
							Name: "series_1",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  25,
									IsNull: false,
								},
								element.NumericElement{
									Value:  40,
									IsNull: false,
								},
							},
							AggregatedMethod: series.Mean,
						},
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Orange",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
							element.StringElement{
								Value:  "Banana",
								IsNull: false,
							},
						},
						AggregatedMethod: series.None,
					},
					{
						Name: "series_1",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  20,
								IsNull: false,
							},
							element.NumericElement{
								Value:  30,
								IsNull: false,
							},
							element.NumericElement{
								Value:  25,
								IsNull: false,
							},
							element.NumericElement{
								Value:  40,
								IsNull: false,
							},
						},
						AggregatedMethod: series.Mean,
					},
				},
				RecordCount: 4,
			},
			wantErr: false,
		},
		{
			name: "fail (columns length mismatch)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
						},
						{
							Name: "series_3",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  25,
									IsNull: false,
								},
								element.NumericElement{
									Value:  40,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				DataFrame: DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Banana",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  25,
									IsNull: false,
								},
								element.NumericElement{
									Value:  40,
									IsNull: false,
								},
							},
						},
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
		{
			name: "fail (series type mismatch)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
							},
						},
					},
					RecordCount: 2,
				},
			},
			args: args{
				DataFrame: DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Banana",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "25",
									IsNull: false,
								},
								element.StringElement{
									Value:  "40",
									IsNull: false,
								},
							},
						},
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.DataFrame.Append(tt.args.DataFrame)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})

	}
}

func TestDataFrame_Delete(t *testing.T) {
	type fields struct {
		DataFrame
	}
	tests := []struct {
		name string
		fields
		want    DataFrame
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
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
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  4,
									IsNull: false,
								},
								element.NumericElement{
									Value:  5,
									IsNull: false,
								},
								element.NumericElement{
									Value:  6,
									IsNull: false,
								},
							},
						},
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name:     "series_1",
						Elements: element.NumericElements{},
					},
					{
						Name:     "series_2",
						Elements: element.NumericElements{},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.Delete()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})
	}
}

func TestDataFrame_Aggregate(t *testing.T) {
	type fields struct {
		DataFrame
	}
	type args struct {
		AggregationConditions
	}
	tests := []struct {
		name string
		fields
		args
		want    DataFrame
		wantErr bool
	}{
		{
			name: "pass (single series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_1",
						Method:     series.Count,
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  3,
								IsNull: false,
							},
						},
						AggregatedMethod: series.Count,
					},
				},
				RecordCount: 1,
			},
			wantErr: false,
		},
		{
			name: "pass (single series with multiple conditions)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_1",
						Method:     series.None,
					},
					{
						ColumnName: "series_1",
						Method:     series.Count,
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
						},
						AggregatedMethod: series.None,
					},
					{
						Name: "series_1",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  3,
								IsNull: false,
							},
						},
						AggregatedMethod: series.Count,
					},
				},
				RecordCount: 1,
			},
			wantErr: false,
		},
		{
			name: "pass (multiple series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
								element.NumericElement{
									Value:  60,
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_1",
						Method:     series.Count,
					},
					{
						ColumnName: "series_2",
						Method:     series.Mean,
					},
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  3,
								IsNull: false,
							},
						},
						AggregatedMethod: series.Count,
					},
					{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  40,
								IsNull: false,
							},
						},
						AggregatedMethod: series.Mean,
					},
				},
				RecordCount: 1,
			},
			wantErr: false,
		},
		{
			name: "fail (aggregated series)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
								element.NumericElement{
									Value:  60,
									IsNull: false,
								},
							},
							AggregatedMethod: series.Count,
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_1",
						Method:     series.Count,
					},
					{
						ColumnName: "series_2",
						Method:     series.Mean,
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
		{
			name: "fail (elements are not same)",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
							},
							AggregatedMethod: series.None,
						},
					},
					RecordCount: 3,
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_1",
						Method:     series.None,
					},
					{
						ColumnName: "series_1",
						Method:     series.Count,
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.DataFrame.Aggregate(tt.args.AggregationConditions)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})

	}
}

func TestDataFrame_DropNA(t *testing.T) {
	type fields struct {
		DataFrame
	}
	tests := []struct {
		name string
		fields
		want    DataFrame
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				DataFrame{
					Columns: []series.Series{
						{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "Apple",
									IsNull: false,
								},
								element.StringElement{
									Value:  "",
									IsNull: true,
								},
								element.StringElement{
									Value:  "Orange",
									IsNull: false,
								},
							},
						},
						{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  20,
									IsNull: false,
								},
								element.NumericElement{
									Value:  30,
									IsNull: false,
								},
								element.NumericElement{
									Value:  0,
									IsNull: true,
								},
							},
						},
					},
					RecordCount: 3,
				},
			},
			want: DataFrame{
				Columns: []series.Series{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "Apple",
								IsNull: false,
							},
						},
					},
					{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  20,
								IsNull: false,
							},
						},
					},
				},
				RecordCount: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.DataFrame.DropNA()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(err != nil, tt.wantErr); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})

	}
}
