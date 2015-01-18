package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	// "github.com/macaron-contrib/binding"
	// "log"
	"strings"
	"time"
)

type Project struct {
	Id          int
	Name        string
	Classes     string // json编码的列表，项目涉及的班级
	GroupNumber int    // 1 表示单独任务， 1以上表示分小组
	Status      bool   `orm:"default(1)"` //是否结束

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

func (p *Project) GetAvaliableTasks() []*Task {
	o := orm.NewOrm()
	var tasks []*Task
	_, err := o.QueryTable("task").Filter(
		"Project", p.Id).Filter(
		"status", 1).Filter(
		"starttime__lte", time.Now()).Filter(
		"endtime__gte", time.Now()).All(&tasks)
	if err != nil {
		fmt.Println(err)
	}
	return tasks
}

type Task struct {
	Id          int
	Type        int      `orm:"default(0)"` //任务类型 1 选题 2 阶段成果上传 3 学生评分 4 调查问卷
	Project     *Project `orm:"rel(fk)"`
	Config      string   `orm:"default('')"` //json编码的字符串用于表示配置
	StartTime   time.Time
	EndTime     time.Time
	Status      bool `orm:"default(1)"`
	scoreConfig ScoreConfig
}

type ScoreConfig []struct {
	StuWeight   int    `json:"stu_weight"`
	Type        int    `json:"type"`
	ScoreWeight int    `json:"score_weight"`
	Desc        string `json:"desc"`
}

func (t *Task) ManReadTaskType() string {
	tasktypemap := map[int]string{
		1: "选题",
		2: "阶段成果上传",
		3: "学生评分",
		4: "调查问卷",
	}
	return tasktypemap[t.Type]
}

func (t *Task) TaskTypeSlug() string {
	tasktypemap := map[int]string{
		1: "chocie",
		2: "upload",
		3: "score",
		4: "question",
	}
	return tasktypemap[t.Type]
}

func (t *Task) GetScoreConfig() ScoreConfig {
	if t.Type == 3 {
		if t.scoreConfig == nil {
			json.NewDecoder(strings.NewReader(t.Config)).Decode(&t.scoreConfig)
		}
	}
	return t.scoreConfig
}

func GetAvaliableTasks() orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable("task").Filter("status", 1).Filter("starttime__lte", time.Now()).Filter("endtime__gte", time.Now())
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

// func (fromGroup *Group) GetJudegGroupScores() []*GroupScore {
// 	var groupscores []*GroupScore
// 	o := orm.NewOrm()
// 	o.QueryTable("group_score").Filter("FromGroup", fromGroup.Id).RelatedSel("JudgeGroup").All(&groupscores)
// 	return groupscores
// }

type FileUploadGroup struct {
	Id       int
	Project  *Project `orm:"rel(fk)"`
	Group    *Group   `orm:"rel(fk)"`
	FilePath string   //文件上传路径
}

type GroupScore struct {
	Id          int
	FromUser    *User    `orm:"rel(fk)"`
	JudgeGroup  *Group   `orm:"rel(fk)"`
	Score       int      //百分制度
	Type        int      //阶段性标志
	ScoreWeight int      `orm:"default(100)"` //这个阶段的比分在总分中占的权重, 百分制。百分之多少
	UserWeight  int      `orm:"default(100)"` //评分用户所占的权重 百分制，百分之多少
	Project     *Project `orm:"rel(fk)"`
	Task        *Task    `orm:"rel(fk)"`
}

func (g *GroupScore) GetJudgeGroup() *Group {
	if g.JudgeGroup != nil {
		return g.JudgeGroup
	}
	o := orm.NewOrm()
	o.Read(g.JudgeGroup)
	return g.JudgeGroup
}

func (g *GroupScore) GetTypeDesc() string {

	if g.Task == nil {
		o := orm.NewOrm()
		err := o.Read(g.Task)
		if err != nil {
			fmt.Println(err)
		}
	}
	// fmt.Println(g.Task.Id)
	// fmt.Println(g.Task.ManReadTaskType())
	for _, config := range g.Task.GetScoreConfig() {
		if config.Type == g.Type {
			return config.Desc
		}
	}
	return ""
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
		new(Group), new(ProjectSubjectChosen), new(Task),
		new(FileUploadGroup), new(QuestionAnswer),
		new(GroupScore), new(UserProject),
	)
}
