package logic

import(
	"github.com/sherlockhua/goproject/short_url/model"
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"crypto/md5"
)

var (
	Db *sqlx.DB
)

type ShortUrl struct {
	Id int64 `db:"id"`
	ShortUlr string `db:"short_url"`
	OriginUrl string `db:"origin_url"`
	HashCode string `db:"hash_code"`
}

func InitDb(dns string) (err error) {
    Db, err = sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Println("connect to msyql failed, ", err)
		return
	}
	
	err = Db.Ping()
    return
}

func Long2Short(req *model.Long2ShortRequest) (response *model.Long2ShortResponse, err error) {
	response = &model.Long2ShortResponse{}
	
	urlMd5 := fmt.Sprintf("%x", md5.Sum([]byte(req.OriginUrl)))
	var short ShortUrl
	err = Db.Get(&short, "select id, short_url, origin_url, hash_code from short_url where hash_code=?", urlMd5)
	fmt.Printf("read db result, short:%#v, err:%v\n", short, err)
	if err == sql.ErrNoRows {
		//没有记录，重新生成一个新的短url
		err = nil
		shortUrl, errRet := generateShortUrl(req, urlMd5)
		if errRet != nil {
			err = errRet
			return
		}

		response.ShortUrl = shortUrl
		return
	}
	if err != nil {
		return
	}

	response.ShortUrl = short.ShortUlr
	return
}


func Short2Long(req *model.Short2LongRequest) (response *model.Short2LongResponse, err error) {
	response = &model.Short2LongResponse{}
	
	
	var short ShortUrl
	err = Db.Get(&short, "select id, short_url, origin_url, hash_code from short_url where short_url=?", req.ShortUrl)
	fmt.Printf("read db result, short:%#v, err:%v\n", short, err)
	if err == sql.ErrNoRows {
		//没有记录，重新生成一个新的短url
		response.Code = 404
		return
	}
	if err != nil {
		response.Code = 500
		return
	}

	response.OriginUrl = short.OriginUrl
	return
}

func generateShortUrl(req *model.Long2ShortRequest, hashcode string)(shortUrl string, err error) {
	result, err := Db.Exec("insert into short_url(origin_url, hash_code)values(?,?)", req.OriginUrl, hashcode)
	if err != nil {
		return
	}
	//0-9a-zA-Z
	insertId, _ := result.LastInsertId()
	shortUrl = transTo62(insertId)
	//更新数据库中的短URL
	_, err = Db.Exec("update short_url set short_url=? where id=?", shortUrl, insertId)
	return
}

func transTo62(id int64) string {
	//1-->1
	//10-->a
	//61-->Z
	//62--> 10
	//0-9a-zA-Z
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortUrl []byte
	for {
		var result byte
		number := id % 62
		//if number >= 0&& number <= 9 {
			result = charset[number]
			shortUrl = append(shortUrl, result)
		//}
		id = id / 62
		if id == 0 {
			break
		}
	}
	return string(shortUrl)
}

func GetLastShortUrl(limit int) (result []*model.ShortUrl, err error) {
	err = Db.Select(&result, "select short_url from short_url order by id desc limit ? ", limit)
	return
}