package services

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileOperations struct {
	C        *gin.Context
	File     *multipart.FileHeader
	Filepath string
}

func NewFileOperations(c *gin.Context) *FileOperations {
	file, err := GetFileFromHeader(c)
	if err != nil {
		log.Println(" New File Operation Failed")
	}
	return &FileOperations{
		C:        c,
		File:     file,
		Filepath: "",
	}
}

// GetFileFromHeader Elementary operation for file processing.
func GetFileFromHeader(c *gin.Context) (*multipart.FileHeader, error) {
	// Get file from payload. Elementary operation
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}
	if file == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No file uploaded",
		})
		return nil, nil
	}
	return file, nil
}

// IsDocTypePDF Checks received files format is PDF
func (f *FileOperations) IsDocTypePDF() bool {
	// Check for the file's MIME type
	header, _ := f.File.Open()
	buffer := make([]byte, 512) // Need only first 512 bytes to sniff the content type
	_, err := header.Read(buffer)
	if err != nil {
		f.C.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read file header",
		})
		log.Fatal("Failed to read file header : ", err)
		return false
	}

	// Check if it's a PDF File.
	contentType := http.DetectContentType(buffer)
	if contentType != "application/pdf" {
		f.C.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Expecting PDF File",
		})
		return false
	}
	return true
}

// GetPDFDoc Returns file from header after checking if its of type PDF
func (f *FileOperations) GetPDFDoc() *FileOperations {

	// Check if Doc Type is PDF
	if !f.IsDocTypePDF() {
		return nil
	}
	log.Println("Document is of type PDF.")
	return f
}

// ConvertPDFToStream  Gets file from request header and converts to stream.
func (f *FileOperations) ConvertPDFToStream() (io.Reader, error) {
	// Open the file from the header
	fileReader, err := f.File.Open()
	if err != nil {
		log.Fatal("Error opening file while trying to convert PDF to streams")
		return nil, err
	}
	defer func(fileReader multipart.File) {
		err := fileReader.Close()
		if err != nil {

		}
	}(fileReader)

	// Reset the reader to the beginning, just in case
	_, err = fileReader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	log.Println("PDF File Converted to Streams.")
	return fileReader, nil
}

// SaveFile Saves file to data folder and returns file path
func (f *FileOperations) SaveFile() string {
	// Construct the path for the new file location and save PDF
	f.Filepath = "data/tempfile_" + f.File.Filename
	if err := f.C.SaveUploadedFile(f.File, f.Filepath); err != nil {
		f.C.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save uploaded file",
		})
		return ""
	}
	return f.Filepath
}
