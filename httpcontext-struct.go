//https://gist.github.com/elithrar/5aef354a54ba71a32e23
//https://elithrar.github.io/article/custom-handlers-avoiding-globals
package main

import (
    "fmt"
    "log"
    "net/http"

    "html/template"

    "github.com/gorilla/sessions"
    "github.com/jmoiron/sqlx"
    "github.com/zenazn/goji/graceful"
    "github.com/zenazn/goji/web"
)

// appContext contains our local context; our database pool, session store, template
// registry and anything else our handlers need to access. We'll create an instance of it
// in our main() function and then explicitly pass a reference to it for our handlers to access.
type appContext struct {
    db        *sqlx.DB
    store     *sessions.CookieStore
    templates map[string]*template.Template
    decoder *schema.Decoder
    // ... and the rest of our globals.
}

// We've turned our original appHandler into a struct with two fields:
// - A function type similar to our original handler type (but that now takes an *appContext)
// - An embedded field of type *appContext
type appHandler struct {
    *appContext
    h func(*appContext, http.ResponseWriter, *http.Request) (int, error)
}

// Our ServeHTTP method is mostly the same, and also has the ability to
// access our *appContext's fields (templates, loggers, etc.) as well.
func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Updated to pass ah.appContext as a parameter to our handler type.
    status, err := ah.h(ah.appContext, w, r)
    if err != nil {
        log.Printf("HTTP %d: %q", status, err)
        switch status {
        case http.StatusNotFound:
            http.NotFound(w, r)
            // And if we wanted a friendlier error page:
            // err := ah.renderTemplate(w, "http_404.tmpl", nil)
        case http.StatusInternalServerError:
            http.Error(w, http.StatusText(status), status)
        default:
            http.Error(w, http.StatusText(status), status)
        }
    }
}

func main() {
    // These are 'nil' for our example, but we'd either assign
    // the values as below or use a constructor function like
    // (NewAppContext(conf config) *appContext) that initialises
    // it for us based on our application's configuration file.
    context := &appContext{db: nil, store: nil} // Simplified for this example

    r := web.New()
    // We pass an instance to our context pointer, and our handler.
    r.Get("/", appHandler{context, IndexHandler})

    graceful.ListenAndServe(":8000", r)
}

func IndexHandler(a *appContext, w http.ResponseWriter, r *http.Request) (int, error) {
    // Our handlers now have access to the members of our context struct.
    // e.g. we can call methods on our DB type via err := a.db.GetPosts()
    fmt.Fprintf(w, "IndexHandler: db is %q and store is %q", a.db, a.store)
    return 200, nil
}
