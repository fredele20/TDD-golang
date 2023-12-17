package driverlicensegenerator_test

import (
	driverlicensegenerator "go-tdd/test-doubles/driver-license-generator"
	"go-tdd/test-doubles/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type DrivingLicenseSuite struct {
	suite.Suite
}

var ctrl *gomock.Controller
var a *mocks.MockApplicant
var l *mocks.MockLogger
var r *mocks.MockRandomNumberGenerator
var lg driverlicensegenerator.NumberGenerator

func (s *DrivingLicenseSuite) SetupTest() {
	ctrl = gomock.NewController(s.T())

	a = mocks.NewMockApplicant(ctrl)
	l = mocks.NewMockLogger(ctrl)
	r = mocks.NewMockRandomNumberGenerator(ctrl)
	lg = *driverlicensegenerator.NewNumberGenerator(l, r)
}

func (s *DrivingLicenseSuite) TestUnderageApplicant() {
	a.EXPECT().IsOver17().Return(false)
	a.EXPECT().HoldsLicense().Return(false)

	l.EXPECT().LogStuff("Underaged Applicant, must be over 17 to get a license").Times(1)

	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Underaged")

}

func (s *DrivingLicenseSuite) TestNoSecondLicense() {
	a.EXPECT().HoldsLicense().Return(true)
	l.EXPECT().LogStuff("Duplicate license attempted, cannot create another license").Times(1)

	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Duplicate license attempted, cannot create another")
}

func (s *DrivingLicenseSuite) TestLicenseGenerator() {
	// a := ValidApplicant{"MDB", "04102023"}
	a.EXPECT().HoldsLicense().Return(false)
	a.EXPECT().IsOver17().Return(true)
	a.EXPECT().GetInitials().Return("MDB")
	a.EXPECT().GetDOB().Return("04102023")

	r.EXPECT().GetRandomNumbers(5).Return("00000")

	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("MDB0410202300000", ln)
}

func (s *DrivingLicenseSuite) TestLicenseGeneratorShorterInitials() {
	// a := ValidApplicant{"MB", "04102023"}
	a.EXPECT().HoldsLicense().Return(false)
	a.EXPECT().IsOver17().Return(true)
	a.EXPECT().GetInitials().Return("MB")
	a.EXPECT().GetDOB().Return("04102023")

	r.EXPECT().GetRandomNumbers(6).Return("000000")

	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("MB04102023000000", ln)
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}