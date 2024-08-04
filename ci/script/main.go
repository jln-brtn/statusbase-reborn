package main

import (
	"encoding/csv"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Owner         string `yaml:"owner"`
	Repo          string `yaml:"repo"`
	Branch        string `yaml:"branch"`
	StatusWebsite struct {
		CName       string `yaml:"cname"`
		LogoURL     string `yaml:"logoUrl"`
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
	} `yaml:"status-website"`
	Groups []Group `yaml:"groups"`
}

type Group struct {
	Name  string `yaml:"name"`
	Slug  string `yaml:"slug"`
	Sites []Site `yaml:"sites"`
}

type Site struct {
	Name string `yaml:"name"`
	Desc string `yaml:"desc"`
	Slug string `yaml:"slug"`
}

type StatusEntry struct {
	Time   time.Time `csv:"time"`
	Status string    `csv:"status"`
}

const (
	maxDays = 45
)

func main() {
	check()
	envGen()
}

func check() {
	config := readConfig("config.yml")
	now := time.Now().UTC()
	for _, group := range config.Groups {
		groupStatus := "success"
		for _, site := range group.Sites {
			siteStatus := checkSite(site, now)
			fmt.Println("Site", site.Name, "has status", siteStatus)
			if siteStatus == "error" {
				groupStatus = "error"
			}
		}
		checkGroup(group, now, groupStatus)
	}
}

func readConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func checkSite(site Site, now time.Time) string {
	url := site.Desc
	resp, err := http.Get(url)
	status := "error"
	if err == nil && resp.StatusCode == http.StatusOK {
		status = "success"
	}
	if resp != nil {
		resp.Body.Close()
	}

	logFilePath := filepath.Join("logs", site.Slug+".csv")
	ensureLogFile(logFilePath)
	writeStatus(logFilePath, now, status)
	return status
}

func checkGroup(group Group, now time.Time, status string) {
	logFilePath := filepath.Join("logs", group.Slug+".csv")
	ensureLogFile(logFilePath)
	writeStatus(logFilePath, now, status)
}

func ensureLogFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		writer.Write([]string{"time", "status"})
		writer.Flush()
	}
}

func writeStatus(path string, now time.Time, status string) {
	entries := readLogEntries(path)
	entries = removeOldDates(entries)

	if now.Day() != entries[len(entries)-1].Time.Day() { //if this is the first run of the day
		yesterday := now.AddDate(0, 0, -1)
		yesterdayAt2359 := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 0, 0, time.UTC)
		entries = append(entries, StatusEntry{yesterdayAt2359, status})
		entries = append(entries, StatusEntry{now, status})
	} else if len(entries) == 0 { //if there is no record before
		entries = append(entries, StatusEntry{now, status})
	} else if entries[len(entries)-1].Status != status { //if the previous record is different from
		entries = append(entries, StatusEntry{now, status})
	}

	writeLogEntries(path, entries)
}

func readLogEntries(path string) []StatusEntry {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var entries []StatusEntry
	reader := csv.NewReader(file)
	reader.Read() // skip header
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		time, err := time.Parse(time.RFC3339, record[0])
		if err != nil {
			panic(err)
		}
		entries = append(entries, StatusEntry{Time: time, Status: record[1]})
	}
	return entries
}

func removeOldDates(entries []StatusEntry) []StatusEntry {
	now := time.Now().UTC()
	dateLimit := now.AddDate(0, 0, -maxDays)

	var filteredEntries []StatusEntry
	for _, entry := range entries {
		if entry.Time.After(dateLimit) || entry.Time.Equal(dateLimit) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries
}

func writeLogEntries(path string, entries []StatusEntry) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"time", "status"})
	for _, entry := range entries {
		writer.Write([]string{entry.Time.Format(time.RFC3339), entry.Status})
	}
	writer.Flush()
}

func envGen() {
	// Read the YAML file
	data, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	// Parse the YAML data
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing YAML file:", err)
		return
	}

	// Create the .env file content
	envContent := fmt.Sprintf(
		"OWNER=%s\nREPO=%s\nCNAME=%s\nLOGO_URL=%s\nNAME=%s\nDESCRIPTION=%s\n",
		config.Owner,
		config.Repo,
		config.StatusWebsite.CName,
		config.StatusWebsite.LogoURL,
		config.StatusWebsite.Name,
		config.StatusWebsite.Description,
	)

	// Write to the .env file
	err = os.WriteFile(".env", []byte(envContent), 0644)
	if err != nil {
		fmt.Println("Error writing .env file:", err)
		return
	}

	fmt.Println(".env file generated successfully")
}
