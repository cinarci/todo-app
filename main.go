package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/cinarci/todo-app/todo"
)

func main() {
	manager := todo.NewManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Yeni ToDo Ekle")
		fmt.Println("2. ToDo Listesini Görüntüle")
		fmt.Println("3. ToDo Tamamla")
		fmt.Println("4. Çıkış")

		fmt.Print("Seçiminiz: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("ToDo Başlığı: ")
			scanner.Scan()
			title := scanner.Text()
			manager.Add(title)
		case "2":
			fmt.Println("ToDo Listesi:")
			for _, todo := range manager.List() {
				status := "Tamamlanmadı"
				if todo.Completed {
					status = "Tamamlandı"
				}
				fmt.Printf("%d - %s [%s]\n", todo.ID, todo.Title, status)
			}
		case "3":
			fmt.Print("Tamamlanacak ToDo ID'si: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err == nil {
				if manager.Complete(id) {
					fmt.Println("ToDo tamamlandı.")
				} else {
					fmt.Println("Geçersiz ID.")
				}
			} else {
				fmt.Println("Geçersiz giriş.")
			}
		case "4":
			fmt.Println("Çıkılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim.")
		}
	}
}
