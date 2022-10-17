package factory

import "fmt"

// 创建接口
type fBook interface {
	Write() string
}

// 创建属性
type fBookAttr struct {
	name string
}

// 创建方法函数
func (f *fBookAttr) Write() string {
	return fmt.Sprintf("write %s", f.name)
}

// 创建中文图书属性
type fChineseBook struct {
	fBookAttr
}

// 调用中文图书
func newChineseBook() fBook {
	return &fChineseBook{fBookAttr{name: "chinese"}}
}

// 创建英文图书属性
type fEnglishBook struct {
	fBookAttr
}

// 调用英文图书
func newEnglishBook() fBook {
	return &fEnglishBook{fBookAttr{name: "english"}}
}

//调度

func GetFBook(fType string) (fBook, error) {
	if fType == "chinese" {
		return newChineseBook(), nil
	}
	if fType == "english" {
		return newEnglishBook(), nil
	}
	return nil, nil
}
