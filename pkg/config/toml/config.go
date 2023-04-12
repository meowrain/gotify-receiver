package toml

import (
	"github.com/pelletier/go-toml/v2"
	"gotify-client/pkg/config"
	"os"
	"path/filepath"
)

func GenerateConfig() error {
	p, _ := filepath.Abs("./config.toml")

	flag := os.O_RDWR
	_, err := os.Stat(p)
	exist := !os.IsNotExist(err)

	if !exist {
		f, err := os.OpenFile(p, flag|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()
		encoder := toml.NewEncoder(f)
		encoder.SetIndentTables(true)
		_ = encoder.Encode(config.DefaultConfig())
		_ = f.Sync()
	}

	return nil
}

func LoadConfig() (*config.Config, error) {
	p, _ := filepath.Abs("./config.toml")

	flag := os.O_RDWR
	_, err := os.Stat(p)
	exist := !os.IsNotExist(err)

	if !exist {
		f, err := os.OpenFile(p, flag|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = f.Close()
		}()
		encoder := toml.NewEncoder(f)
		encoder.SetIndentTables(true)
		_ = encoder.Encode(config.DefaultConfig())
		_ = f.Sync()
	}

	f, err := os.OpenFile(p, flag, 0644)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	c := &config.Config{}
	decoder := toml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
