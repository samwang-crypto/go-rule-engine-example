package models

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/monacohq/go-rule-engine-example/internal/helper"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Result struct {
	mux   sync.RWMutex
	Value map[string]interface{} `json:"value"`
}

func (r *Result) UpdateRestrictions(entityId string, withFiat bool, source map[string][]ForbiddenFeature) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if restriction, ok := source[entityId]; ok {
		for _, feature := range restriction {
			if !withFiat && helper.Contains(fiatFeatures, feature.Feature) {
				continue
			}
			r.Value[feature.Feature] = false
		}
	}
}

func (r *Result) UpdateRestrictionsWithRequiredInfo(entityId string, requiredInfo string, source map[string][]ForbiddenFeature) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if restriction, ok := source[entityId]; ok {
		for _, feature := range restriction {
			if feature.RequiredPersonalInformation == requiredInfo {
				delete(r.Value, feature.Feature)
			}
		}
	}
}

func (r *Result) UpdateRestrictionsWithAppliedAt(entityId string, source map[string][]ForbiddenFeature) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if restriction, ok := source[entityId]; ok {
		for _, feature := range restriction {
			if helper.Contains(skipForAppliedAfter, feature.Feature) {
				delete(r.Value, feature.Feature)
			}
		}
	}
}

// since we are using interface{} and generate results dynamically, we need to
// provide a casting function to convert the interface{} to the proper type
func (r *Result) BooleanEqual(key string, value bool) bool {
	r.mux.RLock()
	defer r.mux.RUnlock()
	if val, ok := r.Value[key].(bool); ok {
		return val == value
	}
	return false
}

// ToJson converts the result to a json string and can be stored in the database as
// static user config data.
func (r *Result) ToJson() string {
	r.mux.RLock()
	defer r.mux.RUnlock()
	jsonStr, err := json.Marshal(r.Value)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}
