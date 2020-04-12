// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func ListBackupPlan(client *Client) error {
	req := client.backupconn.ListBackupPlansRequest(&backup.ListBackupPlansInput{})

	p := backup.NewListBackupPlansPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.BackupPlansList {
			fmt.Println(*r.BackupPlanId)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
