package library

import (
	"errors"
	"go-study/music/src/entity"
)

type MusicManager struct {
	musics []entity.MusicEntity
}

// NewMusicManager 构造函数
func NewMusicManager() *MusicManager {
	// 传入一个数组切片用于存储歌曲列表
	return &MusicManager{make([]entity.MusicEntity, 0)}
}

// Len 获取歌曲列表大小
func (m *MusicManager) Len() int {
	return len(m.musics)
}

// Get 按下标查找歌曲
func (m *MusicManager) Get(index int) (music *entity.MusicEntity, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range")
	}
	return &m.musics[index], nil
}

// Find 按名声查找歌曲
func (m *MusicManager) Find(name string) *entity.MusicEntity {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

// Add 添加歌曲
func (m *MusicManager) Add(music *entity.MusicEntity) {
	m.musics = append(m.musics, *music)
}

// Remove 移除歌曲
func (m *MusicManager) Remove(index int) *entity.MusicEntity {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	rmMusic := &m.musics[index]
	if index < len(m.musics)-1 { // 中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { // 删除唯一元素
		m.musics = make([]entity.MusicEntity, 0)
	} else { //删除最后一个元素
		m.musics = m.musics[:index-1]
	}
	return rmMusic
}
