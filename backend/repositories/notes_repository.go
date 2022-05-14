package repositories

import (
	"moonbrain/models"
)

type NoteRepository struct {
	fakeDb map[string]models.Note
}

func NewNoteRepository() *NoteRepository {
	return &NoteRepository{fakeDb: make(map[string]models.Note)}
}

func (a *NoteRepository) GetNotes() ([]models.Note, error) {
	notes := []models.Note{}

	for _, note := range a.fakeDb {
		notes = append(notes, note)
	}

	return notes, nil
}

func (a *NoteRepository) AddNote(note models.Note) error {
	a.fakeDb[note.ID] = note
	return nil
}

func (a *NoteRepository) UpdateNote(article models.Note) error {
	a.fakeDb[article.ID] = article
	return nil
}

func (a *NoteRepository) BulkUpsert(notes []models.Note) error {
	for _, note := range notes {
		a.fakeDb[note.ID] = note
	}
	return nil
}

func (a *NoteRepository) GetNote(id string) (models.Note, error) {
	article, _ := a.fakeDb[id]

	return article, nil
}
