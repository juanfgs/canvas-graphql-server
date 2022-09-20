package canvas
import (
	"log"
	"fmt"
	"github.com/juanfgs/canvas/internal/canvas/shapes"
	database "github.com/juanfgs/canvas/internal/pkg/db/postgresql"
)

type Canvas struct {
	ID string `json:"id"` 
	Contents []shapes.Rectangle `json:"contents"` 
	Name string `json:"name"` 
}


func (canvas Canvas) Save() string {
	var uuid string
	sql := fmt.Sprintf("INSERT INTO Canvas(Contents,Name) VALUES (%s,%s) RETURNING id", canvas.Contents,canvas.Name)
	
	err := database.Db.QueryRow(sql).Scan(&uuid)
	if err != nil {
		log.Fatal(err)
	}

	return uuid 
}
