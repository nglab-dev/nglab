package conf

import (
	"github.com/xbmlz/ungo/undb"
	"github.com/xbmlz/ungo/unhttp"
	"github.com/xbmlz/ungo/unlog"
)

type Config struct {
	Server unhttp.Config `json:"server" yaml:"server"`
	Log    unlog.Config  `json:"log" yaml:"log"`
	DB     undb.Config   `json:"db" yaml:"db"`
}
