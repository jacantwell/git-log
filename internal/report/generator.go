package report

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"git-log/internal/processing"
	"google.golang.org/genai"
)

func GenerateReport(model string, workLog processing.WorkLog, reportPath string) (string, error) {

	// Use the embedded system prompt
	systemContent := genai.NewContentFromText(SystemPrompt, genai.RoleUser)

	// We check if the file exists. If not (e.g., first run), we use an empty string.
	var reportString string
	reportBytes, err := os.ReadFile(reportPath)
	if err != nil {
		if os.IsNotExist(err) {
			reportString = ""
		} else {
			log.Fatalf("Failed to read report file: %v", err)
			return "", err
		}
	} else {
		reportString = string(reportBytes)
	}

	logBytes, err := json.Marshal(workLog)
	if err != nil {
		fmt.Printf("Error converting workLog to JSON: %v\n", err)
		return "", err
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
		return "", err
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
		return "", err
	}

	// The result.Text() will be the complete, updated Markdown report
	fmt.Println(result.Text())

	return result.Text(), nil
}
