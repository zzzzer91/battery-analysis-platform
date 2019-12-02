package service

import "battery-analysis-platform/pkg/jd"

type JsonServicer interface {
	Do() (*jd.Response, error)
}
