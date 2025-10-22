package main

import (
	"fmt"
	"log"

	dbpackage "github.com/RasagnyaVanga/gorm_activity/database"
)

func main() {
	db, err := dbpackage.NewBlogDB() //db is the interface returned by constructor through which we can access all the methods of db package
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	fmt.Println("Connected to DB")

	//CREATE
	// create_err := db.CreatePost("Change", "Vishwanath", "biography")
	// if create_err != nil {
	// 	fmt.Println(create_err)
	// 	return
	// }

	//UPDATE by id
	// update_err := db.UpdatePost(2, "Colors of life", "drama")
	// if update_err != nil {
	// 	fmt.Println(update_err)
	// 	return
	// }

	//DELETE by title
	// if err := db.DeletePost("Colors of life"); err != nil {
	// 	fmt.Println("Error in deleting: ", err)
	// 	return
	// }

	//SEARCHING
	// posts, err := db.SearchPost("Klaus", "Different ways of life")
	// if err != nil {
	// 	fmt.Println("Error in searching: ", err)
	// 	return
	// }
	// fmt.Println(posts)

	//READ-ALL
	res, err := db.ReadAllPosts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
