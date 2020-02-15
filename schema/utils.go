package schema

import (
	"reflect"
	"regexp"
	"strings"
)

func ParseTagSetting(tags reflect.StructTag) map[string]string {
	setting := map[string]string{}

	for _, value := range strings.Split(tags.Get("gorm"), ";") {
		if value != "" {
			v := strings.Split(value, ":")
			k := strings.TrimSpace(strings.ToUpper(v[0]))

			if len(v) >= 2 {
				setting[k] = strings.Join(v[1:], ":")
			} else {
				setting[k] = k
			}
		}
	}
	return setting
}

func checkTruth(val string) bool {
	if strings.ToLower(val) == "false" {
		return false
	}
	return true
}

func toColumns(val string) (results []string) {
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			results = append(results, strings.TrimSpace(v))
		}
	}
	return
}

func removeSettingFromTag(tag reflect.StructTag, name string) reflect.StructTag {
	return reflect.StructTag(regexp.MustCompile(`(?i)(gorm:.*?)(`+name+`:.*?)(;|("))`).ReplaceAllString(string(tag), "${1}${4}"))
}