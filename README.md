## How to run
untuk run, pertama-tama clone repo ini
```bash
git clone https://github.com/risnandym/test-echo.git
```
kemudian siapkan .env pada file projek yang isinya adalah sebagai berikut (dapat disesuaikan)
```bash
DATABASE_USERNAME=Admin
DATABASE_PASSWORD=
DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_NAME=echo
SECRET=
```
setelah itu buat database bernama "echo" (dapat disesuaikan), setelah membuat database lalu jalankan command berikut:
```bash
go mod download
```
jika muncul beberapa error untuk package, maka jalankan command "go get" contoh:
```bash
 go get -u github.com/swaggo/swag/cmd/swag
```
setelah itu, run programnya dengan command:
```bash
 go run .
```
## How to use
untuk melihat endpoint apasaja yang tersedia, bisa menggunakan link swagger pada local. contoh:
```bash
 http://localhost:8080/swagger/index.html
```
