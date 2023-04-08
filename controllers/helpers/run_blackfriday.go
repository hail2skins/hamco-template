package helpers

import "github.com/russross/blackfriday/v2"

func RunBlackFriday(noteContent string) []byte {
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
