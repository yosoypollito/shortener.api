package links

import (
	"github.com/deta/deta-go/service/base"
	"shortener/api/database"
	"net/http"
	"time"
	"io"
	"strings"
)

type Dates struct {
	Creation time.Time `json:"creation"`
	Modified time.Time `json:"modified"`
}

type Scope string
type Key string

type Link struct{
	ID string `json:"id"`
	Dates Dates `json:"dates"`
	Key Key `json:"key"`
	Scope Scope `json:"scope"`
}

type Body struct{
	Scope Scope `json:"scope" form:"scope" binding:"required,url"`
}

func Get(id string) ([]*Link, error){

	query := base.Query{
		{"id":id},
	}

	results, err := database.GetFromBase[Link]("links", query)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func Create(scope Scope) (string, error){
	db, err := database.Getbase("links")

	if err != nil {
		return "", err
	}

	now := time.Now()

	newLink := Link{
		ID:genId(),
		Dates:Dates{
			Creation: now,
			Modified: now,
			},
		Scope: scope,
	}

	key, err := db.Put(newLink)

	if err != nil {
		return "", err
	}

	return key, nil
}


func genId() string {
	req, err := http.Get("https://www.random.org/strings/?num=1&len=6&digits=off&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	bodystring := string(body)
	return strings.Trim(bodystring, "\n") 
}
