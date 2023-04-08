package helpers

import (
	"html/template"

	"github.com/hail2skins/hamcois-new/models"
)

type NoteView struct {
	models.Note
	Published string
	Content   template.HTML
}

func NotesToNoteViews(notes *[]models.Note) []NoteView {
	var noteViews []NoteView
	for _, note := range *notes {
		published := note.UpdatedAt.Format("Jan 2, 2006")
		content := RenderMarkdownWithGoldmark(note.Content)
		noteView := NoteView{
			Note:      note,
			Published: published,
			Content:   template.HTML(content),
		}
		noteViews = append(noteViews, noteView)
	}
	return noteViews
}
