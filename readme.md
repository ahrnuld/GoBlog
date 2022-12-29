
homepage
individual article 

	// go get -u github.com/gorilla/sessions" to get this package
	// -u = get minor releases
	// go get -u github.com/go-sql-driver/mysql to get this package

admin
    login
    article list
    create article
    edit article
    logout

docker run --link go-mysql:db -p 8080:8080 adminer  

// bezig met https://gowebexamples.com/

// https://go.dev/doc/effective_go, bij 'data'

//docker run --name go-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=dbname -d mysql:latest

// common in GO:
/*
 if err := file.Chmod(0664); err != nil {

	f, err := os.Open(name)
if err != nil {
    return err
}

foreach met range

for key, value := range oldMap {
    newMap[key] = value
}

_ = discard

func (operates on) funcname (params) (returns)

defer func()

parallel assignment is possible

new only allocates

p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer


// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}

organize by component or by kind?
component is preferred?  https://developer20.com/how-to-structure-go-code/

*/

// query := `
	// CREATE TABLE users (
	//     id INT AUTO_INCREMENT,
	//     username TEXT NOT NULL,
	//     password TEXT NOT NULL,
	//     created_at DATETIME,
	//     PRIMARY KEY (id)
	// );`

	// if _, err := db.Exec(query); err != nil {
	// 	log.Fatal(err)
	// }

    
	{ // Insert a new user
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}