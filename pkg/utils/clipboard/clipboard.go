// Copyright 2024 shikong
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Motified By MeowRain on 2024.8.29
// Added copy function

package clipboard

import (
	"gotify-client/pkg/logger"

	"golang.design/x/clipboard"
)

func CopyToClipBoard(str string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	// 使用协程进行复制操作
	go func() {
		isCopy := clipboard.Write(clipboard.FmtText, []byte(str))
		select {
		case <-isCopy:
			log := logger.Log()
			log.Infoln("复制成功")
		}
	}()
	return nil
}
