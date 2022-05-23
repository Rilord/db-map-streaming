package article

import (
	"game-server/internal/core/domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetSuite struct {
	TestSuite
}

func (s *GetSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetSuite) TeardownTest() {
	s.TestSuite.TeardownTest()
}

func (s *GetSuite) TestGet() {
	id := ""
	expected := domain.NewArticle(&id, nil, nil, nil, nil)
	s.mockedRepo.On("Get", mock.Anything).Return(*expected, nil)

	art, err := s.interactor.Get(*expected.ID)

	require.NoError(s.T(), err)
	require.Equal(s.T(), art, *expected)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}
