package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	searchRoot := filepath.Join(os.Getenv("GOPATH"), "src")
	var resultPath string

	searchTerm := os.Args[1]

	repoPaths := digRepoNames(searchRoot)

	var chosenCandidateScore int

	for _, repoPath := range repoPaths {
		candidate := filepath.Base(repoPath)
		score := searchScore(searchTerm, candidate)

		if score > chosenCandidateScore {
			chosenCandidateScore = score
			resultPath = repoPath
		}
	}

	fmt.Println(resultPath)
}

func searchScore(searchTerm, candidate string) int {
	candidateLetters := strings.Split(candidate, "")
	var score int
	for candidateIndex, candidateLetter := range candidateLetters {
		searchLetters := strings.Split(searchTerm, "")
		for searchIndex, searchLetter := range searchLetters {
			if candidateLetter == searchLetter {
				score += 1

				if candidateIndex == searchIndex {
					score += 1
				}
			}
		}
	}

	return score
}

func digRepoNames(dir string) []string {
	var repoNames []string

	info, _ := os.Lstat(dir)

	if !info.IsDir() {
		return repoNames
	}

	isGitRepo, err := findGitRepo(dir)

	if err != nil {
		panic(fmt.Sprintf("error finding go files: %v", err))
	}

	if isGitRepo {
		repoNames = append(repoNames, dir)
		return repoNames
	}

	dirPaths := walk(dir)

	for _, dirPath := range dirPaths {
		subDirPaths := digRepoNames(dirPath)
		repoNames = append(repoNames, subDirPaths...)
	}

	return repoNames
}

func findGitRepo(dir string) (bool, error) {
	var isRepo bool

	f, err := os.Open(dir)
	if err != nil {
		return isRepo, err
	}

	dirNames, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return isRepo, err
	}

	for _, dirName := range dirNames {
		containsGitConfig, err := regexp.MatchString(".git", dirName)

		if err != nil {
			return isRepo, err
		}

		if containsGitConfig {
			isRepo = true
		}
	}

	return isRepo, nil
}

func walk(dir string) []string {
	var dirPaths []string
	dirNames, err := readDirNames(dir)
	if err != nil {
		panic("could not read dirnames")
	}

	for _, dirName := range dirNames {
		dirPath := filepath.Join(dir, dirName)
		dirPaths = append(dirPaths, dirPath)
	}

	return dirPaths
}

func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}
