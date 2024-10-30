package example

import (
	"encoding/json" //JSON 객체를 인코드/디코드
	"errors"        //오류를 생성하는 데 사용
	"fmt"
	"io/ioutil" //입력/출력 유틸리티(REST API에서 데이터 스트림을 읽는 데 사용)
	"net/http"  //HTTP 클라이언트 및 서버 구현(클라이언트 기능만 사용)
	"net/url"   //URL을 처리하는 유틸리티
	//"strings"   //문자열 유틸리티 구현
)

// omdbapi.com API key
const APIKEY = "193ef3a"

// omdbapi.com에서 반환된 JSON 구조체
// 예시의 간략화를 위해 일부 값은 구조체에 매핑하지 않음
type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"ImdbRating"`
	ImdbID     string `json:"ImdbID"`
}

func sendGetRequest(url string) (string, error) {
	// 함수 인자를 통해 전달된 url을 http 모듈의 Get 함수에 전달하며, GET 요청을 수행
	// 응답을 resp 및 err에 저장하고 err가 nil인지 확인
	// 이때 만약 err가 nil이 아니라면, 명시된 오류와 함께 빈 응답 반환
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// io패키지를 이용하여, 응답 분문으로부터 바이트를 읽어 들인다.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 상태코드를 확인하여 200이 아닐 경우, 전체 상태를 설명해는 오류 생성
	// 오류와 응답의 본문 함께 반환
	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}
	// 파이프라인에 오류가 없으므로 본문을 문자열로 아무 문제없이 반환할 수 있다.
	return string(body), nil
}

func SearchByName(name string) (*MovieInfo, error) {
	//쿼리 파라미터 설정
	parms := url.Values{}       //API 쿼리 파라미터 설정
	parms.Set("apikey", APIKEY) // API 키를 쿼리 문자열에 추가
	parms.Set("t", name)        // 검색할 영화 제목 설정

	// API 요청 URL 생성
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()

	// API 호출 및 에러 처리
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}

	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func SearchById(id string) (*MovieInfo, error) {
	parms := url.Values{}
	parms.Set("apikey", APIKEY)
	parms.Set("i", id)

	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	body, err := sendGetRequest(siteURL)

	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func Code300() {
	body, _ := SearchById("tt3896198")
	fmt.Println(body.Title)
	body, _ = SearchByName("Game of")
	fmt.Println(body.Title)
}
