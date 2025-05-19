package model

type KeyWord string

const (
	KeyWord_None   KeyWord = ""
	KeyWord_Var    KeyWord = "var"    // 变量名称
	KeyWord_Type   KeyWord = "type"   // 类型
	KeyWord_Group  KeyWord = "group"  // 分组
	KeyWord_Column KeyWord = "column" // 列
	KeyWord_Sep    KeyWord = "sep"    // 分隔符
	KeyWord_Desc   KeyWord = "desc"   // 描述
	KeyWord_Def    KeyWord = "def"    // 默认值
	KeyWord_Cmt    KeyWord = "cmt"    // 注释
)

// 基本类型定义
type BaseType string

const (
	BaseType_Bool   BaseType = "bool"   // 布尔类型 true、false、0、1都能被识别，大小写不敏感，如True、TRUE也是有效值
	BaseType_Byte   BaseType = "byte"   // byte（uint8_t）
	BaseType_Short  BaseType = "short"  //short（int16_t）
	BaseType_Int    BaseType = "int"    // int（int32_t）
	BaseType_Long   BaseType = "long"   // long（int64_t）
	BaseType_Float  BaseType = "float"  // float（float32）
	BaseType_Double BaseType = "double" // double（float64）
	BaseType_String BaseType = "string" // string（string）
	// BaseType_Text  BaseType = "text"  // text是一个语法糖类型，而不是独立的类型。等价于string#text=1，即包含tag text=1 的string类型
	BaseType_Datetime BaseType = "datetime" //long，值为自UTC 1970-01-01 00:00:00以来的秒数
	// 自定义类型
	BaseType_Enum BaseType = "enum" // 枚举类型
	BaseType_Bean BaseType = "bean" // 复合类型，对应 class或struct。bean支持类型继承和多态
	// 容器类型
	BaseType_List BaseType = "list" // list
	BaseType_Map  BaseType = "map"  // map
	BaseType_Set  BaseType = "set"  // set
)

// enum
// isFlags	bool	是	是否为标志位类型，对应c#的FlagsAttribute语义
//isUniqueItemId	bool	否	枚举值是否唯一
type TypeDefine struct {
	Kind       BaseType // 种类
	ObjectType string   // 对象类型 类型名称

	FieldName string // 字段名
	FieldType string // 字段类型

	Name  string // 标识名
	Value string // 值

	Comment string            // 注释
	Tags    map[string]string // 自定义tag对
	Groups  []string          // 导出分组

	ArraySplitter string // 数组切割
	MakeIndex     bool   // 索引

	DefaultValue string // 默认值"`

}

// InitBuiltinTypes 内建表的列功能
func InitBuiltinTypes() {

	//GlobalData.IndexDefine = []*TypeDefine{
	//	{Kind: TypeUsage_HeaderStruct, ObjectType: "IndexDefine", Name: "表类型", FieldName: "tableType", FieldType: "string"},
	//	{Kind: TypeUsage_HeaderStruct, ObjectType: "IndexDefine", Name: "表文件名", FieldName: "tableFileName", FieldType: "string"},
	//	{Kind: TypeUsage_HeaderStruct, ObjectType: "IndexDefine", Name: "模式", FieldName: "group", FieldType: "string"},
	//	{Kind: TypeUsage_HeaderStruct, ObjectType: "IndexDefine", Name: "注释", FieldName: "comment", FieldType: "string"},
	//}

	//	for _, tf := range []*TypeDefine{
	//
	//		// 类型表类型
	//		{Kind: TypeUsage_Enum, ObjectType: "TypeUsage", Name: "", FieldName: "None", FieldType: "int", Value: "0"},
	//
	//		// 索引表类型
	//		{Kind: TypeUsage_HeaderStruct, ObjectType: "TypeUsage", Name: "", FieldName: "None", FieldType: "int", Value: "0"},
	//		{Kind: TypeUsage_HeaderStruct, ObjectType: "TypeUsage", Name: "", FieldName: "None", FieldType: "int", Value: "0"},
	//		{Kind: TypeUsage_HeaderStruct, ObjectType: "TypeUsage", Name: "", FieldName: "None", FieldType: "int", Value: "0"},
	//		{Kind: TypeUsage_HeaderStruct, ObjectType: "TypeUsage", Name: "", FieldName: "None", FieldType: "int", Value: "0"},
	//
	//		// KV表类型
	//	} {
	//		tf.IsBuiltin = true
	//
	//		typeTab.AddField(tf, nil, 0)
	//	}
	//}
}
