package err

import (
	"fmt"
)

var (
	EXCEED_MAXiMUN_VIDEO_LENGTH_ERROR error = fmt.Errorf("超过视频长度最大限制！")
)
