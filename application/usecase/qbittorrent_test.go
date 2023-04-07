package usecase

import (
	"testing"
)

func TestGetQbDownloadPath(t *testing.T) {
	type args struct {
		animeId int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQbDownloadPath(tt.args.animeId); got != tt.want {
				t.Errorf("GetQbDownloadPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRssPath(t *testing.T) {

}
