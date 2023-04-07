package global

import (
	"github.com/imroc/req/v3"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

var Sqlite *sqlx.DB
var Log *zap.Logger
var QB *req.Client
var Cache *cache.Cache
