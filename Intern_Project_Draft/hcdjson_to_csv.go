package main

import (
	"Intern_Project_Draft/external_functions"
	"fmt"
)

func main() {
	//Picking tar.gz files from the directory
	ext_list_gz := external_functions.ReadCurrentDir(".gz")
	//fmt.Println(ext_list_gz)

	//Untarring the pulled .gz files
	for i := 0; i < len(ext_list_gz); i++ {
		external_functions.Untar(ext_list_gz[i])
		fmt.Println("Done untarring all tarred files.")
	}

	//Time to pick manifest.json and result.json and begin processing
	//Processing on manifest.json

}
