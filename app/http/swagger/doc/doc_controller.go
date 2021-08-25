package doc

import _ "embed"

//go:embed swagger.json
var description string

//go:embed swagger.html
var html string

// @Bean
type Controller struct {
}
