package searcher

import (
	"testing"

	"github.com/j-04/pass-manager/core/searcher"
)

const (
	THRESHOLD int = 2
)

func TestLevenshteinSearcher_Find(t *testing.T) {
	type fields struct {
		threshold int
	}
	type args struct {
		pattern string
		source  string
	}

	fs := &fields{threshold: THRESHOLD}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Exact match",
			fields: *fs,
			args:   args{pattern: "apple", source: "apple"},
			want:   true,
		},
		{
			name:   "One insertion required",
			fields: *fs,
			args:   args{pattern: "apple", source: "aple"},
			want:   true,
		},
		{
			name:   "One deletion required",
			fields: *fs,
			args:   args{pattern: "apple", source: "aple"},
			want:   true,
		},
		{
			name:   "One substitution required",
			fields: *fs,
			args:   args{pattern: "apple", source: "aplee"},
			want:   true,
		},
		{
			name:   "Two insertions requried",
			fields: *fs,
			args:   args{pattern: "apple:", source: "ale"},
			want:   false,
		},
		{
			name:   "Two deletions required",
			fields: *fs,
			args:   args{pattern: "apple", source: "apxxe"},
			want:   true,
		},
		{
			name:   "Two substitutions required",
			fields: *fs,
			args:   args{pattern: "apple", source: "apxxe"},
			want:   true,
		},
		{
			name:   "Exceeds threshold (3 insertions)",
			fields: *fs,
			args:   args{pattern: "apple", source: "a"},
			want:   false,
		},
		{
			name:   "Exceeds threshold (3 substitutions)",
			fields: *fs,
			args:   args{pattern: "apple", source: "aaaaa"},
			want:   false,
		},
		{
			name:   "Empty strings",
			fields: *fs,
			args:   args{pattern: "", source: ""},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := searcher.NewLevenshteinSearcher(tt.fields.threshold)
			got, err := this.Find(tt.args.pattern, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("LevenshteinSearcher.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LevenshteinSearcher.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
