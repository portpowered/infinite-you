package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	publicDocsDir      = "docs"
	publicExamplesDir  = "examples"
	publicFactoryDir   = "factory"
	publicReadmePath   = "README.md"
	publicInitPath     = "pkg/cli/init/init.go"
	internalDocsPrefix = "docs/development/"
)

var couplingTerms = []string{"portos", "port os", "port_os"}

type publicSurfaceTarget struct {
	path      string
	excludeIf func(string) bool
}

type finding struct {
	path    string
	line    int
	term    string
	content string
}

func main() {
	findings, err := scanPublicSurface(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "scan public surface: %v\n", err)
		os.Exit(1)
	}
	if len(findings) == 0 {
		fmt.Println("[agent-factory:public-surface] no Port OS coupling found")
		return
	}

	fmt.Fprintf(os.Stderr, "public-surface Port OS coupling detected (%d findings):\n", len(findings))
	for _, finding := range findings {
		fmt.Fprintf(os.Stderr, "  %s:%d: matched %q in %q\n", finding.path, finding.line, finding.term, finding.content)
	}
	os.Exit(1)
}

func scanPublicSurface(root string) ([]finding, error) {
	targets := []publicSurfaceTarget{
		{path: publicReadmePath},
		{
			path: publicDocsDir,
			excludeIf: func(path string) bool {
				return isInternalDocsPath(path)
			},
		},
		{path: publicExamplesDir},
		{path: publicFactoryDir},
		{path: publicInitPath},
	}

	findings := make([]finding, 0)
	for _, target := range targets {
		targetPath := filepath.Join(root, filepath.FromSlash(target.path))
		info, err := os.Stat(targetPath)
		if err != nil {
			return nil, fmt.Errorf("stat %s: %w", target.path, err)
		}

		if !info.IsDir() {
			fileFindings, err := scanFile(root, target.path)
			if err != nil {
				return nil, err
			}
			findings = append(findings, fileFindings...)
			continue
		}

		err = filepath.WalkDir(targetPath, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}

			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				return fmt.Errorf("rel path for %s: %w", path, err)
			}
			relativePath = filepath.ToSlash(relativePath)

			if target.excludeIf != nil && target.excludeIf(relativePath) {
				if entry.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
			if entry.IsDir() {
				return nil
			}

			fileFindings, err := scanFile(root, relativePath)
			if err != nil {
				return err
			}
			findings = append(findings, fileFindings...)
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("walk %s: %w", target.path, err)
		}
	}

	return findings, nil
}

func scanFile(root string, relativePath string) ([]finding, error) {
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(relativePath)))
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", relativePath, err)
	}
	defer file.Close()

	findings := make([]finding, 0)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		lowerLine := strings.ToLower(line)
		for _, term := range couplingTerms {
			if !strings.Contains(lowerLine, term) {
				continue
			}
			findings = append(findings, finding{
				path:    filepath.ToSlash(relativePath),
				line:    lineNumber,
				term:    term,
				content: strings.TrimSpace(line),
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan %s: %w", relativePath, err)
	}

	return findings, nil
}

func isInternalDocsPath(path string) bool {
	return path == strings.TrimSuffix(internalDocsPrefix, "/") || strings.HasPrefix(path, internalDocsPrefix)
}
