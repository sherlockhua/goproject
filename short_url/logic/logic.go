package logic

import(
	"github.com/sherlockhua/goproject/short_url/model"
)

func Long2Short(req *model.Long2ShortRequest) (response *model.Long2ShortResponse, err error) {
	response = &model.Long2ShortResponse{}
	response.ShortUrl = req.OriginUrl
	return
}