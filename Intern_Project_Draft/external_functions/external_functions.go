package external_functions

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Function to untar files
func Untar(sourcefile string) {
	if sourcefile == "" {
		fmt.Println("Usage : go-untar sourcefile.tar")
		os.Exit(1)
	}

	file, err := os.Open(sourcefile)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {

			fmt.Println(err)
			os.Exit(3)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(4)
		}

		// get the individual filename and extract to the current directory
		filename := header.Name
		s := strings.Split(filename, "/") //Picking the actual json file

		switch header.Typeflag {

		case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				fmt.Println(err)
				os.Exit(5)
			}

		/*
			Note to self:
				1. Find a way to create the directory using the name of the untarred folder.
				2. Then within that directory, create all the required files, manifest.json and results.json

			Current Solution:
				1. Will just pick the result.json and manifest.json files and create those in the project directory
				2. When processing is done on those files, will remove them from the directory
		*/
		case tar.TypeReg:
			//handle normal file
			fmt.Println("Untarring :", filename)
			writer, err := os.Create(s[1])

			if err != nil {
				fmt.Println("Couldn't create file")
				os.Exit(6)
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(s[1], os.FileMode(header.Mode))

			if err != nil {
				fmt.Println(err)
				os.Exit(7)
			}

			writer.Close()

		default:
			fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}
}

//Function to read through current directory and extract all tar.gz/.tar files and put them in a list
func ReadCurrentDir(ext string) []string {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening current directory: %s", err)
	}
	defer file.Close() //Keep the directory opened until all files have being read.

	var ext_list []string

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		if strings.HasSuffix(name, ext) {
			ext_list = append(ext_list, name)
		}
	}
	return ext_list
}
