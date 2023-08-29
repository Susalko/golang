package cashe

import "errors"

type Cashe struct {
	cashMap map[string]interface{}
}

func New() *Cashe {
	return &Cashe{
		cashMap: make(map[string]interface{}),
	}
}

func (c *Cashe) Set(key string, value interface{}) error {
	exist := c.cashMap[key]
	if exist != nil {
		return errors.New("Запись с ключем [" + key + "] существует")
	}
	c.cashMap[key] = value
	return nil
}

func (c *Cashe) Get(key string) (interface{}, error) {
	exist := c.cashMap[key]
	if exist == nil {
		return nil, errors.New("Запись с таким ключем [" + key + "] не существует")
	}
	return c.cashMap[key], nil
}

func (c *Cashe) Delete(key string) error {
	exist := c.cashMap[key]
	if exist == nil {
		return errors.New("Запись с таким ключем [" + key + "] не существует")
	}
	delete(c.cashMap, key)
	return nil
}
