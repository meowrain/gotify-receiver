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
// Added SMS verification code parsing and auto-copy functionality

package parse

import (
	"regexp"
)

func ParseVertificationCode(str string) string {
	re := regexp.MustCompile(`\b\d{4,8}\b`)
	match := re.FindAllString(str, -1)
	if len(match) > 0 {
		return match[0]
	} else {
		return ""
	}
}
