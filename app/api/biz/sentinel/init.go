package sentinel

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
)

func Init() {
	sentinel.InitDefault()
	loadRules()
}
