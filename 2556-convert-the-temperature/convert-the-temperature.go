func convertTemperature(celsius float64) []float64 {
	ans := make([]float64, 2)
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00

	ans[0] = kelvin
	ans[1] = fahrenheit
	return ans
}