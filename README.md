# Git-Log

Automatically generate accomplishment reports from your GitHub commit and PR history using AI.

## What's This About?

I've always been terrible at keeping track of my accomplishments at work. Everyone says you should maintain a log of what you've done for when you need to update your CV or go for a promotion, but I never actually do it. Then I realized that 99% of my work is already tracked on GitHub through my commits and pull requests. So why not just use that data?

This project pulls your GitHub commit and PR history using the GitHub REST API and uses LLMs to turn it into a readable summary of everything you've worked on. It's basically an automated way to document your accomplishments without having to remember to write them down. Plus, knowing that this tool exists has actually inspired me to write better commit messages and create higher quality PRs, since garbage in means garbage out.

## Usage

### Option 1: GitHub Action (Recommended)

The easiest way to use this tool is as a GitHub Action that automatically updates your accomplishment report on a schedule.

#### Quick Start

1. **Create a workflow file** in your repository at `.github/workflows/accomplishment-report.yml`:

```yaml
name: Monthly Accomplishment Report

on:
  schedule:
    - cron: '0 0 1 * *'  # First day of each month at midnight UTC
  workflow_dispatch:  # Allow manual triggering

jobs:
  update-report:
    runs-on: ubuntu-latest
    
    permissions:
      contents: write
    
    steps:
      - name: Checkout
        uses: actions/checkout@v4
      
      - name: Generate Report
        uses: jacantwell/git-log@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          username: ${{ github.actor }}
          google-api-key: ${{ secrets.GOOGLE_API_KEY }}
          model: 'gemini-2.5-flash'
          report-path: 'accomplishments.md'
          lookback_days: 30

      - name: Commit Report
        run: |
          git config user.name "GitHub Actions Bot"
          git config user.email "actions@github.com"
          git add -A
          git diff --quiet && git diff --staged --quiet || (
            git commit -m "Update accomplishment report - $(date +'%Y-%m')"
            git push
          )
```

2. **Add your Google AI Studio API Key**:
   - Go to [Google AI Studio](https://aistudio.google.com/)
   - Create an API key
   - In your GitHub repository, go to Settings > Secrets and variables > Actions
   - Add a new secret named `GOOGLE_API_KEY` with your API key

3. **Done!** The action will run on the first of each month, or you can trigger it manually from the Actions tab.

#### Action Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `github-token` | GitHub token for API access | Yes | - |
| `username` | GitHub username to generate report for | Yes | - |
| `google-api-key` | Google AI Studio API key | Yes | - |
| `days` | Number of days to look back | No | 30 |
| `model` | Google AI model to use | No | `gemini-2.5-flash` |
| `report-path` | Where to save the report | No | `report.md` |


### Option 2: CLI Tool

You can also run this as a standalone CLI tool.

#### Prerequisites

- A GitHub Account with a Personal Access Token (PAT) that has `repo` scope
- A Google AI Studio API Key
- Go 1.24 or later

#### Setup

1. **Clone the repository**:
```bash
git clone https://github.com/jacantwell/git-log.git
cd git-log
```

2. **Create a `.env` file**:
```bash
cp .env.example .env
```

3. **Edit `.env` with your credentials**:
```bash
# Replace with your actual values
USERNAME="your-github-username"
ACCESS_TOKEN="your-github-pat"
GOOGLE_API_KEY="your-google-api-key"

# Optional: customize these
LOKBACK_DAYS=30
MODEL="gemini-2.5-flash"
REPORT_PATH="report.md"
```

4. **Run the script**:
```bash
./run.sh
```

The tool will fetch your GitHub activity, generate a report, and save it to the path specified in `REPORT_PATH`.

## Example Output

The tool generates a structured Markdown report like:

```markdown
# Developer Accomplishment Log

## account-service

### User Account Management API
* Implemented PATCH endpoint for updating user account details with role-based access control (#34)
* Added comprehensive pytest test suite with 120+ tests and 85% coverage requirement (#31)
* Improved error handling with RFC7807-compliant error responses

## billing-system

### Performance Optimization
* Refactored invoice generation to async worker queue using Celery, improving API response latency by 40% (#52)
* Added retry logic and dead-letter handling for failed invoice jobs
* Implemented Prometheus monitoring for queue health metrics

## frontend-portal

### Account Settings UI
* Built new account settings page with reactive form validation and API integration (#77)
* Added comprehensive e2e Playwright tests for settings flows

## 🚧 Work in Progress
* frontend-portal: Began work on new dashboard component (#145)
* billing-system: Exploring webhook integration for payment notifications (#153)
```

## How It Works

1. **Fetch Data**: Uses the GitHub REST API to get all commits and pull requests for your username within the specified time period
2. **Group by Repository**: Organizes all activity by repository for better structure
3. **AI Analysis**: Sends the data to Google's Gemini AI which:
   - Synthesizes commits and PRs into meaningful accomplishments
   - Updates existing report entries (it's a living document, not just a log)
   - Identifies work-in-progress items
   - Generates clear, professional descriptions
4. **Output**: Saves a beautifully formatted Markdown report

## Why This is Useful

- **Performance Reviews**: Have a detailed record of what you've accomplished
- **Resume/CV Updates**: Never forget what you worked on
- **Team Updates**: Share your progress easily
- **Personal Growth**: See how your work evolves over time
- **Better Git Hygiene**: Knowing this tool exists motivates better commit messages

## Notes on AI Models

The default model is `gemini-2.5-flash`, which is fast and accurate.

The free tier of Google AI Studio is generous and sufficient for most personal use.

## Contributing

Issues and pull requests welcome!