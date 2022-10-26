package element

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringListElement_Equal(t *testing.T) {
	type fields struct {
		StringListElement StringListElement
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
			name: "true",
			fields: fields{
				StringListElement: StringListElement{
					"abc",
				},
			},
			args: args{
				element: StringListElement{
					"abc",
				},
			},
			want: true,
		},
		{
			name: "false (type mismatch)",
			fields: fields{
				StringListElement: StringListElement{
					"abc",
				},
			},
			args: args{
				element: StringElement{
					Value:  "abc",
					IsNull: false,
				},
			},
			want: false,
		},
		{
			name: "false (both are NA)",
			fields: fields{
				StringListElement: StringListElement{},
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
			got := tt.fields.StringListElement.Equal(tt.args.element)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})

	}

}

func TestStringListElement_Slice(t *testing.T) {
	type fields struct {
		StringListElement StringListElement
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		fields
		args
		want    StringListElement
		wantErr bool
	}{
		{
			name: "pass (valid range)",
			fields: fields{
				StringListElement: StringListElement{
					"ab", "cd", "ef",
				},
			},
			args: args{
				start: 0,
				end:   2,
			},
			want: StringListElement{
				"ab", "cd",
			},
			wantErr: false,
		},
		{
			name: "fail (negative start)",
			fields: fields{
				StringListElement: StringListElement{
					"ab", "cd", "ef",
				},
			},
			args: args{
				start: -1,
				end:   2,
			},
			want:    StringListElement{},
			wantErr: true,
		},
		{
			name: "fail (start > end)",
			fields: fields{
				StringListElement: StringListElement{
					"ab", "cd", "ef",
				},
			},
			args: args{
				start: 2,
				end:   1,
			},
			want:    StringListElement{},
			wantErr: true,
		},
		{
			name: "pass (start = end)",
			fields: fields{
				StringListElement: StringListElement{
					"ab", "cd", "ef",
				},
			},
			args: args{
				start: 1,
				end:   1,
			},
			want:    StringListElement{},
			wantErr: false,
		},
		{
			name: "pass (end is more than length)",
			fields: fields{
				StringListElement: StringListElement{
					"ab", "cd", "ef",
				},
			},
			args: args{
				start: 1,
				end:   10,
			},
			want: StringListElement{
				"ab", "cd", "ef",
			},
			wantErr: false,
		},
		{
			name: "pass (end is equal to length)",
			fields: fields{
				StringListElement: StringListElement{
					"ab", "cd", "ef",
				},
			},
			args: args{
				start: 1,
				end:   3,
			},
			want: StringListElement{
				"cd", "ef",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.StringListElement.Slice(tt.args.start, tt.args.end)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("StringListElement.Slice() error = %v, wantErr %v", got, tt.wantErr)
			}
		})

	}

}
