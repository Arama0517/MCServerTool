/*
 * Minecraft Server Tool(MCST) is a command-line utility making Minecraft server creation quick and easy for beginners.
 * Copyright (c) 2024-2024 Arama.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package lib_test

import (
	"encoding/json"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Arama-Vanarana/MCServerTool/pkg/lib"
	goversion "github.com/caarlos0/go-version"
)

var URL = url.URL{ // https://ash-speed.hetzner.com/100MB.bin
	Scheme: "https",
	Host:   "ash-speed.hetzner.com",
	Path:   "/100MB.bin",
}

func TestDownload(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过下载")
	}
	if err := lib.Init(goversion.GetVersionInfo()); err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	lib.ConfigsPath = filepath.Join(cwd, "test.json")
	if file, err := os.Create(lib.ConfigsPath); err != nil {
		t.Fatal(err)
	} else {
		jsonData, err := json.MarshalIndent(lib.Config{
			Cores:   map[int]lib.Core{},
			Servers: map[string]lib.Server{},
			Aria2c: lib.Aria2c{
				Path: "unknown",
				Args: []string{
					"--always-resume=false",
					"--max-resume-failure-tries=0",
					"--allow-overwrite=true",
					"--auto-file-renaming=false",
					"--retry-wait=2",
					"--split=16",
					"--max-connection-per-server=16",
					"--min-split-size=1M",
					"--console-log-level=warn",
					"--no-conf=true",
					"--follow-metalink=true",
					"--metalink-preferred-protocol=https",
					"--min-tls-version=TLSv1.2",
					"--continue",
					"--summary-interval=0",
					"--auto-save-interval=0",
				},
			},
			AutoAcceptEULA: false,
		}, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := file.Write(jsonData); err != nil {
			t.Fatal(err)
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	path, err := lib.NewDownloader(URL).Download()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
}

func TestAria2Downlaod(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过下载")
	}
	if err := lib.Init(goversion.GetVersionInfo()); err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	lib.ConfigsPath = filepath.Join(cwd, "test.json")
	if file, err := os.Create(lib.ConfigsPath); err != nil {
		t.Fatal(err)
	} else {
		aria2Name := "aria2c"
		if runtime.GOOS == "windows" {
			aria2Name = "aria2c.exe"
		}
		aria2, err := exec.LookPath(aria2Name)
		if err != nil {
			t.Fatal(err)
		}
		jsonData, err := json.MarshalIndent(lib.Config{
			Cores:   map[int]lib.Core{},
			Servers: map[string]lib.Server{},
			Aria2c: lib.Aria2c{
				Path: aria2,
				Args: []string{
					"--always-resume=false",
					"--max-resume-failure-tries=0",
					"--allow-overwrite=true",
					"--auto-file-renaming=false",
					"--retry-wait=2",
					"--split=16",
					"--max-connection-per-server=8",
					"--min-split-size=5M",
					"--console-log-level=warn",
					"--no-conf=true",
					"--follow-metalink=true",
					"--metalink-preferred-protocol=https",
					"--min-tls-version=TLSv1.2",
					"--continue",
					"--summary-interval=0",
					"--auto-save-interval=0",
				},
			},
			AutoAcceptEULA: false,
		}, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := file.Write(jsonData); err != nil {
			t.Fatal(err)
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	path, err := lib.NewDownloader(URL).Download()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
}
