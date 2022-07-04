package models

import (
	"log"

	"github.com/monacohq/go-rule-engine-example/internal/helper"
)

var FiatFeatures = []string{"sepa_deposit", "au_npp_deposit", "au_bpay_deposit", "ca_interac_etransfer_deposit"}

// This is just temporary, it will be replaced by whitelist
type ForbiddenFeature struct {
	Feature                     string `json:"feature"`
	RequiredPersonalInformation string `json:"required_personal_information"`
}

type EntityRestrictions struct {
	Value map[string][]ForbiddenFeature `json:"value"`
}

func (e *EntityRestrictions) GetRestrictions(entityId string, withFiat bool) []string {
	if restriction, ok := e.Value[entityId]; ok {
		var features []string
		for _, feature := range restriction {
			features = append(features, feature.Feature)
		}
		if withFiat {
			return features
		} else {
			return helper.Difference(features, FiatFeatures)
		}
	}
	return nil
}

func (e *EntityRestrictions) GetRestrictionsWithPersonalInformationRequired(entityId string, requiredInfo string) []string {
	if restriction, ok := e.Value[entityId]; ok {
		var features []string
		for _, feature := range restriction {
			if feature.RequiredPersonalInformation != requiredInfo {
				features = append(features, feature.Feature)
			}
		}
		log.Printf("---- %+v\n", features)
		return features
	}
	return nil
}
