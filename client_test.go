package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//type Rows struct {
//	Rows []UserXML `xml:"row"`
//}
//
//type UserXML struct {
//	Id        int    `xml:"id"`
//	Age       int    `xml:"age"`
//	FirstName string `xml:"first_name" json:"-"`
//	Name      string `xml:"-"`
//	LastName  string `xml:"last_name" json:"-"`
//	About     string `xml:"about"`
//	Gender    string `xml:"gender"`
//}
//
//type ErroReq struct {
//	Error string
//}
//
//func ChoosenSortFunc(orderField string, orderBy int) func(a, b UserXML) int {
//	switch orderField {
//	case "Name", "":
//		return func(a, b UserXML) int {
//			if n := strings.Compare(a.FirstName, b.FirstName); n != 0 {
//				return n * orderBy
//			} else if f := strings.Compare(a.LastName, b.LastName); f != 0 {
//				return f * orderBy
//			}
//			return cmp.Compare(a.Id, b.Id) * orderBy
//		}
//	case "Id":
//		return func(a, b UserXML) int {
//			return cmp.Compare(a.Id, b.Id) * orderBy
//		}
//	case "Age":
//		return func(a, b UserXML) int {
//			return cmp.Compare(a.Age, b.Age) * orderBy
//		}
//	default:
//		return nil
//	}
//}
//
//func SearchHandler(w http.ResponseWriter, r *http.Request) {
//	queryParams := r.URL.Query()
//	//respRow := &Rows{Error: ""}
//	w.Header().Set("Content-Type", "application/json")
//	enc := json.NewEncoder(w)
//	if len(queryParams) == 0 {
//		w.WriteHeader(http.StatusBadRequest)
//		errResp := ErroReq{"ErrorBadQueryParams"}
//		err := enc.Encode(errResp)
//		if err != nil {
//			return
//		}
//		return
//	}
//	queryRaw := queryParams.Get("query")
//	orderField := queryParams.Get("order_field")
//	if orderField != "Name" && orderField != "Id" && orderField != "Age" && orderField != "" {
//		w.WriteHeader(http.StatusBadRequest)
//		errResp := ErroReq{"ErrorBadOrderField"}
//		err := enc.Encode(errResp)
//		if err != nil {
//			return
//		}
//		return
//	}
//	orderBy := queryParams.Get("order_by")
//	order, err := strconv.Atoi(orderBy)
//	if err != nil || order > 1 || order < -1 {
//		w.WriteHeader(http.StatusBadRequest)
//		errResp := ErroReq{"ErrorBadOrderBy"}
//		err := enc.Encode(errResp)
//		if err != nil {
//			return
//		}
//		return
//	}
//	limitStr := queryParams.Get("limit")
//	limit, err := strconv.Atoi(limitStr)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		errResp := ErroReq{"ErrorBadLimit"}
//		err := enc.Encode(errResp)
//		if err != nil {
//			return
//		}
//		return
//	}
//	offsetStr := queryParams.Get("offset")
//	offset, err := strconv.Atoi(offsetStr)
//	if err != nil || offset < 0 {
//		w.WriteHeader(http.StatusBadRequest)
//		errResp := ErroReq{"ErrorBadOffset"}
//		err := enc.Encode(errResp)
//		if err != nil {
//			return
//		}
//		return
//	}
//	rows := &Rows{}
//	xmlfile, err := os.Open("dataset.xml")
//	defer func(xmlfile *os.File) {
//		err := xmlfile.Close()
//		if err != nil {
//			log.Print("Error close xml file")
//			return
//		}
//	}(xmlfile)
//	if err != nil {
//		log.Print("Error open file")
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	data, err := io.ReadAll(xmlfile)
//	if err != nil {
//		log.Print("Error read file")
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	err = xml.Unmarshal(data, rows)
//	if err != nil {
//		log.Print("Error unmarshal data")
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	//_, err = w.Write([]byte(fmt.Sprintf("%#v", rows.Rows)))
//
//	if order != 0 {
//		slices.SortFunc(rows.Rows, ChoosenSortFunc(orderField, order))
//	}
//
//	left := offset
//	right := len(rows.Rows)
//	if limit+offset < len(rows.Rows) {
//		right = limit + offset
//	}
//	w.Write([]byte("["))
//	//testSlice := make([]UserXML, limit)
//	first := true
//	for i := left; i < right; i++ {
//		rows.Rows[i].Name = rows.Rows[i].FirstName + " " + rows.Rows[i].LastName
//		if strings.Contains(rows.Rows[i].Name, queryRaw) || strings.Contains(rows.Rows[i].About, queryRaw) {
//			//w.Write([]byte(fmt.Sprintf("%#v", rows.Rows[i])))
//			if !first {
//				w.Write([]byte(","))
//			}
//			first = false
//			err := enc.Encode(rows.Rows[i])
//			if err != nil {
//				return
//			}
//			//testSlice = append(testSlice, rows.Rows[i])
//		}
//	}
//	w.Write([]byte("]"))
//	//respRow.Rows = testSlice
//	////w.Write([]byte(fmt.Sprintf("%#v", respRow)))
//	//marshall, err := json.Marshal(respRow)
//	//if err != nil {
//	//	log.Print("Error marshall file")
//	//	w.WriteHeader(http.StatusInternalServerError)
//	//	return
//	//}
//	//fmt.Printf("json string:\n\t%s\n", string(marshall))
//	//w.Write(marshall)
//}
//
////func SearchUser() (*[]User, error){
////
////	response, err := http.Get("http://localhost:8080/?order_field=Id&order_by=1&limit=2&offset=1")
////	if err != nil {
////		return nil, errors.New("can't get response from server")
////	}
////	defer response.Body.Close()
////	body, err := io.ReadAll(response.Body)
////	if err != nil {
////		return nil,errors.New("can't read response body")
////	}
////	data := &[]User{}
////	err = json.Unmarshal(body, &data)
////	if err != nil {
////		return nil,errors.New("unmarshall issue")
////	}
////	return data, nil
////

