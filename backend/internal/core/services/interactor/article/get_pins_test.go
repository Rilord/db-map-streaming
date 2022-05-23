package article

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetPinsSuite struct {
	TestSuite
}

func (s *GetPinsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetPinsSuite) TeardownTest() {
	s.TestSuite.TeardownTest()
}

func (s *GetPinsSuite) TestGetPins() {
	s.mockedRepo.On("GetPins", mock.Anything).Return(nil)

	pins, err := s.interactor.Article.GetPins("1234-3242-4354")
	require.NotNil(s.T(), pins)
	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetPinsSuite(t *testing.T) {
	suite.Run(t, new(GetPinsSuite))
}
