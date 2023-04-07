package usecase

import (
	"testing"
	"umaru/application/setting"
)

func TestMain(m *testing.M) {
	setting.QB_DOWNLOAD_PATH = "/downloads"
	setting.QB_RSS_FOLDER = "umaru"
	m.Run()
}
