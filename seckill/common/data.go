package common


type Activity struct {
	ProductId int    `json:"product_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64	 `json:"end_time"`
	Count     int    `json:"count"`
	Status    int	 `json:"status"`
}

type SecKillResult struct {
	UserId int64  `json:"user_id"`
	ProductId int `json:"product_id"`
	Token string  `json:"token"`
	Status int    `json:"status"`
	CurTime int64 `json:"cur_time"`
}

type SecKillRequest struct {
	UserId int64                   `json:"user_id"`
	ProductId int                 `json:"product_id"`
	UserIp string                  `json:"user_ip"`
	CurTime int64					`json:"cur_time"`
	ResultChan chan *SecKillResult  `json:"-"`
}
