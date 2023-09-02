package login

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Msk []byte = []byte("ztyyyds666")

const (
	machineBits  = int64(5)  //机器id位数
	serviceBits  = int64(5)  //服务id位数
	sequenceBits = int64(12) //序列id位数

	maxMachineID  = int64(-1) ^ (int64(-1) << machineBits)  //最大机器id
	maxServiceID  = int64(-1) ^ (int64(-1) << serviceBits)  //最大服务id
	maxSequenceID = int64(-1) ^ (int64(-1) << sequenceBits) //最大序列id

	timeLeft    = uint8(22) //时间id向左移位的量
	machineLeft = uint8(17) //机器id向左移位的量
	serviceLeft = uint8(12) //服务id向左移位的量

	twepoch = int64(1667972427000) //初始毫秒,时间是: Wed Nov  9 13:40:27 CST 2022
)

// -------------------------------------jwt生成token加密------------------------------------------------
type Claim struct {
	ID int64
	jwt.RegisteredClaims
} //创建用户登录标签

// 得到token,c为用户名
func GetToken(b []byte, c int64) (string, error) {
	a := Claim{
		c,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), //token有效时间
			Issuer:    "zty",                                                   //签发人
		},
	} //获取claim实例
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a) //获取token
	return token.SignedString(b)                          //返回加密串
}

// 解析token
func ParseToken(a string) (*jwt.Token, int64, error) {
	claim := &Claim{}
	t, err := jwt.ParseWithClaims(a, claim, func(t *jwt.Token) (interface{}, error) {
		return Msk, nil
	}) //接收前端发来加密字段
	return t, claim.ID, err
}

// ----------------------------------------使用sha256加密密码-----------------------------------------
func Encrypt(password string) string {
	const salt = "2023/08/31" //自定义加盐
	hash := sha256.New()
	hash.Write([]byte(password + salt)) //密码与盐自定义组合
	res := hex.EncodeToString(hash.Sum(nil))
	//fmt.Println(res)
	return res
}

// ----------------------------------------使用雪花算法生成用户id--------------------------------------
type Worker struct {
	sync.Mutex
	lastStamp  int64
	machineID  int64 //机器id,0~31
	serviceID  int64 //服务id,0~31
	sequenceID int64
}

var W Worker

func NewWorker(machineID, serviceID int64) *Worker {
	return &Worker{
		lastStamp:  0,
		machineID:  machineID,
		serviceID:  serviceID,
		sequenceID: 0,
	}
}

func (w *Worker) GetID() int64 {
	//多线程互斥
	w.Lock()
	defer w.Unlock()

	mill := time.Now().UnixMilli()

	if mill == w.lastStamp {
		w.sequenceID = (w.sequenceID + 1) & maxSequenceID
		//当一个毫秒内分配的id数>4096个时，只能等待到下一毫秒去分配。
		if w.sequenceID == 0 {
			for mill > w.lastStamp {
				mill = time.Now().UnixMilli()
			}
		}
	} else {
		w.sequenceID = 0
	}

	w.lastStamp = mill
	//fmt.Println(w.lastStamp - twepoch)
	//fmt.Println((w.lastStamp - twepoch) << timeLeft)
	//fmt.Printf("%b\n", (w.lastStamp-twepoch)<<timeLeft)
	id := (w.lastStamp-twepoch)<<timeLeft | w.machineID<<machineLeft | w.serviceID<<serviceLeft | w.sequenceID
	return id
}
