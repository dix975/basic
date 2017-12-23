package pageable

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"dix975.com/basic/logger"
)

type Pageable struct {
	Page          int
	Size          int
	CurrentCount  int
	SortField     string
	SortDirection string
	values        url.Values
	url           url.URL
}

func (p *Pageable) ApplySort(query *mgo.Query) {

	var sorts []string
	if p.SortField == "" {
		return
	} else {

		if p.SortDirection == "" {
			p.SortDirection = "ASC"
		}

		if p.SortDirection == "DESC" {
			sorts = []string{fmt.Sprintf("-%v", p.SortField)}
		} else {
			sorts = []string{p.SortField}
		}

	}

	logger.Debug.Printf("Applying sort : %v ", sorts)
	query.Sort(sorts...)

}

func (p *Pageable) String() string {

	return fmt.Sprintf("Page : %v Size : %v Sort : %v %v", p.Page, p.Size, p.SortField, p.SortDirection)
}

func (p *Pageable) HasNextPage() bool {

	logger.Debug.Printf("HasNextPage current count [%d] page size [%d]\n", p.CurrentCount, p.Size)
	return p.CurrentCount >= p.Size
}

func (p *Pageable) HasPreviousPage() bool {

	return p.Page > 1
}

func (p *Pageable)NextPageURL() string {

	url := p.url
	q := url.Query()

	if p.HasNextPage() {
		q.Set("page", strconv.Itoa(p.Page + 1))
	}

	url.RawQuery = q.Encode()
	return url.String()
}

func (p *Pageable)PreviousPageURL() string {

	url := p.url
	q := url.Query()

	if p.HasPreviousPage() {

		q.Set("page", strconv.Itoa(p.Page - 1))
	}
	url.RawQuery = q.Encode()
	return url.String()
}

func (p *Pageable)SortURL(field string) string {

	url := p.url
	q := url.Query()


	q.Set("sort", field)

	if field == p.SortField{
		if p.SortDirection == "ASC" {
			q.Set("direction", "DESC")
		}else{
			q.Set("direction", "ASC")
		}
	} else {
		q.Set("direction", "ASC")
	}

	q.Set("page", "1")

	url.RawQuery = q.Encode()
	return url.String()
}

func NewPageFromHttpRequest(request *http.Request, defaultSortField string) *Pageable {

	values, err := url.ParseQuery(request.URL.RawQuery)
	if (err != nil) {
		panic(err)
	}
	logger.Debug.Printf("!!!!!!!!!!!!!!!toto!!!!!!!!!!!!!!!!!!!!!!!!!")
	logger.Info.Println("Url : ", request.URL)
	logger.Trace.Println("Query : ", request.URL.RawQuery)
	logger.Debug.Println("Values from query : ", values)

	pageable := new(Pageable)
	pageable.url = *request.URL

	pageValue := values["page"]
	sizeValue := values["size"]


	//todo default value setup suck

	if len(pageValue) < 1 {
		logger.Debug.Println("Setting default page to 1")
		values["page"] = []string{"1"}
	}
	if len(sizeValue) < 1 {
		logger.Debug.Println("Setting default page to 50")
		values["size"] = []string{"20"}
	}

	pageable.Page, err = strconv.Atoi(values["page"][0])
	if err != nil {
		panic(fmt.Errorf("Failed to parse 'page' paremeter : %v", err))
	}
	pageable.Size, err = strconv.Atoi(values["size"][0])
	if err != nil {
		panic(fmt.Errorf("Failed to parse 'size' paremeter : %v", err))
	}

	sorts := values["sort"]
	if (len(sorts) > 0) {
		pageable.SortField = sorts[0]
	}else{
		pageable.SortField = defaultSortField
		url := pageable.url
		q := url.Query()

			q.Set("sort", defaultSortField)

		pageable.url.RawQuery = q.Encode()

	}
	directions := values["direction"]
	if (len(directions) > 0) {
		pageable.SortDirection = directions[0]
	}

	logger.Debug.Printf("Sort %v %v\n", pageable.SortField, pageable.SortDirection)

	logger.Debug.Println("Page : ", pageable)

	return pageable

}