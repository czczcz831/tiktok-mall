package sentinel

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/util"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type stateChangeTestListener struct{}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	hlog.Errorf("rule.steategy: %+v,rule.resource: %+v, From %s to Closed, time: %d\n", rule.Strategy, rule.Resource, prev.String(), util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	hlog.Errorf("rule.steategy: %+v,rule.resource: %+v, From %s to Open, snapshot: %.2f, time: %d\n", rule.Strategy, rule.Resource, prev.String(), snapshot, util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	hlog.Errorf("rule.steategy: %+v,rule.resource: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, rule.Resource, prev.String(), util.CurrentTimeMillis())
}

func Init() {
	sentinel.InitDefault()
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})
	loadRules()
}
