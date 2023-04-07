// note_view.go is a helper file to create a struct of notes and published dates for use in the templates.
// It is used in the notes_controller.go and homepage_controller.go files.
package helpers

import "github.com/hail2skins/hamcois-new/models"

type NoteView struct {
	models.Note
	Published string
}

func NotesToNoteViews(notes *[]models.Note) []NoteView {
	var noteViews []NoteView
	for _, note := range *notes {
		published := note.UpdatedAt.Format("Jan 2, 2006")
		noteView := NoteView{
			Note:      note,
			Published: published,
		}
		noteViews = append(noteViews, noteView)
	}
	return noteViews
}
