// Code generated by hertz generator.

package virtual_gallery

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	virtual_gallery "nft/hertz_demo/biz/model/nft/ecom/virtual_gallery"
)

// NftMarketSpace .
// @router /api/nft/market_space [GET]
func NftMarketSpace(ctx context.Context, c *app.RequestContext) {
	var err error
	var req virtual_gallery.NftMarketSpaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(virtual_gallery.NftMarketSpaceResp)

	c.JSON(200, resp)
}

// NftHomepage .
// @router /api/nft/homepage [GET]
func NftHomepage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req virtual_gallery.NftHomepageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(virtual_gallery.NftHomepageResp)

	c.JSON(200, resp)
}

// NftMyTab .
// @router /api/nft/my_tab [GET]
func NftMyTab(ctx context.Context, c *app.RequestContext) {
	var err error
	var req virtual_gallery.NftMyTabReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(virtual_gallery.NftMyTabResp)

	c.JSON(200, resp)
}
