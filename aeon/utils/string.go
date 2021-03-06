package utils

import (
	"errors"
	"strconv"
	"strings"
)

// PagingParam 将string类型的分页参数转换成int类型
func PagingParam(pageNoStr, pageSizeStr string) (pageNo, pageSize int, err error) {
	if pageNoStr == "" || pageSizeStr == "" {
		err = errors.New("Missing pageNo or pageSize. ")
		return
	}
	pageNo, err = strconv.Atoi(pageNoStr)
	if err != nil {
		return
	}
	pageSize, err = strconv.Atoi(pageSizeStr)
	return
}

func StringSliceBuilder(strSlice []string) string {
	sb := strings.Builder{}
	for _, s := range strSlice {
		sb.WriteString(s)
	}
	return sb.String()
}
