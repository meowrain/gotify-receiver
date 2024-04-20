package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gopkg.in/toast.v1"
	"gotify-client/internal/client"
	"log"
	"net/http"
	"net/url"
	"time"

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
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			_ = toml.GenerateConfig()
			logger.Log().Fatalf("未找到配置文件, 已生成示例配置文件于运行路径下")
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
	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}

	wsBuilder := func() (*websocket.Conn, error) {
		ws, _, err := dialer.Dial(u.String(), nil)
		return ws, err
	}
	ws, err := wsBuilder()
	if err != nil {
		logger.Log().Fatalf("ws 连接失败: %s, 请检查配置是否有误", err)
	}

	defer func() {
		err := recover()
		if err != nil {
			log.Println("panic:", err)
			ws, _ = wsBuilder()
		}
	}()

	go func() {
		hasError := false

		preHandler := func() {
			if hasError {
				for {
					if ws != nil {
						_ = ws.Close()
					}

					log.Println("尝试断线重连...")
					ws, err = wsBuilder()
					if err != nil {
						log.Println("ws 重连异常:", err)
						hasError = true
						time.Sleep(time.Second * 3)
					} else {
						hasError = false
						break
					}
				}
			}
		}

		handler := func() {
			defer func() {
				_ = ws.Close()
			}()

			preHandler()
			for {
				_, rawMessage, err := ws.ReadMessage()
				if err != nil {
					log.Println("ws 消息接收异常:", err)
					hasError = true
					preHandler()
					continue
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
		}

		handler()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
