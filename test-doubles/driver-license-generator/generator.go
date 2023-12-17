package driverlicensegenerator

import (
	"errors"
	"fmt"
)

//go:generate mockgen -source=generator.go --destination=../mocks/generator.go -package=mocks
type Applicant interface {
	IsOver17() bool
	HoldsLicense() bool
	GetInitials() string
	GetDOB() string
}

type Logger interface {
	LogStuff(v string)
}

type RandomNumberGenerator interface {
	GetRandomNumbers(len int) string
}

type NumberGenerator struct{
	l Logger
	r RandomNumberGenerator
}

func NewNumberGenerator(l Logger, r RandomNumberGenerator) *NumberGenerator {
	return &NumberGenerator{l, r}
}

func (g *NumberGenerator) Generate(a Applicant) (string, error) {
	if a.HoldsLicense() {
		g.l.LogStuff("Duplicate license attempted, cannot create another license")
		return "", errors.New("Duplicate license attempted, cannot create another license")
	}

	if !a.IsOver17() {
		g.l.LogStuff("Underaged Applicant, must be over 17 to get a license")
		return "", errors.New("Underaged Applicant, must be over 17 to get a license")
	}

	n := fmt.Sprintf("%s%s", a.GetInitials(), a.GetDOB())

	num := 16 - len(n)

	return fmt.Sprintf("%s%s", n, g.r.GetRandomNumbers(num)), nil
}
