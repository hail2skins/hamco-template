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
	"github.com/russross/blackfriday/v2"
	"github.com/shurcooL/github_flavored_markdown"
	"github.com/yosssi/gohtml"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	noteViews := helpers.NotesToNoteViews(notes)
	c.HTML(
		http.StatusOK,
		"note/index.html",
		gin.H{
			// Pass the slice of NoteView structs to the template rather than the notes directly
			"notes":     noteViews,
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

	note := models.NotesFind(id)
	published := note.UpdatedAt.Format("Jan 2, 2006")

	htmlContent := renderMarkdown(note.Content)

	c.HTML(
		http.StatusOK,
		"note/show.html",
		gin.H{
			"note":      note,
			"content":   template.HTML(htmlContent),
			"published": published,
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func renderMarkdown(noteContent string) string {
	unsafeHTML := runBlackFriday(noteContent)
	htmlContent := sanitizeHTML(unsafeHTML)
	preWrappedHTML := wrapPreCodeTags(htmlContent)
	gfmHTML := renderGFM(preWrappedHTML)
	highlightedHTML := highlightCodeBlocks(gfmHTML)
	formattedHTML := formatHTML(highlightedHTML)

	return formattedHTML
}

func runBlackFriday(noteContent string) []byte {
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{})
	unsafeHTML := blackfriday.Run([]byte(noteContent),
		blackfriday.WithRenderer(renderer),
		blackfriday.WithExtensions(
			blackfriday.CommonExtensions|
				blackfriday.HardLineBreak|
				blackfriday.AutoHeadingIDs|
				blackfriday.Autolink|
				blackfriday.FencedCode|
				blackfriday.Footnotes,
		),
	)

	return unsafeHTML
}

func sanitizeHTML(unsafeHTML []byte) []byte {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	htmlContent := p.SanitizeBytes(unsafeHTML)

	return htmlContent
}

func wrapPreCodeTags(htmlContent []byte) []byte {
	preWrappedHTML := bytes.ReplaceAll(htmlContent, []byte("<code>"), []byte("<pre><code>"))
	preWrappedHTML = bytes.ReplaceAll(preWrappedHTML, []byte("</code>"), []byte("</code></pre>"))

	return preWrappedHTML
}

func renderGFM(preWrappedHTML []byte) []byte {
	gfmHTML := github_flavored_markdown.Markdown(preWrappedHTML)

	return gfmHTML
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
		note := models.NotesFindByUser(currentUser, id)
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
		note := models.NotesFindByUser(currentUser, id)
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
