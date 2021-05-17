package radarr

import "time"

type Movie struct {
	Title                 string        `json:"title"`
	Originaltitle         string        `json:"originalTitle"`
	Alternatetitles       []interface{} `json:"alternateTitles"`
	Secondaryyearsourceid int           `json:"secondaryYearSourceId"`
	Sorttitle             string        `json:"sortTitle"`
	Sizeondisk            int           `json:"sizeOnDisk"`
	Status                string        `json:"status"`
	Overview              string        `json:"overview"`
	Incinemas             time.Time     `json:"inCinemas"`
	Images                []struct {
		Covertype string `json:"coverType"`
		URL       string `json:"url"`
		Remoteurl string `json:"remoteUrl"`
	} `json:"images"`
	Website             string        `json:"website"`
	Year                int           `json:"year"`
	Hasfile             bool          `json:"hasFile"`
	Youtubetrailerid    string        `json:"youTubeTrailerId"`
	Studio              string        `json:"studio"`
	Path                string        `json:"path"`
	Qualityprofileid    int           `json:"qualityProfileId"`
	Monitored           bool          `json:"monitored"`
	Minimumavailability string        `json:"minimumAvailability"`
	Isavailable         bool          `json:"isAvailable"`
	Foldername          string        `json:"folderName"`
	Runtime             int           `json:"runtime"`
	Cleantitle          string        `json:"cleanTitle"`
	Imdbid              string        `json:"imdbId"`
	Tmdbid              int           `json:"tmdbId"`
	Titleslug           string        `json:"titleSlug"`
	Certification       string        `json:"certification"`
	Genres              []string      `json:"genres"`
	Tags                []interface{} `json:"tags"`
	Added               time.Time     `json:"added"`
	Ratings             struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	Moviefile struct {
		Movieid      int       `json:"movieId"`
		Relativepath string    `json:"relativePath"`
		Path         string    `json:"path"`
		Size         int       `json:"size"`
		Dateadded    time.Time `json:"dateAdded"`
		Scenename    string    `json:"sceneName"`
		Indexerflags int       `json:"indexerFlags"`
		Quality      struct {
			Quality struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				Source     string `json:"source"`
				Resolution int    `json:"resolution"`
				Modifier   string `json:"modifier"`
			} `json:"quality"`
			Revision struct {
				Version  int  `json:"version"`
				Real     int  `json:"real"`
				Isrepack bool `json:"isRepack"`
			} `json:"revision"`
		} `json:"quality"`
		Mediainfo struct {
			Audioadditionalfeatures string  `json:"audioAdditionalFeatures"`
			Audiobitrate            int     `json:"audioBitrate"`
			Audiochannels           float64 `json:"audioChannels"`
			Audiocodec              string  `json:"audioCodec"`
			Audiolanguages          string  `json:"audioLanguages"`
			Audiostreamcount        int     `json:"audioStreamCount"`
			Videobitdepth           int     `json:"videoBitDepth"`
			Videobitrate            int     `json:"videoBitrate"`
			Videocodec              string  `json:"videoCodec"`
			Videofps                float64 `json:"videoFps"`
			Resolution              string  `json:"resolution"`
			Runtime                 string  `json:"runTime"`
			Scantype                string  `json:"scanType"`
			Subtitles               string  `json:"subtitles"`
		} `json:"mediaInfo"`
		Qualitycutoffnotmet bool `json:"qualityCutoffNotMet"`
		Languages           []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"languages"`
		Edition string `json:"edition"`
		ID      int    `json:"id"`
	} `json:"movieFile"`
	ID int `json:"id"`
}

type Health struct {
	Source  string `json:"source"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Wikiurl string `json:"wikiUrl"`
}

