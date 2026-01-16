package knowledge_build

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

type KnCache struct {
	AnyDataAddress string `json:"anydata_address"`
	NetworkID      int    `json:"network_id"`
	//需要删除的资源的缓存
	Cache map[KNResourceType][]string `json:"cache"`
}

func NewKnCache() *KnCache {
	return &KnCache{
		NetworkID: 0,
		Cache:     make(map[KNResourceType][]string),
	}
}

func (k *KnCache) Get(t KNResourceType) []string {
	ds, ok := k.Cache[t]
	if !ok {
		return []string{}
	}
	return ds
}

func (k *KnCache) Set(t KNResourceType, ids []string) {
	k.Cache[t] = ids
}

func (k *KnCache) IsEmpty() bool {
	for _, ids := range k.Cache {
		if len(ids) > 0 {
			return false
		}
	}
	return true
}

func genResName(namePrefix string) string {
	return namePrefix + "_" + lo.RandomString(5, lo.AlphanumericCharset)
}

func parseKnwId(detail string) (int, error) {
	dm := make(map[string]any)
	if err := json.Unmarshal([]byte(detail), &dm); err != nil {
		return 0, fmt.Errorf("parse knwId error %v", err)
	}
	//{"ID":"27","Name":"血缘图谱_i6tav","KnwId":"6","DSIdMap":"{\"9\":6}"}
	d, ok := dm["KnwId"]
	if !ok {
		return 0, errors.New("parse knwId error: empty knwId in config detail")
	}
	id, err := strconv.Atoi(fmt.Sprintf("%v", d))
	if err != nil {
		return 0, fmt.Errorf("parse knwId error: invalid knwId:%v in config detail", d)
	}
	return id, nil
}
