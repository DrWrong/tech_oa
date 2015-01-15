package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id             int
	Name           string
	Slug           string `orm:"unique"`
	Classes        string // json编码的列表，项目涉及的班级
	GroupNumber    int    // 1 表示单独任务， 1以上表示分小组
	TaskTyps       string //任务种类, json编码 1 选题 2 阶段成果上传 3 学生打分 4 调查问卷
	Status         bool   `orm:"default(1)"` //是否结束
	FileUploadConf string //json编码的字符串, 项目对文件上传的需求

	Users           []*User            `orm:"reverse(many)"`
	Groups          []*Group           `orm:"reverse(many)"`
	Subjects        []*Subject         `orm:"null;rel(m2m);rel_through(tech_oa/models.ProjectSubjectChosen)"`
	Files           []*FileUploadGroup `orm:"reverse(many)"`
	Questions       []*Question        `orm:"null;rel(m2m)"`
	QuestionAnswers []*QuestionAnswer  `orm:"reverse(many)"`
}

func (p *Project) GetGroups() []*Group {
	if len(p.Groups) > 0 {
		return p.Groups
	}
	o := orm.NewOrm()
	o.LoadRelated(p, "Groups")
	return p.Groups
}

type Group struct {
	Id                   int
	Project              *Project              `orm:"rel(fk)"`
	GroupLeader          *User                 `orm:"rel(fk)"`
	MembersId            string                //json编码的字符串
	ProjectSubjectChosen *ProjectSubjectChosen `orm:"reverse(one)"`
	Files                []*FileUploadGroup    `orm:"reverse(many)"`
}

func (g *Group) GetDescription() string {
	o := orm.NewOrm()
	err := o.Read(g.GroupLeader)
	if err != nil {
		return fmt.Sprintf("%d", g.Id)
	}
	_, err = o.LoadRelated(g, "ProjectSubjectChosen")
	if err == nil {
		err = o.Read(g.ProjectSubjectChosen.Subject)
	}
	if err != nil {
		return g.GroupLeader.Name
	}

	return g.GroupLeader.Name + g.ProjectSubjectChosen.Subject.Name
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
		new(GroupScore), new(UserProject),
	)
}
