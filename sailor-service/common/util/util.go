package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
)

func Copy(source, dest interface{}) error {
	return copier.Copy(dest, source)
}

func ParseTimeToUnixMilli(dbTime time.Time) (int64, error) {

	timeTemplate := "2006-01-02 15:04:05"
	timeStr := dbTime.String()
	cstLocal, _ := time.LoadLocation("Asia/Shanghai")
	x, err := time.ParseInLocation(timeTemplate, timeStr, cstLocal)
	if err != nil {
		return -1, err
	}
	return x.UnixMilli(), nil
}

// UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

// GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

func PathExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func GetCallerPosition(skip int) string {
	if skip <= 0 {
		skip = 1
	}
	_, filename, line, _ := runtime.Caller(skip)
	projectPath := "af-sailor-service"
	ps := strings.Split(filename, projectPath)
	pl := len(ps)
	return fmt.Sprintf("%s %d", ps[pl-1], line)
}

func RandomInt(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(max)
}

func SliceUnique(s []string) []string {
	m := make(map[string]uint8)
	result := make([]string, 0)
	for _, i := range s {
		_, ok := m[i]
		if !ok {
			m[i] = 1
			result = append(result, i)
		}
	}
	return result
}

func Transfer[T any](d any) (*T, error) {
	result := new(T)
	bts, err := json.Marshal(d)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(bts, result); err != nil {
		return result, err
	}
	return result, nil
}

// IsLimitExceeded total / limit 向上取整是否大于等于 offset，小于则超出总数
func IsLimitExceeded(limit, offset, total float64) bool {
	return math.Ceil(total/limit) < offset
}

func PtrToValue[T any](ptr *T) (res T) {
	if ptr == nil {
		return
	}

	return *ptr
}

func ValueToPtr[T any](v T) *T {
	return &v
}

func CheckKeyword(keyword *string) bool {
	*keyword = strings.TrimSpace(*keyword)
	if len([]rune(*keyword)) > 128 {
		return false
	}
	return regexp.MustCompile("^[a-zA-Z0-9\u4e00-\u9fa5-_]*$").Match([]byte(*keyword))
}

func GenFlowchartVersionName(vid int32) string {
	return fmt.Sprintf("v%d", vid)
}

func NewUUID() string {
	return uuid.NewString()

	// u := uuid.New()
	// buf := make([]byte, 32)
	//
	// hex.Encode(buf[0:8], u[0:4])
	// hex.Encode(buf[8:12], u[4:6])
	// hex.Encode(buf[12:16], u[6:8])
	// hex.Encode(buf[16:20], u[8:10])
	// hex.Encode(buf[20:], u[10:])
	// return string(buf)
}

func CopyMap[K comparable, V any](src map[K]V) map[K]V {
	if src == nil {
		return nil
	}

	dest := make(map[K]V, len(src))
	for k, v := range src {
		dest[k] = v
	}

	return dest
}

func ToInt32s[T ~int32](in []T) []int32 {
	if in == nil {
		return nil
	}

	ret := make([]int32, len(in))
	for i := range in {
		ret[i] = int32(in[i])
	}

	return ret
}

func Decode(dest any, bs []byte) error {
	decoder := json.NewDecoder(bytes.NewBuffer(bs))
	decoder.UseNumber() // 指定使用 Number 类型
	if err := decoder.Decode(&dest); err != nil {
		return fmt.Errorf("decoder error %v", err)
	}
	return nil
}

func CopyUseJson(dest any, src any) error {
	strSrc, ok := src.(string)
	if ok {
		return Decode(dest, []byte(strSrc))
	}
	pSrc, pok := src.(*any)
	if pok {
		pStrSrc, sok := (*pSrc).(string)
		if sok {
			return Decode(dest, []byte(pStrSrc))
		}
	}
	return Decode(dest, lo.T2(json.Marshal(src)).A)
}

func GetJsonInAnswer[T any](answer string) (*T, error) {
	first := strings.Index(answer, "{")
	last := strings.LastIndex(answer, "}")
	if first < 0 || last < 0 {
		return nil, fmt.Errorf("invalid  answer  %v", answer)
	}
	answer = answer[first : last+1]

	result := new(T)
	if err := json.Unmarshal([]byte(answer), result); err != nil {
		return nil, err
	}
	return result, nil
}

// RoundWithPrecision returns the rounded value of x to specified precision, use ROUND_HALF_UP mode.
// For example:
//
//		  Round1(0.6635, 3)   // 0.664
//	   Round1(0.363636, 2) // 0.36
//	   Round1(0.363636, 1) // 0.4
func RoundWithPrecision(x float64, precision int) float64 {
	if precision == 0 {
		return math.Round(x)
	}

	p := math.Pow10(precision)
	if precision < 0 {
		return math.Round(x*p) * math.Pow10(-precision)
	}
	return math.Round(x*p) / p
}

func SliceStr2Int(ss []string) []int {
	ids := make([]int, 0)
	for _, idStr := range ss {
		id, _ := strconv.Atoi(idStr)
		if id <= 0 {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

func SetEnvs(strs string) {
	ps := strings.Split(strs, ";")
	for _, p := range ps {
		kv := strings.Split(p, "=")
		if len(kv) != 2 {
			continue
		}
		if err := os.Setenv(kv[0], kv[1]); err != nil {
			fmt.Printf("set env error %v", err)
		}
	}
}

type BinaryInteger interface {
	int | int32
}

func IntToArray[T BinaryInteger](dataKind T) []T {
	var val T
	array := make([]T, 0, 6)
	for i := 0; i < 6 && dataKind >= 1<<i; i++ {
		val = dataKind & (1 << i)
		if val > 0 {
			array = append(array, val)
		}
	}
	return array
}

func ArrayToInt[T BinaryInteger](ds []T) T {
	var val T
	for _, d := range ds {
		val += 1 << d
	}
	return val
}

func MD5(data []byte) string {
	md5New := md5.New()
	md5New.Write(data)
	return hex.EncodeToString(md5New.Sum(nil))
}

// PreSubString 返回字符串的前n个字符
func PreSubString(s string, n int) string {
	cmts := []rune(s)
	if len(cmts) <= n {
		return s
	}
	return string(cmts[:128])
}
