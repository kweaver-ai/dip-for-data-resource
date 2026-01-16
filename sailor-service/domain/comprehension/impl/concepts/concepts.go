package concepts

var NegativeConcept = "负面(negative，downside)，通常是指不好的，坏的, 消极的一面。数据表有很多字段，结合表名称，每个字段都有一个或者多个功能，这些功能有负面的，有正面的。"

var ConceptSet = map[string][]string{
	"负面支撑": {
		NegativeConcept,
	},
}

//Support 获取概念上的支持
func Support(tag string) []string {
	return ConceptSet[tag]
}
