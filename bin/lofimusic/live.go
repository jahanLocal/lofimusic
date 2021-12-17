package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveRadio struct {
	Slug    string
	Name    string
	Owner   string
	URL     string
	gifURL string
	realGif bool //NOTE - this bool is not yet used because of the naive implementation
	Cards   []string
	Links   []socialLink
	AddedAt time.Time
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		{
		    Slug: "metaboombox",
		    Name: "Metaverse Boom",
		    Owner: "Wu Tang Clan",
		    URL: "https://youtu.be/n4pr7j-kTO0",
		    gifURL: "https://cdn-animation.artstation.com/p/video_sources/000/157/318/video-final0001-0240.mp4",
		    Cards: []string{
                "This is the boombox that rocks the Metaverse",
            	"Discovered high on the slopes of Mt. Fuji by the Wu Tang Clan",
            	"It only rocks for the worthy of spirit.  But it always rocks for them.",
            },
            Links: []socialLink{
                {
                    Slug: "discord",
                    URL:  "https://discord.com/invite/lofigirl",
                },
                {
                    Slug: "NFT",
                    URL: "https://opensea.io/assets/0x495f947276749ce646f68ac8c248420045cb7b5e/107348447574672300898797177560897166546098074247836768685144822490415207809025",
                },
            },
		},
		{
			Slug:  "lofigirl",
			Name:  "Lofi Girl",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/5qap5aO4i9A",
			gifURL: "https://c.tenor.com/dFRmjNDK9TwAAAPo/ghibli.mp4",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is shown studying in Lyon, a city from France where her designer Juan Pablo used to live.",
				"The view through the window depicts the buildings on the slopes of Croix-Rousse, where the bell tower of the Bon-Pasteur church can be spotted.",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://lofigirl.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5qap5aO4i9A",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/lofigirl",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/LofiGirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/lofigirl",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/lofigirl",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/lofigirl",
				},
			},
		},
		{
			Slug:  "chillhop",
			Name:  "Chillhop Racoon",
			Owner: "Chillhop Music",
			URL:   "https://youtu.be/5yx6BWlEVcY",
			gifURL: "https://d.furaffinity.net/art/leto/1633629574/1633629574.leto_chillhop_-_guitar_compressed.gif",
			realGif: false,
			Cards: []string{
			    "Chill out and get me some garbage to eat...",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d?si=3co_7Q6XT0OJZwkBlqWoDQ",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:  "collegemusic",
			Name:  "College Girl",
			Owner: "College Music",
			URL:   "https://youtu.be/MCkTebktHVc",
			gifURL: "https://cdn.dribbble.com/users/330915/screenshots/4840771/2.gif",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/MCkTebktHVc",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/collegemusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/collegemusic",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/collegemusicyt",
				},
			},
		},
		{
			Slug:  "duashouserelax",
			Name:  "Lady Gaga's House Mix",
			Owner: "Lady Gaga",
			URL:   "https://youtu.be/CHmYlCzyZdE",
			gifURL: "https://mir-s3-cdn-cf.behance.net/project_modules/1400_opt_1/5e385214058787.57f85763bc159.gif",
			realGif: true,
			Cards: []string{
			    "Uplifting mix by Lady Gaga",
			    "for VIPs Only",
			    "Should we let you get in?",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/bM0Iw7PPoU4",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/collegemusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/collegemusic",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/collegemusicyt",
				},
			},
		},
		{
		    Slug: "jack-harlow-channel",
		    Name: "Harlow Hour",
		    Owner: "Jack Harlow",
		    URL: "https://www.youtu.be/05689ErDUdM",
		    gifURL: "https://audibletreats.com/wp-content/uploads/2020/06/2db0c64a-8c6b-4cda-b127-5e2b1ccde1e8-1.gif",
		    realGif: true,
		    Cards: []string{
                "All Glorious Hip Hop All the Time",
            },
            Links: []socialLink{
                {
                    Slug: "discord",
                    URL:  "https://discord.com/invite/lofigirl",
                },
                {
                    Slug: "NFT",
                    URL: "https://opensea.io/assets/0x495f947276749ce646f68ac8c248420045cb7b5e/107348447574672300898797177560897166546098074247836768685144822490415207809025",
                },
            },
		},
		{
			Slug:  "lofi-code-beats",
			Name:  "Coding Beats",
			Owner: "Joma Tech",
			URL:   "https://youtu.be/PY8f1Z3nARo",
			gifURL: "https://cdn.dribbble.com/users/720825/screenshots/3253310/media/7988b21c14bfe59f6b0edb72fccfac28.gif",
			realGif: true,
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/PY8f1Z3nARo",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/jomaoppa",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/jomaoppa",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/jomaoppa",
				},
			},
		},
		{
			Slug:  "steezyasfuck-coffee-show",
			Name:  "Steezy Coffee Shop",
			Owner: "STEEZYASFUCK",
			URL:   "https://youtu.be/-5KAN9_CzSA",
			gifURL: "https://cdn.dribbble.com/users/472824/screenshots/15693827/media/40bd1b5274ad5d38593e00317fc83abd.mp4",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.stzzzy.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/-5KAN9_CzSA",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/2s9R059mmdc8kz6lrUqZZd",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/stzzyasfvck/",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/stzzyasfvck",
				},
			},
		},
		{
			Slug:  "legends",
			Name:  "Legends of Rock",
			Owner: "Freddy Mercury",
			URL:   "https://youtu.be/AWw0mJp4B8s",
			gifURL:"https://cdn.dribbble.com/users/1921422/screenshots/5511883/media/3bcad7440135a8609960f46fb4167851.gif",
			realGif: true,
			Cards: []string{
			    "Is this real life...",
			    "Is this just fantasy?",
			    "Caught in a landslide...",
			},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/o33l32ZrIy8",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/artist/1LwjR2mIm78OJRTYdkMLl3",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/closedonsunday",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/closedonsundayy",
				},
			},
		},
		{
			Slug:  "r&bsoul",
			Name:  "R&B (90s-2000s)",
			Owner: "United Masters",
			URL:   "https://youtu.be/xbVO6tuDfcM",
			gifURL: "https://cdn.dribbble.com/users/2936641/screenshots/10481284/media/666b64490c5eb532c3c5200a3951a7b2.gif",
			realGif: true,
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://dreamhopmusic.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/tCs48OFv7xA",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/user/91jfqzlv0htrqrvvc60138qma",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/FxF9kng",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/dreamhopp",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Dreamhopp",
				},
			},
		},
	}

	sort.Slice(radios, func(a, b int) bool {
		return strings.Compare(radios[a].Name, radios[b].Name) < 0
	})

	for _, r := range radios {
		sort.Slice(r.Links, func(a, b int) bool {
			return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
		})
	}

	return radios
}

type socialLink struct {
	Slug string
	URL  string
}

func socialIcon(slug string) string {
	switch slug {
	case "NFT":
	    return cryptoSVG

	case "youtube":
		return youtubeSVG

	case "reddit":
		return redditSVG

	case "facebook":
		return facebookSVG

	case "instagram":
		return instagramSVG

	case "twitter":
		return twitterSVG

	case "spotify":
		return spotifySVG

	case "discord":
		return discordSVG

	case "website":
		return websiteSVG

	default:
		return linkSVG
	}
}
