package article

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
	s.TestSuite.TeardownTest()
}

func (s *AddSuite) TestAdd() {
	expected := domain.NewArticle(nil, nil, nil, nil, nil)
	s.mockedRepo.On("Add", mock.Anything).Return(*expected, nil)

	id, err := s.interactor.Article.Add(*expected)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), id)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
