package config

type logging struct {
	Gin gin
}

type ginLog struct {
	Path string
}

type gin struct {
	Log ginLog
}
