package library

import (
	"go-study/music/src/entity"
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("newMusicManager fail")
		return
	}
	if mm.Len() != 0 {
		t.Error("newMusicManager fail,musics is not empty")
		return
	}

	m0 := &entity.MusicEntity{Id: "1", Name: "年轮", Artist: "张碧晨", Source: "D://musics/1.MP3", Type: "MP3"}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("Add failed")
	}

	music := mm.Find(m0.Name)
	if music == nil || music.Name != m0.Name {
		t.Error("find failed")
	}
	music, _ = mm.Get(0)
	if music == nil {
		t.Error("Get Failed")
	}
	music = mm.Remove(0)
	if music == nil {
		t.Error("Remove Failed")
	}
}
