package graph

import (
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/juanfgs/canvas/graph/generated"
	database "github.com/juanfgs/canvas/internal/pkg/db/postgresql"
	"github.com/stretchr/testify/suite"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
	"os"
	"testing"
)

var Cleaner = dbcleaner.New()

type CanvasSuite struct {
	suite.Suite
	c *client.Client
}

func (suite *CanvasSuite) SetupSuite() {
	dbUsername := os.Getenv("DB_TEST_USERNAME")
	dbPassword := os.Getenv("DB_TEST_PASSWORD")
	dbHost := os.Getenv("DB_TEST_HOST")
	dbName := os.Getenv("DB_TEST_NAME")
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", dbUsername, dbPassword, dbHost, dbName)
	postgresql := engine.NewPostgresEngine(dsn)
	Cleaner.SetEngine(postgresql)
	database.InitDB()
	database.Migrate()
	suite.c = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}})))

}

func (suite *CanvasSuite) SetupTest() {
	Cleaner.Acquire("canvases")
}

func (suite *CanvasSuite) TearDownTest() {
	Cleaner.Clean("canvases")
}

func (suite *CanvasSuite) TestCreateCanvas() {

	var resp struct {
		CreateCanvas struct {
			Name string
		}
	}
	suite.c.MustPost(`mutation { createCanvas(input: {name: "new canvas"}){ name }}`, &resp)
	suite.Equal("new canvas", resp.CreateCanvas.Name)
}

func (suite *CanvasSuite) TestAddShape() {

	var createCanvasResponse struct {
		CreateCanvas struct {
			ID string
		}
	}
	suite.c.MustPost(`mutation { createCanvas(input: {name: "new canvas"}){ id }}`, &createCanvasResponse)
	var addShapeResponse struct {
		AddShape struct {
			ID       string
			Name     string
			Contents []struct {
				X       float64
				Y       float64
				Width   float64
				Height  float64
				Fill    string
				Outline string
			}
		}
	}
	suite.c.MustPost(fmt.Sprintf(`mutation {
  addShape(input: {
    canvasId:"%s"
    x: 11,
    y: 3,
    width:1,
    height:3,
    fill:"B",
    outline:"L",
  }){
    id,
    name,
    contents {
      x,
      y,
      width,
      height,
      fill,
      outline
    }
  }
}`, createCanvasResponse.CreateCanvas.ID), &addShapeResponse)
	suite.Equal(float64(11), addShapeResponse.AddShape.Contents[0].X)
	suite.Equal(float64(3), addShapeResponse.AddShape.Contents[0].Y)
	suite.Equal(float64(1), addShapeResponse.AddShape.Contents[0].Width)
	suite.Equal(float64(3), addShapeResponse.AddShape.Contents[0].Height)
	suite.Equal("B", addShapeResponse.AddShape.Contents[0].Fill)
	suite.Equal("L", addShapeResponse.AddShape.Contents[0].Outline)

}

func (suite *CanvasSuite) TestQueryCanvases() {
	var queryCanvasesResponse struct {
		Canvases []struct {
			ID       string
			Name     string
			Contents []struct {
				X       float64
				Y       float64
				Width   float64
				Height  float64
				Fill    string
				Outline string
			}
		}
	}

	
	suite.c.RawPost(`mutation { createCanvas(input: {name: "new canvas"}){ id }}`)
	suite.c.RawPost(`mutation { createCanvas(input: {name: "new canvas2"}){ id }}`)
	suite.c.RawPost(`mutation { createCanvas(input: {name: "new canvas3"}){ id }}`)
	suite.c.MustPost(`query{
  canvases{
    id,
    name
    contents {
      x,
      y,
      width,
      height,
      fill,
      outline
      
    }
  }
}`, &queryCanvasesResponse)
	
	suite.Equal(3, len(queryCanvasesResponse.Canvases))
}



func TestRunSuite(t *testing.T) {

	suite.Run(t, new(CanvasSuite))
}
