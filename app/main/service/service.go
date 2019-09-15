package service

import "battery-analysis-platform/pkg/jd"

type Servicer interface {
	Do() (*jd.Response, error)
}
