package common

import (
	"time"
)

type Tracker struct {
	user      *XUser
	startTime time.Time
	traceId   string
	agent     string
}

func NewTracker(user *XUser, startTime time.Time, traceId, agent string) *Tracker {
	return &Tracker{
		user:      user,
		startTime: startTime,
		traceId:   traceId,
		agent:     agent,
	}
}

func NewSimpleTracker(traceId string) *Tracker {
	return &Tracker{
		traceId: traceId,
	}
}

func (t *Tracker) GetTraceId() string {
	return t.traceId
}

func (t *Tracker) GetUser() *XUser {
	return t.user
}

func (t *Tracker) GetAgent() string {
	return t.agent
}

func (t *Tracker) GetStartTime() time.Time {
	return t.startTime
}
