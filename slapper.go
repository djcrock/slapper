// Slapper is a web UI for uploading staic websites.
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var dir string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Target directory required")
		fmt.Println("Usage: slapper <targetDirectory>")
		os.Exit(1)
	}

	dir = os.Args[1]
	dirStat, err := os.Stat(dir)
	if err != nil || !dirStat.IsDir() {
		log.Printf("Invalid directory %s", dir)
		os.Exit(2)
	}

	log.Printf("Slapping %s", dir)

	http.HandleFunc("/", handleUploadPage)
	http.HandleFunc("/slap", handleSlap)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleUploadPage serves an upload page to the user.
func handleUploadPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(uploadTemplate))
}

// handleSlap receives a zip file and extracts it into the target directory.
func handleSlap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	buf := new(bytes.Buffer)
	file, header, err := r.FormFile("site")
	if err != nil {
		log.Printf("failed to receive uploaded file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	log.Printf("filename: %s", header.Filename)

	written, err := io.Copy(buf, file)

	log.Printf("received %d bytes", written)
	zipFile := buf.Bytes()
	zipReader, err := zip.NewReader(bytes.NewReader(zipFile), written)
	if err != nil {
		log.Printf("failed to open zip file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Time to slap")

	err = clearDirectory(dir)
	if err != nil {
		log.Printf("failed to clear target directory: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, f := range zipReader.File {
		err = extractFile(f, dir)
		if err != nil {
			log.Printf("failed to extract file %s: %v", f.Name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println("Slapping complete!")

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

// clearDirectory empties the directory identified by dirName of all contents.
func clearDirectory(dirName string) error {
	dir, err := os.Open(dirName)
	if err != nil {
		return err
	}
	defer dir.Close()

	fileNames, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, fileName := range fileNames {
		err = os.RemoveAll(filepath.Join(dirName, fileName))
		if err != nil {
			return err
		}
	}

	return nil
}

// extractFile extracts the given file into the target directory.
func extractFile(f *zip.File, dirName string) error {
	dstName := filepath.Join(dirName, f.Name)
	log.Printf("extracting %s", dstName)
	if f.FileInfo().IsDir() {
		return os.MkdirAll(dstName, os.ModePerm)
	}

	err := os.MkdirAll(filepath.Dir(dstName), os.ModePerm)
	if err != nil {
		return err
	}

	fc, err := f.Open()
	if err != nil {
		return err
	}
	defer fc.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, fc)

	return err
}
