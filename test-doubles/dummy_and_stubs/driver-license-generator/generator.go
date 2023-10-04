package driverlicensegenerator

import "errors"


type Applicant interface {
	IsOver17() bool
	HoldsLicense() bool
}

type NumberGenerator struct {}

func NewNumberGenerator() *NumberGenerator {
	return &NumberGenerator{}
}

func (n *NumberGenerator) Generate(a Applicant) (string, error) {
	if a.HoldsLicense() {
		return "", errors.New("Duplicate license attempted, cannot create another license")
	}
	return "", errors.New("Underaged Applicant, must be over 17 to get a license")
}