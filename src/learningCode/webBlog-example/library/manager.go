package library

import "errors"

type MusicManager struct {
	music []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.music)
}

func (m *MusicManager) Get(index int) (music *MusicManager, err error) {
	if index < 0 || index >= len(m.music) {
		return nil, errors.New("Index out of range.")
	}
	return &m.music[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
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

func (m *MusicManager) Add(music *MusicManager) {
	m.musics = append(m.musics, *music)

}

func (m *MusicManager) Remove(index int) *MusicManager {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)

	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}
