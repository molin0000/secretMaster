package config

import (
	"os"
	"path"
	"path/filepath"
)

var CompetitionPath = path.Join(GetCoolqPath(), "data", "app", "me.cqp.molin.secretMaster", "competition.xlsx")

var PetPath = path.Join(GetCoolqPath(), "data", "app", "me.cqp.molin.secretMaster", "pets.xlsx")

var SkillPath = path.Join(GetCoolqPath(), "data", "app", "me.cqp.molin.secretMaster", "skills.xlsx")

var MissionPath = path.Join(GetCoolqPath(), "data", "app", "me.cqp.molin.secretMaster", "mission")

// var CompetitionPath = "/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/competition.xlsx"

// var PetPath = "/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/pets.xlsx"

// var SkillPath = "/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/skills.xlsx"

// var MissionPath = "/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/mission"

func GetCoolqPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
