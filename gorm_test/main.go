package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string
	Bio   string
	Books []Book // связь "один ко многим"
}

type Book struct {
	gorm.Model
	Title         string
	PublishedYear int
	AuthorID      uint // внешний ключ для связи с Author
}

func addAuthor(db *gorm.DB, reader *bufio.Reader) {
	fmt.Println("Введите имя автора:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	name = strings.TrimSpace(name)

	fmt.Println("Введите биографию автора:")
	bio, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	bio = strings.TrimSpace(bio)

	newAuthor := Author{Name: name, Bio: bio}
	result := db.Create(&newAuthor)
	if result.Error != nil {
		fmt.Println("Ошибка добавления автора:", result.Error)
		return
	}
	fmt.Printf("Автор успешно добавлен с ID: %d\n", newAuthor.ID)
}

func addBook(db *gorm.DB, reader *bufio.Reader) {
	var authorID uint
	for {
		fmt.Print("Введите ID автора: ")
		authorIDStr, _ := reader.ReadString('\n')
		authorIDStr = strings.TrimSpace(authorIDStr)
		id, err := strconv.Atoi(authorIDStr)
		if err != nil {
			fmt.Println("Ошибка: ID должен быть числом. Попробуйте снова")
			continue
		}
		var author Author
		result := db.First(&author, id)
		if result.Error != nil {
			fmt.Printf("Автор с ID %d не найден. Попробуйте снова\n", id)
			continue
		}
		authorID = uint(id)
		break
	}

	fmt.Println("Введите название книги:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	title = strings.TrimSpace(title)

	var year int
	for {
		fmt.Print("Введите год публикации книги: ")
		yearStr, _ := reader.ReadString('\n')
		yearStr = strings.TrimSpace(yearStr)
		y, err := strconv.Atoi(yearStr)
		if err != nil {
			fmt.Println("Ошибка: год должен быть числом. Попробуйте снова")
			continue
		}
		year = y
		break
	}

	newBook := Book{Title: title, PublishedYear: year, AuthorID: authorID}
	result := db.Create(&newBook)
	if result.Error != nil {
		fmt.Println("Ошибка добавления книги:", result.Error)
		return
	}
	fmt.Printf("Книга '%s' успешно добавлена с ID: %d\n", newBook.Title, newBook.ID)
}

func removeEntity(db *gorm.DB, reader *bufio.Reader) {
	fmt.Println("Что хотите удалить?")
	fmt.Println("1. Автор")
	fmt.Println("2. Книга")
	fmt.Print("Введите номер операции: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Print("Введите ID автора для удаления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Ошибка: ID должен быть числом")
			return
		}
		result := db.Delete(&Author{}, id)
		if result.Error != nil {
			fmt.Println("Ошибка удаления автора:", result.Error)
			return
		}
		fmt.Printf("Автор с ID %d удалён\n", id)
	case "2":
		fmt.Print("Введите ID книги для удаления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Ошибка: ID должен быть числом")
			return
		}
		result := db.Delete(&Book{}, id)
		if result.Error != nil {
			fmt.Println("Ошибка удаления книги:", result.Error)
			return
		}
		fmt.Printf("Книга с ID %d удалена\n", id)
	default:
		fmt.Println("Неверный выбор. Операция удаления отменена")
	}
}

func listEntities(db *gorm.DB) {
	var authors []Author
	result := db.Preload("Books").Find(&authors)
	if result.Error != nil {
		fmt.Println("Ошибка получения списка:", result.Error)
		return
	}
	if len(authors) == 0 {
		fmt.Println("База данных пуста")
		return
	}
	fmt.Println("Список авторов и их книг:")
	for _, author := range authors {
		fmt.Printf("Автор: ID=%d, Имя=%s, Биография=%s\n", author.ID, author.Name, author.Bio)
		if len(author.Books) > 0 {
			fmt.Println("  Книги:")
			for _, book := range author.Books {
				fmt.Printf("    ID=%d, Название=%s, Год публикации=%d\n", book.ID, book.Title, book.PublishedYear)
			}
		} else {
			fmt.Println("  Нет книг.")
		}
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	err = db.AutoMigrate(&Author{}, &Book{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Выберите операцию:")
		fmt.Println("1. AddAuthor - добавить автора")
		fmt.Println("2. AddBook - добавить книгу")
		fmt.Println("3. Remove - удалить запись")
		fmt.Println("4. List - вывести список базы данных")
		fmt.Println("5. Exit - выйти")
		fmt.Print("Введите номер операции: ")

		op, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		op = strings.TrimSpace(op)

		switch op {
		case "1", "AddAuthor":
			addAuthor(db, reader)
		case "2", "AddBook":
			addBook(db, reader)
		case "3", "Remove":
			removeEntity(db, reader)
		case "4", "List":
			listEntities(db)
		case "5", "Exit":
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Неверная операция. Попробуйте снова")
		}
		fmt.Println()
	}
}
