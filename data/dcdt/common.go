package dcdt

import "strings"

const (
	// dcdtTickerNumRandChars represents the number of hex-encoded random characters sequence of a ticker
	dcdtTickerNumRandChars = 6
	// separatorChar represents the character that separated the token ticker by the random sequence
	separatorChar = "-"
	// minLengthForTickerName represents the minimum number of characters a token's ticker can have
	minLengthForTickerName = 3
	// maxLengthForTickerName represents the maximum number of characters a token's ticker can have
	maxLengthForTickerName = 10
	// maxLengthDCDTPrefix represents the maximum number of characters a token's prefix can have
	maxLengthDCDTPrefix = 4
	// minLengthDCDTPrefix represents the minimum number of characters a token's prefix can have
	minLengthDCDTPrefix = 1
)

// IsValidPrefixedToken checks if the provided token is valid, and returns if prefix if so
func IsValidPrefixedToken(token string) (string, bool) {
	tokenSplit := strings.Split(token, separatorChar)
	if len(tokenSplit) < 3 {
		return "", false
	}

	prefix := tokenSplit[0]
	if !IsValidTokenPrefix(prefix) {
		return "", false
	}

	tokenTicker := tokenSplit[1]
	if !IsTickerValid(tokenTicker) {
		return "", false
	}

	tokenRandSeq := tokenSplit[2]
	if !(len(tokenRandSeq) >= dcdtTickerNumRandChars) {
		return "", false
	}

	return prefix, true
}

// IsValidTokenPrefix checks if the token prefix is valid
func IsValidTokenPrefix(prefix string) bool {
	prefixLen := len(prefix)
	if prefixLen > maxLengthDCDTPrefix || prefixLen < minLengthDCDTPrefix {
		return false
	}

	for _, ch := range prefix {
		isLowerCaseCharacter := ch >= 'a' && ch <= 'z'
		isNumber := ch >= '0' && ch <= '9'
		isAllowedChar := isLowerCaseCharacter || isNumber
		if !isAllowedChar {
			return false
		}
	}

	return true
}

// IsTickerValid checks if the token ticker is valid
func IsTickerValid(ticker string) bool {
	if !IsTokenTickerLenCorrect(len(ticker)) {
		return false
	}

	for _, ch := range ticker {
		isUpperCaseCharacter := ch >= 'A' && ch <= 'Z'
		isNumber := ch >= '0' && ch <= '9'
		isAllowedChar := isUpperCaseCharacter || isNumber
		if !isAllowedChar {
			return false
		}
	}

	return true
}

// IsTokenTickerLenCorrect checks if the token ticker len is correct
func IsTokenTickerLenCorrect(tokenTickerLen int) bool {
	return !(tokenTickerLen < minLengthForTickerName || tokenTickerLen > maxLengthForTickerName)
}
