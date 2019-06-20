package lexicon

var (
	_ lexicon = (*Lexicon)(nil)
)

type lexicon interface{}

type Lexicon map[string]interface{}
