package article

import (
	"game-server/internal/core/domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UpdateSuite struct {
	TestSuite
}

func (s *UpdateSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *UpdateSuite) TeardownTest() {
	s.TestSuite.TeardownTest()
}

func (s *UpdateSuite) TestUpdate() {

	expected := domain.NewArticle(nil, nil, nil, nil, nil)
	s.mockedRepo.On("Update", mock.Anything).Return(nil)

	err := s.interactor.Article.Update(*expected)
	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}
