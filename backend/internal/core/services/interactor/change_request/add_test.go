package changerequest

import (
	"game-server/internal/core/domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AddSuite struct {
	TestSuite
}

func (s *AddSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *AddSuite) TeardownTest() {
	s.TestSuite.SetupTest()
}

func (s *AddSuite) TestAdd() {
	m := domain.NewArticle("", nil)

	s.mockedRepo.On("Add", mock.Anything).Return(nil)

	err := s.interactor.Add(*m)

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
