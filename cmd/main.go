package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/Crystalix007/shhtime/formats"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"
)

var timeFormatRegexp = regexp.MustCompile(
	`^[0-9 T:-_]*$`,
)

func main() {
	srv, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					format := time.Kitchen

					if predefinedFormat, err := formats.FormatFromSnakeCase(s.User()); err == nil {
						format = string(predefinedFormat)
					} else if timeFormatRegexp.MatchString(s.User()) {
						format = s.User()
					}

					time := time.Now().Format(format)

					fmt.Fprintf(s, "%s\n", time)

					next(s)
				}
			},
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		panic(err)
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		panic(err)
	}
}
