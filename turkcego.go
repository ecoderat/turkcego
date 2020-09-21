package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {

	args := os.Args[1:] //command-line arguments for a filename
	if len(args) == 0 {
		fmt.Println("Provide a file")
		return
	}

	content, err := ioutil.ReadFile(args[0]) //reads the contents of the file
	if err != nil {
		log.Println(err)
	}

	tr := string(content) //byte slice -> string

	equivalent := []map[string]string{
		{"key": "fonksiyon", "value": "func"},
		{"key": "paket", "value": "package"},
		{"key": "döndür", "value": "return"},
		{"key": "aktar", "value": "import"},
		{"key": "sabit", "value": "const"},
		{"key": "değilse", "value": "else"},
		{"key": "tekrarla", "value": "for"},
		{"key": "eğer", "value": "if"},
		{"key": "dönüştür", "value": "map"},
		{"key": "herbirinde", "value": "range"},
		{"key": "yapı", "value": "struct"},
		{"key": "tip", "value": "type"},
		{"key": "değişken", "value": "var"},
		{"key": "sayı", "value": "int"},
		{"key": "Yazdırf", "value": "Printf"},
		{"key": "Yazdır", "value": "Println"},
		{"key": "ondalık", "value": "float64"},
		{"key": "yazı", "value": "string"},
		{"key": "oluştur", "value": "make"},
	}

	for _, v := range equivalent {
		re := regexp.MustCompile(v["key"])
		tr = re.ReplaceAllLiteralString(tr, v["value"])

	}

	en := []byte(tr) //string -> byte slice

	err = ioutil.WriteFile("temp.go", en, 0644) //writes the converted version to a temporary file
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := exec.Command("go", "run", "temp.go").Output() //runs the temporary file
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s", out) //prints the output

	defer exec.Command("rm", "temp.go").Run() //deletes the temporary file

}