type TestCase struct {
	queryParams *SearchRequest
	Result      *SearchResponse
	IsErr       bool
}

const ADDR = "http://localhost:8080"

func TestSearchServer(t *testing.T) {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", SearchHandler)
	//server := http.Server{Handler: mux, Addr: ":8080"}
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}
	testCases := []TestCase{
		{queryParams: &SearchRequest{
			Limit:      2,
			Offset:     1,
			OrderField: "Id",
			OrderBy:    OrderByDesc},
			Result: &SearchResponse{Users: []User{{
				Id:     1,
				Name:   "Hilda Mayer",
				Age:    21,
				About:  "Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.\n",
				Gender: "female",
			}, {
				Id:     2,
				Name:   "Brooks Aguilar",
				Age:    25,
				About:  "Velit ullamco est aliqua voluptate nisi do. Voluptate magna anim qui cillum aliqua sint veniam reprehenderit consectetur enim. Laborum dolore ut eiusmod ipsum ad anim est do tempor culpa ad do tempor. Nulla id aliqua dolore dolore adipisicing.\n",
				Gender: "male",
			}}, NextPage: true},
			IsErr: false},
	}
	ts := httptest.NewServer(http.HandlerFunc(SearchHandler))
	client := SearchClient{URL: ADDR, AccessToken: "12345"}

	for caseNum, testCase := range testCases {
		req := SearchRequest{
			Limit:      testCase.queryParams.Limit,
			Offset:     testCase.queryParams.Offset,
			Query:      testCase.queryParams.Query,
			OrderField: testCase.queryParams.OrderField,
			OrderBy:    testCase.queryParams.OrderBy,
		}
		users, err := client.FindUsers(req)
		if err != nil {
			t.Errorf("Error in find users: %#v", err)
		}
		if !assert.Equal(t, testCase.Result, users) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", caseNum, testCase.Result, users.Users)
		}
	}

	ts.Close()
}
