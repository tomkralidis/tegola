package tegola

import (
    "fmt"
    "log"
    "regexp"
)

type CQLFilter struct {
    Lang string
    FilterText string
}

func NewCQLFilter(filter_lang string, filter_text string) (*CQLFilter, error) {
    cf := CQLFilter{}
    cf.Lang = filter_lang
    cf.FilterText = filter_text

    return &cf, nil
}

func (c *CQLFilter) Parse() error {
    return nil
}

func (c *CQLFilter) ToSQL(sql_statement string) (string, error) {
    var filter_text string
    var value string

    log.Println("Adjusting spatial predicates")
    re := regexp.MustCompile("(?i)(beyond|contains|crosses|disjoint|dwithin|equals|intersects|overlaps|touches)")
	filter_text = re.ReplaceAllString(c.FilterText, "st_$1")

    log.Println("Adjusting geometry representation")
	re2 := regexp.MustCompile(`(?i)(POINT\s?\(.*\d\)|LINESTRING\s?\(.*\d\)|POLYGON\s?\(\(.*\d\)\))`)
 	filter_text = re2.ReplaceAllString(filter_text, "'$1'::geometry")

    log.Println("Assembling final SQL statement")

    matched, err := regexp.MatchString("(?i) WHERE ", sql_statement)

    if err != nil {
        fmt.Errorf("error matching SQL statement")
    }

    if matched {
        log.Println("appending to existing where clause")
        value = sql_statement + " and " + filter_text
    } else {
        log.Println("defining where clause")
        value = sql_statement + " where " + filter_text
    }
    return value, nil
}
