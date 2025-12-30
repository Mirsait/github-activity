# GitHub Activity CLI

## Overview
`github-activity` is a simple command-line application that retrieves and 
displays a GitHub user's recent activity. It uses the GitHub API to fetch events
such as commits, issues, pull requests, and other public actions.

[roadmap.sh](https://roadmap.sh/projects/github-user-activity)

## Features
- Run directly from the command line.
- Accept a GitHub username as an argument.
- Fetch recent activity using the GitHub API.
- Display results in a readable format in the terminal.
- Save fetched activity data locally in a `data` folder located next to the 
application.

## Getting started

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/Mirsait/github-activity.git
cd github-activity
```

Build the project:

```bash
go build -o github-activity
```

## Usage

Run the application from the command line, passing the GitHub username as an 
argument:

```bash
./github-activity <username>
```

Example:

```bash
./github-activity Mirsait
```

This will display the recent public activity of the specified GitHub user.

## Example Output
```
- Pushed 4 commits in Mirsait/task-cli
- Merged 3 pull requests in Mirsait/task-cli
- Opened 3 pull requests in Mirsait/task-cli
- Created 4 branches in Mirsait/task-cli
- Opened 9 pull requests in Mirsait/advent-of-code-2025
....
```

## Notes
- The application uses the public GitHub API
- Only public activity is displayed.

## License

[MIT License](LICENSE) — feel free to use, modify, and distribute.
