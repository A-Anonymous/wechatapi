package controllers
//
//import (
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/go-kit/kit/log"
//	"os"
//	"github.com/boltdb/bolt"
//)
//
//type SearchController struct {
//	beego.Controller
//}
//
//func (rc *SearchController) Get() {
//	ip := rc.GetString("ip")
//	fmt.Println(ip)
//	jsondata, err := Get(ip)
//	fmt.Println(jsondata)
//	if err != nil {
//		fmt.Println(err)
//	}
//	rc.Ctx.WriteString(jsondata)
//}
//func Get(key string)(string, error){
//
//	logger := log.NewLogfmtLogger(os.Stdout)
//	db, err := bolt.Open("mydb.db", 0600, nil)
//	if err != nil {
//		logger.Log("open",err)
//	}
//	defer db.Close()
//	var val string
//	err = db.Update(func(tx *bolt.Tx)  error{
//		b := tx.Bucket([]byte("test"))
//		val= string(b.Get([]byte(key)))
//		return err
//
//	})
//	return val, err
//}
