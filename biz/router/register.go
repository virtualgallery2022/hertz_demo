// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	hello_example "nft/hertz_demo/biz/router/hello/example"
	nft_ecom_virtual_gallery "nft/hertz_demo/biz/router/nft/ecom/virtual_gallery"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	nft_ecom_virtual_gallery.Register(r)
}
