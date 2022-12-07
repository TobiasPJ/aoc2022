package main

type file struct {
	name string
	size int
}

type dir struct {
	name    string
	subdirs []dir
	files   []file
}

func main() {

}
