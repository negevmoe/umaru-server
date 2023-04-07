package global

import (
	"crypto/tls"
	"github.com/imroc/req/v3"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/natefinch/lumberjack"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
	"umaru/application/setting"
)

/**
初始化服务
*/

func Init() {
	initZap()               // 初始化日志库
	initOutputSetting()     // 输出设置
	initSqlite()            // 初始化数据库连接与数据
	initQbittorrentClient() // 初始化qbittorrent客户端
	initCache()             // 初始化缓存
}

func initZap() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "TIME",                         // 时间字段key
		LevelKey:       "LEVEL",                        // 等级字段key
		NameKey:        zapcore.OmitKey,                //
		CallerKey:      "CALLER",                       // 调用函数信息key
		FunctionKey:    zapcore.OmitKey,                // 函数key
		MessageKey:     "MESSAGE",                      // 消息字段key
		StacktraceKey:  "TRACE",                        // 链路字段key
		LineEnding:     zapcore.DefaultLineEnding,      // 行分隔符
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 等级格式
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行消耗时间
		EncodeCaller:   zapcore.FullCallerEncoder,      // 调用函数信息格式
	}

	var encoder zapcore.Encoder
	var level zapcore.Level
	var writer zapcore.WriteSyncer

	if setting.SERVER_DEBUG { // 开发环境
		level = zapcore.DebugLevel                         // [等级]:DEBUG
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // [格式]:text
		writer = zapcore.AddSync(os.Stdout)                // [输出]:控制台
	} else { // 生产环境
		level = zapcore.InfoLevel                       // [等级]:INFO
		encoder = zapcore.NewJSONEncoder(encoderConfig) // [格式]:JSON
		writer = zapcore.AddSync(&lumberjack.Logger{    // [输出]:文件
			Filename:   filepath.Join(setting.SERVER_LOG_DIR, "umaru.Log"), // 文件名
			LocalTime:  true,                                               // 本地时间化
			Compress:   true,                                               // 启用压缩
			MaxSize:    50,                                                 // 切分大小
			MaxAge:     180,                                                // 日志存储时长
			MaxBackups: 30,                                                 // 旧文件保留个数
		})
	}

	core := zapcore.NewCore(encoder, writer, level)
	l := zap.New(core)

	l = l.WithOptions(zap.AddCaller()) // 添加文件行信息
	// DEBUG模式开启堆栈追踪
	if setting.SERVER_DEBUG {
		l = l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))
	}

	Log = l
}

func initSqlite() {
	var err error
	Log.Info("连接至数据库")
	Sqlite, err = sqlx.Open("sqlite3", setting.DB_PATH)
	if err != nil {
		Log.Fatal("连接数据库失败", zap.Error(err))
	}

	Sqlite.SetMaxOpenConns(int(setting.DB_MAX_CONNS))
	Log.Info("数据库连接成功")
}

func initQbittorrentClient() {
	QB = req.C().SetTimeout(10 * time.Second).SetBaseURL(setting.QB_URL + "/api/v2")
	QB.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // 忽略https 安全性
}

func initCache() {
	Cache = cache.New(-1, -1)
}

func initOutputSetting() {
	Log.Info("配置信息",
		zap.Bool("server_debug", setting.SERVER_DEBUG),
		zap.Int64("server_port", setting.SERVER_PORT),
		zap.Int64("token_expiration_time", setting.SERVER_TOKEN_EXPIRATION_TIME),
		zap.String("server_username", setting.SERVER_USERNAME),
		zap.String("server_password", setting.SERVER_PASSWORD),
		zap.String("server_log_dir", setting.SERVER_LOG_DIR),
		zap.String("source_path", setting.SOURCE_PATH),
		zap.String("media_path", setting.MEDIA_PATH),
		zap.String("db_path", setting.DB_PATH),
		zap.Int64("db_max_conns", setting.DB_MAX_CONNS),
		zap.String("qb_url", setting.QB_URL),
		zap.String("qb_username", setting.QB_USERNAME),
		zap.String("qb_password", setting.QB_PASSWORD),
		zap.String("qb_category", setting.QB_CATEGORY),
		zap.String("qb_rss_folder", setting.QB_RSS_FOLDER),
		zap.String("qb_download_path", setting.QB_DOWNLOAD_PATH),
	)
}
