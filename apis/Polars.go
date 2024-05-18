/*
 * MCSCS can be used to easily create, launch, and configure a Minecraft server.
 * Copyright (C) 2024 Arama
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

// 极星云镜像

package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Arama-Vanarana/MCSCS-Go/lib"
)

type Polars struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type ParsedPolars map[string]Polars

func GetPolarsDatas() (ParsedPolars, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "mirror.polars.cc",
		Path:   "/api/query/minecraft/core",
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return ParsedPolars{}, err
	}
	req.Header.Set("User-Agent", "MCSCS-Golang/"+lib.VERSION)
	resp, err := client.Do(req)
	if err != nil {
		return ParsedPolars{}, err
	}
	defer resp.Body.Close()
	var data []Polars
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ParsedPolars{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ParsedPolars{}, err
	}
	parsedData := ParsedPolars{}
	for i := 0; i < len(data); i++ {
		data := data[i]
		parsedData[data.Name] = Polars{
			ID:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			Icon:        data.Icon,
		}
	}
	return parsedData, nil
}

type PolarsCores struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DownloadURL string `json:"downloadUrl"`
	Type        int    `json:"type"`
}

type ParsedPolarsCores map[string]PolarsCores

func GetPolarsCoresDatas(ID int) (ParsedPolarsCores, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "mirror.polars.cc",
		Path:   fmt.Sprintf("/api/query/minecraft/core/%d", ID),
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return ParsedPolarsCores{}, err
	}
	req.Header.Set("User-Agent", "MCSCS-Golang/"+lib.VERSION)
	resp, err := client.Do(req)
	if err != nil {
		return ParsedPolarsCores{}, err
	}
	defer resp.Body.Close()
	var data []PolarsCores
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ParsedPolarsCores{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ParsedPolarsCores{}, err
	}
	parsedData := ParsedPolarsCores{}
	for i := 0; i < len(data); i++ {
		data := data[i]
		parsedData[data.Name] = PolarsCores{
			ID:          data.ID,
			Name:        data.Name,
			DownloadURL: data.DownloadURL,
			Type:        data.Type,
		}
	}
	return parsedData, nil
}

// DownloadURL 和 Name 参数从 GetPolarsCoresDatas 获取
func DownloadPolarsServer(DownloadURL string, Name string) (string, error) {
	path, err := lib.Download(DownloadURL, Name)
	if err != nil {
		return "", err
	}
	return path, nil
}
