# Go Coding Challenge: Resume Parser using ChatGPT API

## Overview

Your task is to develop a Go service that uses the [ChatGPT API](https://platform.openai.com/docs/guides/chat) to parse resumes and extract skills and experience level. The service will expose a REST API accepting POST requests to upload a resume file, process it through the ChatGPT API, and store the parsed information in a PostgreSQL database.

## Requirements

1. **GitHub**: Create a specific branch with your name for this task.
2. **API Endpoint**: Create a single POST endpoint `/upload-resume` for resume file uploads.
3. **File Handling**: Validate the uploaded file format (e.g., PDF, DOCX).
4. **ChatGPT API**: Use the ChatGPT API to extract information. Craft prompts to identify the candidate's skills and experience level.
5. **PostgreSQL**: Store the parsed data in a PostgreSQL database.
6. **Testing**: Unit tests for the code.
7. **Dockerization**: `Dockerfile` and `docker-compose.yml` for containerization.
8. **Code Quality**: Go idioms, choice of libraries, performance, and readability.

## Deliverables

1. Go program for the REST API and ChatGPT interaction.
2. Unit test file.
3. `Dockerfile`.
4. `docker-compose.yml`.
5. `INSTRUCTIONS.md` for build/run steps.

## Evaluation Criteria

- Go idioms and best practices
- Code quality
- Library choice
- Performance
- Readability
- Test coverage

## Bonus

- Rate-limiting for the API.
- Database migrations via [golang-migrate](https://github.com/golang-migrate/migrate).
- Use of [sqlc](https://github.com/kyleconroy/sqlc) for database interactions.

## ChatGPT Prompts

For the ChatGPT API, consider using prompts formatted as follows to extract the desired information:

- **Skills**: "List the IT skills mentioned in this resume: {Resume_Content}"
- **Experience Level**: "How many years of experience does the candidate have in the IT field as mentioned in this resume? {Resume_Content}"

You can encode the skills as a JSON array and experience level as an integer value.

While the above requirements outline a comprehensive task, it's understood that time constraints or other factors may limit what you can achieve. What's most important is not the full completion of all requirements, but rather your ability to deliver a functional solution and thoughtfully argue your design choices and implementation. Your approach to problem-solving and decision-making are as much a part of this exercise as the code itself.
