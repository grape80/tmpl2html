package tmpl2html

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	OPT_BASEDIR  = "basedir"
	OPT_MAKEDEPS = "makedeps"
	OPT_VERBOSE  = "verbose"
)

const (
	OPT_BASEDIR_DEFAULT  = "./"
	OPT_MAKEDEPS_DEFAULT = false
	OPT_VERBOSE_DEFAULT  = false
)

type converter struct {
	stdout   io.Writer
	stderr   io.Writer
	makedeps bool
	verbose  bool
	basedir  string
	gotmpl   string
	partials []string
}

func (c *converter) run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s %v\n", ERR_RUNTIME, r)
		}
	}()

	if err = c.parseTemplate(); err != nil {
		return err
	}

	if err = c.execute(); err != nil {
		return err
	}

	if c.makedeps {
		return c.createMakeDeps()
	}

	return err
}

const (
	partialPrefix = `{{template "`
	partialSuffix = `}}`
)

func (c *converter) parseTemplate() (err error) {

	f, err := os.Open(gotmpl)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text()) // to clean
		if strings.HasPrefix(l, partialPrefix) && strings.HasSuffix(l, partialSuffix) {
			p := filepath.Join(c.basedir, strings.Split(l, `"`)[1]) // prefix”path"suffix -> [prefix path suffix]
			c.partials = append(c.partials, p)
		}
	}

	return err
}

func (c *converter) execute() (err error) {

	files := []string{}
	files = append(files, c.gotmpl)
	files = append(files, c.partials...)

	t := template.Must(template.ParseFiles(files...))

	var buf bytes.Buffer
	if err = t.Execute(&buf, nil); err != nil {
		return err
	}

	_, err = fmt.Fprintf(c.stdout, "%s", &buf)

	return err
}

const (
	depsPerm fs.FileMode = 0644
)

func (c *converter) createMakeDeps() (err error) {
	var buf bytes.Buffer

	buf.WriteString(c.gotmpl + ":")
	for _, p := range c.partials {
		buf.WriteString(" " + p)
	}
	buf.WriteString("\n")

	return os.WriteFile(c.gotmpl+".deps", buf.Bytes(), depsPerm)
}
