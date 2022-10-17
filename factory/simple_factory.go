package factory

//简单工厂
//创建图书接口

type Book interface {
	Write() string
}

//创建中文图书

type chineseBook struct {
	name string
}

func (c *chineseBook) Write() string {
	return "write chinese book"
}

//创建英文图书

type englishBook struct {
	name string
}

func (c *englishBook) Write() string {
	return "write english book"
}

//判断返回

func NewBook(t string) Book {
	if t == "chinese" {
		return &chineseBook{}
	} else if t == "english" {
		return &englishBook{}
	}
	return nil
}
