package data_subject

type SubjectDomainObjectTypeString string

const (
	StringRoot               SubjectDomainObjectTypeString = "root"
	StringSubjectDomainGroup SubjectDomainObjectTypeString = "subject_domain_group" // 主题域分组
	StringSubjectDomain      SubjectDomainObjectTypeString = "subject_domain"       // 主题域
	StringBusinessObject     SubjectDomainObjectTypeString = "business_object"      // 业务对象
	StringBusinessActivity   SubjectDomainObjectTypeString = "business_activity"    // 业务活动
	StringLogicEntity        SubjectDomainObjectTypeString = "logic_entity"         // 逻辑实体
	StringAttribute          SubjectDomainObjectTypeString = "attribute"            // 属性
)

const (
	SubjectDomainGroup int8 = 1 + iota
	SubjectDomain
	BusinessObject
	BusinessActivity
	LogicEntity
	Attribute

	SubjectDomainObjectTypeRoot int8 = -1
)

var (
	SubjectDomainObjectTypeStringToObjectType = map[SubjectDomainObjectTypeString]int8{
		StringRoot:               SubjectDomainObjectTypeRoot,
		StringSubjectDomainGroup: SubjectDomainGroup,
		StringSubjectDomain:      SubjectDomain,
		StringBusinessObject:     BusinessObject,
		StringBusinessActivity:   BusinessActivity,
		StringLogicEntity:        LogicEntity,
		StringAttribute:          Attribute,
	}

	SubjectDomainObjectTypeToObjectTypeString = map[int8]SubjectDomainObjectTypeString{
		SubjectDomainObjectTypeRoot: StringRoot,
		SubjectDomainGroup:          StringSubjectDomainGroup,
		SubjectDomain:               StringSubjectDomain,
		BusinessObject:              StringBusinessObject,
		BusinessActivity:            StringBusinessActivity,
		LogicEntity:                 StringLogicEntity,
		Attribute:                   StringAttribute,
	}
)

func SubjectDomainObjectStringToInt(s string) int8 {
	return SubjectDomainObjectTypeStringToObjectType[SubjectDomainObjectTypeString(s)]
}
func SubjectDomainObjectIntToString(i int8) string {
	return string(SubjectDomainObjectTypeToObjectTypeString[i])
}
