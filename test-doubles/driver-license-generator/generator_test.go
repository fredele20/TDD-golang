package driverlicensegenerator_test

import (
	driverlicensegenerator "go-tdd/test-doubles/driver-license-generator"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DrivingLicenseSuite struct {
	suite.Suite
}

func (s *DrivingLicenseSuite) TestUnderageApplicant() {
	a := UnderageApplicant{}
	l := &SpyLogger{}
	r := FakeRand{}

	lg := driverlicensegenerator.NewNumberGenerator(l, r)
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Underaged")

	s.Equal(1, l.callCount)
	s.Contains(l.lastMessage, "Underaged")
}

func (s *DrivingLicenseSuite) TestNoSecondLicense() {
	a := LicenseHolderApplicant{}
	l := &SpyLogger{}
	r := FakeRand{}

	lg := driverlicensegenerator.NewNumberGenerator(l, r)
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Duplicate")

	s.Equal(1, l.callCount)
	s.Contains(l.lastMessage, "Duplicate")
}

func (s *DrivingLicenseSuite) TestLicenseGenerator() {
	a := ValidApplicant{"MDB", "04102023"}
	l := &SpyLogger{}
	r := FakeRand{}

	lg := driverlicensegenerator.NewNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("MDB0410202300000", ln)
}

func (s *DrivingLicenseSuite) TestLicenseGeneratorShorterInitials() {
	a := ValidApplicant{"MB", "04102023"}
	l := &SpyLogger{}
	r := FakeRand{}

	lg := driverlicensegenerator.NewNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("MB04102023000000", ln)
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

type ValidApplicant struct {
	initials string
	dob      string
}

func (v ValidApplicant) IsOver17() bool {
	return true
}

func (v ValidApplicant) HoldsLicense() bool {
	return false
}

func (v ValidApplicant) GetInitials() string {
	return v.initials
}

func (v ValidApplicant) GetDOB() string {
	return v.dob
}

type UnderageApplicant struct{}

// GetDOB implements driverlicensegenerator.Applicant.
func (u UnderageApplicant) GetDOB() string {
	return ""
}

// getInitials implements driverlicensegenerator.Applicant.
func (u UnderageApplicant) GetInitials() string {
	return ""
}

func (u UnderageApplicant) IsOver17() bool {
	return false
}

func (u UnderageApplicant) HoldsLicense() bool {
	return false
}

type LicenseHolderApplicant struct{}

// GetDOB implements driverlicensegenerator.Applicant.
func (l LicenseHolderApplicant) GetDOB() string {
	return ""
}

func (l LicenseHolderApplicant) IsOver17() bool {
	return true
}

func (l LicenseHolderApplicant) HoldsLicense() bool {
	return true
}

func (l LicenseHolderApplicant) GetInitials() string {
	return ""
}

type FakeRand struct {}

func (f FakeRand) GetRandomNumbers(len int) string {
	return strings.Repeat("0", len)
}

type SpyLogger struct {
	callCount   int
	lastMessage string
}

func (s *SpyLogger) LogStuff(v string) {
	s.callCount++
	s.lastMessage = v
}
