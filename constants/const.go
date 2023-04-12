package constants

const (
	ConfigFileName = "config"
	ConfigType     = "toml"
)

var ConfigPaths = []string{
	".",
	"./conf",
	"/config",
}
