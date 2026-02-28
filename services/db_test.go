package services

import (
	"context"
	"go-bookman-app/models"
	"go-bookman-app/testhelpers"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	service     *BookDBService
	ctx         context.Context
}

func (suite *BookRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	service := InitBookDBService(suite.pgContainer.DB)
	//TODO: migrate DB
	suite.service = &service
}

func (suite *BookRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *BookRepoTestSuite) TestInsertBook() {
	t := suite.T()

	//TODO: use faker
	book, err := suite.service.InsertBook(models.Book{
		ID:          1,
		Title:       "sample book",
		Description: "book description",
		Publisher:   "sample publisher",
	})

	assert.NoError(t, err)
	assert.NotNil(t, book.ID)
}

func (suite *BookRepoTestSuite) TestGetBookByID() {
	t := suite.T()

	book, err := suite.service.GetBookByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, book.ID, 1)
	assert.Equal(t, book.Title, "sample book")
	assert.Equal(t, book.Description, "book description")
	assert.Equal(t, book.Publisher, "sample publisher")
}

func TestBookRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepoTestSuite))
}
