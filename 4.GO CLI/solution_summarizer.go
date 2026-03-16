// Go version: 1.26.1
// API Documentation: https://huggingface.co/docs/inference-providers/en/index
// Model used: openai/gpt-oss-120b:fastest

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// 1. estructuras para el formato de Chat Completions
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func main() {
	apiToken := os.Getenv("HF_TOKEN")
	// ejecutar primero powerhshell para crear la variable en la sesion
	// $env:HF_TOKEN="hf_tu_token_aqui_123"
	// 1. CLI Argument Parsing
	if apiToken == "" {
		fmt.Println("Error: Missing HF_TOKEN environment variable.")
		os.Exit(1)
	}
	var summaryType string
	var inputFile string

	// se define flag para el tipo de resumen y el archivo de entrada
	flag.StringVar(&summaryType, "type", "", "Summary type: short, medium, or bullet (required)")
	flag.StringVar(&summaryType, "t", "", "Summary type (shorthand, required)")
	flag.StringVar(&inputFile, "input", "", "Path to the text file to summarize")

	flag.Parse()

	// If --input is not explicitly used, fallback to the first positional argument.
	// This satisfies the requirement: `go run solution_summarizer.go -t short article.txt`
	if inputFile == "" && flag.NArg() > 0 {
		inputFile = flag.Arg(0)
	}

	// Validation
	if inputFile == "" {
		fmt.Println("Error: Missing input file.")
		fmt.Println("Usage: go run solution_summarizer.go [--input file.txt] [-t/--type short|medium|bullet]")
		os.Exit(1)
	}

	summaryType = strings.ToLower(summaryType)
	if summaryType == "" {
		fmt.Println("Error: Missing required flag -t (or --type).")
		fmt.Println("Usage: go run solution_summarizer.go --input file.txt -t/--type short|medium|bullet")
		os.Exit(1)
	}

	if summaryType != "short" && summaryType != "medium" && summaryType != "bullet" {
		fmt.Println("Error: Invalid summary type. Must be 'short', 'medium', or 'bullet'.")
		os.Exit(1)
	}

	// 2. Read File Contents
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", inputFile, err)
		os.Exit(1)
	}
	text := string(content)

	// 3. Prompt Engineering
	var instruction string
	switch summaryType {
	case "short":
		instruction = `Write a concise summary (1-2 sentences) of the following text. You must always start your response with the exact phrase "The article discusses":\n\n`
	case "medium":
		instruction = `Write a comprehensive paragraph summary of the following text. You must always start your response with the exact phrase "The article discusses":\n\n`
	case "bullet":
		instruction = "Summarize the following text as a list of bullet points. Start each point with a dash (-):\n\n"
	}

	// Mistral specific instruction format for better results
	prompt := fmt.Sprintf("%s%s", instruction, text)

	// 4. API Request Setup
	apiURL := "https://router.huggingface.co/v1/chat/completions"

	reqBody := ChatRequest{
		Model: "openai/gpt-oss-120b:fastest",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens: 300,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("Internal error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	// 5. Execute API Call
	// Added a 15-second timeout as a best practice for CLI tools
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Network error while connecting to the GenAI API: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	// 6. Graceful Error Handling
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API Error (Status %d): Model might be loading or busy.\n", resp.StatusCode)
		fmt.Printf("Raw details: %s\n", string(bodyBytes))
		os.Exit(1)
	}

	// 7. Parse and Output (Navegando por la estructura de Choices -> Message -> Content)
	var apiResponse ChatResponse
	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil || len(apiResponse.Choices) == 0 {
		fmt.Println("Error parsing API response. The model may have returned an unexpected format.")
		os.Exit(1)
	}

	// Imprimir limpiamente el resultado en consola
	fmt.Println(strings.TrimSpace(apiResponse.Choices[0].Message.Content))
}
