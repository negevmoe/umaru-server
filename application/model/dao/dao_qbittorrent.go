package dao

type QbLog struct {
	Id        int    `json:"id"`        // ID
	Message   string `json:"message"`   // 消息
	Timestamp int64  `json:"timestamp"` // 时间
	Type      int    `json:"type"`      // 类型 Log::NORMAL: 1, Log::INFO: 2, Log::WARNING: 4, Log::CRITICAL: 8
}

type QbCategory struct {
}

type RuleDef struct {
	Enabled          bool     `json:"enabled"`          // 是否启用
	MustContain      string   `json:"mustContain"`      // 必须包含
	MustNotContain   string   `json:"mustNotContain"`   // 必须不包含
	UseRegex         bool     `json:"useRegex"`         // 启用正则
	EpisodeFilter    string   `json:"episodeFilter"`    // 剧集过滤
	SmartFilter      bool     `json:"smartFilter"`      // 是否开启智能剧集过滤
	AffectedFeeds    []string `json:"affectedFeeds"`    // 应用rss
	IgnoreDays       int      `json:"ignoreDays"`       // 忽略多少天前的
	LastMatch        string   `json:"lastMatch"`        // 最后匹配
	AddPaused        bool     `json:"addPaused"`        // true: 添加后不会立即下载
	AssignedCategory string   `json:"assignedCategory"` // 下载分类
	SavePath         string   `json:"savePath"`         // 下载后的保存路径
}

type QbittorrentRule struct {
	AffectedFeeds    []string `json:"affectedFeeds"`    // 影响的rss url
	AssignedCategory string   `json:"assignedCategory"` // 种子下载分类
	Enabled          bool     `json:"enabled"`          // 是否启用
	EpisodeFilter    string   `json:"episodeFilter"`    // 剧集过滤
	IgnoreDays       int      `json:"ignoreDays"`       // 忽略指定时间后的匹配项
	LastMatch        string   `json:"lastMatch"`        // 上次匹配时间
	MustContain      string   `json:"mustContain"`      // 必须包含
	MustNotContain   string   `json:"mustNotContain"`   // 必须不包含
	SavePath         string   `json:"savePath"`         // 默认下载路径
	SmartFilter      bool     `json:"smartFilter"`      // 智能剧集过滤
	UseRegex         bool     `json:"useRegex"`         // 是否使用正则表达式
}
