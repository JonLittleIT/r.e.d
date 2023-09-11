package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/projectdiscovery/httpx/v2/pkg/httpx"
	"github.com/projectdiscovery/nuclei/v2/pkg/template"
	"github.com/projectdiscovery/subfinder/v2/pkg/subscraping"
)

func main() {
	// Defining the target domain
	domain := "amazon.com"

	// Running the subdomain finder
	subdomains, err := subscraping.Subdomains(domain, false)
	if err != nil {
		log.Fatal("Error running subscraping.Subdomains:", err)
	}

	// Saving the output to a file
	subdomainsFile, err := os.Create("subdomains.txt")
	if err != nil {
		log.Fatal("Error creating subdomains.txt:", err)
	}
	defer subdomainsFile.Close()

	for _, subdomain := range subdomains {
		_, err := subdomainsFile.WriteString(subdomain + "\n")
		if err != nil {
			log.Println("Error writing to subdomains.txt:", err)
		}
	}

	// Reading the subdomains from the file
	subdomainList, err := subscraping.ReadSubdomainList("subdomains.txt")
	if err != nil {
		log.Fatal("Error reading subdomain list from subdomains.txt:", err)
	}

	// Using httpx to check each subdomain
	for _, subdomain := range subdomainList {
		url := subdomain
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		resp, err := httpx.Get(url)
		if err != nil {
			log.Println("Error making HTTP request:", err)
			continue
		}

		// Saving the output to a file
		httpxOutputFile, err := os.Create("httpx_output.txt")
		if err != nil {
			log.Println("Error creating httpx_output.txt:", err)
			continue
		}
		defer httpxOutputFile.Close()

		_, err = httpxOutputFile.WriteString(resp.String())
		if err != nil {
			log.Println("Error writing to httpx_output.txt:", err)
		}
	}

	// Defining the rate limit and user agent
	options := template.Options{
		RateLimit: 5,
		UserAgent: "amazonvrpresearcher_crashoverrid3",
	}

	// Running nuclei with the given options
	err = template.Run("httpx_output.txt", "nuclei_output.txt", options)
	if err != nil {
		log.Println("Error running nuclei:", err)
	}

	// Defining the Slack webhook URL
	webhookURL := "https://hooks.slack.com/services/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"

	// Reading the output from the nuclei file
	nucleiOutputFile, err := os.Open("nuclei_output.txt")
	if err != nil {
		log.Println("Error opening nuclei_output.txt:", err)
	}
	defer nucleiOutputFile.Close()

	output, err := ioutil.ReadAll(nucleiOutputFile)
	if err != nil {
		log.Println("Error reading from nuclei_output.txt:", err)
	}

	// Sending the output to the Slack webhook
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(output))
	if err != nil {
		log.Println("Error sending data to Slack webhook:", err)
	}
	defer resp.Body.Close()
}
