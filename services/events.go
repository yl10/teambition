package services

import (
	"encoding/json"
	"time"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

// DefaultEventsService 默认 TBEvents 接口实现
type DefaultEventsService struct {
	cli    *util.HttpClient
	apiURL string
}

// NewDefaultEventsService 创建实例
func NewDefaultEventsService(cli *util.HttpClient, apiURL string) (*DefaultEventsService, error) {
	events := &DefaultEventsService{cli: cli, apiURL: apiURL}
	return events, nil
}

// AddEvents 新增日程
func (t *DefaultEventsService) AddEvents(newEvents *model.NewEvents) (*model.Events, error) {
	reqURL := t.apiURL + "api/events"
	res, err := t.cli.Post(reqURL, newEvents)
	if err != nil {
		return nil, err
	}

	var events model.Events
	if err = json.Unmarshal(res, &events); err != nil {
		return nil, util.TransErr2XErr(err)
	}

	return &events, nil
}

// DeleteEvents 删除日程
func (t *DefaultEventsService) DeleteEvents(eventsID string, occurrenceDate ...time.Time) error {
	if eventsID == "" {
		return util.NewError("删除日程失败, 没有日程ID")
	}

	reqURL := t.apiURL + "api/events/" + eventsID

	var params map[string]string

	if len(occurrenceDate) > 0 {
		params = make(map[string]string)
		params["occurrenceDate"] = occurrenceDate[0].Format(time.RFC3339)
	}

	_, err := t.cli.Delete(reqURL, params)
	return err
}
