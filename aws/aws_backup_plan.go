// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListBackupPlan(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Backupconn.ListBackupPlansRequest(&backup.ListBackupPlansInput{})

	var result []terraform.Resource

	p := backup.NewListBackupPlansPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.BackupPlansList {

			t := *r.CreationDate
			result = append(result, terraform.Resource{
				Type:      "aws_backup_plan",
				ID:        *r.BackupPlanId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
