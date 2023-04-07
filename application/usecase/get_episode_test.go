package usecase

import "testing"

func TestGetEpisode(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{"", args{title: "[桜都字幕组] 不死者之王 第四季 / OVERLORD Ⅳ [05][1080p][简体内嵌]"}, 5, true},
		{"", args{title: "[桜都字幕组] 不死者之王 第季 / OVERLORD Ⅳ [ 05 ][1080p][简体内嵌]"}, 5, true},
		{"", args{title: "[桜都字幕组] 不死者之王 第十三季 / OVERLORD Ⅳ [ 5 ][1080p][简体内嵌]"}, 5, true},
		{"", args{title: "[桜都字幕组] 不死者之王 第三四二季 / OVERLORD Ⅳ [5][1080p][简体内嵌]"}, 5, true},
		{"", args{title: "[桜都字幕组] 不死者之王 第三十四季 / OVERLORD Ⅳ [15][1080p][简体内嵌]"}, 15, true},
		{"", args{title: "[桜都字幕组] 不死者之王 第四期 / OVERLORD Ⅳ [ 35 ][1080p][简体内嵌]"}, 35, true},
		{"", args{title: "[桜都字幕组] 不死者之王 s4 / OVERLORD Ⅳ 15 [1080p][简体内嵌]"}, 15, true},
		{"", args{title: "[桜都字幕组] 不死者之王 s12 / OVERLORD Ⅳ [1080p][简体内嵌]"}, 0, false},
		{"", args{title: "[桜都字幕组] 不死者之王 S4 / OVERLORD Ⅳ 06 [1080p][简体内嵌]"}, 6, true},
		{"", args{title: "[桜都字幕组] 不死者之王 s04 / OVERLORD Ⅳ 8 [1080p][简体内嵌]"}, 8, true},
		{"第x集", args{title: "[桜都字幕组] 不死者之王 s04 / OVERLORD Ⅳ 第08集 [1080p][简体内嵌]"}, 8, true},
		{"第x话", args{title: "[桜都字幕组] 不死者之王 s04 / OVERLORD Ⅳ 第08话 [1080p][简体内嵌]"}, 8, true},
		{"第x话 #2", args{title: "[桜都字幕组] 不死者之王 s04 / OVERLORD Ⅳ 第382话 [1080p][简体内嵌]"}, 382, true},
		{"中文数字1", args{title: "[桜都字幕组] 不死者之王 s04 / OVERLORD Ⅳ 第八话 [1080p][简体内嵌]"}, 0, false},
		{"中文数字2", args{title: "[桜都字幕组] 不死者之王 s04 / OVERLORD Ⅳ 第八集 [1080p][简体内嵌]"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetEpisode(tt.args.title)
			if got != tt.want {
				t.Errorf("GetEpisode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetEpisode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
