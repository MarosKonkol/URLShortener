 I've createad simple URL shortener. For the code itself I used these libraries: 

 import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/zpnk/go-bitly"
)

The last one is bitly library, which is simple bitly API client for shortening URLs.
I used functiion:
func (client *Links) Shorten(longURL string) to shorten my submited URL.

I have also created simple form page with where user can submit their URL they want to shorten.

Unfortunately, I wasn't able to get the shortened URL to be printed on website so it is only valiable via terminal.

Server runs on localhost with port :8080.

on localhost:8080/index is a form page, where user submits their URL, it is then posted to the server, server then executes the **shorten** function and respond to the request to the terminal, where user can see long url, hashes and short url.

Server also prints submited URL to the page.
