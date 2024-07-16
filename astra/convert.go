package astra

import "encoding/json"

func BoolValue(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func Float64Value(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func StringSlice(s *[]string) []string {
	if s == nil {
		return nil
	}
	return *s
}

func AvailableRegionCombinationSlice(v *[]AvailableRegionCombination) []AvailableRegionCombination {
	if v == nil {
		return nil
	}
	return *v
}

func DatabaseSlice(v *[]Database) []Database {
	if v == nil {
		return nil
	}
	return *v
}

// all currently available versions of github.com/deepmap/oapi-codegen/cmd/oapi-codegen (up to 1.16)
// fail to handle 'oneOf' Swagger clause, parsing response manually
func ParseSecureBundles(response *GenerateSecureBundleURLResponse) (*SecureBundles, error) {
	var destObj CredsURL
	err := json.Unmarshal(response.Body, &destObj)
	if err == nil {
		return &SecureBundles{destObj}, nil
	} else {
		var destArr SecureBundles
		err = json.Unmarshal(response.Body, &destArr)
		if err == nil {
			return &destArr, nil
		}
	}
	return nil, err
}
