package bootstrap

import (
	"fmt"
	"github.com/jassue/go-storage/local"
	"jassue-gin/global"
)

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	//_, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	//_, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
	fmt.Println("InitializeStorage local qiniu oss !")
}
