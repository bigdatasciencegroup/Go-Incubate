package main

import (
	"net/http"
	// _ "github.com/pion/webrtc/v2"
)

func main() {

	// Everything below is the pion-WebRTC API, thanks for using it ❤️.
	// Create a MediaEngine object to configure the supported codec
	// m := webrtc.MediaEngine{}

	// if err := prometheus.Register(prommod.NewCollector("sfu-ws")); err != nil {
	// 	panic(err)
	// }

	// port := flag.String("p", "8443", "https port")
	// flag.Parse()

	// Websocket handle func
	// http.HandleFunc("/ws", room)

	// // Html handle func
	// http.HandleFunc("/", web)

	http.Handle("/", http.StripPrefix("", http.FileServer(http.Dir("./static"))))
	// fmt.Println("Web listening :" + *port)
	http.ListenAndServe(":8888", nil)

	// Support https, so we can test by lan
	// fmt.Println("Web listening :" + *port)
	// panic(http.ListenAndServeTLS(":"+*port, "cert.pem", "key.pem", nil))

}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 		return
// 	}
// 	renderTemplate(w, "view", p)
// }
