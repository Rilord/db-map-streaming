package changerequest

import (
	repo "game-server/internal/adapter/repository/mock"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	interactor *Interactor

	mockedRepo *repo.ChangeRequestRepository
}

func (s *TestSuite) SetupTest() {

	s.interactor = NewInteractor()
	require.NotNil(s.T(), s.interactor)

	s.mockedRepo = s.interactor.Request.(*repo.ChangeRequestRepository)
}

func (s *TestSuite) TeardownTest() {

}
