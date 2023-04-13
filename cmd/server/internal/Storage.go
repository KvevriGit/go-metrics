package internal

import (
	"errors"
	"strconv"
	"strings"
)

type MemStorage struct {
	Values map[string]float64
}

func (s MemStorage) SaveMetric(URL string) error {
	c := strings.Split(URL, "/")
	var err error
	switch c[2] {
	case "gauge":
		s.Values[c[3]], err = strconv.ParseFloat(c[4], 64)
		return err
	case "counter":
		v, err := strconv.ParseFloat(c[4], 64)
		s.Values[c[3]] += v
		return err
	case "":
		return errors.New("no name")
	}
	return nil
}
