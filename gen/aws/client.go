// +build codegen

package aws

import (
	"bytes"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jckuester/awsls/gen/util"
)

// GenerateClient returns Go code that initializes all AWS API clients.
func GenerateClient(outputPath string, services []string) {
	err := util.WriteGoFile(
		filepath.Join(outputPath, "client.go"),
		util.CodeLayout,
		"",
		"aws",
		clientGoCode(services),
	)

	if err != nil {
		panic(err)
	}
}

func clientGoCode(services []string) string {
	var buf bytes.Buffer
	err := clientTmpl.Execute(&buf, services)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var clientTmpl = template.Must(template.New("client").Funcs(
	template.FuncMap{
		"Title": strings.Title,
	}).Parse(`import (
"context"
"fmt"
"github.com/apex/log"
"github.com/aws/aws-sdk-go-v2/aws/external"
"github.com/aws/aws-sdk-go-v2/service/sts"
{{range .}}"github.com/aws/aws-sdk-go-v2/service/{{.}}"
{{end}})

type Client struct {
	AccountID string
	Region string
	Profile string
	{{ range . }}{{ . | Title }}conn *{{.}}.Client
{{end}}}

func NewClient(configs ...external.Config) (*Client, error) {
	cfg, err := external.LoadDefaultAWSConfig(configs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
	{{ range . }}{{ . | Title }}conn: {{.}}.New(cfg),
	{{end}}}

	profile, _, err := external.GetSharedConfigProfile(configs)
	if err != nil {
		return nil, err
	}

	client.Profile = profile
	client.Region = cfg.Region

	log.WithFields(log.Fields{
		"profile": profile,
		"region":  cfg.Region,
	}).Debugf("created new instance of AWS client")

	return client, nil
}

// SetAccountID populates the AccountID field of the client.
func (client *Client) SetAccountID() error {
	req := client.Stsconn.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	resp, err := req.Send(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get caller identity: %s", err)
	}

	client.AccountID = *resp.Account

	return nil
}
`))
