package entities

type Config struct {
	Port      int       `json:"port"`
	MongoDB   MongoDB   `json:"mongodb"`
	SqlServer SqlServer `json:"sqlserver"`
}

type SqlServer struct {
    User    string `json:"user"`
    Pass    string `json:"pass"`
    DB      string `json:"db"`
    Options string `json:"options"`
}

type MongoDB struct {
    Url string `json:"url"`
    DB  string `json:"db"`
}