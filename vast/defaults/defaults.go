package defaults

import "time"

// MaxVideoDuration stands for the max video duration with a default of 180 seconds.
const MaxVideoDuration = 180 * time.Second

// MinVideoDuration stands for the min video duration with a default of 1 second.
const MinVideoDuration = 1 * time.Second

// MaxVideoWidth stands for the max value of video width with a default of 5000.
const MaxVideoWidth = 5000

// MinVideoWidth stands for the min value of video width with a default of 1.
const MinVideoWidth = 1

// MaxVideoHeight stands for the max video height with a default of 5000.
const MaxVideoHeight = 5000

// MinVideoHeight stands for the min video height with a default 1.
const MinVideoHeight = 1

// SupportedMimeTypes is used for default supported video format.
var SupportedMimeTypes = []string{"video/mp4", "application/x-mpegURL", "application/vnd.apple.mpegurl"}

// SupportedStreamingADSuffix is the only supported format suffix for streaming ad.
const SupportedStreamingADSuffix = ".m3u8"
