Connecting Go to Databases: A Complete Guide
=============================================

Go has excellent support for databases through the standard library and third-party packages. The most common approach is using the database/sql package along with drivers for specific databases.

1. The database/sql Package
---------------------------
The database/sql package is part of Go's standard library. It provides a generic interface for working with databases. However, you still need a driver specific to your database (MySQL, PostgreSQL, SQLite, etc.).

Why use database/sql?
- It handles connection pooling automatically
- It provides a clean, abstract interface regardless of which database you're using
- It prevents SQL injection when used with parameterized queries
- It handles resource management (closing connections, etc.)


2. Installing Database Drivers
------------------------------
Go doesn't include database drivers in the standard library to keep it modular. You need to import a driver for your specific database.

Common Drivers:
- MySQL: github.com/go-sql-driver/mysql
- PostgreSQL: github.com/lib/pq
- SQLite: github.com/mattn/go-sqlite3
- MongoDB (not SQL, but popular): go.mongodb.org/mongo-driver/mongo

Installation:
bash
go get github.com/go-sql-driver/mysql
go get github.com/lib/pq
go get github.com/mattn/go-sqlite3


3. Basic Database Connection Pattern
-------------------------------------
Here's how to connect to a database in Go:

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" // Import the driver (blank identifier)
)

func main() {
    // Open doesn't actually connect; it just validates the arguments
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database_name")
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close() // Always close the connection when done

    // Ping actually establishes the connection
    err = db.Ping()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }

    fmt.Println("Successfully connected to MySQL!")
}

Important: The blank identifier (_) imports the driver for its side effects (registering itself with database/sql) without giving us a name to use it directly.


4. Connection String Formats
----------------------------
Different databases have different connection string formats:

MySQL:
user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local

PostgreSQL:
user=username password=password dbname=dbname sslmode=disable

SQLite:
test.db (just a file path)


5. CRUD Operations
------------------

CREATE (Insert Data):
result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "John", "john@example.com")
if err != nil {
    log.Fatal(err)
}
id, _ := result.LastInsertId() // Get the auto-generated ID

READ (Query Data):
// Single row query
var name string
var email string
row := db.QueryRow("SELECT name, email FROM users WHERE id = ?", 1)
err := row.Scan(&name, &email)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Name: %s, Email: %s\n", name, email)

// Multiple rows query
rows, err := db.Query("SELECT name, email FROM users")
if err != nil {
    log.Fatal(err)
}
defer rows.Close() // IMPORTANT: Always close rows!

for rows.Next() {
    var name string
    var email string
    if err := rows.Scan(&name, &email); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Name: %s, Email: %s\n", name, email)
}

// Always check for errors from rows.Next()
if err := rows.Err(); err != nil {
    log.Fatal(err)
}

UPDATE:
result, err := db.Exec("UPDATE users SET email = ? WHERE name = ?", "newemail@example.com", "John")
if err != nil {
    log.Fatal(err)
}
rowsAffected, _ := result.RowsAffected()
fmt.Printf("Updated %d rows\n", rowsAffected)

DELETE:
result, err := db.Exec("DELETE FROM users WHERE id = ?", 1)
if err != nil {
    log.Fatal(err)
}


6. Using Prepared Statements
----------------------------
Prepared statements are pre-compiled SQL templates that can be executed repeatedly. They provide:
- Better performance for repeated queries
- Protection against SQL injection

// Create once, use many times
stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Execute multiple times with different values
for _, user := range users {
    _, err := stmt.Exec(user.Name, user.Email)
    if err != nil {
        log.Fatal(err)
    }
}


7. Transaction Handling
------------------------
When you need multiple operations to succeed or fail together, use transactions:

tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback() // Rollback if not committed

_, err = tx.Exec("INSERT INTO orders (product, price) VALUES (?, ?)", "Widget", 19.99)
if err != nil {
    log.Fatal(err)
}

_, err = tx.Exec("UPDATE inventory SET count = count - 1 WHERE product = ?", "Widget")
if err != nil {
    log.Fatal(err)
}

// Commit the transaction
err = tx.Commit()
if err != nil {
    log.Fatal(err)
}

Important: Always use defer tx.Rollback() right after Begin(), even before checking the error. This ensures the rollback happens if something goes wrong before Commit().


8. Connection Pooling
---------------------
database/sql automatically manages a connection pool. You can configure it:

// Set maximum number of open connections
db.SetMaxOpenConns(25)

// Set maximum number of idle connections
db.SetMaxIdleConns(25)

