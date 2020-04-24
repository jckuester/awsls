// +build codegen

package aws

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jckuester/awsls/gen/util"
)

// GenerateClient writes Go code to initialize all AWS API Clients
// used by Terraform AWS resources.
func GenerateClient(outputPath string, services []string) error {
	err := os.MkdirAll(outputPath, 0775)
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", err)
	}

	err = util.WriteGoFile(
		filepath.Join(outputPath, "client.go"),
		util.CodeLayout,
		"",
		"aws",
		clientGoCode(services),
	)

	if err != nil {
		return fmt.Errorf("failed to write Go code to file: %s", err)
	}

	return nil
}

func clientGoCode(services []string) string {
	var buf bytes.Buffer
	err := clientTmpl.Execute(&buf, services)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var clientTmpl = template.Must(template.New("client").Parse(`import (
"context"
"github.com/aws/aws-sdk-go-v2/aws/external"
"github.com/aws/aws-sdk-go-v2/service/sts"
{{range .}}"github.com/aws/aws-sdk-go-v2/service/{{.}}"
{{end}})
type Client struct {
accountid string
{{range .}}{{.}}conn *{{.}}.Client
{{end}}}

func  NewClient() (*Client, error) {
cfg, err := external.LoadDefaultAWSConfig()
if err != nil {
	panic("failed to load config, " + err.Error())
}



client := &Client{
{{range .}}{{.}}conn: {{.}}.New(cfg),
{{end}}}

stsconn := sts.New(cfg)
req := stsconn.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
resp, err := req.Send(context.Background())
if err != nil {
	return nil, err
}
client.accountid = *resp.Account

return client, nil
}
`))
