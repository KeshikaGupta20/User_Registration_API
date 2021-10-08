package main

import (
	"fmt"
	"register/template"

	"github.com/gofiber/fiber"
	 db "register/dao"
)

func main() {

	app := fiber.New()
	client , ctx, cancel, err := db.connect("mongodb+srv://keshika:Keshika@cluster0.nlwsi.mongodb.net")
	if err != nil{
		panic(err)
	}

	defer close(client, ctx. cancel)

	var details 
	details := ContactDetails{
		Email:   r.FormValue("email"),
		Name:    r.FormValue("name"),
		Address: r.FormValue("address"),
	}

	insertOneResult, err := db.insertOne(client, ctx, "gfg",
		"marks", details)
		if err != nil{
			panic(err)
		}
     fmt.Println("Result of InsertOne")
	 fmt.Println(insertOneResult.InsertedID)
	 app.Listen(":3001")





}
