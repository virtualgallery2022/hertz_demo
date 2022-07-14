package config

import (
	"context"
	"nft/hertz_demo/biz/model/db_model"
	"nft/hertz_demo/consts"
	"testing"
)

func TestDBConnection(t *testing.T) {
	ctx := context.Background()
	InitDB()
	demo, err := db_model.GetDemoRecord(ctx, consts.Db, "DEMO")
	if err != nil {
		t.Errorf("err=%+v", err)
		return
	}
	t.Logf("demo=%+v", demo)
}
