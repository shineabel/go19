package music

import "errors"

type Music struct {
	Name string
	Type string
	Source string
}

type MusicManager struct {
	musicList []Music
}

func NewMusicManager() *MusicManager  {

	return &MusicManager{
		musicList:make([]Music,0),
	}
}

func (m *MusicManager) Len()  int {
	return len(m.musicList)
}

func (m *MusicManager) Get(index int) (*Music, error)  {

	if len(m.musicList) <0 || index > len(m.musicList) {
		return nil, errors.New("out of range")
	}
	return &m.musicList[index], nil
}

func (m *MusicManager) Find(name string) *Music  {

	if len(m.musicList) == 0 {
		return nil
	}

	for _, mu := range  m.musicList {
		if mu.Name == name {
			return &mu
		}
	}
	return nil
}

func (m *MusicManager) Add(mu *Music)  {
	m.musicList = append(m.musicList,*mu)
}




func (m *MusicManager) Delete(index int) *Music  {

	if len(m.musicList) == 0 || index > len(m.musicList ) {
		return nil
	}
	dm := &m.musicList[index]

	m.musicList = append(m.musicList[:index], m.musicList[index+ 1:]...)


	return dm
}
