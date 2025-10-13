package main

import (
	"fmt"
	"log"
)

func main() {
	database, err := NewBlogDB()
	if err != nil {
		log.Fatal(err)
	}
	var db BlogManager
	db = database
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	fmt.Println("Connected to DB:", *database)

	//CREATE
	create_err := db.CreateBlog("Different ways of life", "Klaus", "drama")
	if create_err != nil {
		fmt.Println(create_err)
		return
	}

	//UPDATE by id
	// update_err := db.UpdateTable(4, "Colors of life", "drama")
	// if update_err != nil {
	// 	fmt.Println(update_err)
	// 	return
	// }

	//DELETE by title
	// if err := db.DeleteBlog("Change"); err != nil {
	// 	fmt.Println("Error in deleting: ", err)
	// 	return
	// }
	// db.DeleteBlog("Home")

	//SEARCHING
	// posts, err := db.SearchBlog("Klaus", "Colors of life")
	// if err != nil {
	// 	fmt.Println("Error in searching: ", err)
	// 	return
	// }
	// fmt.Println(posts)

	//READ-ALL
	res, err := db.ReadAllBlogs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
