package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"belajar-golang/models"
	"belajar-golang/storage"
)

// main adalah titik masuk program.
// Ia memeriksa argumen, lalu mengdelegasikan ke handler perintah spesifik.
func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		if err := handleAdd(args); err != nil {
			fmt.Println("Error:", err)
		}
	case "edit":
		if err := handleEdit(args); err != nil {
			fmt.Println("Error:", err)
		}
	case "list":
		if err := handleList(); err != nil {
			fmt.Println("Error:", err)
		}
	case "done":
		if err := handleDoneDelete(args, true); err != nil {
			fmt.Println("Error:", err)
		}
	case "delete":
		if err := handleDoneDelete(args, false); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		printUsage()
	}
}

// printUsage menampilkan petunjuk pemakaian singkat ke terminal.
func printUsage() {
	exe := filepath.Base(os.Args[0])
	commands := []struct {
		usage string
		desc  string
	}{
		{"add \"nama tugas\"", "Tambah tugas baru"},
		{"edit <id> \"nama tugas baru\"", "Ubah isi tugas"},
		{"list", "Daftar semua tugas"},
		{"done <id>", "Tandai tugas selesai"},
		{"delete <id>", "Hapus tugas"},
	}

	fmt.Println("Penggunaan:")
	for _, cmd := range commands {
		fmt.Printf("  %s %-25s %s\n", exe, cmd.usage, cmd.desc)
	}
}

// handleAdd memproses perintah `add`.
func handleAdd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("format: todo add \"nama tugas\"")
	}
	title := strings.Join(args, " ")

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	newID := nextID(tasks)
	tasks = append(tasks, models.Task{
		ID:        newID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err := storage.SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("✓ Task #%d ditambahkan: %s\n", newID, title)
	return nil
}

// handleEdit memproses perintah `edit`.
func handleEdit(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("format: todo edit <id> \"tugas baru\"")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("id harus berupa angka")
	}

	newTitle := strings.Join(args[1:], " ")

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	index := findTaskIndex(tasks, id)
	if index == -1 {
		return fmt.Errorf("task #%d tidak ditemukan", id)
	}

	tasks[index].Title = newTitle
	tasks[index].UpdatedAt = time.Now()

	if err := storage.SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("✎ Task #%d berhasil diubah menjadi: %s\n", id, newTitle)
	return nil
}

// handleList memproses perintah `list`.
func handleList() error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas ✨")
		return nil
	}

	fmt.Println("Daftar tugas:")
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "✓"
		}
		fmt.Printf("  [%s] %d. %s\n", status, t.ID, t.Title)
	}
	return nil
}

// handleDoneDelete memproses `done` dan `delete`.
//
// Param markDone menentukan mode:
//   - true  → tandai selesai
//   - false → hapus
func handleDoneDelete(args []string, markDone bool) error {
	if len(args) == 0 {
		cmd := "done"
		if !markDone {
			cmd = "delete"
		}
		return fmt.Errorf("format: todo %s <id>", cmd)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("id harus berupa angka")
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	index := findTaskIndex(tasks, id)
	if index == -1 {
		return fmt.Errorf("task #%d tidak ditemukan", id)
	}

	if markDone {
		tasks[index].Done = true
		fmt.Printf("✓ Task #%d ditandai selesai\n", id)
	} else {
		fmt.Printf("✗ Task #%d dihapus\n", id)
		tasks = append(tasks[:index], tasks[index+1:]...)
	}

	return storage.SaveTasks(tasks)
}

// nextID mengembalikan id baru yang belum terpakai pada slice tasks.
func nextID(tasks []models.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

// findTaskIndex mencari indeks task berdasarkan id.
// Mengembalikan -1 jika tidak ketemu.
func findTaskIndex(tasks []models.Task, id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}
