package arrays

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	hashmap := map[string]bool{}
	for _, email := range emails {
		unique := getUniqueEmail(email)
		if unique == "" {
			continue
		}
		fmt.Println(unique)
		hashmap[unique] = true
	}
	return len(hashmap)
}

func getUniqueEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	uniqueName := getUniqueName(parts[0])
	validDomain := isDomainValid(parts[1])
	if !validDomain {
		return ""
	}
	return uniqueName + "@" + parts[1]
}

func getUniqueName(name string) string {
	plusSplit := strings.Split(name, "+")
	name = plusSplit[0]
	dotSplit := strings.Split(name, ".")
	name = strings.Join(dotSplit, "")
	return name
}

func isDomainValid(domain string) bool {
	parts := strings.Split(domain, ".")
	if parts[len(parts)-1] != "com" {
		return false
	}
	return true
}
