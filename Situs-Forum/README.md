### Fitur yang Akan Dibuat

Proyek ini akan mengimplementasikan dua fitur utama: **Membership** dan **Posting Forum**.

#### 1. Membership
Fitur ini mencakup sistem otentikasi pengguna untuk memastikan hanya pengguna yang terdaftar yang dapat mengakses fitur-fitur tertentu.

* **Sign Up**: Pengguna dapat membuat akun baru.
* **Login**: Pengguna yang sudah terdaftar dapat masuk ke sistem.
    * **Access Token**: Token yang digunakan untuk otentikasi dalam setiap permintaan API.
    * **Refresh Token**: Token yang digunakan untuk memperbarui Access Token saat masa berlakunya habis.

#### 2. Posting Forum
Fitur ini memungkinkan pengguna untuk berinteraksi dalam forum, seperti membuat postingan, komentar, dan memberikan 'like'.

* **Middleware**: Logika yang akan dijalankan sebelum permintaan ke endpoint tertentu untuk memeriksa otentikasi pengguna.
* **Create Post**: Endpoint untuk membuat postingan baru di forum.
* **Create Comment**: Endpoint untuk menambahkan komentar pada postingan yang sudah ada.
* **Like/Unlike**: Endpoint untuk memberikan atau menghapus 'like' pada postingan atau komentar.
* **Get All Post**: Endpoint untuk mengambil semua postingan yang ada.
* **Get Post by ID**: Endpoint untuk mengambil detail sebuah postingan berdasarkan ID-nya.

### Entity Relationship Diagram (ERD)

| Nama Tabel | Kolom | Tipe Data | Keterangan |
| :--- | :--- | :--- | :--- |
| **Users** | `id` | `INT` | Primary Key, Not Null |
| | `email` | `VARCHAR(100)` | Not Null |
| | `password` | `VARCHAR(500)` | Not Null |
| | `created_at` | `TIMESTAMP` | Not Null |
| | `updated_at` | `TIMESTAMP` | Not Null |
| | `created_by` | `LONGTEXT` | Not Null |
| | `updated_by` | `LONGTEXT` | Not Null |
| **Posts** | `id` | `INT` | Primary Key, Not Null |
| | `user_id` | `INT` | Foreign Key (FK1) |
| | `post_title` | `VARCHAR(500)` | Not Null |
| | `post_content` | `LONGTEXT` | Not Null |
| | `hashtag` | `LONGTEXT` | Not Null |
| | `created_at` | `TIMESTAMP` | Not Null |
| | `updated_at` | `TIMESTAMP` | Not Null |
| | `created_by` | `LONGTEXT` | Not Null |
| | `updated_by` | `LONGTEXT` | Not Null |
| **Comments** | `id` | `INT` | Primary Key, Not Null |
| | `post_id` | `INT` | Foreign Key (FK1) |
| | `comments` | `LONGTEXT` | Not Null |
| | `created_at` | `TIMESTAMP` | Not Null |
| | `updated_at` | `TIMESTAMP` | Not Null |
| | `created_by` | `LONGTEXT` | Not Null |
| | `updated_by` | `LONGTEXT` | Not Null |
| **User Activities** | `id` | `INT` | Primary Key, Not Null |
| | `post_id` | `INT` | Foreign Key (FK1) |
| | `is_liked` | `BOOLEAN` | Not Null |
| | `created_at` | `TIMESTAMP` | Not Null |
| | `updated_at` | `TIMESTAMP` | Not Null |
| | `created_by` | `LONGTEXT` | Not Null |
| | `updated_by` | `LONGTEXT` | Not Null |