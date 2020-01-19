package main

// BoolToFloat64 converts bools to their 0, 1 float representation
func BoolToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
