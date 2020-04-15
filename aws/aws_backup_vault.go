// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func ListBackupVault(client *Client) ([]Resource, error) {
	req := client.backupconn.ListBackupVaultsRequest(&backup.ListBackupVaultsInput{})

	var result []Resource

	p := backup.NewListBackupVaultsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.BackupVaultList {

			t := *r.CreationDate
			result = append(result, Resource{
				Type: "aws_backup_vault",
				ID:   *r.BackupVaultName,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
