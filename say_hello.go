package example_go_module

//misal kita menambah parameter difunction sayhello, artinya semua org yg menggil SayHello sebelumnya akan error karna tidak menambahkan argument
//ini bisa kita anggap sebagai major changes biarpun perubahannya hanya 1 parameter namum code sebelumnya rusak
func SayHello(name string) string {
	return "Hello World" + name
}