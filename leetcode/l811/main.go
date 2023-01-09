package l811

import "fmt"

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)

	for _, cpdomain := range cpdomains {
		domains, count := parse(cpdomain)
		for _, domain := range domains {
			m[domain] = m[domain] + count
		}
	}

	res := make([]string, 0, len(m))
	for k, v := range m {
		res = append(res, fmt.Sprintf("%v %s", v, k))
	}
	return res
}

func parse(cpdomain string) (domains []string, count int) {
	n := 0
	ind := 0

	for cpdomain[ind] != ' ' {
		n = n*10 + int(cpdomain[ind]-'0')
		ind++
	}

	domains = make([]string, 0, 8)

	for ind < len(cpdomain) {
		if cpdomain[ind] == '.' || cpdomain[ind] == ' ' {
			domains = append(domains, cpdomain[ind+1:])
		}
		ind++
	}
	return domains, n
}
