package hello

import (
    "fmt"
    "net/http"
    "appengine"
    "appengine/user"
)

func init() {
    http.HandleFunc("/", handler)
}

func root(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, guestbookForm)
}

const guestbookForm = `
<html>
  <body>
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`


func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
      url, err := user.LoginURL(c, r.URL.String())
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      w.Header().Set("Location", url)
      w.WriteHeader(http.StatusFound)
      return
    }
    fmt.Fprintf(w, "Hello, %v!", u)
}
