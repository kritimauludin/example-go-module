package example_go_module

/**
Membuat Module

● Untuk membuat module baru, kita bisa menggunakan perintah : go mod init nama-module
● Go-Lang akan secara otomatis membuat file go.mod yang berisikan informasi nama module dan
juga versi Go-Lang yang kita gunakan
*/

//misal kita menambah parameter difunction sayhello, artinya semua org yg menggil SayHello sebelumnya akan error karna tidak menambahkan argument
//ini bisa kita anggap sebagai major changes biarpun perubahannya hanya 1 parameter namum code sebelumnya rusak

func SayHello(name string) string {
	return "Hello World" + name
}
