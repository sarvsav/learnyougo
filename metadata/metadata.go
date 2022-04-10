package metadata

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func problemCount(rootDir string) (int, error) {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		return 0, err
	}
	return len(files), nil
}

type ProblemConfig struct {
	Title    string
	FileData fileData
}

type fileData struct {
	Filename string
	Problem  string
	Hint     string
	Language []string
}

func ProblemList(rootDir string) ([]string, error) {
	var config ProblemConfig
	var data []string
	problems, err := problemCount(rootDir)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= problems; i++ {
		_, err := toml.DecodeFile(rootDir+fmt.Sprintf("%d", i)+"/problem.toml", &config)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		data = append(data, config.Title)
	}
	return data, err
}

func ProblemDescription(rootDir string, num int) (string, error) {
	var config ProblemConfig
	basePath := rootDir + fmt.Sprintf("%d", num) + "/"
	_, err := toml.DecodeFile(basePath+"problem.toml", &config)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	description, err := os.ReadFile(basePath + config.FileData.Problem + ".en" + ".md")
	return string(description), err
}

func ProblemHint(rootDir string, num int) (string, error) {
	var config ProblemConfig
	basePath := rootDir + fmt.Sprintf("%d", num) + "/"
	_, err := toml.DecodeFile(basePath+"problem.toml", &config)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	hint, err := os.ReadFile(basePath + config.FileData.Hint + ".en" + ".md")
	return string(hint), err
}
