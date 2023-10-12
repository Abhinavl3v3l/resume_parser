package services

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"

	"code.sajari.com/docconv"
)

// CheckPromptFormat Checks if the pattern of prompt is in the way we expect it to be
func CheckPromptFormat(candidateData string) (string, bool) {
	//log.Println(" Received Prompt : ", candidateData)
	// / Define the regular expression pattern
	pattern := `\{"Email": "([^@]+?@[^.]+?\.[a-z]+?)", "Skills": \["[^"]*?(?:",\s*"[^"]*?)*?"\], "Experience Level": (?:[0-9]|[1-4][0-9]|50)\}`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Find the match
	match := regex.FindString(candidateData)

	if match != "" {
		// Extract the matched JSON
		log.Println("Matched JSON:", match)

		// Now you can parse the JSON
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(match), &data); err != nil {
			log.Println("Error parsing JSON:", err)
		} else {
			skills := data["Skills"].([]interface{})
			experienceLevel := data["Experience Level"].(float64)
			log.Println("Skills:", skills)
			log.Println("Experience Level:", experienceLevel)
		}

		return match, true

	} else {
		log.Println("Received JSON does not match the expected format.")
		return "", false
	}

}

// ConvertSavedPDFToText takes a file path of a saved PDF, converts it to text using docconv, and returns the text.
func ConvertSavedPDFToText(filePath string) (string, error) {
	// Open the file from the provided path
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing file.")
		}
	}(file)

	// Use docconv to convert the file to text
	text, _, err := docconv.ConvertPDF(file)
	if err != nil {
		return "", err
	}

	return text, nil
}

// ConvertStreamToText takes an io.Reader of a PDF stream, converts it to text using docconv, and returns the text.
func ConvertStreamToText(pdfStream io.Reader) (string, error) {
	log.Println("Converting Streams to Text using DOCCONV ")
	// Use docconv to convert the stream to text
	text, _, err := docconv.ConvertPDF(pdfStream)
	if err != nil {
		return "", err
	}

	return text, nil
}
