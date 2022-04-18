package service

import (
	"errors"
	"fmt"
	"github.com/Ascemme/microservice.users.status/pkg/model"
	"sort"
)

func (s Service) sortingSlicePage(value int, array []model.Page) (int, error) {
	i := sort.Search(len(array), func(i int) bool { return value <= array[i].Id })
	if i < len(array) && array[i].Id == value {
		return i, nil
	} else {
		return 0, errors.New(fmt.Sprintf("not id found in slice %v with value %d", array, value))
	}
}

func (s Service) sortingSlicePost(value int, array []model.Post) (int, error) {
	i := sort.Search(len(array), func(i int) bool { return value <= array[i].Id })
	if i < len(array) && array[i].Id == value {
		return i, nil
	} else {
		return 0, errors.New(fmt.Sprintf("not id found in slice %v with value %d", array, value))
	}
}

func (s Service) removeIndexPage(slice []model.Page, index int) []model.Page {
	return append(slice[:index], slice[index+1:]...)
}
func (s Service) removeIndexPost(slice []model.Post, index int) []model.Post {
	return append(slice[:index], slice[index+1:]...)
}
