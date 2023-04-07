## 创建番剧

`setting.QB_DOWNLOAD_PATH/数据库番剧ID` 为qb下载的视频存放路径,手动添加种子,rss订阅,下载的视频都存放在此处
bangumiID 唯一 可以为空 空表示为手动添加的番剧
番剧名+季 唯一

## 硬链接

`setting.QB_DOWNLOAD_PATH/数据库番剧ID/原文件` ==> `setting.MEDIA_PATH/分类名称/番剧名称/S季/重命名文件`

## 删除番剧

会从数据库中删除数据
会删除硬链接的视频文件
会删除qb中的rss连接
会删除qb中的下载规则
不会qb已经下载的视频文件与种子列表中的数据

## 更新番剧

```json
{
  "id": 0,
  "title": "",
  "season": 0,
  "category_id": 0,
  "total": 0,
  "play_time": 0
}
```

可以更新 标题,季,分类,总集数,分类,放送时间
其中
只更新标题时,创建 新番剧名称/S季 的文件夹, 原番剧名称/S季 目录移动到 新番剧名称/S季 ; 更新标题时要判断
新标题+原季是否存在    
`setting.MEDIA_PATH/分类/原番剧名称/S季/重命名文件` ==> `setting.MEDIA_PATH/分类/新番剧名称/S季/重命名文件`
只更新季时, 新建季文件夹, 原季文件夹 移动到 新季文件夹; 更新季时要判断 标题+新季是否存在
`setting.MEDIA_PATH/分类/番剧名称/原季/重命名文件` ==> `setting.MEDIA_PATH/分类/番剧名称/新季/重命名文件`

标题与季更新时, 创建 新番剧名称/新季 的文件夹, 原番剧名称/原季 目录移动到

## 更新分类

原分类下的所有番剧都移动到新文件夹下
`setting.MEDIA_PATH/原分类/番剧名称/S季/重命名文件` ==> `setting.MEDIA_PATH/新分类/番剧名称/S季/重命名文件`

## 删除分类

默认的未分类禁止删除
原分类下的所有番剧都移动到默认文件夹下
`setting.MEDIA_PATH/原分类/番剧名称/S季/重命名文件` ==> `setting.MEDIA_PATH/未分类/番剧名称/S季/重命名文件`

## 刮削插件

使用bangumi

```
https://jellyfin-plugin-bangumi.pages.dev/repository.json
```