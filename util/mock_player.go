package util

func RandomNamePlayer() string {
	return RandomString(10)
}

func RandomPasswordPlayer() string {
	return RandomString(10)
}

func RandomEmailPlayer() string {
	return RandomString(10) + "@gmail.com"
}

func RandomCashPlayer() string {
	return "1000.00"
}

func RandomFullNamePlayer() string {
	return RandomString(10)
}
