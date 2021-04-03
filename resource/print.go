package resource

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/apex/log"
	"github.com/jckuester/awstools-lib/terraform"
)

func PrintResources(resources []terraform.Resource, hasAttrs map[string]bool, attributes []string) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	printHeader(w, attributes)

	for _, r := range resources {
		profile := `N/A`
		if r.Profile != "" {
			profile = r.Profile
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s", r.Type, r.ID, profile, r.Region)

		if r.CreatedAt != nil {
			fmt.Fprintf(w, "\t%s", r.CreatedAt.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Fprint(w, "\tN/A")
		}

		for _, attr := range attributes {
			v := "N/A"

			_, ok := hasAttrs[attr]
			if ok {
				var err error
				v, err = GetAttribute(attr, &r)
				if err != nil {
					log.WithFields(log.Fields{
						"type": r.Type,
						"id":   r.ID}).WithError(err).Debug("failed to get attribute")
					v = "error"
				}
			}

			fmt.Fprintf(w, "\t%s", v)
		}

		fmt.Fprintf(w, "\t\n")
	}

	w.Flush()
	fmt.Println()
}

func PrintResourcesAsJSON(resources []terraform.Resource, hasAttrs map[string]bool, attributes []string) {
	results := make([]map[string]string, 0)

	for _, r := range resources {
		profile := `N/A`
		if r.Profile != "" {
			profile = r.Profile
		}

		res := map[string]string{}
		res["type"] = r.Type
		res["id"] = r.ID
		res["profile"] = profile
		res["region"] = r.Region

		if r.CreatedAt != nil {
			res["created_at"] = r.CreatedAt.Format("2006-01-02 15:04:05")
		} else {
			res["created_at"] = "N/A"
		}

		for _, attr := range attributes {
			v := "N/A"

			_, ok := hasAttrs[attr]
			if ok {
				var err error
				v, err = GetAttribute(attr, &r)
				if err != nil {
					log.WithFields(log.Fields{
						"type": r.Type,
						"id":   r.ID}).WithError(err).Debug("failed to get attribute")
					v = "error"
				}
			}

			res[strings.ToLower(attr)] = v
		}

		results = append(results, res)
	}

	jsonStream, err := json.Marshal(results)
	if err != nil {
		log.WithError(err).Debug("failed to marshal output to JSON")
	}

	fmt.Println(string(jsonStream))
}

func printHeader(w *tabwriter.Writer, attributes []string) {
	fmt.Fprintf(w, "TYPE\tID\tPROFILE\tREGION\tCREATED")

	for _, attribute := range attributes {
		fmt.Fprintf(w, "\t%s", strings.ToUpper(attribute))
	}

	fmt.Fprintf(w, "\t\n")
}
