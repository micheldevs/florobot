package services

import (
	"log"
	"strings"
	"time"

	"github.com/micheldevs/florobot/dao"
	"github.com/micheldevs/florobot/models"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func ScrapeMoviesWebBulkMovies() {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", 4444)
	if err != nil {
		if service, err = selenium.NewChromeDriverService("./chromedriver", 4444); err != nil {
			log.Fatalln("Error:", err)
		}
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		// comment these lines for testing
		"--headless",
		"--no-sandbox",
		"--user-agent=Mozilla/5.0 (Linux; Android 13; SM-G998B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.5790.171 Mobile Safari/537.36",
		"--disable-dev-shm-usage",
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	driver.SetPageLoadTimeout(15 * time.Second)

	err = driver.Get("https://example.org/movies/")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	time.Sleep(time.Duration(20) * time.Second)

	if acceptCookiesBtn, err := driver.FindElement(selenium.ByXPATH, "//*[@id='onetrust-accept-btn-handler']"); err == nil {
		acceptCookiesBtn.Click()
	}

	moviesCategoriesBtn, err := driver.FindElements(selenium.ByXPATH, "//button/span[text() = 'Current movies' or text() = 'Upcoming movies']/parent::button")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	var moviesUrls = []string{}
	for _, moviesCategoryBtn := range moviesCategoriesBtn {
		moviesCategoryBtn.Click()
		time.Sleep(time.Duration(5) * time.Second)

		moviesListLis, err := driver.FindElements(selenium.ByXPATH, "//ul[@class='v-film-list-grid']/li")
		if err != nil {
			log.Fatalln("Error:", err)
		}

		for _, movieLi := range moviesListLis {
			movieAElement, err := movieLi.FindElement(selenium.ByXPATH, "./div/a")
			if err != nil {
				log.Fatalln("Error:", err)
			}

			url, _ := movieAElement.GetAttribute("href")

			if existMovie, _ := dao.ExistMovie(strings.ReplaceAll(url[strings.LastIndex(url, "/")-10:], "/", "")); !existMovie {
				moviesUrls = append(moviesUrls, url)
			}
		}
	}

	for _, movieUrl := range moviesUrls {
		err := driver.Get(movieUrl)
		if err != nil {
			log.Fatalln("Error:", err)
		}
		time.Sleep(time.Duration(5) * time.Second)

		titleWe, _ := driver.FindElement(selenium.ByXPATH, "//h1[@class='v-film-title__text']")
		title, _ := titleWe.Text()
		durationWe, _ := driver.FindElement(selenium.ByXPATH, "//div[@class='v-description-list-item v-film-runtime']/dd/div/span")
		duration, _ := durationWe.Text()
		releaseWe, _ := driver.FindElement(selenium.ByXPATH, "//div[@class='v-description-list-item v-film-release-date']/dd/div/span")
		release, _ := releaseWe.Text()
		releaseDate, _ := time.Parse("YYYY/MM/DD", release)
		genresWe, _ := driver.FindElements(selenium.ByXPATH, "//div[@class='v-description-list-item v-film-genres__list']/dd")
		genres, _ := genresWe[0].Text()
		for _, genreWe := range genresWe[1:] {
			genre, _ := genreWe.Text()
			genres += ", " + genre
		}
		actorsWe, _ := driver.FindElement(selenium.ByXPATH, "//div[@class='v-description-list-item v-film-actors']/dd/div/span")
		if actorsWe == nil {
			continue
		}
		actors, _ := actorsWe.Text()
		directorWe, _ := driver.FindElement(selenium.ByXPATH, "//div[@class='v-description-list-item v-film-directors']/dd/div/span")
		director := ""
		if directorWe != nil {
			director, _ = directorWe.Text()
		}
		synopsisWe, _ := driver.FindElement(selenium.ByXPATH, "//div[@class='v-description-list-item v-film-synopsis']/dd/div/span")
		synopsis := ""
		if synopsisWe != nil {
			synopsis, _ = synopsisWe.GetAttribute("innerText")
		}
		posterWe, _ := driver.FindElement(selenium.ByXPATH, "//div[@class='v-film-image v-film-image--media-type-poster']/div/img")
		poster, _ := posterWe.GetAttribute("src")

		dao.InsertMovie(&models.Movie{
			ExtId:        strings.ReplaceAll(movieUrl[strings.LastIndex(movieUrl, "/")-10:], "/", ""),
			ExtUrl:       movieUrl,
			Title:        title,
			PosterUrl:    poster,
			Synopsis:     synopsis,
			Actors:       actors,
			Director:     director,
			Duration:     duration,
			Genres:       genres,
			PremiereDate: releaseDate,
		})

	}
	driver.Close()
}
