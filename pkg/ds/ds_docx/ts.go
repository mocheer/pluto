package ds_docx

import (
	"archive/zip"

	"github.com/mocheer/pluto/pkg/ds/ds_xml"
)

type WordprocessingML struct {
	reader             *zip.Reader
	ContentTypes       *ds_xml.Document   // [Content_Types].xml
	RelsRoot           *ds_xml.Document   // _rels/.rels
	DocPropsCore       *ds_xml.Document   // docProps/core.xml
	DocPropsApp        *ds_xml.Document   // docProps/app.xml
	WordDocument       *ds_xml.Document   // word/document.xml
	WordStyles         *ds_xml.Document   // word/styles.xml
	WordNumbering      *ds_xml.Document   // word/numbering.xml
	WordFontTable      *ds_xml.Document   // word/fontTable.xml
	WordSettings       *ds_xml.Document   // word/settings.xml
	WordWebSettings    *ds_xml.Document   // word/webSettings.xml
	WordComments       *ds_xml.Document   // word/comments.xml
	WordCommentsExt    *ds_xml.Document   // word/commentsExtended.xml
	WordPeople         *ds_xml.Document   // word/people.xml
	WordFootnotes      *ds_xml.Document   // word/footnotes.xml
	WordEndnotes       *ds_xml.Document   // word/endnotes.xml
	WordHeaders        []*ds_xml.Document // word/header1.xml, header2.xml, ...
	WordFooters        []*ds_xml.Document // word/footer1.xml, footer2.xml, ...
	WordTheme          *ds_xml.Document   // word/theme/theme1.xml
	WordDocumentRels   *ds_xml.Document   // word/_rels/document.xml.rels
	CustomXmlItems     []*ds_xml.Document // customXml/item1.xml, ...
	CustomXmlItemProps []*ds_xml.Document // customXml/itemProps1.xml, ...
}
