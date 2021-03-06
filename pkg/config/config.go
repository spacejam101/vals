package config

import (
	"fmt"
	"github.com/variantdev/vals/pkg/api"
)

type MapConfig struct {
	M map[string]interface{}
}

var _ api.StaticConfig = &MapConfig{}

func (m MapConfig) String(path ...string) string {
	var cur interface{}

	cur = m.M

	for i, k := range path {
		switch typed := cur.(type) {
		case map[string]interface{}:
			cur = typed[k]
		case map[interface{}]interface{}:
			cur = typed[k]
		default:
			return ""
		}
		if i == len(path)-1 {
			if cur == nil {
				return ""
			}
			return fmt.Sprintf("%v", cur)
		}
	}

	panic("invalid state")
}

func (m MapConfig) StringSlice(path ...string) []string {
	var cur interface{}

	cur = m.M

	for i, k := range path {
		switch typed := cur.(type) {
		case map[string]interface{}:
			cur = typed[k]
		case map[interface{}]interface{}:
			cur = typed[k]
		default:
			return nil
		}
		if i == len(path)-1 {
			if cur == nil {
				return nil
			}
			switch ary := cur.(type) {
			case []string:
				return ary
			case []interface{}:
				ss := make([]string, len(ary))
				for i := range ary {
					ss[i] = fmt.Sprintf("%v", ary[i])
				}
				return ss
			default:
				panic(fmt.Errorf("unexpected type: value=%v, type=%T", ary, ary))
			}
		}
	}

	panic("invalid state")
}

func (m MapConfig) Config(path ...string) api.StaticConfig {
	return Map(m.Map(path...))
}

func (m MapConfig) Exists(path ...string) bool {
	var cur interface{}
	var ok bool

	cur = m.M

	for _, k := range path {
		switch typed := cur.(type) {
		case map[string]interface{}:
			cur, ok = typed[k]
			if !ok {
				return false
			}
		case map[interface{}]interface{}:
			cur, ok = typed[k]
			if !ok {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func (m MapConfig) Map(path ...string) map[string]interface{} {
	var cur interface{}

	cur = m.M

	for _, k := range path {
		switch typed := cur.(type) {
		case map[string]interface{}:
			cur = typed[k]
		case map[interface{}]interface{}:
			cur = typed[k]
		default:
			return nil
		}
	}

	switch typed := cur.(type) {
	case map[string]interface{}:
		return typed
	case map[interface{}]interface{}:
		strmap := map[string]interface{}{}
		for k, v := range typed {
			strmap[fmt.Sprintf("%v", k)] = v
		}
		return strmap
	default:
		return nil
	}

	panic("invalid state")
}

func Map(m map[string]interface{}) MapConfig {
	return MapConfig{
		M: m,
	}
}
