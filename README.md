# Ctrl

Controller helpers for JSON webservices that use Gorilla Mux.

# Install

        $ go get github.com/johnny-morrice/ctrl

# Usage suggestion

        func (s *server) endpoint(w http.ResponseWriter, r *http.Request) {
                withctrl(w, r, s.action)
        }

        func (s *server) action(c ctrl.C) error {
                ok := doStuff()

                if !ok {
                        // Send 500 to browser.
                        return c.InternalError()
                }

                return c.ServeJson("alright!")
        }

        func withctrl(w http.ResponseWriter, r *http.Request, handler func (ctrl.C) error) {
            c := ctrl.New(w, r)

            err := handler(c)

            if err != nil {
                log.Println(err)
            }
        }
