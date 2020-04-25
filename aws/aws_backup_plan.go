// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func ListBackupPlan(client *Client) ([]Resource, error) {
	req := client.backupconn.ListBackupPlansRequest(&backup.ListBackupPlansInput{})

	var result []Resource

	p := backup.NewListBackupPlansPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.BackupPlansList {

			t := *r.CreationDate
			result = append(result, Resource{
				Type:   "aws_backup_plan",
				ID:     *r.BackupPlanId,
				Region: client.backupconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
