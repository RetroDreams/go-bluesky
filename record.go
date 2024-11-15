package bluesky

import (
	"context"

	"github.com/bluesky-social/indigo/repo"
)

func (c *Client) GetRecord(did, rpath string) string {
	ctx := context.Background()
	repo := repo.NewRepo(ctx, did, nil)
	record, _, err := repo.GetRecord(ctx, rpath)
	if err != nil {
		panic(err)
	}
	return record.String()
}
