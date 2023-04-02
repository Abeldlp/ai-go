# Go-AI

Go-AI is a project that allows you to interact with the ChatGPT language model in your terminal. It's a command-line tool that you can use to generate text, answer questions, or chat with an AI.

## Installation

To use Go-AI, you need to have Go installed on your machine. Once you have Go installed, you can install Go-AI by running the following command:

```bash
$ git clone https://github.com/username/go-ai.git
$ cd go-ai
$ ./install
```

The install script will compile the binary and copy it to your $GOPATH/bin directory. Make sure that $GOPATH/bin is in your $PATH environment variable.

## Usage

Before using you need to export your API key as an environment variable:
(e.g., ~/.bashrc, ~/.zshrc)

```bash
export OPENAI_API_KEY="your_api_key_here"
```

To start using Go-AI, simply type ai followed by your prompt:

```bash
$ ai "Hello, how are you?"
```
