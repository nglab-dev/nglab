package conf

import (
	"github.com/xbmlz/ungo/server"
)

type Config struct {
	Server server.Config `json:"server" yaml:"server"`
}
