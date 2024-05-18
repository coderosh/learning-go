package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gobeam/stringy"
)

func MangaSearchUpdate(key string, m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	if key == "enter" {
		m.step = MangaSearchLoading
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func MangaSelectUpdate(key string, m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	if key == "enter" {
		m.selectedManga = &m.mangas[m.mangaList.Index()]

		m.step = ChaptersFindLoading

		return m, m.mangaList.NewStatusMessage("") // to force update
	}

	var cmd tea.Cmd
	m.mangaList, cmd = m.mangaList.Update(msg)

	return m, cmd
}

func ChapterSelectUpdate(key string, m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	if key == "enter" {
		m.selectedChapter = &m.chapters[m.chapterList.Index()]

		m.step = ChapterSelectLoading
		m.chapterImages = GetChapterImagesUrl(m.selectedChapter)

		return m, m.chapterList.NewStatusMessage("") // force update
	}

	var cmd tea.Cmd
	m.chapterList, cmd = m.chapterList.Update(msg)

	return m, cmd
}

func MangaSearchLoadingUpdate(m *Model) (tea.Model, tea.Cmd) {
	m.mangas = SearchManga(m.textInput.Value())

	items := []list.Item{}

	for i := range m.mangas {
		items = append(items, &m.mangas[i])
	}

	m.mangaList.SetItems(items)
	m.mangaList.Title = "Mangas"

	m.step = MangaSelect

	return m, nil
}

func ChaptersFindLoadingUpdate(m *Model) (tea.Model, tea.Cmd) {
	m.chapters = GetChapters(m.selectedManga)

	items := []list.Item{}

	for i := range m.chapters {
		items = append(items, &m.chapters[i])
	}

	m.chapterList.SetItems(items)
	m.chapterList.Title = "Chapters"

	m.step = ChapterSelect

	return m, nil
}

func ChapterSelectLoadingUpdate(m *Model) (tea.Model, tea.Cmd) {
	name := stringy.New(m.selectedManga.Name + m.selectedChapter.Name).SnakeCase()
	CreatePdfFromImages(m.chapterImages, name.ToLower())

	m.step = MangaSearch

	return m, nil
}
