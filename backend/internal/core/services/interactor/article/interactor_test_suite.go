package article

import (
	repo "game-server/internal/adapter/repository/mock"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	interactor *Interactor

	mockedRepo *repo.ArticleRepository
}

func (s *TestSuite) SetupTest() {

	s.interactor = NewInteractor()
	require.NotNil(s.T(), s.interactor)

	s.mockedRepo = s.interactor.Article.(*repo.ArticleRepository)
}

func (s *TestSuite) TeardownTest() {

}
