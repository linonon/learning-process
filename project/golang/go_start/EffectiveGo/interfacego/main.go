package main

import (
	"fmt"
	"gostart/EffectiveGo/interfacego/mock"
	"gostart/EffectiveGo/interfacego/real"
	"time"
)

const url = "http://www.google.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.google.com")
}

func post(poster Poster) {
	url := "http://www.google.com"
	poster.Post(url, map[string]string{
		"name":     "linonon",
		"learning": "go",
	})
}

// 組合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Printf("%T, %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this fake."}
	r = &retriever
	inspect(r)
	// fmt.Println(download(r))

	r = &real.Retriever{
		UserAgent: "linonon",
		Timeout:   time.Minute,
	}
	inspect(r)
	// fmt.Println(download(r))

	// Type assertion
	// 通過這種方法，判斷獲得不同的內容。
	realRetriever := r.(*real.Retriever)
	fmt.Println("realRetriever.Timeout: ", realRetriever.Timeout)

	fmt.Println("Try assesion")
	fmt.Println(session(&retriever))
}
