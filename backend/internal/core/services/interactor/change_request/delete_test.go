package changerequest

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DeleteSuite struct {
	TestSuite
}

func (s *DeleteSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *DeleteSuite) TeardownTest() {
	s.TestSuite.SetupTest()
}

func (s *DeleteSuite) TestDelete() {

	requestId := ""
	s.mockedRepo.On("Delete", mock.Anything).Return(nil)

	err := s.interactor.Delete(requestId)

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
