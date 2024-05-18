package main

import "github.com/charmbracelet/lipgloss"

func centerAlign(m *Model, value string) string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		value,
	)
}

func MangaSearchView(m *Model) string {
	return centerAlign(
		m,
		lipgloss.JoinVertical(
			lipgloss.Left,
			"Search Manga",
			m.styles.WithBorder.Render(m.textInput.View()),
		),
	)
}

func MangaSelectView(m *Model) string {
	return m.mangaList.View()
}

func ChapterSelectView(m *Model) string {
	return m.chapterList.View()
}

func LoadingView(m *Model) string {
	return centerAlign(
		m,
		m.styles.WithBorder.Render("Loading..."),
	)
}

func MangaSearchLoadingView(m *Model) string {
	return centerAlign(
		m,
		m.styles.WithBorder.Render("Searching Manga"),
	)
}

func ChaptersFindLoadingView(m *Model) string {
	return centerAlign(
		m,
		m.styles.WithBorder.Render("Finding Chapters"),
	)
}

func ChapterSelectLoadingView(m *Model) string {
	return centerAlign(
		m,
		m.styles.WithBorder.Render("Loading Chapter"),
	)
}
