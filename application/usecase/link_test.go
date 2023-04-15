package usecase

import "testing"

func TestGetLinkDir(t *testing.T) {
	type args struct {
		category string
		title    string
		season   int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{"剧场版", "天气之子", -1}, "/media/剧场版/天气之子", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLinkDir(tt.args.category, tt.args.title, tt.args.season)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinkDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLinkDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLinkPath(t *testing.T) {
	type args struct {
		category string
		title    string
		season   int64
		episode  int64
		ext      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{"剧场版", "天气之子", -1, -1, ".mkv"}, "/media/剧场版/天气之子/天气之子.mkv", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLinkPath(tt.args.category, tt.args.title, tt.args.season, tt.args.episode, tt.args.ext)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinkPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLinkPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
