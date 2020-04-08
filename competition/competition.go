package competition

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/molin0000/secretMaster/config"
)

// var competitionPath = path.Join("data", "app", "me.cqp.molin.secretMaster", "competition.xlsx")

var competitionPath = config.CompetitionPath

// var competitionPath = "/Users/molin/coolq/data/app/me.cqp.molin.secretMaster/competition.xlsx"

var questions []*Question

type Question struct {
	Question string
	Option   string
	Answer   string
	Time     uint64
	Author   string
}

func atoi(msg string) uint64 {
	n, _ := strconv.Atoi(msg)
	return uint64(n)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func init() {
	fmt.Println("加载竞赛表格...")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("加载竞赛表格...失败")
			fmt.Println(err)
		}
	}()

	if !fileExists(competitionPath) {
		fmt.Println("文件不存在")
	}

	f, err := excelize.OpenFile(competitionPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := f.GetRows("竞赛答题搜集表")
	for i, row := range rows {
		if i == 0 {
			continue
		}

		q := &Question{
			Question: row[1],
			Option:   row[2],
			Answer:   row[3],
			Time:     atoi(row[4]),
			Author:   row[5],
		}

		questions = append(questions, q)
	}

	fmt.Printf("共加载竞赛题：%d\n", len(questions))
}

func GetRandomQuestion() *Question {
	length := len(questions)
	if length == 0 {
		return nil
	}
	q := questions[rand.Intn(length)]
	fmt.Println("q:", q)
	return q
}

func GetQuestionCount() int {
	return len(questions)
}
