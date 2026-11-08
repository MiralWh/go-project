package works

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite" // слепая инициализация (работает тот самый init()) но работем через абстракцию database/sql
)

// AuthResult — результат аутентификации
type AuthResult int

const (
	UserNotFound  AuthResult = iota // Пользователь не найден
	WrongPassword                   // Пароль неверный
	Success                         // Пароль верный
)

func (r AuthResult) String() string {
	switch r {
	case UserNotFound:
		return "Пользователь не найден"
	case WrongPassword:
		return "Пароль не найден"
	case Success:
		return "Пароль верный"
	default:
		return "Неизвестный результат"
	}
}

func InitDB(path string) (*sql.DB, error) {

	memType := ":memory:"

	if path != "" {

		memType = path
	}

	db, err := sql.Open("sqlite", memType)
	if err != nil {
		return nil, fmt.Errorf("открытие БД: %w", err)
	}

	db.SetMaxOpenConns(2)

	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)

	if err != nil {
		return nil, fmt.Errorf("создание таблицы: %w", err)
	}

	if err := seedUsers(db, 5); err != nil {
		return nil, fmt.Errorf("заполнение данными: %w", err)
	}

	return db, nil
}

func seedUsers(db *sql.DB, n int) error {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range n {
		email := fmt.Sprintf("user%d_%d@example.com", i, r.Intn(10000))
		plainPassword := fmt.Sprintf("pass%d_%d", i, r.Intn(100000))

		hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

		if err != nil {
			return fmt.Errorf("хеширование пароля: %w", err)
		}

		_, err = db.Exec(
			"INSERT INTO users (email, password) VALUES (?, ?)",
			email, string(hash),
		)

		if err != nil {
			return fmt.Errorf("вставка пользователя: %w", err)
		}
	}

	return nil
}

func Authenticate(db *sql.DB, email, password string) (AuthResult, error) {

	var storedHash string

	err := db.QueryRow(
		"SELECT password FROM users WHERE email = ?", email,
	).Scan(&storedHash)

	if err == sql.ErrNoRows {
		return UserNotFound, nil
	}

	if err != nil {
		return UserNotFound, fmt.Errorf("запрос пользователя: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))

	if err != nil {
		return WrongPassword, nil
	}

	return Success, nil
}

func Work12() {

	fmt.Println("Добро пожаловать на 12-e задание по языку \"Go\"")

	tmpFile, err := os.CreateTemp("", "app-debug-*.db")
	if err != nil {
		log.Fatalf("Не удалось создать временный файл: %v", err)
	}

	dbPath := tmpFile.Name()
	tmpFile.Close()

	defer func() {
		fmt.Printf("\n[Удаление] Удаляем временную базу данных: %s\n", dbPath)
		if err := os.Remove(dbPath); err != nil {
			log.Printf("Ошибка при удалении файла базы: %v", err)
		}
	}()

	fmt.Printf("[Отладка] База данных создана! Путь для вьюера:\n%s\n\n", dbPath)

	db, err := InitDB(dbPath)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	var demoEmail, demoHash string

	err = db.QueryRow("SELECT email, password FROM users LIMIT 1").Scan(&demoEmail, &demoHash)

	if err != nil {
		log.Fatalf("Получение демо-пользователя: %v", err)
	}

	cases := []struct {
		email    string
		password string
	}{
		{demoEmail, "laskmdlkmnewou, mc`po91-30-o32-9"}, // белиберда
		{"czxacoweuweuigwkuhjr@example.com", "any"},     // белиберда
	}

	for _, c := range cases {

		result, err := Authenticate(db, c.email, c.password)

		if err != nil {
			log.Printf("Ошибка аутентификации: %v", err)
			continue
		}

		fmt.Printf("Email: %s | Пароль: %s >>> %s\n", c.email, c.password, result)
	}

	const testEmail = "admin@example.com"
	const testPassword = "secret123"

	hash, _ := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.DefaultCost)

	_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", testEmail, string(hash))

	if err != nil {
		log.Printf("Ошибка добавления валидного пользователя: %v", err)
	}

	// тут можно не обрабатывать, ошибки в самом Authenticate
	result, _ := Authenticate(db, testEmail, testPassword)

	fmt.Printf("Email: %s | Пароль: %s >>> %s\n", testEmail, testPassword, result)
}
