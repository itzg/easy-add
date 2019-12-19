package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"github.com/itzg/go-flagsfiller"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	version string
	commit  string
)

var args struct {
	From    string `usage:"[URL] of a tar.gz archive to download"`
	File    string `usage:"The [path] to executable to extract within archive"`
	To      string `usage:"The [path] where executable will be placed" default:"/usr/local/bin"`
	Mkdirs  bool   `usage:"Attempt to create the directory path specified by to"`
	Version bool   `usage:"Show version and exit"`
}

func main() {

	err := flagsfiller.Parse(&args)
	if err != nil {
		log.Fatal(err)
	}

	if args.Version {
		fmt.Printf("version=%s, commit=%s\n", version, commit)
		return
	}

	if args.From == "" || args.File == "" {
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "from and file are required")
		flag.Usage()
		os.Exit(2)
	}

	log.SetOutput(os.Stdout)

	if !isTarGz(args.From) {
		log.Fatal("Only supports processing tar-gzipped files with tar.gz or tgz suffix")
	}

	if args.Mkdirs {
		err := os.MkdirAll(args.To, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("I! Retrieving %s", args.From)
	client, err := setupHttpClient()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get(args.From)
	if err != nil {
		log.Fatal(err)
	}
	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		outFilePath, err := processTarGz(resp.Body, args.File, args.To)
		if err != nil {
			log.Fatalf("E! %v", err)
		}
		log.Printf("I! Extracted file to %s", outFilePath)
	} else {
		log.Fatalf("E! Failed to retrieve archive: %s", resp.Status)
	}
}

func setupHttpClient() (*http.Client, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		log.Printf("W! %v", err)
		certPool = x509.NewCertPool()
	}
	for _, pem := range extraCerts {
		if !certPool.AppendCertsFromPEM([]byte(pem)) {
			return nil, errors.New("Unable to add Github CA cert")
		}
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: certPool},
		},
	}

	return client, nil
}

func processTarGz(reader io.Reader, file string, to string) (string, error) {

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return "", fmt.Errorf("failed to read gzip content: %w", err)
	}

	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			return "", errors.New("unable to find requested file in archive")
		}

		if header.Name == file {
			return extractExe(tarReader, file, to, header.FileInfo())
		}
	}
}

func extractExe(reader io.Reader, filename string, to string, fileInfo os.FileInfo) (string, error) {
	outPath := path.Join(to, path.Base(filename))

	file, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, fileInfo.Mode())
	if err != nil {
		return "", fmt.Errorf("unable to create destination file: %w", err)
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		return "", fmt.Errorf("unable to copy extracted file content: %w", err)
	}

	return outPath, nil
}

func isTarGz(url string) bool {
	url = strings.ToLower(url)
	return strings.HasSuffix(url, ".tar.gz") || strings.HasSuffix(url, ".tgz")
}
