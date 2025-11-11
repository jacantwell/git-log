package report

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func GenerateReport(model string) error {

	systemPromptBytes, err := os.ReadFile("internal/analyser/system_prompt.md")
	if err != nil {
		log.Fatalf("Failed to read system_prompt.md: %v", err)
		return err
	}
	systemPromptString := string(systemPromptBytes)
	systemContent := genai.NewContentFromText(systemPromptString, genai.RoleUser)

	// We check if the file exists. If not (e.g., first run), we use an empty string.
	var reportString string
	reportBytes, err := os.ReadFile("report.md")
	if err != nil {
		if os.IsNotExist(err) {
			reportString = ""
		} else {
			log.Fatalf("Failed to read report.md: %v", err)
			return err
		}
	} else {
		reportString = string(reportBytes)
	}

	// Load the work log. This file should exist.
	logBytes, err := os.ReadFile("work_log.json")
	if err != nil {
		log.Fatalf("Failed to read work_log.json: %v", err)
		return err
	}
	logString := string(logBytes)

	userPromptTemplate := `Here is the existing accomplishment report and the work log for the past month.

Please update and merge the report according to your system instructions.

report.md:

%s

work_log.json:

%s
`
	// Use fmt.Sprintf to "paste" the file contents into the template
	finalUserPromptString := fmt.Sprintf(userPromptTemplate, reportString, logString)

	// --- 4. Set up the API Client and Config ---
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
		return err
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: systemContent,
	}

	userMessage := genai.Text(finalUserPromptString)

	result, err := client.Models.GenerateContent(
		ctx,
		model,
		userMessage,
		config,
	)
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
		return err
	}

	// The result.Text() will be the complete, updated Markdown report
	fmt.Println(result.Text())

	return nil
}
