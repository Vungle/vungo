package defaults

import "time"

const MAX_VIDEO_DURATION = 60 * time.Second
const MIN_VIDEO_DURATION = 1 * time.Second

const MAX_VIDEO_WIDTH = 5000
const MIN_VIDEO_WIDTH = 1

const MAX_VIDEO_HEIGHT = 5000
const MIN_VIDEO_HEIGHT = 1

var SUPPORTED_MIME_TYPES = []string{"video/mp4"}
