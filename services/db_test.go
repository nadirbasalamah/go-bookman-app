package services

import (
	"context"
	"fmt"
	"go-bookman-app/models"
	"go-bookman-app/testhelpers"
	"log"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	service     *BookDBService
	ctx         context.Context
}

type BookRecord struct {
	Title       string `faker:"word"`
	Description string `faker:"sentence"`
	Publisher   string `faker:"name"`
}

func (suite *BookRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	pgContainer.DB.AutoMigrate(&models.Book{})

	suite.pgContainer = pgContainer
	service := InitBookDBService(suite.pgContainer.DB)
	suite.service = &service
}

func (suite *BookRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *BookRepoTestSuite) TestInsertBook() {
	t := suite.T()

	book, err := insertSampleBook(suite)

	assert.NoError(t, err)
	assert.NotZero(t, book.ID)
	assert.NotNil(t, book.Title)
	assert.NotNil(t, book.Description)
	assert.NotNil(t, book.Publisher)
}

func (suite *BookRepoTestSuite) TestGetBookByID() {
	t := suite.T()

	id := strconv.Itoa(1)

	_, err := insertSampleBook(suite)

	book, err := suite.service.GetBookByID(id)

	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.NotZero(t, book.ID)
	assert.NotNil(t, book.Title)
	assert.NotNil(t, book.Description)
	assert.NotNil(t, book.Publisher)
}

func TestBookRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepoTestSuite))
}

func insertSampleBook(suite *BookRepoTestSuite) (models.Book, error) {
	record := BookRecord{}
	err := faker.FakeData(&record)
	if err != nil {
		fmt.Println(err)
	}

	return suite.service.InsertBook(models.Book{
		Title:       record.Title,
		Description: record.Description,
		Publisher:   record.Publisher,
	})
}
