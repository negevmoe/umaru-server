package usecase

import "testing"

func TestGetSeason(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
		wantOk  bool
	}{
		{"中文", args{title: "OVERLORD 第四季"}, 4, true},
		{"中文", args{title: "OVERLORD 第四期"}, 4, true},
		{"中文无空格", args{title: "OVERLORD第四季"}, 0, false},
		{"中文无信息", args{title: "OVERLORD 第季"}, 0, false},
		{"中文 >10", args{title: "OVERLORD 第十三季"}, 0, false},

		{"数字", args{title: "OVERLORD s4"}, 4, true},
		{"数字 >10", args{title: "OVERLORD s12"}, 12, true},
		{"数字 >100", args{title: "OVERLORD s12232"}, 12232, true},
		{"数字 无空格", args{title: "OVERLORDs12232"}, 0, false},
		{"数字 无信息", args{title: "OVERLORD s"}, 0, false},
		{"数字 空", args{title: "OVERLORD []"}, 0, false},

		{"罗马", args{title: "OVERLORD IV"}, 4, true},
		{"缺失", args{title: "OVERLORD"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := GetSeason(tt.args.title)
			if gotRes != tt.wantRes {
				t.Errorf("GetSeason() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetSeason() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_findNumberSeason(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
		wantOk  bool
	}{
		{"数字", args{title: "OVERLORD s4"}, 4, true},
		{"数字 大于10", args{title: "OVERLORD s12"}, 12, true},
		{"数字 大于100", args{title: "OVERLORD s12232"}, 12232, true},
		{"数字 无空格", args{title: "OVERLORDs12232"}, 0, false},
		{"数字 无信息", args{title: "OVERLORD s"}, 0, false},
		{"数字 空", args{title: "OVERLORD []"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := findNumberSeason(tt.args.title)
			if gotRes != tt.wantRes {
				t.Errorf("findNumberSeason() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("findNumberSeason() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_findRomeSeason(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
		wantOk  bool
	}{
		{"罗马", args{title: "OVERLORD IV"}, 4, true},
		{"罗马", args{title: "OVERLORD 第IV季"}, 4, true},
		{"罗马", args{title: "OVERLORD 第IV期"}, 4, true},
		{"罗马 防读到集", args{title: "OVERLORD 第IV集"}, 0, false},

		{"罗马 数字格式错误", args{title: "OVERLORD IVV"}, 0, false},
		{"罗马 不支持的季", args{title: "OVERLORD XII"}, 0, false},
		{"罗马 缺失", args{title: "OVERLORD"}, 0, false},
		{"罗马 无空格", args{title: "OVERLORDVI"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := findRomeSeason(tt.args.title)
			if gotRes != tt.wantRes {
				t.Errorf("findRomeSeason() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("findRomeSeason() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_findZhSeason(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
		wantOk  bool
	}{
		{"数字", args{"[桜都字幕组] 不死者之王 第4季 / OVERLORD Ⅳ [05][1080p][简体内嵌]"}, 0, false},
		{"罗马", args{"[桜都字幕组] OVERLORD Ⅳ [05][1080p][简体内嵌]"}, 0, false},
		{"中文", args{"[桜都字幕组] 不死者之王 第四季 / OVERLORD Ⅳ [05][1080p][简体内嵌]"}, 4, true},
		{"中文", args{"[桜都字幕组] 不死者之王 第四期 / OVERLORD Ⅳ [05][1080p][简体内嵌]"}, 4, true},
		{"中文缺失", args{"[桜都字幕组] 不死者之王 第期 / OVERLORD Ⅳ [05][1080p][简体内嵌]"}, 0, false},
		{"中文缺失", args{"[桜都字幕组] 不死者之王 / OVERLORD Ⅳ [05][1080p][简体内嵌]"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := findZhSeason(tt.args.title)
			if gotRes != tt.wantRes {
				t.Errorf("findZhSeason() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("findZhSeason() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
