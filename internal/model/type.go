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

type TypeUsage int32

const (
	TypeUsage_None         TypeUsage = iota //
	TypeUsage_HeaderStruct                  // 表头
	TypeUsage_Enum                          // 枚举
)

type BaseType string

const (
	BaseType_Int    BaseType = "int"
	BaseType_Float  BaseType = "float"
	BaseType_String BaseType = "string"
	BaseType_Bool   BaseType = "bool"
	BaseType_List   BaseType = "list"
	BaseType_Map    BaseType = "map"
	BaseType_Class  BaseType = "class"
	BaseType_Enum   BaseType = "enum"
)

type TypeDefine struct {
	Kind          TypeUsage // 种类
	ObjectType    string    // 对象类型
	Name          string    // 标识名
	FieldName     string    // 字段名
	FieldType     string    // 字段类型
	Value         string    // 值
	ArraySplitter string    // 数组切割
	MakeIndex     bool      // 索引
	Tags          []string  // 标记
	IsBuiltin     bool      // ",omitempty"`
	UeType        string    // UE类型"`
	DefaultValue  string    // 默认值"`
}
