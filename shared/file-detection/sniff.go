package detection

import (
	"bytes"
	"encoding/binary"
)

const sniffLen = 128

type sniffSig interface {
	// match returns the MIME type of the data, or "" if unknown.
	match(data []byte, firstNonWS int) string
}

var sniffSignatures = []sniffSig{
	// adding here

	// JPEG 2000
	&maskedSig{
		mask: []byte("\x00\x00\x00\x00\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
		pat:  []byte("\x00\x00\x00\x00\x6A\x50\x20\x20\x00\x00\x00\x00\x00\x00\x00\x00\x66\x74\x79\x70\x6A\x70\x32\x20"),
		ct:   "image/jp2",
	},
	&maskedSig{
		mask: []byte("\x00\x00\x00\x00\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
		pat:  []byte("\x00\x00\x00\x00\x6A\x50\x20\x20\x00\x00\x00\x00\x00\x00\x00\x00\x66\x74\x79\x70\x6A\x70\x78\x20"),
		ct:   "image/jpx",
	},

	// Portable any map
	&exactSig{[]byte("\x50\x34\x0A"), "image/x-portable-bitmap"},
	&exactSig{[]byte("\x50\x35\x0A"), "image/x-portable-graymap"},
	&exactSig{[]byte("\x50\x36\x0A"), "image/x-portable-pixmap"},
	&exactSig{[]byte("\x50\x37\x0A"), "image/x-portable-arbitrarymap"},
	&exactSig{[]byte("\x03\x37\x3A"), "image/x-portable-bitmap"},

	// TIFF
	&exactSig{[]byte("\x49\x20\x49"), "image/tiff"},
	&exactSig{[]byte("\x49\x49\x2A\x00"), "image/tiff"},
	&exactSig{[]byte("\x4D\x4D\x00\x2A"), "image/tiff"},
	&exactSig{[]byte("\x4D\x4D\x00\x2B"), "image/tiff"},

	// default of http go package
	htmlSig("<!DOCTYPE HTML"),
	htmlSig("<HTML"),
	htmlSig("<HEAD"),
	htmlSig("<SCRIPT"),
	htmlSig("<IFRAME"),
	htmlSig("<H1"),
	htmlSig("<DIV"),
	htmlSig("<FONT"),
	htmlSig("<TABLE"),
	htmlSig("<A"),
	htmlSig("<STYLE"),
	htmlSig("<TITLE"),
	htmlSig("<B"),
	htmlSig("<BODY"),
	htmlSig("<BR"),
	htmlSig("<P"),
	htmlSig("<!--"),

	&maskedSig{mask: []byte("\xFF\xFF\xFF\xFF\xFF"), pat: []byte("<?xml"), skipWS: true, ct: "text/xml; charset=utf-8"},

	&exactSig{[]byte("%PDF-"), "application/pdf"},
	&exactSig{[]byte("%!PS-Adobe-"), "application/postscript"},

	// UTF BOMs.
	&maskedSig{mask: []byte("\xFF\xFF\x00\x00"), pat: []byte("\xFE\xFF\x00\x00"), ct: "text/plain; charset=utf-16be"},
	&maskedSig{mask: []byte("\xFF\xFF\x00\x00"), pat: []byte("\xFF\xFE\x00\x00"), ct: "text/plain; charset=utf-16le"},
	&maskedSig{mask: []byte("\xFF\xFF\xFF\x00"), pat: []byte("\xEF\xBB\xBF\x00"), ct: "text/plain; charset=utf-8"},

	&exactSig{[]byte("GIF87a"), "image/gif"},
	&exactSig{[]byte("GIF89a"), "image/gif"},
	&exactSig{[]byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A"), "image/png"},
	&exactSig{[]byte("\xFF\xD8\xFF"), "image/jpeg"},
	&exactSig{[]byte("BM"), "image/bmp"},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF"),
		pat:  []byte("RIFF\x00\x00\x00\x00WEBPVP"),
		ct:   "image/webp",
	},
	&exactSig{[]byte("\x00\x00\x01\x00"), "image/vnd.microsoft.icon"},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF"),
		pat:  []byte("RIFF\x00\x00\x00\x00WAVE"),
		ct:   "audio/wave",
	},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF"),
		pat:  []byte("FORM\x00\x00\x00\x00AIFF"),
		ct:   "audio/aiff",
	},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF\xFF"),
		pat:  []byte(".snd"),
		ct:   "audio/basic",
	},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF\xFF\xFF"),
		pat:  []byte("OggS\x00"),
		ct:   "application/ogg",
	},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
		pat:  []byte("MThd\x00\x00\x00\x06"),
		ct:   "audio/midi",
	},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF"),
		pat:  []byte("ID3"),
		ct:   "audio/mpeg",
	},
	&maskedSig{
		mask: []byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\xFF\xFF\xFF\xFF"),
		pat:  []byte("RIFF\x00\x00\x00\x00AVI "),
		ct:   "video/avi",
	},
	&exactSig{[]byte("\x1A\x45\xDF\xA3"), "video/webm"},
	&exactSig{[]byte("\x52\x61\x72\x20\x1A\x07\x00"), "application/x-rar-compressed"},
	&exactSig{[]byte("\x50\x4B\x03\x04"), "application/zip"},
	&exactSig{[]byte("\x1F\x8B\x08"), "application/x-gzip"},

	mp4Sig{},

	textSig{}, // should be last
}

