package article

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetCommentsSuite struct {
	TestSuite
}

func (s *GetCommentsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetCommentsSuite) TeardownTest() {
	s.TestSuite.TeardownTest()
}

func (s *GetCommentsSuite) TestGetComments() {
	s.mockedRepo.On("GetComments", mock.Anything).Return(nil)

	comments, err := s.interactor.Article.GetComments("1234-3242-4354")
	require.NotNil(s.T(), comments)
	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedRepo.AssertExpectations(s.T()))
}

func TestGetCommentsSuite(t *testing.T) {
	suite.Run(t, new(GetCommentsSuite))
}
