# GOLANG DATABASE

## List
- Pengenalan Golang Database
- Package database
- Membuat koneksi database
- Eksekusi perintah SQL
- SQL Injection
- Prepare Statement
- Database Transaction

## Pengenalan Golang Database
- secara default memiliki sebuah package bernama `database`
- https://golang.org/s/sqldrivers
- go get -u github.com/go-sql-driver/mysql

## Koneksi ke database
- menggunakan object `sql.DB` menggunakan function `sql.Open(driver, dataSourceName)`
- jangan lupa close koneksinya

## Database pooling
- management koneksi

## Insert, Update, Delete
- disarankan untuk menggunakan `context`
- namun untuk query jangan menggunakan context karena context tidak mengembalikan value. gunakan `QueryContext`

## Nullable data

## Prepare Statement
Koneksi yg sama pakai parameter berbeda, biar nggak manggil koneksi baru terus. Biar nggak nanya terus ke db pool.

## Database Transaction
## Repository Pattern
To seperate business logic

## Slide
<a href="https://docs.google.com/presentation/d/15pvN3L3HTgA9aIMNkm03PzzIwlff0WDE6hOWWut9pg8/edit?usp=sharing" target="_blank">Here</a>