type exactSig struct {
	sig []byte
	ct  string
}

func (e *exactSig) match(data []byte, firstNonWS int) string {
	if bytes.HasPrefix(data, e.sig) {
		return e.ct
	}
	return ""
}

type maskedSig struct {
	mask, pat []byte
	skipWS    bool
	ct        string
}

func (m *maskedSig) match(data []byte, firstNonWS int) string {
	if m.skipWS {
		data = data[firstNonWS:]
	}
	if len(m.pat) != len(m.mask) {
		return ""
	}
	if len(data) < len(m.mask) {
		return ""
	}
	for i, mask := range m.mask {
		db := data[i] & mask
		if db != m.pat[i] {
			return ""
		}
	}
	return m.ct
}

type htmlSig []byte

func (h htmlSig) match(data []byte, firstNonWS int) string {
	data = data[firstNonWS:]
	if len(data) < len(h)+1 {
		return ""
	}
	for i, b := range h {
		db := data[i]
		if 'A' <= b && b <= 'Z' {
			db &= 0xDF
		}
		if b != db {
			return ""
		}
	}
	// Next byte must be space or right angle bracket.
	if db := data[len(h)]; db != ' ' && db != '>' {
		return ""
	}
	return "text/html; charset=utf-8"
}

var mp4ftype = []byte("ftyp")
var mp4 = []byte("mp4")

type mp4Sig struct{}

func (mp4Sig) match(data []byte, firstNonWS int) string {
	// https://mimesniff.spec.whatwg.org/#signature-for-mp4
	// c.f. section 6.2.1
	if len(data) < 12 {
		return ""
	}
	boxSize := int(binary.BigEndian.Uint32(data[:4]))
	if boxSize%4 != 0 || len(data) < boxSize {
		return ""
	}
	if !bytes.Equal(data[4:8], mp4ftype) {
		return ""
	}
	for st := 8; st < boxSize; st += 4 {
		if st == 12 {
			// minor version number
			continue
		}
		if bytes.Equal(data[st:st+3], mp4) {
			return "video/mp4"
		}
	}
	return ""
}

type textSig struct{}

func (textSig) match(data []byte, firstNonWS int) string {
	// c.f. section 5, step 4.
	for _, b := range data[firstNonWS:] {
		switch {
		case b <= 0x08,
			b == 0x0B,
			0x0E <= b && b <= 0x1A,
			0x1C <= b && b <= 0x1F:
			return ""
		}
	}
	return "text/plain; charset=utf-8"
}

// DetectContentType return a internet media type
func DetectContentType(data []byte) string {
	if len(data) > sniffLen {
		data = data[:sniffLen]
	}

	// Index of the first non-whitespace byte in data.
	firstNonWS := 0
	for ; firstNonWS < len(data) && isWS(data[firstNonWS]); firstNonWS++ {
	}

	for _, sig := range sniffSignatures {
		if ct := sig.match(data, firstNonWS); ct != "" {
			return ct
		}
	}
	return "application/octet-stream"
}

func isWS(b byte) bool {
	switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
		return true
	}
	return false
}
