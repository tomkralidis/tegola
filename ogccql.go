package tegola

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

func (c *CQLFilter) ToSQL() error {
    return nil
}
