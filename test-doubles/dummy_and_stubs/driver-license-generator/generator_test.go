package driverlicensegenerator_test

import (
	driverlicensegenerator "go-tdd/test-doubles/dummy_and_stubs/driver-license-generator"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DrivingLicenseSuite struct {
	suite.Suite
}

func (s *DrivingLicenseSuite) TestUnderageApplicant() {
	a := UnderageApplicant{}

	lg := driverlicensegenerator.NewNumberGenerator()
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Underaged")
}

func (s *DrivingLicenseSuite) TestNoSecondLicense() {
	a := LicenseHolderApplicant{}

	lg := driverlicensegenerator.NewNumberGenerator()
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Duplicate")
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

type UnderageApplicant struct {}

func (u UnderageApplicant) IsOver17() bool {
	return false
}

func (u UnderageApplicant) HoldsLicense() bool {
	return false
}

type LicenseHolderApplicant struct {}

func (l LicenseHolderApplicant) IsOver17() bool {
	return true
}

func (l LicenseHolderApplicant) HoldsLicense() bool {
	return true
}
