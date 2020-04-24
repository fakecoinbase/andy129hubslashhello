package model

type Config struct {
	KafkaConfig `ini:"kafka"`
	ESConfig `ini:"es"`
}
type KafkaConfig struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
	ChanSize int64 `ini:"chan_size"`
}

type ESConfig struct {
	Address string `ini:"address"`
	Index string `ini:"index"`
	GoroutineNum int `ini:"goroutine_num"`
}
