package dao

type Cron struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Next int64  `json:"next"`
	Prev int64  `json:"prev"`
	Sep  string `json:"sep"`
}

type CronCache struct {
	Id   int
	Name string
	Sep  string
	Func func()
}
