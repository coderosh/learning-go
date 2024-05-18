package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Step int

const (
	MangaSearch Step = iota
	MangaSearchLoading

	MangaSelect
	ChaptersFindLoading

	ChapterSelect
	ChapterSelectLoading
)

type Model struct {
	width  int
	height int
	styles *Styles

	mangas          []Manga
	chapters        []Chapter
	chapterImages   []string
	selectedManga   *Manga
	selectedChapter *Chapter

	textInput   textinput.Model
	mangaList   list.Model
	chapterList list.Model

	step Step
}

type Styles struct {
	WithBorder lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}

	s.WithBorder = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("24")).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).
		Width(80)

	return s
}

func New() Model {
	styles := DefaultStyles()
	return Model{
		styles:        styles,
		chapterImages: []string{},
		mangas:        []Manga{},
		chapters:      []Chapter{},
	}
}

func (m *Model) Init() tea.Cmd {
	m.textInput = textinput.New()
	m.textInput.Focus()
	m.mangaList = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	m.chapterList = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var key string

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.mangaList.SetSize(m.width, m.height)
		m.chapterList.SetSize(m.width, m.height)

	case tea.KeyMsg:
		key = msg.String()

		if key == "ctrl+c" {
			return m, tea.Quit
		}
	}

	switch m.step {
	case MangaSearch:
		return MangaSearchUpdate(key, m, msg)
	case MangaSelect:
		return MangaSelectUpdate(key, m, msg)
	case ChapterSelect:
		return ChapterSelectUpdate(key, m, msg)
	case MangaSearchLoading:
		return MangaSearchLoadingUpdate(m)
	case ChaptersFindLoading:
		return ChaptersFindLoadingUpdate(m)
	case ChapterSelectLoading:
		return ChapterSelectLoadingUpdate(m)
	}

	return m, nil
}

func (m Model) View() string {
	switch m.step {
	case MangaSearch:
		return MangaSearchView(&m)
	case MangaSearchLoading:
		return MangaSearchLoadingView(&m)
	case MangaSelect:
		return MangaSelectView(&m)
	case ChaptersFindLoading:
		return ChaptersFindLoadingView(&m)
	case ChapterSelect:
		return ChapterSelectView(&m)
	case ChapterSelectLoading:
		return ChapterSelectLoadingView(&m)
	}

	return LoadingView(&m)
}

func main() {
	model := New()

	p := tea.NewProgram(&model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
