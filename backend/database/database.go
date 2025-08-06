package database

import (
	"database/sql"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "todo.db"
	dbPath = "todo.db"
)

func GetDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	var baseDir string
	switch runtime.GOOS {
	case "windows":
		baseDir = filepath.Join(currentUser.HomeDir, "AppData", "Roaming")
		break
	case "darwin":
		baseDir = filepath.Join(currentUser.HomeDir, "Library", "Application Support")
		break
	case "linux":
		baseDir = filepath.Join(currentUser.HomeDir, ".config")
		break
	default:
		return "", os.ErrNotExist
	}

	appDir := filepath.Join(baseDir, dbPath)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}

	return appDir, nil
}

func GetDb() *sql.DB {
	dir, err := GetDir()
	if err != nil {
		panic(err)
	}

	dbPath := filepath.Join(dir, dbName)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	if err := initTables(db); err != nil {
		panic(err)
	}

	return db
}

func initTables(db *sql.DB) error {
	todoTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FAULSE,
		favourite BOOLEAN DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		due_date DATETIME,
		priority INTEGER DEFAULT 0
	);`

	focusTable := `
	CREATE TABLE IF NOT EXISTS focus_session (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		todo_id INTEGER,
		start_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		end_time DATETIME,
		duration_minutes INTEGER DEFAULT 25,
		completed BOOLEAN DEFAULT FALSE,
		FOREIGN KEY (todo_id) REFERENCES todos(id)
	);`
	if _, err := db.Exec(todoTable); err != nil {
		return err
	}
	if _, err := db.Exec(focusTable); err != nil {
		return err
	}
	return nil
}
