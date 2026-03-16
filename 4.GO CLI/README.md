# Tasks
Create a Go command-line application that summarizes the contents of a text file using a free, public GenAI API (such as HuggingFace Inference API or any other public endpoint).

* The CLI must be written in Go.
* The CLI must accept:
  * --input or a positional argument: path to the text file to summarize.
  * --type or -t : summary type, one of short , medium , or bullet .
  * The CLI must call a free, public GenAI API for summarization (e.g., HuggingFace Inference API).
* The prompt sent to the API should be engineered to match the summary type:
  * short : a concise summary (1-2 sentences)
  * medium : a paragraph summary
  * bullet : a list of bullet points
* The CLI should output the summary to stdout.
* The CLI should handle API errors gracefully and print user-friendly messages.
* Document the Go version used in your code comments.

For the solution, it was used https://huggingface.co/ and the documentation API can be found here: https://huggingface.co/docs/inference-providers/en/index and the code is in the file called **solution_summarizer.go**

use:
```
go run solution_summarizer.go --input article.txt --type bullet
```
```
go run solution_summarizer.go -t short article.txt
```
