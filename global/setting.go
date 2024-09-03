package global

import (
	"GPT_cli/setting"
	"log"
)

var (
	TokenSetting *setting.Settings
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Models", &TokenSetting)
	if err != nil {
		return err
	}
	return nil
}
