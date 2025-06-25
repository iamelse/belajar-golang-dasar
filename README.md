# 📋 Todo CLI – Golang + JSON

Todo CLI adalah aplikasi baris-perintah sederhana untuk mencatat, mengubah, menampilkan, serta menandai selesai daftar tugas (to-do list).  
Data disimpan lokal dalam satu berkas **JSON**, sehingga ringan dan tanpa ketergantungan basis data eksternal.

---

## ✨ Fitur

| Perintah       | Deskripsi                                  | Contoh                                              |
| -------------- | ------------------------------------------ | --------------------------------------------------- |
| `add`          | Menambah tugas baru                        | `todo add "Belajar unit test"`                      |
| `edit`         | Mengubah judul tugas                       | `todo edit 3 "Refactor handler"`                    |
| `list`         | Menampilkan seluruh tugas                  | `todo list`                                         |
| `done`         | Menandai tugas selesai                     | `todo done 2`                                       |
| `delete`       | Menghapus tugas                            | `todo delete 5`                                     |

---

## 🏗️ Struktur Proyek

```bash
├── main.go # Titik masuk aplikasi & routing CLI
├── models
│ └── task.go # Definisi struct Task
├── storage
│ └── storage.go # Load/Save JSON ke disk
└── tasks.json # Dibuat otomatis saat runtime
```


## 🚀 Instalasi & Build

```bash
# Klon repositori
git clone https://github.com/iamelse/belajar-golang-dasar.git

cd belajar-golang-dasar

# Build biner
go build -o todo .
```

## 📝 Cara Pakai

```bash
# 1. Tambah tugas baru
go run main.go add "Menulis dokumentasi"

# 2. Lihat daftar tugas
go run main.go list

# 3. Tandai selesai
go run main.go done 1

# 4. Edit tugas
go run main.go edit 1 "Menulis README.md"

# 5. Hapus tugas
go run main.go delete 1
```

