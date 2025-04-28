package endpoints

import (
    "net/http"
    "os"
    "io"
    "fmt"
)

func PostFile(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Create a new file on the server
	outFile, err := os.Create("uploaded_file.dat")
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Copy the request body (byte stream) to the file
	_, err = io.Copy(outFile, r.Body)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "File uploaded successfully")
}


func GetFile(w http.ResponseWriter, r *http.Request) {
	// Path to the file you want to serve
	filePath := "C:\\Users\\pe190\\Desktop\\GitHub\\KaraokeGameBackend\\files\\thisworks.txt"
	fileName := "test.txt" // The name client sees when downloading

	// Set headers to trigger download
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	// Serve the file
	http.ServeFile(w, r, filePath)
}
