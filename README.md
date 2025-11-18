# Git-Log

An intelligent CLI tool that automatically transforms your GitHub commit history and pull requests into professional accomplishment summaries using AI. Never manually track your work achievements again - let your git history speak for itself with AI-powered analysis and reporting.

## Technical Implementation

### Backend
- Language: Go 1.24
- GitHub Integration: GitHub REST API (Commits & Pull Requests)
- AI/ML: Google Gemini AI (gemini-2.5-flash)
- Key Features: Commit aggregation, PR analysis, repository grouping, AI-powered report generation

### Infrastructure
- Deployment: Standalone CLI tool
- Configuration: Environment variables with .env file support
- Output: Markdown formatted reports
- CI/CD: Ready for GitHub Actions integration

## Key Achievements

- **Automated Activity Tracking**: Built a comprehensive system that fetches all commits and pull requests from GitHub's REST API, eliminating manual tracking of professional accomplishments.

- **Intelligent Data Processing**: Developed a sophisticated pipeline that groups commits and PRs by repository, filters relevant activity, and structures data for optimal AI analysis.

- **AI-Powered Report Generation**: Integrated Google Gemini AI to analyze git history and generate professional, readable summaries of technical work and achievements.

- **Smart Commit Message Analysis**: Leverages existing commit messages and PR descriptions to automatically document work, encouraging better documentation practices through garbage-in-garbage-out principle.

- **Flexible Configuration**: Implemented environment-based configuration allowing customization of timeframes, AI models, and report output paths.

- **Repository-Based Grouping**: Organizes work activity by repository with comprehensive statistics including total repositories, commits, and pull requests over specified time periods.

## Highlights

### Report Generation Pipeline
1. User configures credentials and settings via environment variables
2. CLI tool connects to GitHub REST API using personal access token
3. Fetches all commits and pull requests for specified user and timeframe
4. Processes and filters activity data, grouping by repository
5. Structures data with metadata (dates, statistics, repository info)
6. Google Gemini AI analyzes the work log using custom system prompt
7. Generates professional markdown report highlighting key accomplishments
8. Saves report to configurable file path for easy sharing

### Data Processing Pipeline
1. GitHub API queries with date-based filtering
2. Commit and PR data extraction with full metadata
3. Repository-based grouping and aggregation
4. Statistical summary generation (totals, date ranges)
5. JSON serialization for AI consumption
6. Iterative report updates with existing report merging
7. Markdown output formatting

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


