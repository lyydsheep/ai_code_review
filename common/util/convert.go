package util

import (
	"errors"
	"github.com/faiz/llm-code-review/common/enum"
	"github.com/jinzhu/copier"
	"regexp"
	"time"
)

func Convert(to, from any) error {
	return copier.CopyWithOption(to, from, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type is not time.TIme")
					}
					return val.Format(enum.TimeFormatHyphenedYMDHIS), nil
				},
			},
			{
				SrcType: copier.String,
				DstType: time.Time{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(string)
					if !ok {
						return nil, errors.New("src type is not string")
					}
					if !ok {
						return nil, errors.New("src type is not time format string")
					}
					pattern := `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$` // YYYY-MM-DD HH:MM:SS
					matched, _ := regexp.MatchString(pattern, val)
					if matched {
						return time.Parse(enum.TimeFormatHyphenedYMDHIS, val)
					}
					return nil, errors.New("src type is not time format string")
				},
			},
		},
	})
}
