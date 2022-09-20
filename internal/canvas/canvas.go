package canvas

import (
	"encoding/json"
	"fmt"
	"github.com/juanfgs/canvas/internal/canvas/shapes"
	database "github.com/juanfgs/canvas/internal/pkg/db/postgresql"
	"log"
)

type Canvas struct {
	ID       string               `json:"id"`
	Contents shapes.RectangleList `json:"contents"`
	Name     string               `json:"name"`
}

func (canvas Canvas) Create() string {
	var uuid string
	sql := fmt.Sprintf("INSERT INTO Canvases(Contents,Name) VALUES ($1,$2)  RETURNING id ")

	jsonContents, err := json.Marshal(canvas.Contents)
	if err != nil {
		log.Fatal(err)
	}
	err = database.Db.QueryRow(sql, jsonContents, canvas.Name).Scan(&uuid)
	if err != nil {
		log.Fatal(err)
	}

	return uuid
}

func (canvas Canvas) Save() error {
	sql := fmt.Sprintf("UPDATE Canvases SET Name = $1, Contents = $2  WHERE id = $3 RETURNING Id,Name,Contents ")

	jsonContents, err := json.Marshal(canvas.Contents)
	if err != nil {
		log.Fatal(err)
	}
	err = database.Db.QueryRow(sql, canvas.Name, jsonContents, canvas.ID).Scan(&canvas.ID,&canvas.Name,&canvas.Contents)
	if err != nil {
		log.Fatal(err)
	}

	return nil 
}

func (canvas *Canvas) Get(uuid string) error {
	sql := "SELECT id,name,contents FROM Canvases WHERE id = $1"

	err := database.Db.QueryRow(sql, uuid).Scan(&canvas.ID, &canvas.Name, &canvas.Contents)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetAll() []*Canvas {
	var canvases []*Canvas
	rows, err := database.Db.Query("SELECT id,name,contents FROM Canvases")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var canvas = &Canvas{}
		rows.Scan(&canvas.ID,&canvas.Name,&canvas.Contents)
		if err != nil {
			log.Fatal(err)
		}
		canvases = append(canvases, canvas)
	}


	return canvases 
}
