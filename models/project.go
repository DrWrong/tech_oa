package models

import (
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id             int
	Name           string
	Slug           string `orm:"unique"`
	Classes        string // json编码的列表，项目涉及的组号
	GroupNumber    int    // 1 表示单独任务， 1以上表示分小组
	TaskTyps       string //任务种类, json编码 1 选题 2 阶段成果上传 3 学生打分 4 调查问卷
	Status         bool   `orm:"default(1)"` //是否结束
	FileUploadConf string //json编码的字符串, 项目对文件上传的需求

	Users           []*User            `orm:"rel(m2m);rel_table(rel_user_project)"`
	Groups          []*Group           `orm:"reverse(many)"`
	Subjects        []*Subject         `orm:"null;rel(m2m);rel_through(tech_oa/models.ProjectSubjectChosen)"`
	Files           []*FileUploadGroup `orm:"reverse(many)"`
	Questions       []*Question        `orm:"null;rel(m2m)"`
	QuestionAnswers []*QuestionAnswer  `orm:"reverse(many)"`
}

type Group struct {
	Id                   int
	Project              *Project              `orm:"rel(fk)"`
	GroupLeader          *User                 `orm:"rel(fk)"`
	MembersId            string                //json编码的字符串
	ProjectSubjectChosen *ProjectSubjectChosen `orm:"reverse(one)"`
	Files                []*FileUploadGroup    `orm:"reverse(many)"`
}

type FileUploadGroup struct {
	Id       int
	Project  *Project `orm:"rel(fk)"`
	Group    *Group   `orm:"rel(fk)"`
	FilePath string   //文件上传路径
}

type GroupScore struct {
	Id         int
	FromGroup  *Group `orm:"rel(fk)"`
	JudgeGroup *Group `orm:"rel(fk)"`
	Score      int
}

type QuestionAnswer struct {
	Id       int
	Question *Question `orm:"rel(fk)"`
	Project  *Project  `orm:"rel(fk)"`
	Group    *Group    `orm:"rel(fk)"`
	Answer   string
}

func init() {

	orm.RegisterModel(
		new(Project), new(User), new(Subject), new(Question),
		new(Group), new(ProjectSubjectChosen),
		new(FileUploadGroup), new(QuestionAnswer),
		new(GroupScore),
	)
}
