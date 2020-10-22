package history

import (
	"net/url"
	"sort"
	"strings"
)

// Uniq 对history去重
func Uniq(in []Entry) []Entry {
	if len(in) == 0 {
		return []Entry{}
	}

	entryByKey := make(map[string]Entry, len(in))
	for i := len(in) - 1; i >= 0; i-- { // 反向遍历，后面的权重低，所以会被前面的覆盖
		key := getKeyFromEntry(in[i])
		entryByKey[key] = in[i]
	}

	// map => slice ordered by xxx DESC
	out := make([]Entry, 0, len(entryByKey))
	for _, entry := range entryByKey {
		out = append(out, entry)
	}
	sort.Sort(ByDesc(out))

	// fmt.Printf("in %+v => out %+v\n", in, out)
	// fmt.Printf("in %d => out %d\n", len(in), len(out))
	return out
}

func getKeyFromEntry(entry Entry) string {
	origURL := entry.URL
	u, err := url.Parse(origURL)
	if err != nil {
		return origURL
	}

	key := u.Host + u.Path + "?" + u.RawQuery // 不要#hash
	if u.Host == "bytedance.feishu.cn" {
		if strings.HasPrefix(u.Path, "/docs") || strings.HasPrefix(u.Path, "/sheets") {
			key = u.Host + u.Path // 不要?qeury和#hash
		}
	}
	// fmt.Println(origURL, key)

	return key
}
