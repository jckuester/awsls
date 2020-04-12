// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func ListBackupVault(client *Client) error {
	req := client.backupconn.ListBackupVaultsRequest(&backup.ListBackupVaultsInput{})

	p := backup.NewListBackupVaultsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.BackupVaultList {
			fmt.Println(*r.BackupVaultName)

			fmt.Printf("CreatedAt: %s\n", *r.CreationDate)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
