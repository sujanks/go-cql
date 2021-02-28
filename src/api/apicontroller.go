package api

import "github.com/sujanks/go-cql/src/model"

type Controller interface {
	Query(string) *model.Container
}
