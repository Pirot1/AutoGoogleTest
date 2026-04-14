# AutoGoogleDocs Solver

![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue?logo=go)
![Ollama](https://img.shields.io/badge/Ollama-Llama3.1-black?logo=ollama)
![Browser Automation](https://img.shields.io/badge/go--rod-v0.116-green)
![License](https://img.shields.io/badge/License-MIT-yellow)
![Status](https://img.shields.io/badge/Status-Active-brightgreen)


A robust automation tool designed to solve Google Forms by leveraging local AI (Ollama). This script automatically parses questions, analyzes them using an AI model, and selects the correct answers with high precision.

---

## 🚀 Key Features
* **Intelligent Parsing**: Uses go-rod to navigate Google Forms and extract questions.

* **Smart Filtering**: Automatically identifies and skips non-question elements (headers, names, group inputs) to prevent answer offsets.

* **AI Integration**: Seamlessly processes questions via local LLMs using Ollama.

* **Batch Processing**: Processes questions in chunks of 5 to balance performance and stability.

* **Real-time Execution**: Clicks answers immediately upon receiving AI responses, ensuring high reliability

---

## 🛠Tech Stack
* **Language**: Go
* **Browser Automation**: [go-rod](https://github.com/go-rod/rod)
* **AI Engine**: [Ollama(llama3.1)](https://ollama.com/library/llama3.1)
  
---

## ⚙️ Setup & Usage

1) **Requirements**
    * Go (version 1.20+)
    * Ollama (with a model pulled, e.g., ollama run llama3)
2) **Installation**
```bash
# Clone the repository
git clone <your-repository-url>
cd AutoGoogleDocs

# Install dependencies
go mod tidy
```
3) **Running the Bot**
Make sure the Ollama server is running in the background:
```bash
ollama serve
```
Then run the application:
```bash
go run main.go
```

---

## ⚠️ Project Structure
```text
.
├── cmd/
│   └── bot/
│       └── main.go           # Entry point (Application startup)
├── pkg/
│   ├── ai/
│   │   └── ollama.go         # API-based answer selection (multiple choice 1-4)
│   ├── browser/
│   │   └── init.go           # Browser initialization (Headless mode, etc.)
│   └── parser/
│       ├── login.go          # Site login and session management
│       └── test.go           # Answering question and question parser
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
├── .gitignore                # Files excluded from version control
└── README.md                 # Project documentation (You are here)
```
---
## 💡 Best Practices

* **XPath Selectors**: If Google updates the layout of their forms, verify the XPath selectors in `parser/`.

* **Environment Variables**: Do not hardcode credentials. It is recommended to use an `.env` file for your Gmail login and password.

* **Headless Mode**: If you don't need to see the browser, you can configure rod to `run` in `--headless` mode.

---

## 🛡️ Troubleshooting
* **"Target index invalid"**: Usually occurs if the script tries to click a non-question element. The current version includes auto-filtering to prevent this.

* **Dependencies**: If you encounter `undefined` errors, ensure you run g`o mod tidy` to sync the `rod` library and its `lib/input` submodule.