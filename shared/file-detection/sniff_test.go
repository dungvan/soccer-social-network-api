package detection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sniffTests = []struct {
	desc        string
	data        []byte
	contentType string
}{
	// Some nonsense.
	{"129 bytes data", []byte("5DWPkb5HysD1z6Dz68krJgqWY18M7aH12Anp8kFSju5w87hdmlIdBiVmUvQjJMRVmYHZQemdzC6FncRukg5hdmfmbOakx3MdB9tvZcfUXcLstFGdt1xQThuhJHdfbQkQgq"), "text/plain; charset=utf-8"},
	{"Empty", []byte{}, "text/plain; charset=utf-8"},
	{"Binary", []byte{1, 2, 3}, "application/octet-stream"},

	{"HTML document #1", []byte(`<HtMl><bOdY>blah blah blah</body></html>`), "text/html; charset=utf-8"},
	{"HTML document #2", []byte(`<HTML></HTML>`), "text/html; charset=utf-8"},
	{"HTML document #3 (leading whitespace)", []byte(`   <!DOCTYPE HTML>...`), "text/html; charset=utf-8"},
	{"HTML document #4 (leading CRLF)", []byte("\r\n<html>..."), "text/html; charset=utf-8"},

	{"Plain text", []byte(`This is not HTML. It has ☃ though.`), "text/plain; charset=utf-8"},

	{"XML", []byte("\n<?xml!"), "text/xml; charset=utf-8"},

	// Image types.
	{"GIF 87a", []byte(`GIF87a`), "image/gif"},
	{"GIF 89a", []byte(`GIF89a...`), "image/gif"},

	// Audio types.
	{"MIDI audio", []byte("MThd\x00\x00\x00\x06\x00\x01"), "audio/midi"},
	{"MP3 audio/MPEG audio", []byte("ID3\x03\x00\x00\x00\x00\x0f"), "audio/mpeg"},
	{"WAV audio #1", []byte("RIFFb\xb8\x00\x00WAVEfmt \x12\x00\x00\x00\x06"), "audio/wave"},
	{"WAV audio #2", []byte("RIFF,\x00\x00\x00WAVEfmt \x12\x00\x00\x00\x06"), "audio/wave"},
	{"AIFF audio #1", []byte("FORM\x00\x00\x00\x00AIFFCOMM\x00\x00\x00\x12\x00\x01\x00\x00\x57\x55\x00\x10\x40\x0d\xf3\x34"), "audio/aiff"},

	{"OGG audio", []byte("OggS\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x7e\x46\x00\x00\x00\x00\x00\x00\x1f\xf6\xb4\xfc\x01\x1e\x01\x76\x6f\x72"), "application/ogg"},
	{"Must not match OGG", []byte("owow\x00"), "application/octet-stream"},
	{"Must not match OGG", []byte("oooS\x00"), "application/octet-stream"},
	{"Must not match OGG", []byte("oggS\x00"), "application/octet-stream"},

	// Video types.
	{"MP4 video", []byte("\x00\x00\x00\x18ftypmp42\x00\x00\x00\x00mp42isom<\x06t\xbfmdat"), "video/mp4"},
	{"AVI video #1", []byte("RIFF,O\n\x00AVI LISTÀ"), "video/avi"},
	{"AVI video #2", []byte("RIFF,\n\x00\x00AVI LISTÀ"), "video/avi"},
}

func TestDetectContentType(t *testing.T) {
	for _, tt := range sniffTests {
		ct := DetectContentType(tt.data)
		assert.Equal(t, tt.contentType, ct)
	}
}

