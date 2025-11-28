package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func mkDir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}

func getText(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return ""
	}
	return string(data)
}

func cp(path1 string, path2 string, year string, day string) {
	file := getText(path1)
	file = strings.ReplaceAll(file, "%annee", year)
	file = strings.ReplaceAll(file, "%jour", day)
	err := os.WriteFile(path2, []byte(file), 0644)
	if err != nil {
		fmt.Println("Error writing to destination file:", err)
		return
	}
}

func createFile(path string, content io.Reader) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, content)
	if err != nil {
		panic(err)
	}
}

func createFile2(path string, data []byte) {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	pathFolder := fmt.Sprintf("%s/%s", args[1], args[2])
	mkDir(pathFolder)
	cp("templates/day.go", fmt.Sprintf("%s/day%s.go", pathFolder, args[2]), args[1], args[2])
	cp("templates/day_test.go", fmt.Sprintf("%s/day%s_test.go", pathFolder, args[2]), args[1], args[2])
	if args[2][0] == '0' {
		args[2] = string(args[2][1])
	}
	urlInput := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", args[1], args[2])
	req, err := http.NewRequest("GET", urlInput, nil)
	if err != nil {
		panic(err)
	}
	err = godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	sessionValue := os.Getenv("SESSION_COOKIE")
	if sessionValue == "" {
		panic("SESSION_COOKIE environment variable is not set")
	}
	sessionCookie := &http.Cookie{
		Name:  "session",
		Value: sessionValue,
	}
	req.AddCookie(sessionCookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Échec du téléchargement : " + resp.Status)
	}
	contentBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	createFile2(fmt.Sprintf("%s/input.txt", pathFolder), contentBytes)
	var dirPathInput, _ = filepath.Abs(fmt.Sprintf("../../advent-of-code-inputs/%s", pathFolder))
	mkDir(dirPathInput)
	createFile2(fmt.Sprintf("%s/input.txt", dirPathInput), contentBytes)
	var pathGit, _ = filepath.Abs("../../advent-of-code-inputs")
	err = os.Chdir(pathGit)
	if err != nil {
		fmt.Println("Erreur lors du changement de répertoire :", err)
		return
	}
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = pathFolder
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error adding files to git:", err)
		return
	}

	cmd = exec.Command("git", "commit", "-m", fmt.Sprintf("Add input for year %s day %s", args[1], args[2]))
	cmd.Dir = pathFolder
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error committing files to git:", err)
		return
	}

	cmd = exec.Command("git", "push")
	cmd.Dir = pathFolder
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error pushing files to git:", err)
		return
	}
}
