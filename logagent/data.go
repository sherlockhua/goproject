package main

type LogConfig struct {
	Topic string `json:"topic"`
	LogPath string `json:"log_path"`
	Service string `json:"service"`
	SendRate int `json:"send_rate"`
}