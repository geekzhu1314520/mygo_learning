package util

type CommonItem struct {
	Type  int32       `json:"type"`
	Title string      `json:"title,omitempty"`
	Avatar string     `json:"avatar"`
	Route string      `json:"route,omitempty"`
	RouteText string `json:"route_text"`
	Info  interface{} `json:"info,omitempty"`
}

func GenCommonItemInner(Type int32, title string, route string, router_text string, info interface{}) *CommonItem {
	return &CommonItem{
		Type:  Type,
		Title: title,
		Route: route,
		RouteText: router_text,
		Info:  info,
	}
}

func GenCommonItem(Type int32, title string, route string, info interface{}) *CommonItem {
	return GenCommonItemInner(Type, title, route, "", info)
}
