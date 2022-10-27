package dataframe

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hrbrain/goban/element"
	"github.com/hrbrain/goban/series"
)

func TestGroups_Update(t *testing.T) {
	type fields struct {
		groups Groups
	}
	type args struct {
		seriesName series.Name
		element    element.Element
		df         DataFrame
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Groups
	}{
		{
			name: "pass",
			fields: fields{
				groups: []Group{
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element1",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								series.Series{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element1",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element1",
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
											Value:  1,
											IsNull: false,
										},
									},
								},
							},
						},
					},
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element2",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								series.Series{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element2",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element2",
											IsNull: false,
										},
									},
								},
								series.Series{
									Name: "series_2",
									Elements: element.NumericElements{
										element.NumericElement{
											Value:  2,
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
					},
				},
			},
			args: args{
				seriesName: "series_1",
				element: element.StringElement{
					Value: "element2",
				},
				df: DataFrame{
					Columns: Columns{
						series.Series{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "element1",
									IsNull: false,
								},
								element.StringElement{
									Value:  "element1",
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
									Value:  3,
									IsNull: false,
								},
							},
						},
					},
				},
			},
			want: []Group{
				{
					SeriesName: "series_1",
					Element: element.StringElement{
						Value:  "element1",
						IsNull: false,
					},
					DataFrame: DataFrame{
						Columns: Columns{
							series.Series{
								Name: "series_1",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "element1",
										IsNull: false,
									},
									element.StringElement{
										Value:  "element1",
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
										Value:  1,
										IsNull: false,
									},
								},
							},
						},
					},
				},
				{
					SeriesName: "series_1",
					Element: element.StringElement{
						Value:  "element2",
						IsNull: false,
					},
					DataFrame: DataFrame{
						Columns: Columns{
							series.Series{
								Name: "series_1",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "element1",
										IsNull: false,
									},
									element.StringElement{
										Value:  "element1",
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
										Value:  3,
										IsNull: false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "pass (new group)",
			fields: fields{
				groups: []Group{
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element1",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								series.Series{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element1",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element1",
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
											Value:  1,
											IsNull: false,
										},
									},
								},
							},
						},
					},
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element2",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								series.Series{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element2",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element2",
											IsNull: false,
										},
									},
								},
								series.Series{
									Name: "series_2",
									Elements: element.NumericElements{
										element.NumericElement{
											Value:  2,
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
					},
				},
			},
			args: args{
				seriesName: "series_1",
				element: element.StringElement{
					Value:  "element3",
					IsNull: false,
				},
				df: DataFrame{
					Columns: Columns{
						series.Series{
							Name: "series_1",
							Elements: element.StringElements{
								element.StringElement{
									Value:  "element3",
									IsNull: false,
								},
								element.StringElement{
									Value:  "element3",
									IsNull: false,
								},
							},
						},
						series.Series{
							Name: "series_2",
							Elements: element.NumericElements{
								element.NumericElement{
									Value:  2,
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
			},
			want: []Group{
				{
					SeriesName: "series_1",
					Element: element.StringElement{
						Value:  "element1",
						IsNull: false,
					},
					DataFrame: DataFrame{
						Columns: Columns{
							series.Series{
								Name: "series_1",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "element1",
										IsNull: false,
									},
									element.StringElement{
										Value:  "element1",
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
										Value:  1,
										IsNull: false,
									},
								},
							},
						},
					},
				},
				{
					SeriesName: "series_1",
					Element: element.StringElement{
						Value:  "element2",
						IsNull: false,
					},
					DataFrame: DataFrame{
						Columns: Columns{
							series.Series{
								Name: "series_1",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "element2",
										IsNull: false,
									},
									element.StringElement{
										Value:  "element2",
										IsNull: false,
									},
								},
							},
							series.Series{
								Name: "series_2",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  2,
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
				},
				{
					SeriesName: "series_1",
					Element: element.StringElement{
						Value:  "element3",
						IsNull: false,
					},
					DataFrame: DataFrame{
						Columns: Columns{
							series.Series{
								Name: "series_1",
								Elements: element.StringElements{
									element.StringElement{
										Value:  "element3",
										IsNull: false,
									},
									element.StringElement{
										Value:  "element3",
										IsNull: false,
									},
								},
							},
							series.Series{
								Name: "series_2",
								Elements: element.NumericElements{
									element.NumericElement{
										Value:  2,
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
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.groups.Update(tt.args.seriesName, tt.args.element, tt.args.df)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf(diff)
			}
		})

	}
}

func TestGroups_Aggregate(t *testing.T) {
	type fields struct {
		Groups Groups
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
			name: "pass　(count)",
			fields: fields{
				Groups: []Group{
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_1",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element1",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element1",
											IsNull: false,
										},
									},
									AggregatedMethod: series.None,
								},
							},
							RecordCount: 2,
						},
					},
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_2",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element2",
											IsNull: false,
										},
									},
									AggregatedMethod: series.None,
								},
							},
							RecordCount: 1,
						},
					},
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
				Columns: Columns{
					{
						Name: "series_1",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  2,
								IsNull: false,
							},
							element.NumericElement{
								Value:  1,
								IsNull: false,
							},
						},
						AggregatedMethod: series.Count,
					},
				},
				RecordCount: 2,
			},
			wantErr: false,
		},
		{
			name: "pass　(mean with multiple columns)",
			fields: fields{
				Groups: []Group{
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_1",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_1",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element_1",
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
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_2",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_2",
											IsNull: false,
										},
									},
								},
								{
									Name: "series_2",
									Elements: element.NumericElements{
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
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_1",
						Method:     series.None,
					},
					{
						ColumnName: "series_2",
						Method:     series.Mean,
					},
				},
			},
			want: DataFrame{
				Columns: Columns{
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "element_1",
								IsNull: false,
							},
							element.StringElement{
								Value:  "element_2",
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
						AggregatedMethod: series.Mean,
					},
				},
				RecordCount: 2,
			},
			wantErr: false,
		},
		{
			name: "pass　(sum with multiple columns)",
			fields: fields{
				Groups: []Group{
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_1",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_1",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element_1",
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
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_2",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_2",
											IsNull: false,
										},
									},
								},
								{
									Name: "series_2",
									Elements: element.NumericElements{
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
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_2",
						Method:     series.Sum,
					},
				},
			},
			want: DataFrame{
				Columns: Columns{
					{
						Name: "series_2",
						Elements: element.NumericElements{
							element.NumericElement{
								Value:  50,
								IsNull: false,
							},
							element.NumericElement{
								Value:  40,
								IsNull: false,
							},
						},
						AggregatedMethod: series.Sum,
					},
				},
				RecordCount: 2,
			},
			wantErr: false,
		},
		{
			name: "pass (columns and conditions order mismatch)",
			fields: fields{
				Groups: []Group{
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_1",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_1",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element_1",
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
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_2",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_2",
											IsNull: false,
										},
									},
								},
								{
									Name: "series_2",
									Elements: element.NumericElements{
										element.NumericElement{
											Value:  40,
											IsNull: false,
										},
									},
								},
							},
							RecordCount: 1,
						},
					},
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_2",
						Method:     series.Mean,
					},
					{
						ColumnName: "series_1",
						Method:     series.None,
					},
				},
			},
			want: DataFrame{
				Columns: Columns{
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
						AggregatedMethod: series.Mean,
					},
					{
						Name: "series_1",
						Elements: element.StringElements{
							element.StringElement{
								Value:  "element_1",
								IsNull: false,
							},
							element.StringElement{
								Value:  "element_2",
								IsNull: false,
							},
						},
						AggregatedMethod: series.None,
					},
				},
				RecordCount: 2,
			},
			wantErr: false,
		},
		{
			name: "fail　(not grouped)",
			fields: fields{
				Groups: []Group{
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_1",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_1",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element_1",
											IsNull: false,
										},
									},
								},
							},
						},
					},
					{
						SeriesName: "series_1",
						Element: element.StringElement{
							Value:  "element_2",
							IsNull: false,
						},
						DataFrame: DataFrame{
							Columns: Columns{
								{
									Name: "series_1",
									Elements: element.StringElements{
										element.StringElement{
											Value:  "element_2",
											IsNull: false,
										},
										element.StringElement{
											Value:  "element_3",
											IsNull: false,
										},
									},
								},
							},
						},
					},
				},
			},
			args: args{
				AggregationConditions{
					{
						ColumnName: "series_1",
						Method:     series.None,
					},
				},
			},
			want:    DataFrame{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.Groups.Aggregate(tt.args.AggregationConditions)
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
