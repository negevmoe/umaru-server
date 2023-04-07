package tool

import "testing"

func TestPathJoin(t *testing.T) {
	type args struct {
		p []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"", args{p: []string{"/downloads", "123", "222"}}, "/downloads/123/222"},
		{"", args{p: []string{"./downloads/dir", "123", "222"}}, "./downloads/dir/123/222"},
		{"", args{p: []string{"downloads", "123", "222"}}, "downloads/123/222"},
		{"", args{p: []string{"/downloads", "/123", "/222"}}, "/downloads/123/222"},
		{"", args{p: []string{"/downloads", "////123", "////222"}}, "/downloads/123/222"},
		{"", args{p: []string{"////downloads///", "//123///", "//222////"}}, "/downloads/123/222"},
		{"", args{p: []string{"\\downloads\\", "\\123\\", "\\222\\"}}, "\\downloads\\/\\123\\/\\222\\"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := PathJoin(tt.args.p...); gotRes != tt.wantRes {
				t.Errorf("PathJoin() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
