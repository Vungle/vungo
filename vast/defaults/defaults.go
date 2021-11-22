package defaults

import "time"

// MaxVideoDuration stand for the max video duration with a default of 60 seconds
const MaxVideoDuration = 60 * time.Second

// MinVideoDuration stand for the min video duration with a default of 1 second
const MinVideoDuration = 1 * time.Second

// MaxVideoWidth stand for the max value of video width with a default of 5000
const MaxVideoWidth = 5000

// MinVideoWidth stand for the min value of video width with a default of 1
const MinVideoWidth = 1

// MaxVideoHeight stand for the max video height with a default of 5000
const MaxVideoHeight = 5000

// MinVideoHeight stand for the min video height with a default 1
const MinVideoHeight = 1

// SupportedMineTypes is used for default supported video format
var SupportedMineTypes = []string{"video/mp4"}
