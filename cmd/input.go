package cmd

import "io/ioutil"

type FileInput []string

func FileInput_func(filename string) ([]byte, error) {
	file_byteS, err := ioutil.ReadFile(filename)
	return file_byteS, err
}