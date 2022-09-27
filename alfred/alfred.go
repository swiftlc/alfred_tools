package alfred

import (
	"io/ioutil"
	"os"
	"path"
	"time"
)

var CstTimezone = time.FixedZone("CST", 8*3600)

const (
	DefaultTimeFmt = `2006-01-02 15:04:05`
)

func FormateTime(sec int64) string {
	return time.Unix(sec, 0).In(CstTimezone).Format(DefaultTimeFmt)
}

func WalkDir(filePath string, level int, handler func(p string)) {
	if level == 0 {
		return
	}

	infos, err := ioutil.ReadDir(filePath)
	if err != nil {
		return
	}

	for _, v := range infos {
		if v.IsDir() {
			filePath := path.Join(filePath, v.Name())
			exists, _ := PathExists(path.Join(filePath, ".git"))
			if exists {
				handler(filePath)
			} else {
				WalkDir(filePath, level-1, handler)
			}
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
