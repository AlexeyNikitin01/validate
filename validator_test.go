package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		checkErr func(err error) bool
	}{
		{
			name: "valid struct with tagged fields",
			args: args{
				v: struct{
					ID        int64  `json:"id"`
					Title     string `json:"title"`
					Text      string `json:"text"`
					AuthorID  int64  `json:"author_id"`
					Published bool   `json:"published"`
				}{
					Title:       "asdf",
					Text:      "asfdasdf",
				},
			},
			wantErr: false,
		},
		{
			name: "wrong in",
			args: args{
				v: struct{
					ID        int64  `json:"id"`
					Title     string `json:"title"`
					Text      string `json:"text"`
					AuthorID  int64  `json:"author_id"`
					Published bool   `json:"published"`
				}{
					Title:       "asdf",
					Text:      "",
				},
			},
			wantErr: true,
			checkErr: func(err error) bool {
				assert.Len(t, err.(ValidationErrors), 1)
				return true
			},
		},
		{
			name: "wrong min",
			args: args{
				v: struct{
					ID        int64  `json:"id"`
					Title     string `json:"title"`
					Text      string `json:"text"`
					AuthorID  int64  `json:"author_id"`
					Published bool   `json:"published"`
				}{
					Title:       "",
					Text:      "",
				},
			},
			wantErr: true,
			checkErr: func(err error) bool {
				assert.Len(t, err.(ValidationErrors), 2)
				return true
			},
		},
		{
			name: "wrong max",
			args: args{
				v: struct{
					ID        int64  `json:"id"`
					Title     string `json:"title"`
					Text      string `json:"text"`
					AuthorID  int64  `json:"author_id"`
					Published bool   `json:"published"`
				}{
					Title:       "",
					Text:      "asdf",
				},
			},
			wantErr: true,
			checkErr: func(err error) bool {
				assert.Len(t, err.(ValidationErrors), 1)
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.args.v)
			if tt.wantErr {
				assert.Error(t, err)
				assert.True(t, tt.checkErr(err), "test expect an error, but got wrong error type")
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
