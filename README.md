# Rent-Book-App
Adalah sebuah mini project RestAPI untuk sewa menyewa buku, dengan menggunakan bahasa GO dan echo framework

# DokumentasiAPI

server = localhost:9090/

1. Authentication

   a) POST ==> /login
   
       dengan memasukkan email dan password
    
2. Users

   a) POST ==> /users
    
       membuat user baru, dengan menginputkan: name, email, password
      
   b) GET ==> /users/:userId
   
       untuk mendapatkan user sesuai dengan userId
        
   c) GET ==> /users
   
       untuk mendapatkan semua users
        
   d) DELETE ==> /users
   
       untuk menghapus data user
        
   e) PUT ==> /users
   
       untuk mengupdate data pada user
    
    
3. Books

   a) POST ==> /books
    
       untuk menambahkan buku baru, dengan menginputkan: title, author
      
   b) GET ==> /books/:bookId
   
       untuk mendapatkan book sesuai dengan bookId
        
   c) GET ==> /books
   
       untuk mendapatkan semua books
        
   d) DELETE ==> /books/:bookId
   
       untuk menghapus book sesuai dengan bookId
        
   e) PUT ==> /books/:bookId
   
       untuk mengupdate books sesuai dengan bookId
        
4) Rent

  a) POST ==> /rents
    
       untuk menambahkan rent baru, dengan menginputkan: bookId
      
  b) GET ==> /rents
   
       untuk mendapatkan semua rents
        
  c) PUT ==> /rents/:rentId
   
       untuk mengupdate tanggal kembalian rent sesuai dengan rentId
        
# Run Local

Cloning Project

      & git clone https://github.com/NANDA-FATIH-BE-1/Rent-Book-App.git

# Author

@pramudya3
