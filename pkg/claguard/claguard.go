package claguard

// Check checks if the own claims contain any of the wanted claims
// Returns true if any of the wanted claims is found in the own claims
// Returns false if none of the wanted claims is found in the own claims
func Check(ownClaims []string, wantClaims []string) bool {
	for _, c := range wantClaims {
		if CheckSingle(ownClaims, c) {
			return true
		}
	}
	return false
}

// CheckSingle checks if the own claims contain the wanted claim
// Returns true if the wanted claim is found in the own claims
// Returns false if the wanted claim is not found in the own claims
func CheckSingle(ownClaims []string, wantClaim string) bool {
	for _, r := range ownClaims {
		if r == wantClaim {
			return true
		}
	}
	return false
}
