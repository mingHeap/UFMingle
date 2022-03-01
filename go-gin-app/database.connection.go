package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func ConnectDB() error {
	//Open the database, and if it does not exist, create
	db, err := sql.Open("sqlite3", "./UFMingle.db")
	if err != nil {
		return err
	}

	DB = db
	fmt.Println("Successfully connected to the database")
	return nil
}

func createArticleTable() error {
	//create the table of article
	sqlArticleTable := `
		CREATE TABLE IF NOT EXISTS articles(
			id 			INTEGER PRIMARY KEY AUTOINCREMENT,
			author 		TEXT 		NOT NULL,
			title 		TEXT 		NOT NULL,
			content 	TEXT 		NOT NULL
    );
    `
	_, err := DB.Exec(sqlArticleTable)
	if err != nil {
		return err
	}

	//Test: Insert some articles
	//article_1 := article{
	//	Author: "loveLadyGaga",
	//	Title:  "[seeking for a woman]Where is my true love",
	//	Content: "name: Mike.B\n " +
	//		"age:23\n" +
	//		"height:175c\n" +
	//		"I like music and I wanna meet a girl who loves music too",
	//}
	//
	//article_2 := article{
	//	Author: "123handsomeGator",
	//	Title:  "[seeking for a man]I am Rose, looking for my Jack",
	//	Content: "name: Zoey.J\n" +
	//		"age:20\n" +
	//		"height:170cm\n" +
	//		"I was born in a happy family and I prefers boys who have a good sense of humor.\n",
	//}
	//
	//article_3 := article{
	//	Author: "_gators",
	//	Title:  "[seeking for a man]Spotlights on me~",
	//	Content: "name: Lily.J\n" +
	//		"age:27\n" +
	//		"height:172cm\n" +
	//		"I am now a teacher in a high school. I have a dog and tow cats. I am looking for a handsome boyfriend~\n",
	//}
	//stmt, err := DB.Prepare("INSERT INTO articles(author, title, content) values(?,?,?)")
	//checkErr(err)
	//_, err = stmt.Exec(article_1.Author, article_1.Title, article_1.Content)
	//checkErr(err)
	//_, err = stmt.Exec(article_2.Author, article_2.Title, article_2.Content)
	//checkErr(err)
	//_, err = stmt.Exec(article_3.Author, article_3.Title, article_3.Content)
	//checkErr(err)

	fmt.Println("Create table articles successfully")
	return nil
}

func createUserTable() error {
	//create the table of user
	sqlUserTable := `
		CREATE TABLE IF NOT EXISTS users(
			username 	TEXT PRIMARY KEY	NOT NULL,
			password 	TEXT 				NOT NULL
	);
    `
	_, err := DB.Exec(sqlUserTable)
	if err != nil {
		return err
	}

	//Test: Insert some users
	//var userList = []user{
	//	user{Username: "user1", Password: "pass1"},
	//	user{Username: "user2", Password: "pass2"},
	//	user{Username: "user3", Password: "pass3"},
	//}
	//
	//stmt, err := DB.Prepare("INSERT INTO users(username, password) values(?,?)")
	//checkErr(err)
	//_, err = stmt.Exec(userList[0].Username, userList[0].Password)
	//checkErr(err)
	//_, err = stmt.Exec(userList[1].Username, userList[1].Password)
	//checkErr(err)
	//_, err = stmt.Exec(userList[2].Username, userList[2].Password)
	//checkErr(err)

	fmt.Println("Create table users successfully")
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}