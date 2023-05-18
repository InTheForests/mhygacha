package config

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"os"
)

var Config Proto

type Proto struct {
	OpenAuthKey bool `json:"open_auth_key"`
}

func ConfigInit(staticFiles embed.FS) {
	configPath := "./config.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Println("配置文件不存在，已生成配置文件!")
		configData, _ := fs.ReadFile(staticFiles, "static/config.json")
		err = os.WriteFile(configPath, configData, 0644)
		configNoteData, _ := fs.ReadFile(staticFiles, "static/configNote.txt")
		err = os.WriteFile("./configNote.txt", configNoteData, 0644)
	}
	configData, _ := os.ReadFile(configPath)
	err := json.Unmarshal(configData, &Config)
	if err != nil {
		log.Println("配置文件解析错误，程序可能无法正常运行....")
		return
	}
	log.Println("配置文件读取成功!")
}
