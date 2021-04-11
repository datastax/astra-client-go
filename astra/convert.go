package astra

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
