package main

import "github.com/imafaz/B9CA/controller"

func main() {
	cn := controller.NewBindManager(`/etc/bind/zones`, `/etc/bind/named.conf.local`)
	cn.AddDomain(`afaz.me`)
	cn.AddRecord(`afaz.me`, controller.A, "www", "8.8.8.8", 8600)
}
