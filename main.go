package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	barWidth = 40
	minFreq  = 1000.0
	maxFreq  = 5000.0
)

func main() {
	freqs, err := getFreqs()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	m, M := getRange(freqs)
	m = math.Min(m, minFreq)
	M = math.Max(M, maxFreq)

	for _, f := range freqs {
		w := int(math.Round((f - m) / (M - m) * barWidth))
		bar := strings.Repeat(".", w) + strings.Repeat(" ", barWidth-w)
		fmt.Println(bar + "| " + strconv.FormatFloat(f, 'f', 0, 64))
	}
}

func getRange(freqs []float64) (m, M float64) {
	m = 99999.9
	for _, f := range freqs {
		if f > M {
			M = f
		}
		if f < m {
			m = f
		}
	}
	return
}

func getFreqs() (freqs []float64, err error) {
	var f *os.File
	f, err = os.Open("/proc/cpuinfo")
	defer f.Close()
	if err != nil {
		return
	}

	var buf []byte
	buf, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}

	pattern := regexp.MustCompile(`^cpu MHz\s+:\s+([0-9.]+).*$`)

	var freq float64
	lines := bytes.Split(buf, []byte{'\n'})
	for _, line := range lines {
		matches := pattern.FindStringSubmatch(string(line))
		if len(matches) == 2 {
			freq, err = strconv.ParseFloat(matches[1], 64)
			if err != nil {
				return
			}
			freqs = append(freqs, freq)
		}
	}

	return
}
