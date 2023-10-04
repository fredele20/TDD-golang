package driverlicensegenerator_test

import (
	driverlicensegenerator "go-tdd/test-doubles/driver-license-generator"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DrivingLicenseSuite struct {
	suite.Suite
}

func (s *DrivingLicenseSuite) TestUnderageApplicant() {
	a := UnderageApplicant{}
	l := &SpyLogger{}

	lg := driverlicensegenerator.NewNumberGenerator(l)
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Underaged")

	s.Equal(1, l.callCount)
	s.Contains(l.lastMessage, "Underaged")
}

func (s *DrivingLicenseSuite) TestNoSecondLicense() {
	a := LicenseHolderApplicant{}
	l := &SpyLogger{}

	lg := driverlicensegenerator.NewNumberGenerator(l)
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Duplicate")

	s.Equal(1, l.callCount)
	s.Contains(l.lastMessage, "Duplicate")
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

type UnderageApplicant struct{}

func (u UnderageApplicant) IsOver17() bool {
	return false
}

func (u UnderageApplicant) HoldsLicense() bool {
	return false
}

type LicenseHolderApplicant struct{}

func (l LicenseHolderApplicant) IsOver17() bool {
	return true
}

func (l LicenseHolderApplicant) HoldsLicense() bool {
	return true
}

type SpyLogger struct {
	callCount   int
	lastMessage string
}

func (s *SpyLogger) LogStuff(v string) {
	s.callCount++
	s.lastMessage = v
}
