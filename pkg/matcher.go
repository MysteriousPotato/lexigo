package lexigo

import (
	"golang.org/x/text/language"
	"strings"
	"sync"
)

type Matcher struct {
	internal  language.Matcher
	matches   map[string]language.Tag
	mu        sync.RWMutex
	supported []language.Tag
}

func NewMatcher(supported []language.Tag, opts ...language.MatchOption) *Matcher {
	internal := language.NewMatcher(supported, opts...)
	return &Matcher{
		internal:  internal,
		matches:   map[string]language.Tag{},
		mu:        sync.RWMutex{},
		supported: supported,
	}
}

func (m *Matcher) addMatch(lang string, match language.Tag) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.matches[lang] = match
}

func (m *Matcher) getMatch(lang string) (language.Tag, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	match, ok := m.matches[lang]

	return match, ok
}

func (m *Matcher) MatchAll(langs ...language.Tag) language.Tag {
	for _, lang := range langs {
		for _, tag := range m.supported {
			if tag == lang {
				return tag
			}
		}
	}

	var str strings.Builder
	for _, l := range langs {
		str.WriteString(l.String())
	}

	if match, ok := m.getMatch(str.String()); ok {
		return match
	}

	match, _, _ := m.internal.Match(langs...)
	for {
		for _, tag := range m.supported {
			if tag == match {
				m.addMatch(str.String(), tag)
				return tag
			}
		}
		match = match.Parent()
		if match == language.Und {
			m.addMatch(str.String(), m.supported[0])
			return m.supported[0]
		}
	}
}

func (m *Matcher) Match(lang language.Tag) language.Tag {
	for _, tag := range m.supported {
		if tag == lang {
			return tag
		}
	}

	if match, ok := m.getMatch(lang.String()); ok {
		return match
	}

	match, _, _ := m.internal.Match(lang)
	for {
		for _, tag := range m.supported {
			if tag == match {
				m.addMatch(lang.String(), tag)
				return tag
			}
		}
		match = match.Parent()
		if match == language.Und {
			m.addMatch(lang.String(), m.supported[0])
			return m.supported[0]
		}
	}
}
