package pwndoctor

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func DoImport(auditName string, findingsDir string) {
	allAudits, err := pwndocAPI.GetAudits()
	if err != nil {
		fmt.Printf("Error getting audits from pwndoc: %s", err)
	}

	auditExists := false
	var auditID string
	for _, audit := range allAudits.Data {
		if audit.Name == auditName {
			auditExists = true
			auditID = audit.ID
			break
		}
	}

	if !auditExists {
		fmt.Printf("[!] Error: Audit %s not found in Pwndoc! Please enter a valid audit name and try again", auditName)
		os.Exit(1)
	}

	fmt.Printf("\n[+] Audit found in PwnDoc! Importing findings from %s into audit %s", findingsDir, auditName)
	err = ImportFindings(auditID, findingsDir)
	if err != nil {
		log.Fatal("[-] Error import audit response body (import audit): ", err)
	}
	fmt.Println("\n[+] Done importing audit info...")
	fmt.Println("\n[+] Done importing audit info...")

}

func ImportFindings(auditID string, findingsDir string) error {
	auditFilePath := filepath.Join(findingsDir, "audit-findings")
	files, err := os.ReadDir(auditFilePath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	//Only selecting entries that are JSON files
	var fileNames []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			fileNames = append(fileNames, filepath.Join(auditFilePath, file.Name()))
		}
	}

	for _, finding := range fileNames {
		data, err := os.ReadFile(finding)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", finding, err)
			continue
		}

		byteData := bytes.NewReader(data)

		APIPostResponse, err := pwndocAPI.InsertFinding(auditID, byteData)
		if err != nil {
			fmt.Printf("\n[-] Error: An issue occurred inserting the finding %s: %v\n", finding, err)
			break
		}
		fmt.Printf("\n[+] POST Response is %s", string(APIPostResponse))

	}
	return nil

}
