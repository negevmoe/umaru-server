package handler

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"strconv"
	"umaru/application/model/dao"
)

var CronServer = &CronServerT{
	C: cron.New(cron.WithSeconds()),
	M: cache.New(-1, -1),
}

type CronServerT struct {
	C *cron.Cron
	M *cache.Cache
}

func (c *CronServerT) cacheKey(id cron.EntryID) string {
	return strconv.Itoa(int(id))
}

func (c *CronServerT) Start() {
	c.C.Start()
}

func (c *CronServerT) Stop() {
	c.C.Stop()
}

// Add 添加定时任务
func (c *CronServerT) Add(sep string, name string, f func()) error {
	id, err := c.C.AddFunc(sep, f)
	if err != nil {
		return err
	}

	entry := c.C.Entry(id)
	err = c.M.Add(c.cacheKey(id), dao.CronCache{
		Id:   int(entry.ID),
		Name: name,
		Sep:  sep,
		Func: f,
	}, -1)
	if err != nil {
		return err
	}

	return nil
}

func (c *CronServerT) Init() {
	err := c.Add("0 */10 * * * ?", "硬连接", Link)
	if err != nil {
		log.Fatal("硬连接定时任务创建失败", zap.Error(err))
	}
	log.Info("硬连接定时任务创建成功")

	err = c.Add("0 */10 * * * ?", "清理媒体文件夹空目录", ClearEmptyDir)
	if err != nil {
		log.Fatal("清理空目录定时任务创建失败", zap.Error(err))
	}
	log.Info("清理空目录定时任务创建成功")

	//err = c.Add("*/10 * * * * ?", "测试任务", func() {
	//	fmt.Println("每10秒一次的定时任务", time.Now().Format("2006-01-02 15:04:05"))
	//})
	//if err != nil {
	//	log.Fatal("测试任务添加失败", zap.Error(err))
	//}
	//log.Info("测试任务添加成功")
}

// List 获取定时任务列表
func (c *CronServerT) List() []dao.Cron {
	list := c.C.Entries()
	res := make([]dao.Cron, 0, len(list))

	for _, item := range list {
		a, found := c.M.Get(c.cacheKey(item.ID))
		if found {
			it := a.(dao.CronCache)
			var prev int64 = 0
			if item.Prev.Unix() > 0 {
				prev = item.Prev.Unix()
			}
			res = append(res, dao.Cron{
				Id:   it.Id,
				Name: it.Name,
				Next: item.Next.Unix(),
				Prev: prev,
				Sep:  it.Sep,
			})
		}
	}
	return res
}

// UpdateCron 更新定时任务时间
func (c *CronServerT) UpdateCron(id cron.EntryID, sep string) error {

	// 检查定时任务是否存在
	entry := c.C.Entry(id)
	if !entry.Valid() {
		return errors.New("定时任务不存在")
	}
	// 获取定时任务实例
	a, found := c.M.Get(c.cacheKey(id))
	if !found {
		return errors.New("定时任务不存在")
	}
	it := a.(dao.CronCache)

	// 重新添加定时任务
	err := c.Add(sep, it.Name, it.Func)
	if err != nil {
		return err
	}

	// 删除定时任务
	c.C.Remove(id)
	c.M.Delete(c.cacheKey(id))

	return nil
}

// RunIt 立即运行一次定时任务
func (c *CronServerT) RunIt(id cron.EntryID) (err error) {
	entry := c.C.Entry(id)
	if !entry.Valid() {
		return errors.New("任务不存在")
	}
	entry.Job.Run()
	return nil
}

func (c *CronServerT) Get(id cron.EntryID) (res dao.Cron, err error) {
	entry := c.C.Entry(id)
	if !entry.Valid() {
		err = errors.New("定时任务不存在")
		return
	}
	a, found := c.M.Get(c.cacheKey(id))
	if !found {
		err = errors.New("定时任务不存在")
		return
	}
	ch := a.(dao.CronCache)
	res.Id = ch.Id
	res.Name = ch.Name
	res.Next = entry.Next.Unix()
	res.Sep = ch.Sep
	var prev int64 = 0
	if entry.Prev.Unix() > 0 {
		prev = entry.Prev.Unix()
	}
	res.Prev = prev
	return
}
