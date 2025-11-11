# Git-Log

## What's This About?

I've always been terrible at keeping track of my accomplishments at work. Everyone says you should maintain a log of what you've done for when you need to update your CV or go for a promotion, but I never actually do it. Then I realized that 99% of my work is already tracked on GitHub through my commits and pull requests. So why not just use that data?

This project pulls your GitHub commit and PR history using the GitHub REST API and uses LLMs to turn it into a readable summary of everything you've worked on. It's basically an automated way to document your accomplishments without having to remember to write them down. Plus, knowing that this tool exists has actually inspired me to write better commit messages and create higher quality PRs, since garbage in means garbage out.

## Usage

### Prerequisites

You will need the following accounts and keys:

- A GitHub Account with a Personal Access Token (PAT). The PAT must have the necessary read permissions (repo or similar scope) to access your commit and PR history.

- A Google AI Studio API Key. You can use smaller models on Google AI Studio for free.

### Setup and Configuration

- Create the Environment File (.env)
    
    Create a file named .env in the root of the project directory and populate it with your credentials:

    ```
    # Replace with your actual GitHub username
    USERNAME="YOUR_GITHUB_USERNAME" 

    # Replace with your GitHub Personal Access Token (PAT)
    ACCESS_TOKEN="YOUR_GITHUB_PAT"

    # Replace with your Google AI Studio API Key
    GOOGLE_API_KEY="YOUR_GOOGLE_API_KEY"
    ```

- Run the Script

    Execute the provided shell script to generate your report:
    
    ```bash
    ./run.sh
    ```

    The script will fetch your history, generate the summaries using the LLM, and save the output.

### Advanced Configuration (Optional)

You can customize the report generation by editing the config/config.go file. The following settings are typically available to adjust:

- Days: Change the number of days the tool will search back for commits and PRs.

- Model: Specify a different Google AI model to use for the summarization.

- ReportPath: Define the file path where the final summary report will be saved.