func Test_isWS(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test byte escape sequences \\t (horizontal tab)",
			args{
				0x09,
			},
			true,
		},
		{
			"test byte escape sequences \\n (new line)",
			args{
				0x0A,
			},
			true,
		},
		{
			"test byte escape sequences 0x0C (new page)",
			args{
				0x0C,
			},
			true,
		},
		{
			"test byte escape sequences \\r (carriage return)",
			args{
				0x0D,
			},
			true,
		},
		{
			"test byte escape sequences space",
			args{
				0x20,
			},
			true,
		},
		{
			"test byte is not escape sequences",
			args{
				0x3A,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isWS(tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_exactSig_match(t *testing.T) {
	type fields struct {
		sig []byte
		ct  string
	}
	type args struct {
		data       []byte
		firstNonWS int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"test with right prefix of image/x-portable-bitmap",
			fields{
				[]byte("\x50\x34\x0A"),
				"image/x-portable-bitmap",
			},
			args{
				[]byte("\x50\x34\x0A"),
				0,
			},
			"image/x-portable-bitmap",
		},
		{
			"test with wrong prefix",
			fields{
				[]byte("\x50\x34\x0A"),
				"image/x-portable-bitmap",
			},
			args{
				[]byte("\x50\x50\x50"),
				0,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &exactSig{
				sig: tt.fields.sig,
				ct:  tt.fields.ct,
			}
			got := e.match(tt.args.data, tt.args.firstNonWS)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_maskedSig_match(t *testing.T) {
	type fields struct {
		mask   []byte
		pat    []byte
		skipWS bool
		ct     string
	}
	type args struct {
		data       []byte
		firstNonWS int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"test with right type of image/jp2",
			fields{
				mask: []byte("\x00\x00\x00\x00\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
				pat:  []byte("\x00\x00\x00\x00\x6A\x50\x20\x20\x00\x00\x00\x00\x00\x00\x00\x00\x66\x74\x79\x70\x6A\x70\x32\x20"),
				ct:   "image/jp2",
			},
			args{
				[]byte("\x00\x00\x00\x0C\x6A\x50\x20\x20\x0D\x0A\x87\x0A\x00\x00\x00\x1C\x66\x74\x79\x70\x6A\x70\x32\x20"),
				0,
			},
			"image/jp2",
		},
		{
			"test no match byte",
			fields{
				mask: []byte("\x00\x00\x00\x00\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
				pat:  []byte("\x00\x00\x00\x00\x6A\x50\x20\x20\x00\x00\x00\x00\x00\x00\x00\x00\x66\x74\x79\x70\x6A\x70\x32\x20"),
				ct:   "image/jp2",
			},
			args{
				[]byte("\x00\x00\x00\x0C\x6A\x51\x21\x21\x0D\x0A\x87\x0A\x00\x00\x00\x1C\x66\x74\x79\x70\x6A\x70\x32\x20"),
				0,
			},
			"",
		},
		{
			"test length of data less than mask",
			fields{
				mask: []byte("\x00\x00\x00\x00\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
				pat:  []byte("\x00\x00\x00\x00\x6A\x50\x20\x20\x00\x00\x00\x00\x00\x00\x00\x00\x66\x74\x79\x70\x6A\x70\x32\x20"),
				ct:   "image/jp2",
			},
			args{
				[]byte("\x6A\x50\x20\x20\x0D\x0A\x87\x0A\x00\x00\x00\x1C\x66\x74\x79\x70\x6A\x70\x32\x20"),
				0,
			},
			"",
		},
		{
			"test length of mask not equal to length of pat",
			fields{
				mask: []byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
				pat:  []byte("\x00\x00\x00\x00\x6A\x50\x20\x20\x00\x00\x00\x00\x00\x00\x00\x00\x66\x74\x79\x70\x6A\x70\x32\x20"),
				ct:   "image/jp2",
			},
			args{
				[]byte("\x00\x00\x00\x0C\x6A\x50\x20\x20\x0D\x0A\x87\x0A\x00\x00\x00\x1C\x66\x74\x79\x70\x6A\x70\x32\x20"),
				0,
			},
			"",
		},
		{
			"test with skipWS",
			fields{
				mask:   []byte("\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
				pat:    []byte("\x00\x00\x00\x00\x6A\x50\x20\x20\x00\x00\x00\x00\x00\x00\x00\x00\x66\x74\x79\x70\x6A\x70\x32\x20"),
				skipWS: true,
				ct:     "image/jp2",
			},
			args{
				[]byte("\x00\x00\x00\x0C\x6A\x50\x20\x20\x0D\x0A\x87\x0A\x00\x00\x00\x1C\x66\x74\x79\x70\x6A\x70\x32\x20"),
				0,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &maskedSig{
				mask:   tt.fields.mask,
				pat:    tt.fields.pat,
				skipWS: tt.fields.skipWS,
				ct:     tt.fields.ct,
			}
			got := m.match(tt.args.data, tt.args.firstNonWS)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_htmlSig_match(t *testing.T) {
	type args struct {
		data       []byte
		firstNonWS int
	}
	tests := []struct {
		name string
		h    htmlSig
		args args
		want string
	}{
		{
			"test with length of data less than length of htmlSig",
			htmlSig("<!DOCTYPE HTML"),
			args{
				[]byte("<!DOCTYPE HTM"),
				0,
			},
			"",
		},
		{
			"test with data is string with no open tag html (< or >)",
			htmlSig("<!DOCTYPE HTML"),
			args{
				[]byte("DOCTYPE HTML "),
				0,
			},
			"",
		},
		{
			"test with wrong tag",
			htmlSig("<!DOCTYPE HTML"),
			args{
				[]byte("<!HOCTYPE HTML "),
				0,
			},
			"",
		},
		{
			"test with not space and not right angle bracket finish",
			htmlSig("<!DOCTYPE HTML"),
			args{
				[]byte("<!DOCTYPE HTML-"),
				0,
			},
			"",
		},
		{
			"test success type of text/html",
			htmlSig("<!DOCTYPE HTML"),
			args{
				[]byte("<!DOCTYPE HTML "),
				0,
			},
			"text/html; charset=utf-8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.h.match(tt.args.data, tt.args.firstNonWS)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_mp4Sig_match(t *testing.T) {
	type args struct {
		data       []byte
		firstNonWS int
	}
	tests := []struct {
		name string
		m    mp4Sig
		args args
		want string
	}{
		{
			"right mp4 MIME type",
			mp4Sig{},
			args{
				[]byte("\x00\x00\x00\x18\x66\x74\x79\x70\x6D\x70\x34\x32\x00\x00\x00\x00\x69\x73\x6F\x6D\x6D\x70\x34\x32\x00\x00\x18\xB2\x6D\x6F\x6F\x76"),
				0,
			},
			"video/mp4",
		},
		{
			"right data length less than 12",
			mp4Sig{},
			args{
				[]byte("\x00\x00\x00\x18\x66\x74\x79"),
				0,
			},
			"",
		},
		{
			"boxsize mod 4 not equal 0",
			mp4Sig{},
			args{
				[]byte("\x00\x00\x00\x17\x66\x74\x79\x70\x6D\x70\x34\x32\x00\x00\x00\x00\x69\x73\x6F\x6D\x6D\x70\x34\x32\x00\x00\x18\xB2\x6D\x6F\x6F\x76"),
				0,
			},
			"",
		},
		{
			"length of data less than box size",
			mp4Sig{},
			args{
				[]byte("\x00\x00\x00\x30\x66\x74\x79\x70\x6D\x70\x34\x32\x00\x00\x00\x00\x69\x73\x6F\x6D\x6D\x70\x34\x32\x00\x00\x18\xB2\x6D\x6F\x6F\x76"),
				0,
			},
			"",
		},
		{
			"data[4:8] not equal ftype",
			mp4Sig{},
			args{
				[]byte("\x00\x00\x00\x18\x65\x74\x79\x70\x6D\x70\x34\x32\x00\x00\x00\x00\x69\x73\x6F\x6D\x6D\x70\x34\x32\x00\x00\x18\xB2\x6D\x6F\x6F\x76"),
				0,
			},
			"",
		},
		{
			"wrong mp4 MIME type",
			mp4Sig{},
			args{
				[]byte("\x00\x00\x00\x18\x66\x74\x79\x70\x6E\x71\x35\x32\x00\x00\x00\x00\x69\x73\x6F\x6D\x6E\x71\x35\x32\x00\x00\x18\xB2\x6D\x6F\x6F\x76"),
				0,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mp4Sig{}
			got := m.match(tt.args.data, tt.args.firstNonWS)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_textSig_match(t *testing.T) {
	type args struct {
		data       []byte
		firstNonWS int
	}
	tests := []struct {
		name string
		t    textSig
		args args
		want string
	}{
		{
			"test not text file",
			textSig{},
			args{
				[]byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A"),
				0,
			},
			"",
		},
		{
			"test right text file",
			textSig{},
			args{
				[]byte("This is test plain file"),
				0,
			},
			"text/plain; charset=utf-8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := textSig{}
			got := ts.match(tt.args.data, tt.args.firstNonWS)
			assert.Equal(t, tt.want, got)
		})
	}
}
