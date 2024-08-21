package model

type Dict struct {
	BaseModel
	Type   int    `json:"type"`   //字典类型，1 样本类型 2 实验方法 3 结果单位 4 标本性状 5 禁止打印原因
	Name   string `json:"name"`   //字典名称
	Alias  string `json:"alias"`  //别名
	Remark string `json:"remark"` //备注
	Sort   int    `json:"sort"`   //排序
}

type Dicts []Dict

func (d *Dict) TableName() string {
	return "base_dict"
}

func (d Dicts) ToNames() []string {
	names := make([]string, len(d))
	for i, item := range d {
		names[i] = item.Name
	}
	return names
}
