// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func ListBackupPlan(client *Client) {
	req := client.backupconn.ListBackupPlansRequest(&backup.ListBackupPlansInput{})

	p := backup.NewListBackupPlansPaginator(req)
	fmt.Println("")
	fmt.Println("aws_backup_plan:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.BackupPlansList {
			fmt.Println(*r.BackupPlanId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_backup_plan: %s", err)
	}

}
