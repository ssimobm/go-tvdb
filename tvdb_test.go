package tvdb

import (
	"testing"
)

var (
	tvdb = New(APIKey)
)

func TestLogin(t *testing.T) {
	if err := tvdb.Login(); err != nil {
		t.Error(err)
	}
}

func TestRefreshToken(t *testing.T) {
	if err := tvdb.RefreshToken(); err != nil {
		t.Error(err)
	}
}

func TestLanguages(t *testing.T) {
	languages, err := tvdb.Languages()

	if err != nil {
		t.Error(err)
	}

	for _, language := range languages {
		if language.Name == "English" {
			return
		}
	}

	t.Error("No 'English' language could be found.")
}

func TestLanguageByID(t *testing.T) {
	language, err := tvdb.LanguageByID(7)

	if err != nil {
		t.Error(err)
	}

	if language.Name != "English" {
		t.Error("Unexpected language.")
	}
}

func TestSearchSeriesParams(t *testing.T) {
	_, err := tvdb.SearchSeriesParams()

	if err != nil {
		t.Error(err)
	}
}

func TestSearchSeries(t *testing.T) {
	params := map[string]string{
		"name": "The Simpsons",
	}
	seriesList, err := tvdb.SearchSeries(params, "en")

	if err != nil {
		t.Error(err)
	}

	for _, series := range seriesList {
		if series.SeriesName == "The Simpsons" {
			return
		}
	}

	t.Error("No 'The Simpsons' title could be found.")
}

// TestGetSeriesByID tests the GetSeriesByID function.
func TestGetSeriesByID(t *testing.T) {
	series, err := tvdb.GetSeriesByID(71663, "en")

	if err != nil {
		t.Error(err)
	}

	if series.SeriesName != "The Simpsons" {
		t.Error("ID lookup for '71663' failed.")
	}
}

func TestGetSeriesEpisodes(t *testing.T) {
	series, err := tvdb.GetSeriesByID(71663, "en")

	if err != nil {
		t.Error(err)
	}

	episodes, err := series.Episodes()

	if err != nil {
		t.Error(err)
	}

	if len(episodes) < 649 {
		t.Error("Can't retrive all episodes.")
	}
}

func TestGetSeriesEpisodesPage(t *testing.T) {
	series, err := tvdb.GetSeriesByID(71663, "en")

	if err != nil {
		t.Error(err)
	}

	episodes, err := series.EpisodesPage(5)

	if err != nil {
		t.Error(err)
	}

	if len(episodes) != 100 {
		t.Error("Episodes page lookup error.")
	}
}

func TestGetSeriesEpisodeSummary(t *testing.T) {
	series, err := tvdb.GetSeriesByID(71663, "en")

	if err != nil {
		t.Error(err)
	}

	summary, err := series.SeriesEpisodeSummary()

	if err != nil {
		t.Error(err)
	}

	if len(summary.AiredSeasons) < 27 {
		t.Error("Invalid series episode summary.")
	}
}

func TestGetSeriesActors(t *testing.T) {
	series, err := tvdb.GetSeriesByID(71663, "en")

	if err != nil {
		t.Error(err)
	}

	actors, err := series.Actors()

	if err != nil {
		t.Error(err)
	}

	for _, actor := range actors {
		if actor.Name == "Nancy Cartwright" {
			return
		}
	}

	t.Error("No 'Nany Cartwright' actor could be found.")
}

func TestGetSeriesImages(t *testing.T) {
	series, err := tvdb.GetSeriesByID(71663, "en")

	if err != nil {
		t.Error(err)
	}

	_, err = series.Images()

	if err != nil {
		t.Error(err)
	}
}

func TestGetEpisode(t *testing.T) {
	episode, err := tvdb.GetEpisodeByID(55561, "en")

	if err != nil {
		t.Error(err)
	}

	if episode.EpisodeName != "Bart's Girlfriend" {
		t.Error("ID lookup for '55561' failed.")
	}
}

/*
// TestGetSeriesByIMDBID tests the GetSeriesByIMDBID function.
func TestGetSeriesByIMDBID(t *testing.T) {
	series, err := GetSeriesByIMDBID("tt0096697")

	if err != nil {
		t.Error(err)
	}

	if series.SeriesName != "The Simpsons" {
		t.Error("IMDb ID lookup for 'tt0096697' failed.")
	}
}

// TestSearchSeries tests the SearchSeries function.
func TestSearchSeries(t *testing.T) {
	seriesList, err := SearchSeries("The Simpsons", 5)

	if err != nil {
		t.Error(err)
	}

	for _, series := range seriesList.Series {
		if series.SeriesName == "The Simpsons" {
			return
		}
	}

	t.Error("No 'The Simpsons' title could be found.")
}

// TestSeriesGetDetail tests the Series type's GetDetail function.
func TestSeriesGetDetail(t *testing.T) {
	series, err := GetSeriesByID(71663)

	if err != nil {
		t.Error(err)
	}

	if series.Seasons != nil {
		t.Error("series.Seasons should be nil.")
	}

	series.GetDetail()

	if series.Seasons == nil {
		t.Error("series.Seasons should not be nil.")
	}
}

// TestSeriesListGetDetail tests the SeriesList type's GetDetail function.
func TestSeriesListGetDetail(t *testing.T) {
	seriesList, err := tvdb.GetSeries("The Simpsons")

	if err != nil {
		t.Error(err)
	}

	for _, series := range seriesList.Series {
		if series.Seasons != nil {
			t.Error("series.Seasons should be nil.")
		}
	}

	seriesList.GetDetail()

	for _, series := range seriesList.Series {
		if series.Seasons == nil {
			t.Error("series.Seasons should not be nil.")
		}

		// Need to check that a value not present in GetSeries result is now
		// available
		if series.Poster == "" {
			t.Error("series.poster should not be empty.")
		}
	}
}
*/
