package internal

import (
	"errors"
	"strconv"
	"strings"
)

type MemStorage struct {
	Values map[string]float64
}

func Length(c []string) func(min int) bool {
	l := len(c)
	return func(min int) bool {
		return l >= min
	}
}

func TrimF(s string, prefix string, end string) (purified string) {
	s = strings.TrimPrefix(s, prefix)
	purified = strings.TrimRight(s, end)
	return
}

func ContainsF(carr []string) func(t string) bool {
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

var HasKnownTypes = ContainsF(KnownTypes)

func canParse(s string, v *float64) bool {
	var err error
	*v, err = strconv.ParseFloat(s, 64)
	if err == nil {
		return true
	}
	return false
}

func (s MemStorage) SaveMetric(URL string) error {
	URL = TrimF(URL, "/update/", "/")
	c := strings.Split(URL, "/")
	lnF := Length(c)
	len1 := lnF(1)
	len2 := lnF(2)
	len3 := lnF(3)
	typeIsKnown := HasKnownTypes(c[0])
	var v float64
	switch true {
	case len3 && typeIsKnown && canParse(c[2], &v):
		println(v)
		return nil
	case len3 && typeIsKnown:
		return ErrMap[400]
	case len2 && typeIsKnown:
		return ErrMap[404]
	case len2 && c[0] != "":
		return ErrMap[501]
	case len1:
		return ErrMap[404]
	}
	return nil
}
