/*
App Store Connect API — devices pagination helper.

Hand-written (not generated): convenience that follows the response cursor to
return every device across all pages.
*/

package openapi

import (
	"context"
	"net/http"
	"net/url"
)

// devicesPageLimit is Apple's maximum page size for the devices collection.
const devicesPageLimit = 200

// DevicesGetAll fetches every device across all pages, following links.next.
// It pages at the maximum size to minimise round-trips. Returns the aggregated
// list, the last HTTP response, and the first error encountered.
func (a *DevicesAPIService) DevicesGetAll(ctx context.Context) ([]Device, *http.Response, error) {
	var all []Device
	var lastResp *http.Response
	cursor := ""

	for {
		req := a.DevicesGetCollection(ctx).Limit(devicesPageLimit)
		if cursor != "" {
			req = req.Cursor(cursor)
		}

		resp, httpResp, err := req.Execute()
		lastResp = httpResp
		if err != nil {
			return all, lastResp, err
		}
		all = append(all, resp.GetData()...)

		links := resp.GetLinks()
		next := links.GetNext()
		if next == "" {
			return all, lastResp, nil
		}
		parsed, err := url.Parse(next)
		if err != nil {
			return all, lastResp, nil
		}
		cursor = parsed.Query().Get("cursor")
		if cursor == "" {
			return all, lastResp, nil
		}
	}
}
