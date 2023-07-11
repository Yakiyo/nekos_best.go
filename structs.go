package nekos_best

// All image categories of nekos.best
var image_categories = []string{"kitsune", "neko", "husbando", "waifu"}

// All gif categories of nekos.best
var gif_categories = []string{
	"baka",
	"bite",
	"blush",
	"bored",
	"cry",
	"cuddle",
	"dance",
	"facepalm",
	"feed",
	"happy",
	"highfive",
	"hug",
	"kiss",
	"laugh",
	"pat",
	"pout",
	"shrug",
	"slap",
	"sleep",
	"smile",
	"smug",
	"stare",
	"think",
	"thumbsup",
	"tickle",
	"wave",
	"wink",
	"kick",
	"handhold",
	"punch",
	"shoot",
	"yeet",
	"poke",
	"nod",
	"nom",
	"nope",
	"handshake",
	"lurk",
	"nibble",
	"peck",
	"yawn",
}

var categories = append(gif_categories, image_categories...)

// A nekos.best response struct
//
// Url is always present. For gif endpoints, Source_url and Anime_name exists,
// and for image endpoints Artist_href and Artist_name exists
type NBResponse struct {
	Url         string
	Artist_href string
	Artist_name string
	Source_url  string
	Anime_name  string
}

// A nekos.best buffer response
//
// Used to represent result of an image along with its associated `NBResponse`
type NBBufferResponse struct {
	Data []byte
	Body NBResponse
}

// Result returned from nekos.best api responses
type fullNBResponse struct {
	Results []NBResponse
}
