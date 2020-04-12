// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func ListBackupVault(client *Client) {
	req := client.backupconn.ListBackupVaultsRequest(&backup.ListBackupVaultsInput{})

	p := backup.NewListBackupVaultsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_backup_vault:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.BackupVaultList {
			fmt.Println(*r.BackupVaultName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_backup_vault: %s", err)
	}

}
