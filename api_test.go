package api

import (
	"net/http"
	"os"
	"testing"

	"github.com/PGITAb/an-example-http-api/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var BASE_URL string

var client *http.Client
var db *gorm.DB

func TestMain(m *testing.M) {

	// load env
	godotenv.Load()
	BASE_URL = os.Getenv("BASE_URL")
	client = &http.Client{}

	// run tests
	exitVal := m.Run()

	// clean up

	// exit tests
	os.Exit(exitVal)
}

func TestDBConnection(t *testing.T) {

	dsn := os.Getenv("POSTGRES_DSN")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	t.Log("connected to database:")
}

func TestSetupDB(t *testing.T) {
	db.AutoMigrate(&model.Album{})
	albums := []model.Album{
		{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	db.Create(&albums)
}

func TestGetAlbums(t *testing.T) {

	url := BASE_URL + `/ablums`
	res, err := client.Get(url)
	if err != nil {
		t.Fatalf(`%v error = %v`, t.Name(), err)
	}

	assert.Equal(t, 200, res.Status)
}
