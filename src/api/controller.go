package api

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/sujanks/go-cql/src/config"
	"github.com/sujanks/go-cql/src/model"
	"regexp"
	"strings"
)

type ApiController struct {
	session *gocql.Session
}

var instance *ApiController

type QueryModel struct {
	SelectFields string
	Keyspace     string
	Table        string
	Raw          string
}

const (
	QueryRegex = `^select(?P<select>.*)\sfrom\s(?P<kt>.*)?`
	LimitRegex = `limit\s\d{1,}`
)

func NewContoller() Controller {
	if instance == nil {
		session := config.InitCluster()
		instance = &ApiController{
			session: session,
		}
	}
	return instance
}

func (a ApiController) Query(s string) *model.Container {
	queryModel := Validate(s)
	return getResult(a.session, queryModel)
}

func getResult(session *gocql.Session, qModel *QueryModel) *model.Container {
	hiddenFields := [2]string{"amount", "dob"}
	sliceMap, _ := session.Query(qModel.Raw).Iter().SliceMap()
	resMap := make([]map[string]interface{}, 0)
	for _, m := range sliceMap {
		for _, v := range hiddenFields {
			if m[v] != nil {
				hideField := fmt.Sprintf("%v", m[v])
				if len(hideField) > 0 {
					m[v] = "***"
				}
			}
		}
		resMap = append(resMap, m)
	}
	return &model.Container{
		Data: resMap,
	}
}

func Validate(s string) *QueryModel {
	qString := strings.ToLower(s)
	regExp, _ := regexp.Compile(QueryRegex)
	match := regExp.FindStringSubmatch(qString)
	result := make(map[string]string)
	if len(match) > 0 {
		for i, name := range regExp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		keyTable := strings.Split(result["kt"], ".")
		return &QueryModel{
			SelectFields: result["select"],
			Keyspace:     keyTable[0],
			Table:        keyTable[1],
			Raw:          s,
		}
	}
	return nil

}
