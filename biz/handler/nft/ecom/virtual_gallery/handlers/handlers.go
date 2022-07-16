package handlers

import (
	"context"
	"github.com/KArtorias/logs"
	"nft/hertz_demo/biz/model/nft/ecom/virtual_gallery"
	"nft/hertz_demo/biz/service"
)

func MarketSpace(ctx context.Context, req virtual_gallery.NftMarketSpaceReq) (resp *virtual_gallery.NftMarketSpaceResp, err error) {
	normalEditions := make([]*virtual_gallery.NftAttributes, 0)
	limitedEditions := make([]*virtual_gallery.NftAttributes, 0)
	eachNormalEdition := &virtual_gallery.NftAttributes{
		ID:       1,
		ImageURL: "https://unsplash.com/photos/MLvCvS5eX-Y",
		ShaHash:  "sha256_test",
		Group:    0,
		Theme:    virtual_gallery.NftTheme_Unknown,
		Edition:  virtual_gallery.NftEdtion_Normal,
		Price:    "$10",
	}
	eachLimitedEdition := &virtual_gallery.NftAttributes{
		ID:       1,
		ImageURL: "https://unsplash.com/photos/MLvCvS5eX-Y",
		ShaHash:  "sha256_test",
		Group:    0,
		Theme:    virtual_gallery.NftTheme_Unknown,
		Edition:  virtual_gallery.NftEdtion_Limited,
		Price:    "$100",
	}
	normalEditions = append(normalEditions, eachNormalEdition)
	limitedEditions = append(limitedEditions, eachLimitedEdition)
	resp = &virtual_gallery.NftMarketSpaceResp{
		NormalEdition:  normalEditions,
		LimitedEdition: limitedEditions,
	}
	return
}

func HomePage(ctx context.Context, req virtual_gallery.NftHomepageReq) (resp *virtual_gallery.NftHomepageResp, err error) {
	attributions := make([]*virtual_gallery.NftAttributes, 0)
	banners := make([]string, 0)
	banners = append(banners, "https://unsplash.com/photos/zFIEshS7La8")
	eachAttri := &virtual_gallery.NftAttributes{
		ID:       1,
		ImageURL: "https://unsplash.com/photos/MLvCvS5eX-Y",
		ShaHash:  "sha256_test",
		Group:    0,
		Theme:    virtual_gallery.NftTheme_Unknown,
		Edition:  virtual_gallery.NftEdtion_Normal,
		Price:    "$10",
	}
	attributions = append(attributions, eachAttri)
	resp = &virtual_gallery.NftHomepageResp{
		HomepageEdition: attributions,
		Banners:         banners,
		Activities:      nil,
	}
	return
}

func MyTab(ctx context.Context, req virtual_gallery.NftMyTabReq) (resp *virtual_gallery.NftMyTabResp, err error) {
	authorEntity := service.MakeOne(ctx, 0)
	personalInfo, err := authorEntity.GetPersonalInfo()
	if err != nil {
		logs.CtxError(ctx, "authorEntity.GetPersonalInfo() failed. err=%+v", err)
		return
	}
	pResp := &virtual_gallery.PersonalInfo{
		PhoneNumber:  "123456",
		IdentityHash: "",
		Name:         personalInfo.NAME,
		NickName:     "test demo",
	}
	resp = &virtual_gallery.NftMyTabResp{
		PersonalInfo:  pResp,
		SavingAccount: 0,
	}
	return resp, nil
}
