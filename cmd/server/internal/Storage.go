package internal

import (
	"errors"
	"strconv"
	"strings"
)

type MemStorage struct {
	Values map[string]float64
}

func InitStorage() MemStorage {
	return MemStorage{Values: make(map[string]float64)}
}

func length(c []string) func(min int) bool {
	l := len(c)
	return func(min int) bool {
		return l >= min
	}
}

func trim(s string, prefix string, end string) (purified string) {
	s = strings.TrimPrefix(s, prefix)
	purified = strings.TrimRight(s, end)
	return
}

func contains(carr []string) func(t string) bool {
	return func(t string) bool {
		for _, v := range carr {
			if v == t {
				return true
			}
		}
		return false
	}
}

type ErrorCollection struct {
	Errors []error
}

var ErrMap = map[int]error{
	400: errors.New("value not parsable"),
	404: errors.New("no value"),
	501: errors.New("no known type"),
}

var KnownTypes = []string{"counter", "gauge"}

var IsItKnownType = contains(KnownTypes)

func canParse(s string, v *float64) bool {
	var err error
	*v, err = strconv.ParseFloat(s, 64)
	if err == nil {
		return true
	}
	return false
}

func ProcessURL(URL string) (value float64, name string, err error) {
	trimmedURL := trim(URL, "/update/", "/")
	dividedURL := strings.Split(trimmedURL, "/")
	isUrlLengthMoreOrEqual := length(dividedURL)
	LenMoreOrEqual1 := isUrlLengthMoreOrEqual(1)
	LenMoreOrEqual2 := isUrlLengthMoreOrEqual(2)
	LenMoreOrEqual3 := isUrlLengthMoreOrEqual(3)
	typeIsKnown := IsItKnownType(dividedURL[0])
	switch true {
	case LenMoreOrEqual3 && typeIsKnown && canParse(dividedURL[2], &value):
		return value, dividedURL[1], nil
	case LenMoreOrEqual3 && typeIsKnown:
		return value, dividedURL[1], ErrMap[400]
	case LenMoreOrEqual2 && typeIsKnown:
		return value, dividedURL[1], ErrMap[404]
	case LenMoreOrEqual2 && dividedURL[0] != "":
		return value, dividedURL[1], ErrMap[501]
	case LenMoreOrEqual1:
		return value, name, ErrMap[404]
	}
	return value, name, nil
}

func (s MemStorage) SaveMetric(value float64, name string) error {
	s.Values[name] = value
	return nil
}
