package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"github.com/seqyuan/annogene/io/fastq"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func usage() {
	fmt.Printf("    -inFQ          faseq.gz\n")
	fmt.Printf("    -o             outdir\n")
	os.Exit(1)
}

func main() {
	infq := flag.String("inFQ", "", "fastq.gz")
	outdir := flag.String("o", "", "outdir")
	flag.Parse()
	if *infq == "" || *outdir == "" {
		usage()
	}

	file, err := os.Open(*infq)
	check(err)
	gz, err := gzip.NewReader(file)
	check(err)

	defer file.Close()
	defer gz.Close()

	r := fastq.NewReader(gz)
	sc := fastq.NewScanner(r)

	outfqgz := fmt.Sprintf("%s/%s_clean.fq.gz", *outdir, filepath.Base(*infq))

	fo, err := os.OpenFile(outfqgz, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	check(err)
	ogz := gzip.NewWriter(fo)
	check(err)

	defer fo.Close()
	defer ogz.Close()

	w := fastq.NewWriter(ogz)

	i := 0
	for sc.Next() {
		i += 1
		c := i % 2
		if c == 0{
			_, eer := w.Write(sc.Seq())
			check(eer)
		}

	}
	
}

