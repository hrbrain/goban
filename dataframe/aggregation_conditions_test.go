package dataframe

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hrbrain/goban/series"
)

func TestAggregationConditions_GetCondition(t *testing.T) {
	type fields struct {
		AggregationConditions
	}
	type args struct {
		index int
	}
	tests := []struct {
		name string
		fields
		args
		want    AggregationCondition
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				AggregationConditions{
					{
						Method: series.Count,
					}, {
						Method: series.Mean,
					},
				},
			},
			args: args{
				index: 1,
			},
			want: AggregationCondition{
				Method: series.Mean,
			},
			wantErr: false,
		},
		{
			name: "fail (index is less than 0)",
			fields: fields{
				AggregationConditions{
					{
						Method: series.Count,
					}, {
						Method: series.Mean,
					},
				},
			},
			args: args{
				index: -1,
			},
			want:    AggregationCondition{},
			wantErr: true,
		},
		{
			name: "fail (index is equal to length)",
			fields: fields{
				AggregationConditions{
					{
						Method: series.Count,
					}, {
						Method: series.Mean,
					},
				},
			},
			args: args{
				index: 2,
			},
			want:    AggregationCondition{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.GetCondition(tt.args.index)
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
