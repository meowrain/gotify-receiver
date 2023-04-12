package client

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gopkg.in/toast.v1"
	"gotify-client/internal/client"
	"log"
	"net/url"

	"gotify-client/constants"
	"gotify-client/pkg/config"
	"gotify-client/pkg/config/toml"
	"gotify-client/pkg/logger"
	"os"
	"os/signal"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var notification = toast.Notification{
	AppID:               "Gotify Client",
	Title:               "",
	Message:             "",
	Icon:                "",
	ActivationType:      "",
	ActivationArguments: "",
	Actions:             nil,
	Audio:               "",
	Loop:                false,
	Duration:            "",
}

var conf = new(config.Config)

func Main() {
	viper.SetConfigName(constants.ConfigFileName)
	viper.SetConfigType(constants.ConfigType)
	for _, path := range constants.ConfigPaths {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_ = toml.GenerateConfig()
			logger.Log().Fatalf("未找到配置文件, 已生成示例配置文件于运行路径下")
		} else {
			logger.Log().Fatalf("配置解析失败 %s", err)
		}
	}

	err := viper.Unmarshal(conf)
	if err != nil {
		logger.Log().Fatalf("配置文件解析失败: %s, 请检查配置是否有误", err)
	}

	u := url.URL{
		Scheme: "ws",
		Host:   conf.Server.Addr,
		Path:   "/stream",
	}
	params := url.Values{}
	params.Add("token", conf.Server.UserToken)
	u.RawQuery = params.Encode()

	logger.Log().Infof("url %s", u.String())
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logger.Log().Fatalf("ws 连接失败: %s, 请检查配置是否有误", err)
	}
	defer func() {
		_ = ws.Close()
	}()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, rawMessage, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("recv: %s", rawMessage)
			data := &client.GotifyMessage{}
			_ = json.Unmarshal(rawMessage, data)
			message := fmt.Sprintf("%s", data.Message)

			// windows 下 默认GBK中文编码转换
			retTitle, _ := simplifiedchinese.GBK.NewEncoder().String(data.Title)
			retMessage, _ := simplifiedchinese.GBK.NewEncoder().String(message)

			notification.Title = retTitle
			notification.Message = retMessage
			_ = notification.Push()
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
