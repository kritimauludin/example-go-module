// Major Upgrade

// ● Major upgrade biasanya terjadi dikarenakan ada perubahan pada isi kode program kita, sehingga
// membuatnya tidak backward compatible
// ● Sebaiknya hal ini sebisa mungkin dihindari
// ● Namun jika tidak bisa dihindari, strategy terbaik adalah merubah nama module

// biasanya programer go akan merubah nama module dengan menambah keyword diakhir nama modul menggunakan nama version
// harapannya jika ada yg melakukan get kemodule tanpa (v2) akan dapat versi 1 
//sehingga jika ingin dapat v2 perlu menambah keyword dibelakangnya 

// old
// module github.com/kritimauludin/example-go-module
module github.com/kritimauludin/example-go-module/v2

go 1.17
