package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "sort"
)

func main() {
    generateLandingPage()
}

func generateLandingPage() {
    // Read all markdown files in the content dir
    files, err := os.ReadDir(filepath.Join("web", "content"))
    if err != nil {
        log.Fatal(err)
    }

    // Create a list of markdown filenames
    var markdownFiles []string
    for _, file := range files {
        if strings.HasSuffix(file.Name(), ".md") && file.Name() != "root.md" {
            markdownFiles = append(markdownFiles, file.Name())
        }
    }

    markdownContent := generateMarkdownContent(markdownFiles)

    outputFile := "web/content/root.md"
    err = os.WriteFile(outputFile, []byte(markdownContent), 0755)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Markdown file generated and saved to %s\n", outputFile)
}

func generateMarkdownContent(files []string) string {
    // Sort files alphabetically
    sort.Strings(files)

    var content strings.Builder

    content.WriteString("# All projects\n\n")
    content.WriteString("|      |                                   |\n")
    content.WriteString("|-------|-----------------------------------|\n")

    for _, file := range files {
        // Get file modification time
        fileInfo, err := os.Stat("web/content/" + file)
        if err != nil {
            log.Fatal(err)
        }
        modTime := fileInfo.ModTime().Year()

        link := fmt.Sprintf("| %d | [%s][] |\n", modTime, strings.TrimSuffix(file, ".md"))
        content.WriteString(link)
    }

    content.WriteString("\n")

    // Append file links
    for _, file := range files {
        link := fmt.Sprintf("[%s]: %s\n", strings.TrimSuffix(file, ".md"), strings.TrimSuffix(file, ".md"))
        content.WriteString(link)
    }

    return content.String()
}
