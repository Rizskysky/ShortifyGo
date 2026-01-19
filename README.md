
# 🚀 Shortify - URL Shortener (Go Backend Workshop)

Selamat datang di project **Shortify**!  
Di workshop ini, kamu akan berperan sebagai **Backend Engineer** untuk membangun layanan pemendek URL (seperti *bit.ly*) menggunakan bahasa pemrograman **Go (Golang)** dan framework **Gin**.

Project ini tidak hanya sekadar "jalan", tapi juga menggunakan struktur **Clean Architecture** yang merupakan standar industri di perusahaan teknologi besar (Unicorn/Decacorn).

---

## 🛠️ 1. Persiapan (Prerequisites)

Sebelum mulai coding, pastikan "senjata" kamu sudah siap.

### A. Install Golang
Wajib terinstall agar komputer kamu mengerti bahasa Go.
1.  Buka website resmi: [**go.dev/dl**](https://go.dev/dl/).
2.  Download installer sesuai OS kamu (Windows/macOS/Linux).
3.  Install (klik *Next* sampai selesai).
4.  **PENTING: Verifikasi Instalasi**
    Buka terminal (CMD / PowerShell / Terminal), lalu ketik:
    ```bash
    go version
    ```
    *Jika muncul tulisan seperti `go version go1.22.x ...`, berarti kamu siap!*

### B. Tools Pendukung (Rekomendasi)
* **Text Editor:** [VS Code](https://code.visualstudio.com/).
    * *Tips:* Install extension **"Go"** di VS Code agar coding lebih mudah (auto-complete).
* **API Tester:** [Postman](https://www.postman.com/downloads/) (untuk mencoba kirim data ke server).

---

## 🏃‍♂️ 2. Cara Menjalankan Project

Ikuti langkah ini untuk menghidupkan server Shortify di laptopmu:

### Langkah 1: Download Library
Karena kita menggunakan Framework **Gin**, kita perlu mendownload "bahan-bahannya" dulu. Buka terminal di folder project ini, lalu ketik:

```bash
go mod tidy

```

_(Tunggu sampai proses download selesai)_

### Langkah 2: Jalankan Server

Ketik perintah "sakti" ini untuk menjalankan aplikasi:

Bash
```
go run cmd/main.go

```

### Langkah 3: Cek Hasil

Jika berhasil, terminal akan menampilkan tulisan:

> `Listening and serving HTTP on :8080`

Sekarang servermu sudah hidup di alamat `http://localhost:8080`! 🎉

----------

## 🏰 3. Penjelasan Arsitektur (Analogi Restoran)

Mungkin kamu bingung melihat banyak folder. Tenang, kita menggunakan **Clean Architecture**.

Bayangkan aplikasi ini adalah sebuah **Restoran Mewah**:

### 📂 1. Entity (`internal/entity/`)

-   **Analogi:** **Bahan Makanan / Buku Menu.**
    
-   **Fungsi:** Menjelaskan bentuk data dasar.
    
-   **Contoh:** Di sini kita mendefinisikan bahwa sebuah "URL" itu harus punya `ID` (kode pendek) dan `Original` (link asli).
    

### 📂 2. Repository (`internal/repository/`)

-   **Analogi:** **Kulkas / Gudang.**
    
-   **Fungsi:** Tempat menyimpan dan mengambil data.
    
-   **Logika:** Repository tidak peduli datanya untuk apa, tugasnya hanya: _"Simpan ini!"_ atau _"Ambilkan itu!"_. (Di project ini, kita simpan di Memori RAM).
    

### 📂 3. Usecase (`internal/usecase/`)

-   **Analogi:** **Koki (Chef).**
    
-   **Fungsi:** Otak dari aplikasi (Business Logic).
    
-   **Logika:** Di sinilah proses memasak terjadi.
    
    -   User minta pendekin link? -> Cek dulu link-nya kosong gak? -> Generate kode unik -> Suruh Repository simpan.
        

### 📂 4. Handler (`internal/handler/`)

-   **Analogi:** **Pelayan (Waiter).**
    
-   **Fungsi:** Melayani tamu (User/Frontend).
    
-   **Logika:**
    
    -   Menerima pesanan (Request JSON).
        
    -   Meneruskan pesanan ke Koki (Usecase).
        
    -   Mengantar makanan jadi ke tamu (Response JSON).
        

----------

## 🧪 4. Cara Testing (API Endpoints)

Gunakan **Postman** untuk mencoba fitur yang sudah kamu buat.

### 1. Membuat Short URL (POST)

-   **URL:** `http://localhost:8080/shorten`
    
-   **Method:** `POST`
    
-   **Body (Raw JSON):**
    
    JSON
    
    ```
    {
      "url": "[https://www.youtube.com/watch?v=dQw4w9WgXcQ](https://www.youtube.com/watch?v=dQw4w9WgXcQ)"
    }
    
    ```
    
-   **Response:** Kamu akan dapat kode pendek, misal `AbCd1`.
    

### 2. Mengakses Link (GET)

-   **URL:** `http://localhost:8080/AbCd1` (Ganti dengan kode yang kamu dapat)
    
-   **Method:** `GET`
    
-   **Hasil:** Browser/Postman akan otomatis diarahkan (Redirect) ke link asli (YouTube).
    

----------

## 📚 5. Sumber Belajar (References)

Struktur yang kita pakai ini bukan sembarangan, tapi mengacu pada best practice dunia:

1.  **Clean Architecture (The Concept)**
    
    -   _Pencetus:_ Robert C. Martin (Uncle Bob).
        
    -   _Penjelasan:_ Memisahkan aplikasi menjadi layer agar mudah dirawat.
        
    -   [📖 Baca Blog Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
        
2.  **Golang Project Layout**
    
    -   _Standar:_ Struktur folder `cmd`, `internal`, `pkg` adalah standar tidak tertulis komunitas Go.
        
    -   [📖 Lihat Standar Layout Go](https://github.com/golang-standards/project-layout)
        
3.  **Gin Web Framework**
    
    -   _Tools:_ Framework yang kita pakai di file `handler`. Terkenal sangat cepat.
        
    -   [📖 Dokumentasi Gin](https://gin-gonic.com/)
        

----------

### 🔥 Challenge (PR)

Sudah selesai? Coba tantangan ini di rumah:

1.  Tambahkan validasi: Link yang diinput harus diawali `http://` atau `https://`.
    
2.  Tambahkan fitur: Hitung berapa kali link tersebut di-klik (`clicks`).
    

_Happy Coding! Error adalah guru terbaik._ 🚀