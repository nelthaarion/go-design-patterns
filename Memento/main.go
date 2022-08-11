package main

import (
	"log"
	"time"
)

type IMusicPlayer interface {
	CreateCheckPoint() *CheckPoint
	RestoreCheckPoint(InMemoryGit, int)
}

type CheckPoint struct {
	Time time.Time
}

type MusicPlayer struct {
	Current time.Time
}

func (m *MusicPlayer) CreateCheckPoint() *CheckPoint {
	return &CheckPoint{Time: time.Now()}
}
func (m *MusicPlayer) RestoreCheckPoint(img *InMemoryGit, number int) {
	if number >= len(img.CheckPoints) {
		log.Println("wrong checkpoint")
	}
	m.Current = img.CheckPoints[m][number].Time
}

type InMemoryGit struct {
	CheckPoints map[*MusicPlayer][]*CheckPoint
}

func (i *InMemoryGit) AddCheckPoint(c *CheckPoint, m *MusicPlayer) {
	i.CheckPoints[m] = append(i.CheckPoints[m], c)
}

func main() {
	player := &MusicPlayer{}
	inMemoryGit := &InMemoryGit{CheckPoints: map[*MusicPlayer][]*CheckPoint{
		player: *new([]*CheckPoint),
	}}
	cp := player.CreateCheckPoint()
	inMemoryGit.AddCheckPoint(cp, player)
	log.Println("a checkpoint created")

	time.Sleep(time.Second * 3)
	player.Current = time.Now()
	log.Printf("Current state is:%v\n", player.Current)
	time.Sleep(time.Second * 2)
	log.Println("Let's restore last state")
	time.Sleep(time.Second * 2)
	player.RestoreCheckPoint(inMemoryGit, 0)
	log.Println("state restored ")
	time.Sleep(time.Second * 2)
	log.Printf("current state is:%v", player.Current)

}
