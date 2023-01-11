### Listing Category

* [1] File [tag.go](./tag.go)
  * Test : `go run tag.go`
* [2.a] Struktur Table :
    ```
    Table Category
    Nama Field  | Tipe Data    | Description (Optional)
    id         | int(unsigned) | Primary(key)
    category   | Varchar(125)  |
    main_name  | Varchar(125)  |
    child_name | Varchar(125)  |
    ```
* [2.b] List Category : [category.json](./category.json)
* [2.c] Code to Listing Category
  * Test : `go run category.go`
  * Result: 
    ```
    Buku
       Arsitektur & Desain
          Buku Bangunan
          Buku Codes & Standars
       Buku Hukum
          Buku Gender & Hukum
          Buku Hukum Dagang
    Rumah Tangga
       Dekorasi
          Cover Kipas Angin
          Cover Kursi
          Hiasan Dinding
       Furniture
          Bedside Table
          Cermin Badan
          Kasur
    ```
