package cmd

import (
	"bytes"
	"fmt"
	"github.com/romie-gr/romie/internal/archive"
	"github.com/romie-gr/romie/internal/scraper"
	"github.com/romie-gr/romie/internal/utils"
	"github.com/romie-gr/romie/pkg/websites/emulatorgames"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install new game ROMs",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the current filepath where the binary of romie is running
		// TODO: To read the config file and use the correct PATH
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		// ---- Code Duplication with search.go ---- //
		// EmulatorGames.Net
		if !utils.FileExists(emulatorgames.DBFile) {
			log.Fatalf("There is no database for EmulatorGames.Net: %s", emulatorgames.DBFile)
		}

		jsonToEmuDB(emulatorgames.DBFile)

		var foundGames []scraper.Rom

		notFound := true

		for _, rom := range emulatorgames.Roms {
			if utils.StringContains(rom.Name, Title) {
				notFound = false
				foundGames = append(foundGames, rom)
			}
		}

		if notFound {
			log.Fatal("No games matching your title")
		}

		log.Infof("Installing %d games ...\n", len(foundGames))

		for i, game := range foundGames {
			dirPath := filepath.Join(path, game.Name)
			log.Debugf("Checking if folder exists: %s\n", dirPath)

			if utils.FolderExists(dirPath) {
				log.Errorf("%s is already installed. Skip downloading ...\n", game.Name)
				continue
			}

			log.Debugf("Folder doesn't exist. Creating it now!\n")
			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				log.Errorf("couldn't create %s directory to save the game. %q\n", dirPath, err)
				continue
			}

			// If CTRL+C is pressed, handle this with grace
			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c
				cleanup(dirPath)
				os.Exit(1)
			}()

			filePath := filepath.Join(dirPath, filepath.Base(game.DownloadLink))
			log.Debugf("The game will be saved into: \"%s\"", filePath)

			if err := downloadFile(game.Name, game.DownloadLink, dirPath, i+1, len(foundGames)); err != nil {
				log.Errorf("Failed to download \"%s\" game!\n", game.Name)
				log.Errorf("Error: %q\n", err)
				os.RemoveAll(dirPath)
				continue
			}

			// Extract
			log.Infof("Extracting the compressed archive ...")
			if err := archive.Extract(filePath); err != nil {
				log.Errorf("Failed to extract the compressed archive %s. Error: %q\n", filePath, err)
				os.RemoveAll(dirPath)
				continue
			}
			log.Infof("done!\n\n")
		}

	},
}

func init() {
	installCmd.Flags().StringVarP(&Title, "title", "t", "", "Title of the game you want to install")
	_ = installCmd.MarkFlagRequired("title")
	rootCmd.AddCommand(installCmd)
}

func cleanup(dirPath string) {
	// In case the user types 'CTRL+C' during installing (downloading, extracting, etc)
	log.Debugf("\nCTRL+C signal detected!!! Cleaning up ...\n")
	os.RemoveAll(dirPath)
}

func printDownloadPercent(done chan int64, savePath string, totalFilesize int64) {
	var stop = false
	var totalMB int64

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(savePath)
			if err != nil {
				log.Error(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Error(err)
			}

			currentDownloadedSize := fi.Size()

			if currentDownloadedSize == 0 {
				currentDownloadedSize = 1
			}

			var percent = float64(currentDownloadedSize) / float64(totalFilesize) * 100
			currentMB := currentDownloadedSize / 1024 / 1024
			totalMB = totalFilesize / 1024 / 1024

			fmt.Printf("\rDownload progress: %d%% - %d/%d MB", int64(percent), currentMB, totalMB)
		}


		if stop {
			break
		}

		time.Sleep(time.Second)
	}
	fmt.Printf("\rDownload progress: 100%% - %d/%d MB\n", totalMB, totalMB)
}

func downloadFile(gameName, url, dest string, current, total int) error {

	file := path.Base(url)

	log.Printf("[%d/%d] - Downloading: \"%s\"\tfrom\t\"%s\"\n", current, total, gameName, url)

	var path bytes.Buffer
	path.WriteString(dest)
	path.WriteString("/")
	path.WriteString(file)

	start := time.Now()

	out, err := os.Create(path.String())

	if err != nil {
		log.Debugf("Failed at: 'os.Create(path.String()'\n")
		return err
	}

	defer out.Close()

	headResp, err := http.Head(url)

	if err != nil {
		log.Debugf("Failed at: 'http.Head(url)'\n")
		return err
	}

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))

	if err != nil {
		log.Debugf("Failed at: 'strconv.Atoi(headResp.Header.Get(\"Content-Length\"))'\n")
		return err
	}

	done := make(chan int64)

	go printDownloadPercent(done, path.String(), int64(size))

	resp, err := http.Get(url)

	if err != nil {
		log.Debugf("Failed at: ''http.Get(url)\n")
		return err
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		log.Debugf("Failed at: 'io.Copy(out, resp.Body)'")
		return err
	}

	done <- n

	elapsed := time.Since(start)
	log.Infof("Download completed in %s\n", elapsed)
	return nil
}
