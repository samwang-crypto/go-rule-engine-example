package models

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Result struct {
	Value map[string]interface{} `json:"value"`
}

// since we are using interface{} and generate results dynamically, we need to
// provide a casting function to convert the interface{} to the proper type
func (r *Result) BooleanEqual(key string, value bool) bool {
	if val, ok := r.Value[key].(bool); ok {
		return val == value
	}
	return false
}

// ToJson converts the result to a json string and can be stored in the database as
// static user config data.
func (r *Result) ToJson() string {
	jsonStr, err := json.Marshal(r.Value)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}
