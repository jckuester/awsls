// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafGeoMatchSet(client *Client) ([]Resource, error) {
	req := client.Wafconn.ListGeoMatchSetsRequest(&waf.ListGeoMatchSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.GeoMatchSets) > 0 {
		for _, r := range resp.GeoMatchSets {

			result = append(result, Resource{
				Type:   "aws_waf_geo_match_set",
				ID:     *r.GeoMatchSetId,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
