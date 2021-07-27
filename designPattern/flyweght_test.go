package designPattern

import (
	"fmt"
	"testing"
	"time"
)

/**
享元模式
使用享元模式进行优化，将其中固定不变的部分设计成共享对象（享元，flyweight），这样就能节省大量的系统内存和CPU。
享元模式摒弃了在每个对象中保存所有数据的方式， 通过共享多个对象所共有的相同状态， 让你能在有限的内存容量中载入更多对象
 */

type TeamId =uint8

const (
	Warrior TeamId = iota
	Laker
)

type Player struct {
	Name string
	Team TeamId
}

type Team struct {
	Id TeamId
	Name string
	Players []*Player
}

type Match struct {
	Date time.Time
	LocalTeam *Team
	VisitorTeam *Team
	LocalScore uint8
	VisitorScore uint8
}

func (m *Match) ShowResult() {
	fmt.Printf("%s VS %s - %d:%d\n", m.LocalTeam.Name, m.VisitorTeam.Name,
		m.LocalScore, m.VisitorScore)
}

type teamFactory struct {
	teams map[TeamId]*Team
}


func (t *teamFactory) TeamOf(id TeamId) *Team {
	team, ok := t.teams[id]
	if !ok {
		team = createTeam(id)
		t.teams[id] = team
	}
	return team
}

var factory = &teamFactory{teams: make(map[TeamId]*Team)}

func FlyWeightFactory() *teamFactory {
	return factory
}

func createTeam(id TeamId) *Team {
	switch id {
	case Warrior:
		w := &Team{
			Id:      Warrior,
			Name:    "Golden State Warriors",
		}
		curry := &Player{
			Name: "Stephen Curry",
			Team: Warrior,
		}
		thompson := &Player{
			Name: "Klay Thompson",
			Team: Warrior,
		}
		w.Players = append(w.Players, curry, thompson)
		return w
	case Laker:
		l := &Team{
			Id:      Laker,
			Name:    "Los Angeles Lakers",
		}
		james := &Player{
			Name: "LeBron James",
			Team: Laker,
		}
		davis := &Player{
			Name: "Anthony Davis",
			Team: Laker,
		}
		l.Players = append(l.Players, james, davis)
		return l
	default:
		fmt.Printf("Get an invalid team id %v.\n", id)
		return nil
	}
}

func TestFlyWeight(t *testing.T) {
	game1 := &Match{
		Date:         time.Date(2020, 1, 10, 9, 30, 0, 0, time.Local),
		LocalTeam:    FlyWeightFactory().TeamOf(Warrior),
		VisitorTeam:  FlyWeightFactory().TeamOf(Laker),
		LocalScore:   102,
		VisitorScore: 99,
	}
	game1.ShowResult()
	game2 := &Match{
		Date:         time.Date(2020, 1, 12, 9, 30, 0, 0, time.Local),
		LocalTeam:    FlyWeightFactory().TeamOf(Laker),
		VisitorTeam:  FlyWeightFactory().TeamOf(Warrior),
		LocalScore:   110,
		VisitorScore: 118,
	}
	game2.ShowResult()
	// 两个Match的同一个球队应该是同一个实例的
	if game1.LocalTeam != game2.VisitorTeam {
		t.Errorf("Warrior team do not use flyweight pattern")
	}
}