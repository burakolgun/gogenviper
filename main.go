package gogenviper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type CfgFileWatcher[T any] struct {
	viper *viper.Viper
	model T
}

func Init[T any](path string, fileName string, fileType string, model T) (*CfgFileWatcher[T], error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(model)

	if err != nil {
		return nil, err
	}

	return &CfgFileWatcher[T]{viper: v, model: model}, nil
}

func (w CfgFileWatcher[T]) refresh() error {
	err := w.viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = w.viper.Unmarshal(w.model)

	if err != nil {
		return err
	}

	return nil
}

func (w CfgFileWatcher[T]) Watch() {
	w.viper.WatchConfig()

	w.viper.OnConfigChange(func(in fsnotify.Event) {
		err := w.refresh()
		if err != nil {
			log.Println("Error on refreshing application toggles due to file change, error: ", err)
			return
		}
		log.Println(fmt.Sprintf("Application configuration %T file changed", w.model))
	})
}
