package sr

import (
	"github.com/gin-gonic/gin"
)

type AddProto struct {
	Authkey   string `json:"authkey"`
	Lang      string `json:"lang"`
	GachaID   string `json:"gacha_id"`
	GachaType string `json:"gacha_type"`
	UID       string `json:"uid"`
	ItemID    string `json:"item_id"`
	Time      string `json:"time"`
	Name      string `json:"name"`
	ItemType  string `json:"item_type"`
	RankType  string `json:"rank_type"`
	ID        int    `json:"id"`
}

func GaChaLogAdd(c *gin.Context) {
	c.String(200, "OK")
}
