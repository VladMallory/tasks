package main

import (
	"encoding/json"
	"log"
	"os"
	"sync/atomic"
	"time"
)

type Config struct {
	DBHost string `json:"db_host"`
	DBPort int    `json:"db_port"`
	Debug  bool   `json:"debug"`
}

var config atomic.Value // хранит *Config

func Load(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var newCfg Config
	if err := json.Unmarshal(file, &newCfg); err != nil {
		return err
	}

	config.Store(&newCfg) // атомарная замена

	return nil
}

func Get() *Config {
	return config.Load().(*Config)
}

func main() {
	// начальная загрузка
	if err := Load("config.json"); err != nil {
		log.Fatalf("load error: %v", err)
	}
	// несколько читателей
	for i := range 5 {
		go func(id int) {
			for {
				cfg := Get()
				log.Printf("worker %d: %+v\n", id, cfg)
				time.Sleep(2 * time.Second)
			}
		}(i)
	}

	// имитация редкой перезагрузки
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if err := Load("config.json"); err != nil {
				log.Println("reload error:", err)
			}
		}
	}()

	select {} // блокируем main
}
