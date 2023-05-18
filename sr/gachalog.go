package sr

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mhygacha/config"
	"mhygacha/global"
	"net/http"
	"strconv"
)

func GaChaLog(c *gin.Context) {
	Authkey := c.Query("authkey")      //登陆密钥
	Page := c.Query("page")            //第几页
	Size := c.Query("size")            //每页几个
	GaChaType := c.Query("gacha_type") //卡池ID
	EndID := c.Query("end_id")         //截止ID
	BeginID := c.Query("begin_id")     //起始ID

	var (
		rows     *sql.Rows
		err      error
		EndIDint int
	)

	if BeginID != "" {
		BeginIDint, _ := strconv.Atoi(BeginID)
		Sizeint, _ := strconv.Atoi(Size)
		EndIDint = BeginIDint + Sizeint + 1
	} else {
		EndIDint, _ = strconv.Atoi(EndID)
	}
	var authkeyCondition, idCondition string
	var args []interface{}

	if config.Config.OpenAuthKey {
		authkeyCondition = "Authkey = ? AND "
		args = append(args, Authkey)
	}
	if EndIDint > 0 {
		idCondition = "Id < ? AND "
		args = append(args, EndIDint)
	}
	args = append(args, GaChaType, Size)
	rows, err = global.SQLDB.Query(fmt.Sprintf(`SELECT Uid, GachaId, GachaType, ItemId, Count, Time, Name, Lang, ItemType, RankType, Id 
         FROM srgachalog 
         WHERE %s GaChaType = ? 
         ORDER BY Id DESC 
         LIMIT ?`, authkeyCondition+idCondition), args...)

	if err != nil {
		log.Printf("查询数据出错: %v", err)
		return
	}

	defer rows.Close()

	var ListData []GaChaLogDataList
	for rows.Next() {
		var data GaChaLogDataList
		if err := rows.Scan(&data.Uid, &data.GachaId, &data.GachaType, &data.ItemId, &data.Count, &data.Time, &data.Name, &data.Lang, &data.ItemType, &data.RankType, &data.Id); err != nil {
			log.Printf("插入数据时: %v", err)
			return
		}
		ListData = append(ListData, data)
	}
	if ListData == nil {
		ListData = []GaChaLogDataList{}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error reading rows: %v", err)
		return
	}

	//输出数据
	c.JSON(http.StatusOK, GaChaLogProto{
		Retcode: 0,
		Message: "OK",
		Data: GaChaLogData{
			Page:           Page,
			Size:           Size,
			List:           ListData,
			Region:         "prod_gf_cn",
			RegionTimeZone: 8,
		},
	})
}

// GachaLogProto /*
type GaChaLogProto struct {
	Retcode int          `json:"retcode"`
	Message string       `json:"message"`
	Data    GaChaLogData `json:"data"`
}
type GaChaLogData struct {
	Page           string             `json:"page"`
	Size           string             `json:"size"`
	List           []GaChaLogDataList `json:"list"`
	Region         string             `json:"region"`
	RegionTimeZone int                `json:"region_time_zone"`
}

type GaChaLogDataList struct {
	Uid       string `json:"uid"`
	GachaId   string `json:"gacha_id"`
	GachaType string `json:"gacha_type"`
	ItemId    string `json:"item_id"`
	Count     string `json:"count"`
	Time      string `json:"time"`
	Name      string `json:"name"`
	Lang      string `json:"lang"`
	ItemType  string `json:"item_type"`
	RankType  string `json:"rank_type"`
	Id        string `json:"id"`
}
