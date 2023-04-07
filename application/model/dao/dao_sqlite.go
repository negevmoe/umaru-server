package dao

type Anime struct {
	Id         int64  `db:"id" json:"id"`                   // ID
	BangumiId  int64  `db:"bangumi_id" json:"bangumi_id"`   // bangumi的ID
	CategoryId int64  `db:"category_id" json:"category_id"` // 类别 1:tv,2:剧场版,3:OVA
	Title      string `db:"title" json:"title"`             // 标题
	Season     int64  `db:"season" json:"season"`           // 第几季
	Cover      string `db:"cover" json:"cover"`             // 封面图
	Total      int64  `db:"total" json:"total"`             // 总集数
	RssUrl     string `db:"rss_url" json:"rss_url"`         // rss链接
	RssPath    string `db:"rss_path" json:"-"`              // qb rss path
	PlayTime   int64  `db:"play_time" json:"play_time"`     // 放送时间
	CreateTime int64  `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64  `db:"update_time" json:"update_time"` // 更新时间
}

type Category struct {
	Id         int64  `db:"id" json:"id"`                   // ID
	Name       string `db:"name" json:"name"`               // 分类名称
	Origin     int64  `db:"origin" json:"origin"`           // 分类来源 1为内置 2为自定义
	CreateTime int64  `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64  `db:"update_time" json:"update_time"` // 更新时间
}

type Rule struct {
	Id             int64  `db:"id" json:"id"`                             // ID
	Name           string `db:"name" json:"name"`                         // 名称
	MustContain    string `db:"must_contain" json:"must_contain"`         // 必须包含
	MustNotContain string `db:"must_not_contain" json:"must_not_contain"` // 必须不包含
	UseRegex       int    `db:"use_regex" json:"use_regex"`               // 正则表达式 1:true 2:false 默认2
	EpisodeFilter  string `db:"episode_filter" json:"episode_filter"`     // 剧集过滤
	SmartFilter    int    `db:"smart_filter" json:"smart_filter"`         // 智能剧集过滤 1:true 2:false 默认2
	CreateTime     int64  `db:"create_time" json:"create_time"`           // 创建时间
	UpdateTime     int64  `db:"update_time" json:"update_time"`           // 更新时间
}

type AnimeInfoView struct {
	Id           int64  `db:"id" json:"id"`                       // ID
	BangumiId    int64  `db:"bangumi_id" json:"bangumi_id"`       // bangumi的ID
	CategoryId   int64  `db:"category_id" json:"category_id"`     // 分类ID
	CategoryName string `db:"category_name" json:"category_name"` // 分类名称
	Title        string `db:"title" json:"title"`                 // 标题
	Season       int64  `db:"season" json:"season"`               // 第几季
	Cover        string `db:"cover" json:"cover"`                 // 封面图
	Total        int64  `db:"total" json:"total"`                 // 总集数
	RssUrl       string `db:"rss_url" json:"rss_url"`             // rss链接
	RssPath      string `db:"rss_path" json:"-"`                  // qb rss path
	PlayTime     int64  `db:"play_time" json:"play_time"`         // 放送时间
	CreateTime   int64  `db:"create_time" json:"create_time"`     // 创建时间
	UpdateTime   int64  `db:"update_time" json:"update_time"`     // 更新时间
}

type Video struct {
	AnimeId    int64  `json:"anime_id"`    // 所属番剧ID
	Path       string `json:"path"`        // 存储路径
	Filename   string `json:"filename"`    // 文件名
	Size       int64  `json:"size"`        // 文件大小
	UpdateTime int64  `json:"update_time"` // 最后修改时间
}
