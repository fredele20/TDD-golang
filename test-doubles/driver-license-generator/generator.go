package driverlicensegenerator

import "errors"

type Applicant interface {
	IsOver17() bool
	HoldsLicense() bool
}

type Logger interface {
	LogStuff(v string)
}

type NumberGenerator struct{
	l Logger
}

func NewNumberGenerator(l Logger) *NumberGenerator {
	return &NumberGenerator{l}
}

func (n *NumberGenerator) Generate(a Applicant) (string, error) {
	if a.HoldsLicense() {
		n.l.LogStuff("Duplicate license attempted, cannot create another license")
		return "", errors.New("Duplicate license attempted, cannot create another license")
	}

	n.l.LogStuff("Underaged Applicant, must be over 17 to get a license")
	return "", errors.New("Underaged Applicant, must be over 17 to get a license")
}