// Set maximum connection lifetime
db.SetConnMaxLifetime(time.Hour)


9. Handling NULL Values
------------------------
SQL NULL values require special handling in Go because they don't map directly to Go types:

import "database/sql"

var name sql.NullString // Use sql.NullString instead of string

row := db.QueryRow("SELECT name FROM users WHERE id = 1")
err := row.Scan(&name)
if err != nil {
    log.Fatal(err)
}

// Check if NULL before using
if name.Valid {
    fmt.Println("Name:", name.String)
} else {
    fmt.Println("Name is NULL")
}

// Alternative: Use pointers to pointers or custom types
// Or better: design your schema to avoid NULLs when possible


10. Best Practices and Tips
---------------------------

ALWAYS:
- Use parameterized queries (?) to prevent SQL injection
- Close rows with defer rows.Close()
- Check for errors after every database operation
- Use transactions for multiple related operations
- Set connection pool limits appropriate for your application
- Handle NULL values properly using sql.NullString, sql.NullInt64, etc.
- Close the database connection with defer db.Close()

NEVER:
- Concatenate user input into SQL strings
- Forget to check for errors (even when using QueryRow)
- Leave connections open
- Query in a loop if you can use batch operations

Trick: Use context.Context for timeouts and cancellation:

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, "SELECT * FROM users")
if err != nil {
    // Handle timeout or cancellation
}


11. Working with Different Databases
------------------------------------

SQLite (for development/testing):
- No server needed, just a file
- Great for local development
- Use github.com/mattn/go-sqlite3

PostgreSQL (recommended for production):
- Most feature-rich open source database
- Use github.com/lib/pq (pure Go, no C dependencies)
- Supports advanced features like JSONB, arrays, etc.

MySQL (popular, widely used):
- Use github.com/go-sql-driver/mysql
- Good for web applications
- MariaDB is a drop-in replacement


12. Practical Example: Full CRUD API
------------------------------------

package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/myapp")
    if err != nil {
        log.Fatal(err)
    }
    
    // Test connection
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
    
    // Configure pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        users = append(users, u)
    }
    
    json.NewEncoder(w).Encode(users)
}

func main() {
    initDB()
    defer db.Close()
    
    http.HandleFunc("/users", getUsers)
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


13. ORM vs. Raw SQL
-------------------
You have two main approaches:

Raw SQL (using database/sql):
- More control over queries
- Better performance
- Portable across databases
- Requires writing more boilerplate

ORM (Object-Relational Mapping):
- GORM (github.com/jinzhu/gorm) - Most popular
- SQLx (github.com/jmoiron/sqlx) - Extension of database/sql
- XORM (github.com/go-xorm/xorm)

Recommendation:
- For simple applications: Use raw SQL with database/sql
- For complex applications: Consider GORM or SQLx
- For learning: Stick with raw SQL to understand what's happening


14. Testing with Databases
---------------------------
For testing, use an in-memory database:

func TestUserRepository(t *testing.T) {
    // Use SQLite in memory for testing
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // Create tables
    _, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
    if err != nil {
        t.Fatal(err)
    }

    // Now run your tests against this db
}


15. Common Pitfalls to Avoid
-----------------------------

1. Forgetting to check for sql.ErrNoRows:
   row := db.QueryRow("SELECT * FROM users WHERE id = ?", 999)
   err := row.Scan(&user)
   if err == sql.ErrNoRows {
       // Handle "not found" case - this is NOT an error!
   } else if err != nil {
       // This IS an actual error
   }

2. Not closing database rows:
   Always use defer rows.Close() after Query()
   Even if you break out of a loop early, rows.Next() might not have been called
   The defer ensures resources are cleaned up

3. Connection exhaustion:
   Make sure you're calling Close() on connections and rows
   Set appropriate limits on the connection pool

4. SQL injection vulnerabilities:
   NEVER do this: fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", userInput)
   ALWAYS do this: db.Query("SELECT * FROM users WHERE name = ?", userInput)


Summary
-------
Go's database/sql package provides a powerful, standard interface for working with databases. The key points to remember:

1. Import a driver using blank identifier
2. Use sql.Open() to create a database handle
3. Use db.Ping() to verify the connection
4. Use parameterized queries for security
5. Always handle errors and close resources
6. Use transactions for multi-step operations
7. Leverage connection pooling for performance
8. Use context for timeouts and cancellation

With these fundamentals, you can connect Go to virtually any database and build robust, efficient data-driven applications!

