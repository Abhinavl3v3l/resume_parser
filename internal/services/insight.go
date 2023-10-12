package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/now"
	openai "github.com/sashabaranov/go-openai"
)

func cleanseText(respObj openai.ChatCompletionResponse) string {
	cleansedtext, check := CheckPromptFormat(respObj.Choices[0].Message.Content)
	if !check {
		log.Fatal(" Incorrect prompt text format")
	}
	return cleansedtext
}

func InsightStream(pdfStream io.Reader) (string, error) {
	// Extract Text from the file.
	res, err := ConvertStreamToText(pdfStream)
	if err != nil {
		log.Fatal(" Stream to text conversion failed ", err)
		return "", err
	}
	// Strips unnecessary spaces and newlines
	cleanedContent := strings.Join(strings.Fields(res), " ")
	// fmt.Println(cleanedContent) // cleaned text.
	currentDate := now.New(time.Now()).Format("January 2, 2006")
	day := "todays date is " + currentDate

	// Open AI Setup
	value, exists := os.LookupEnv("OPENAI_API_KEY")
	if !exists {
		log.Println("No OpenAI API key found!")
	}
	client := openai.NewClient(value)

	respObj, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Your task involves analyzing candidates' resumes to identify their technical skills and calculate their years of experience in the IT field, IT framework,It Tools,etc.Excluding internships experience. You will receive textual representations of these resumes and are required to provide three key-value pairs as output in json format: 'Email'(first encountered emailid),'Skills' (a list of technical skills mentioned in the resume) and 'Experience Level' (the number of years of IT experience in number). Please ensure that the 'Experience Level' only returns the numerical value of years of experience.Wiht no extra detials ensure that the results are in the following JSON format: {Email:[Email ID],Skills: [List of technical skills],Experience Level: [Number of years of IT experience in number]} ALWAYS" + day,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: cleanedContent,
				},
			},
			Temperature: 0.2,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	// log.Println("CLEANSED PROMPT TEXT", respObj.Choices[0].Message.Content)
	cleansedtext := cleanseText(respObj)

	return cleansedtext, err
}

//func InsightFile(filePath string) (string, error) {
//	// Extract Text from the file.
//	res, err := ConvertSavedPDFToText(filePath)
//	if err != nil {
//		log.Fatal(" Pdf to text conversion failed ", err)
//		return "", err
//	}
//	// Strips uncessary spaces and newlines
//	cleanedContent := strings.Join(strings.Fields(res), " ")
//	// fmt.Println(cleanedContent) // cleaned text.
//	currentDate := now.New(time.Now()).Format("January 2, 2006")
//	day := "todays date is " + currentDate
//
//	// Open AI Setup
//	value, exists := os.LookupEnv("OPENAI_API_KEY")
//	if !exists {
//		log.Println("No OpenAI API key found!")
//	}
//	client := openai.NewClient(value)
//
//	respObj, err := client.CreateChatCompletion(
//		context.Background(),
//		openai.ChatCompletionRequest{
//			Model: openai.GPT4,
//			Messages: []openai.ChatCompletionMessage{
//				{
//					Role:    openai.ChatMessageRoleSystem,
//					Content: "Your task involves analyzing candidates' resumes to identify their technical skills and calculate their years of experience in the IT field, IT framework,It Tools,etc.Excluding internships experience. You will receive textual representations of these resumes and are required to provide three key-value pairs as output in json format: 'Email'(first encountered emailid),'Skills' (a list of technical skills mentioned in the resume) and 'Experience Level' (the number of years of IT experience). Please ensure that the 'Experience Level' only returns the numerical value of years of experience.Please ensure that the results are in the following JSON format: {Email:[Email ID],Skills: [List of technical skills],Experience Level: [Number of years of IT experience]} ALWAYS" + day,
//				},
//				{
//					Role:    openai.ChatMessageRoleUser,
//					Content: cleanedContent,
//				},
//			},
//			Temperature: 0.2,
//		},
//	)
//
//	if err != nil {
//		fmt.Printf("ChatCompletion error: %v\n", err)
//	}
//	cleansedtext := cleanseText(respObj)
//	return cleansedtext, err
//}
