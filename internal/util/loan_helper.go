package util

func GenerateBasicInterest(otr, interestRate, timeOccuring, admin float64) float64 {
	return (otr * interestRate * timeOccuring) + admin
}
