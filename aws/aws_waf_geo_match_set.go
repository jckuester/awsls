// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafGeoMatchSet(client *Client) error {
	req := client.wafconn.ListGeoMatchSetsRequest(&waf.ListGeoMatchSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.GeoMatchSets) > 0 {
		for _, r := range resp.GeoMatchSets {
			fmt.Println(*r.GeoMatchSetId)
		}
	}

	return nil
}
