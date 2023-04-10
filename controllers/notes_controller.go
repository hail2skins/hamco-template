package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/alecthomas/chroma/quick"
	"github.com/hail2skins/hamcois-new/controllers/helpers"
	"github.com/hail2skins/hamcois-new/models"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yosssi/gohtml"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes, err := models.NotesAll()
	if err != nil {
		// Handle the error, e.g., log it and return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching notes"})
		return
	}
	noteViews := helpers.NotesToNoteViews(&notes)
	// Render the content of each note using Goldmark
	for i := range noteViews {
		truncatedContent := helpers.TruncateWords(string(noteViews[i].Content), 25) // Limit to 25 words, for example
		noteViews[i].Content = template.HTML(truncatedContent)
	}
	c.HTML(
		http.StatusOK,
		"note/index.html",
		gin.H{
			// Pass the slice of NoteView structs to the template rather than the notes directly
			"notes":     noteViews,
			"title":     "All the thoughts",
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"note/new.html",
		gin.H{
			"title":     "New Note",
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func NotesCreate(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		// Can do this through a bind.  See go-gin-bootcamp repo for actual code
		title := c.PostForm("title")
		content := c.PostForm("content")
		// call the model to create the note
		models.NotesCreate(currentUser, title, content)
		// redirect to the notes index
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func NotesShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing note id: %v\n", err)
	}

	note, _ := models.NotesFind(id)
	published := note.UpdatedAt.Format("Jan 2, 2006")

	htmlContent := renderMarkdown(note.Content)

	c.HTML(
		http.StatusOK,
		"note/show.html",
		gin.H{
			"note":      note,
			"content":   template.HTML(htmlContent),
			"published": published,
			"title":     note.Title,
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func renderMarkdown(noteContent string) string {
	unsafeHTML := parseMarkdown(noteContent)
	htmlContent := sanitizeHTML(unsafeHTML)
	preWrappedHTML := wrapPreCodeTags(htmlContent)
	highlightedHTML := highlightCodeBlocks(preWrappedHTML)
	formattedHTML := formatHTML(highlightedHTML)

	return formattedHTML
}

func parseMarkdown(noteContent string) []byte {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(noteContent), &buf); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func sanitizeHTML(unsafeHTML []byte) []byte {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("chroma")
	p.AllowElements("blockquote")
	htmlContent := p.SanitizeBytes(unsafeHTML)

	return htmlContent
}

func wrapPreCodeTags(htmlContent []byte) []byte {
	preWrappedHTML := bytes.ReplaceAll(htmlContent, []byte("<code>"), []byte("<pre><code>"))
	preWrappedHTML = bytes.ReplaceAll(preWrappedHTML, []byte("</code>"), []byte("</code></pre>"))

	return preWrappedHTML
}

func highlightCodeBlocks(gfmHTML []byte) []byte {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(gfmHTML))
	if err != nil {
		return gfmHTML
	}

	doc.Find("pre code").Each(func(i int, s *goquery.Selection) {
		code := s.Text()

		var b bytes.Buffer
		err := quick.Highlight(&b, code, "go", "html", "monokai")
		if err == nil {
			s.SetHtml(b.String())
		}
	})

	highlightedHTML, _ := doc.Html()

	return []byte(highlightedHTML)
}

func formatHTML(highlightedHTML []byte) string {
	formattedHTML := gohtml.Format(string(highlightedHTML))

	return formattedHTML
}

func NotesEditPage(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		idStr := c.Param("id")
		//fmt.Printf("ID string: %s\n", idStr) // Debugging statement
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing note id: %v\n", err)
		}
		//fmt.Printf("Parsed ID: %d\n", id) // Debugging statement
		note, _ := models.NotesFindByUser(currentUser, id)
		c.HTML(
			http.StatusOK,
			"note/edit.html",
			gin.H{
				"note":      note,
				"logged_in": c.MustGet("logged_in").(bool),
			},
		)
	}
}

func NotesUpdate(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing note id: %v\n", err)
		}
		note, _ := models.NotesFindByUser(currentUser, id)
		title := c.PostForm("title")
		content := c.PostForm("content")
		note.Update(title, content)
		c.Redirect(http.StatusMovedPermanently, "/notes/"+idStr)
	}
}

func NotesDelete(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing note id: %v\n", err)
		}
		models.NotesMarkDelete(currentUser, id)
		c.Redirect(http.StatusMovedPermanently, "/notes")
	}
}
