package element

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringElement_Equal(t *testing.T) {
	type fields struct {
		stringElement StringElement
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
				stringElement: StringElement{
					Value:  "abc",
					IsNull: false,
				},
			},
			args: args{
				element: StringElement{
					Value:  "abc",
					IsNull: false,
				},
			},
			want: true,
		},
		{
			name: "fail (type mismatch)",
			fields: fields{
				stringElement: StringElement{
					Value:  "abc",
					IsNull: false,
				},
			},
			args: args{
				element: NumericElement{
					Value:  99,
					IsNull: false,
				},
			},
			want: false,
		},
		{
			name: "fail (different value)",
			fields: fields{
				stringElement: StringElement{
					Value:  "abc",
					IsNull: false,
				},
			},
			args: args{
				element: StringElement{
					Value:  "ddd",
					IsNull: false,
				},
			},
			want: false,
		},
		{
			name: "fail (both elements are null)",
			fields: fields{
				stringElement: StringElement{
					Value:  "xxx",
					IsNull: true,
				},
			},
			args: args{
				element: StringElement{
					Value:  "",
					IsNull: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.stringElement.Equal(tt.args.element)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})

	}

}

func TestStringElement_Split(t *testing.T) {
	type fields struct {
		stringElement StringElement
	}

	type args struct {
		separator string
		limit     int
	}

	tests := []struct {
		name string
		fields
		args
		want    StringListElement
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				stringElement: StringElement{
					Value:  "a,b,c",
					IsNull: false,
				},
			},
			args: args{
				separator: ",",
				limit:     3,
			},
			want: StringListElement{"a", "b", "c"},
		},
		{
			name: "pass (limit is less than the number of values)",
			fields: fields{
				stringElement: StringElement{
					Value:  "a,b,c",
					IsNull: false,
				},
			},
			args: args{
				separator: ",",
				limit:     2,
			},
			want: StringListElement{"a", "b,c"},
		},
		{
			name: "pass (limit is more than the number of values)",
			fields: fields{
				stringElement: StringElement{
					Value:  "a,b,c",
					IsNull: false,
				},
			},
			args: args{
				separator: ",",
				limit:     10,
			},
			want: StringListElement{"a", "b", "c"},
		},
		{
			name: "pass (limit is zero)",
			fields: fields{
				stringElement: StringElement{
					Value:  "a,b,c",
					IsNull: false,
				},
			},
			args: args{
				separator: ",",
				limit:     0,
			},
			want: StringListElement{"a,b,c"},
		},
		{
			name: "pass (limit is one)",
			fields: fields{
				stringElement: StringElement{
					Value:  "a,b,c",
					IsNull: false,
				},
			},
			args: args{
				separator: ",",
				limit:     1,
			},
			want: StringListElement{"a,b,c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.stringElement.Split(tt.args.separator, tt.args.limit)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
