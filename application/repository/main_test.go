package repository

import (
	"github.com/jmoiron/sqlx"
	"os"
	"testing"
	"umaru/application/global"
	"umaru/application/setting"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	setting.SERVER_DEBUG = true
	setting.SERVER_PORT = 8081
	setting.SERVER_TOKEN_EXPIRATION_TIME = 36000
	setting.SERVER_USERNAME = "admin"
	setting.SERVER_PASSWORD = "admin"
	setting.SERVER_LOG_DIR = "D:/home/negevmoe/umaru-server/.log"
	setting.SOURCE_PATH = "D:/home/negevmoe/umaru-server/docker/qbittorrent/downloads"
	setting.MEDIA_PATH = "D:/home/negevmoe/umaru-server/docker/jellyfin/media"
	setting.DB_PATH = "./test.db"
	setting.DB_MAX_CONNS = 30
	setting.QB_URL = "http://localhost:9999"
	setting.QB_USERNAME = "admin"
	setting.QB_PASSWORD = "adminadmin"
	setting.QB_CATEGORY = "umaru"
	setting.QB_DOWNLOAD_PATH = "/downloads"

	global.Init()
	db = global.Sqlite
	defer db.Close()

	Init()
	m.Run()

	os.Remove("test.db")

}
