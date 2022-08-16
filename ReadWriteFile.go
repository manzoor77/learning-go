package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("myfile.text") //create a new file myfile
	if err != nil {
		panic(err)
	}
	//Write data in file
	data := []byte("Hello World\n")             //byte array to create a string
	file.Write(data)                            //write data using byte
	file.WriteString("I am Doing GOPractice\n") // write data without byte
	file.Close()
	//Read Data From File
	//Rfile, err := os.Open("myfile.text")
	if err != nil {
		panic(err)
	}
	//Rdata := make([]byte, 35) //The data read from the file will be of the size of the byte array
	//Rfile.Seek(6, 0)          //I want to seek the file pointer to the 5th byte of the data
	//Seek has two arguments, first one is the position of the file pointer to read data
	//From where in the file to read from,0-begining,1-at current position,2-form end

	abc, err := ioutil.ReadFile("myfile.text") //read the data from file to the byte array 'data'
	if err != nil {
		panic(err)
	}
	fmt.Println(string(abc))
}
