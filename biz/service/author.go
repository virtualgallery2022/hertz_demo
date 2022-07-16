package service

import (
	"context"
	"github.com/KArtorias/logs"
	"nft/hertz_demo/biz/model/db_model"
	"nft/hertz_demo/consts"
)

type AuthorEntity struct {
	ctx      context.Context
	authorID int64
}

func MakeOne(ctx context.Context, authorID int64) *AuthorEntity {
	authorEntity := &AuthorEntity{
		ctx:      ctx,
		authorID: authorID,
	}
	return authorEntity
}

func (h *AuthorEntity) GetPersonalInfo() (personalInfo *db_model.DemoStruct, err error) {
	personalInfo, err = db_model.GetDemoRecord(h.ctx, consts.Db, "DEMO")
	if err != nil {
		logs.CtxError(h.ctx, "GetDemoRecord failed.")
		return
	}
	return
}
