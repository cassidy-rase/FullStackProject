package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	// driver to tell mysql is being used
	_ "github.com/go-sql-driver/mysql"
)

// port the api is running on
var port = "8080"

// global db variable as a pointer. pointer is used so there is no need to create copies of the db
var db *sql.DB

// User struct to hold fields from db
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// Product struct to hold fields from db
type Product struct {
	ProductID     int64  `json:"product_id"`
	ProductTypeID int64  `json:"product_typeid"`
	ProductName   string `json:"product_name"`
	ProductPrice  int64  `json:"product_price"`
}

func main() {
	database, err := sql.Open("mysql", "root:mysql123@tcp(mysql-dev:3306)/database?charset=utf8")
	if err != nil {
		panic(err)
	}

	db = database

	defer db.Close()

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/update", handleUpdate)
	http.HandleFunc("/delete", handleDelete)
	http.HandleFunc("/createproduct", createProduct)
	http.HandleFunc("/updateproduct", updateProduct)
	http.HandleFunc("/deleteproduct", deleteProduct)

	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// USER TABLE
func handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		users := []User{}

		enableCors(&w)

		query := `SELECT id, first_name, last_name, email FROM user`

		rows, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {

			var user User

			err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

			if err != nil {
				fmt.Println(err)
				return
			}

			users = append(users, user)
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(users)
	}

	if r.Method == http.MethodPost {

		var user User

		enableCors(&w)

		json.NewDecoder(r.Body).Decode(&user)

		query := `INSERT INTO user (first_name, last_name, email) values (?,?,?)`

		res, err := db.Exec(query, user.FirstName, user.LastName, user.Email)

		if err != nil {
			fmt.Println(err)
			return
		}

		id, err := res.LastInsertId()

		if err != nil {
			fmt.Println(err)
			return
		}

		user.ID = id

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(user)
	}
}

// USER TABLE

// func handleRequest(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {

// 		users := []User{}

// 		enableCors(&w)

// 		query := `SELECT id, first_name, last_name, email FROM user`

// 		rows, err := db.Query(query)

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		for rows.Next() {

// 			var user User

// 			err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			users = append(users, user)
// 		}

// 		w.WriteHeader(http.StatusOK)

// 		json.NewEncoder(w).Encode(users)
// 	}

// 	if r.Method == http.MethodPost {

// 		var user User

// 		enableCors(&w)

// 		json.NewDecoder(r.Body).Decode(&user)

// 		query := `INSERT INTO user (first_name, last_name, email) values (?,?,?)`

// 		res, err := db.Exec(query, user.FirstName, user.LastName, user.Email)

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		id, err := res.LastInsertId()

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		user.ID = id

// 		w.WriteHeader(http.StatusCreated)

// 		json.NewEncoder(w).Encode(user)
// 	}
// }

var saveID int64

func handleUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		enableCors(&w)

		query := `SELECT id FROM user WHERE email="jdoe@gmail.com"`

		rows, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {

			var user User

			err := rows.Scan(&user.ID)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(user)
			//users = append(users, user)
			saveID = user.ID
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println(saveID)
	}

	if r.Method == http.MethodPost {

		enableCors(&w)

		var user User

		json.NewDecoder(r.Body).Decode(&user)

		query := `UPDATE user SET first_name = ?, last_name = ?, email = ? WHERE id = ?`

		_, err := db.Query(query, user.FirstName, user.LastName, user.Email, 1)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		enableCors(&w)

		query := `SELECT id FROM user WHERE first_name="Cayley"`

		rows, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {

			var user User

			err := rows.Scan(&user.ID)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(user)
			//users = append(users, user)
			saveID = user.ID
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println(saveID)
	}

	if r.Method == http.MethodPost {

		var user User

		enableCors(&w)

		json.NewDecoder(r.Body).Decode(&user)

		query := `DELETE FROM user WHERE id = 11`

		_, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}

// PRODUCT TABLE
func createProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		enableCors(&w)

		products := []Product{}

		// query := `SELECT product_id, product_typeid, product_name, product_price FROM product`
		query := `SELECT * FROM database.product`

		rows, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {

			var product Product

			err := rows.Scan(&product.ProductID, &product.ProductTypeID, &product.ProductName, &product.ProductPrice)

			if err != nil {
				fmt.Println(err)
				return
			}

			products = append(products, product)
		}

		fmt.Println(products)

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(products)
	}

	if r.Method == http.MethodPost {

		var product Product

		enableCors(&w)

		json.NewDecoder(r.Body).Decode(&product)

		query := `INSERT INTO product (product_typeid, product_name, product_price) values (?,?,?)`

		res, err := db.Exec(query, product.ProductTypeID, product.ProductName, product.ProductPrice)

		if err != nil {
			fmt.Println(err)
			return
		}

		id, err := res.LastInsertId()

		if err != nil {
			fmt.Println(err)
			return
		}

		product.ProductID = id

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(product)
	}
}

var saveProductID int64

func updateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		enableCors(&w)

		query := `SELECT id FROM product WHERE product_name="Aloe And Kendi Oil"`

		rows, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {

			var product Product

			err := rows.Scan(&product.ProductID)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(product)
			//users = append(users, user)
			saveProductID = product.ProductID
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println(saveProductID)
	}

	if r.Method == http.MethodPost {

		enableCors(&w)

		var product Product

		json.NewDecoder(r.Body).Decode(&product)

		query := `UPDATE product SET product_typeid = ?, product_name = ?, product_price = ? WHERE product_id = ?`

		_, err := db.Query(query, product.ProductTypeID, product.ProductName, product.ProductPrice, 1)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		enableCors(&w)

		query := `SELECT product_id FROM product WHERE product_name="Lengthening Mascara"`

		rows, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {

			var product Product

			err := rows.Scan(&product.ProductID)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(product)
			saveProductID = product.ProductID
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println(saveProductID)
	}

	if r.Method == http.MethodPost {

		var user User

		enableCors(&w)

		json.NewDecoder(r.Body).Decode(&user)

		query := `DELETE FROM product WHERE product_id = 12`

		_, err := db.Query(query)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}
