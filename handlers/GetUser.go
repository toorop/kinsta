package handlers

import (
	"fmt"
	"github.com/toorop/kinsta/services/insta"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/toorop/goinsta"
)

// GetUser retourne les info d'un user
func GetUser(c echo.Context) (err error) {

	var username string
	var nbImages uint64

	// bind params
	username = c.QueryParam("username")
	if username == "" {
		return c.String(http.StatusBadRequest, "no username")
	}

	nbImagesStr := c.QueryParam("images")
	if nbImagesStr != "" {
		nbImages, err = strconv.ParseUint(nbImagesStr, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("bad parameters for image: %s is not an uint", nbImagesStr))
		}
	} else {
		nbImages = 5
	}

	now := time.Now()
	response := GetUserResponse{
		Date:        now.Format("2006-01-02 15:04:05"),
		Time:        now.Unix(),
		Username:    username,
		URL:         fmt.Sprintf("https://www.instagram.com/%s/", username),
		ImagesLimit: nbImages,
	}

	user, err := insta.Client.Profiles.ByName(username)
	if err != nil {
		// not found ?
		if strings.Contains(err.Error(), "found") {
			return c.String(http.StatusNotFound, fmt.Sprintf("user %s nor found", username))
		}
		return c.String(http.StatusInternalServerError, fmt.Sprintf("insta.Client.Profiles.ByName(%s) failed: %v", username, err))
	}

	response.User = *user
	response.Followers = user.FollowerCount

	feed := user.Feed()
	feed.Next()

	count := uint64(0)
	for _, post := range feed.Items {
		if post.MediaType == 1 || post.MediaType == 8 {
			count++
			// si carroussel
			image := Image{
				ID:      post.ID,
				URL:     fmt.Sprintf("https://www.instagram.com/p/%s/", post.Code),
				Caption: post.Caption.Text,
			}
			if post.MediaType == 1 {
				image.Thumbail = post.Images.GetLower()
				image.Src = post.Images.GetBest()
			} else {
				image.Thumbail = post.CarouselMedia[0].Images.GetLower()
				image.Src = post.CarouselMedia[0].Images.GetBest()
			}
			response.Images = append(response.Images, image)
			if count == nbImages {
				break
			}

		}
	}
	return c.JSON(http.StatusOK, response)

}

type GetUserResponse struct {
	Date        string       `json:"date"`
	Time        int64        `json:"time"`
	Username    string       `json:"username"`
	URL         string       `json:"url"`
	Followers   int          `json:"followers"`
	ImagesLimit uint64       `json:"imagesLimit"`
	User        goinsta.User `json:"user"`
	Images      []Image      `json:"images"`
}

type Image struct {
	ID       string `json:"id"`
	Thumbail string `json:"thumbail"`
	Src      string `json:"src"`
	URL      string `json:"url"`
	Caption  string `json:"caption"`
}