type Indexer struct {
	Enablerss               bool   `json:"enableRss"`
	Enableautomaticsearch   bool   `json:"enableAutomaticSearch"`
	Enableinteractivesearch bool   `json:"enableInteractiveSearch"`
	Supportsrss             bool   `json:"supportsRss"`
	Supportssearch          bool   `json:"supportsSearch"`
	Protocol                string `json:"protocol"`
	Priority                int    `json:"priority"`
	Name                    string `json:"name"`
	Fields                  []struct {
		Order    int    `json:"order"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		Helptext string `json:"helpText"`
		Value    string `json:"value"`
		Type     string `json:"type"`
		Advanced bool   `json:"advanced"`
	} `json:"fields"`
	Implementationname string `json:"implementationName"`
	Implementation     string `json:"implementation"`
	Configcontract     string `json:"configContract"`
	Infolink           string `json:"infoLink"`
	Tags               []struct {
	} `json:"tags"`
	ID int `json:"id"`
}

type Disk struct {
	Path       string `json:"path"`
	Label      string `json:"label"`
	Freespace  int64  `json:"freeSpace"`
	Totalspace int64  `json:"totalSpace"`
}

type UI struct {
	Firstdayofweek           int    `json:"firstDayOfWeek"`
	Calendarweekcolumnheader string `json:"calendarWeekColumnHeader"`
	Movieruntimeformat       string `json:"movieRuntimeFormat"`
	Shortdateformat          string `json:"shortDateFormat"`
	Longdateformat           string `json:"longDateFormat"`
	Timeformat               string `json:"timeFormat"`
	Showrelativedates        bool   `json:"showRelativeDates"`
	Enablecolorimpairedmode  bool   `json:"enableColorImpairedMode"`
	Movieinfolanguage        int    `json:"movieInfoLanguage"`
	ID                       int    `json:"id"`
}

type Host struct {
	Bindaddress               string `json:"bindAddress"`
	Port                      int    `json:"port"`
	Sslport                   int    `json:"sslPort"`
	Enablessl                 bool   `json:"enableSsl"`
	Launchbrowser             bool   `json:"launchBrowser"`
	Authenticationmethod      string `json:"authenticationMethod"`
	Analyticsenabled          bool   `json:"analyticsEnabled"`
	Username                  string `json:"username"`
	Password                  string `json:"password"`
	Loglevel                  string `json:"logLevel"`
	Consoleloglevel           string `json:"consoleLogLevel"`
	Branch                    string `json:"branch"`
	Apikey                    string `json:"apiKey"`
	Sslcertpath               string `json:"sslCertPath"`
	Sslcertpassword           string `json:"sslCertPassword"`
	Urlbase                   string `json:"urlBase"`
	Updateautomatically       bool   `json:"updateAutomatically"`
	Updatemechanism           string `json:"updateMechanism"`
	Updatescriptpath          string `json:"updateScriptPath"`
	Proxyenabled              bool   `json:"proxyEnabled"`
	Proxytype                 string `json:"proxyType"`
	Proxyhostname             string `json:"proxyHostname"`
	Proxyport                 int    `json:"proxyPort"`
	Proxyusername             string `json:"proxyUsername"`
	Proxypassword             string `json:"proxyPassword"`
	Proxybypassfilter         string `json:"proxyBypassFilter"`
	Proxybypasslocaladdresses bool   `json:"proxyBypassLocalAddresses"`
	Certificatevalidation     string `json:"certificateValidation"`
	Backupfolder              string `json:"backupFolder"`
	Backupinterval            int    `json:"backupInterval"`
	Backupretention           int    `json:"backupRetention"`
	ID                        int    `json:"id"`
}

type Naming struct {
	Renamemovies             bool   `json:"renameMovies"`
	Replaceillegalcharacters bool   `json:"replaceIllegalCharacters"`
	Colonreplacementformat   string `json:"colonReplacementFormat"`
	Standardmovieformat      string `json:"standardMovieFormat"`
	Moviefolderformat        string `json:"movieFolderFormat"`
	Includequality           bool   `json:"includeQuality"`
	Replacespaces            bool   `json:"replaceSpaces"`
	ID                       int    `json:"id"`
}
