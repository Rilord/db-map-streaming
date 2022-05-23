package article

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetChangeRequestSuite struct {
	TestSuite
}

func (s *GetChangeRequestSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetChangeRequestSuite) TeardownTest() {
	s.TestSuite.TeardownTest()
}

func (s *GetChangeRequestSuite) TestGetChangeRequest() {
	s.mockedRepo.On("GetChangeRequest", mock.Anything).Return(nil)

	requests, err := s.interactor.Article.GetChangeRequests("1234-3242-4354")
	require.NotNil(s.T(), requests)
	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetChangeRequestSuite(t *testing.T) {
	suite.Run(t, new(GetChangeRequestSuite))
}
