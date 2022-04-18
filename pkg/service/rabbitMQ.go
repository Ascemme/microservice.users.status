package service

import (
	"encoding/json"
	"errors"
	"github.com/Ascemme/microservice.users.status/pkg/model"
	"github.com/streadway/amqp"
	"log"
)

func (s *Service) ServiceMq(ch <-chan amqp.Delivery) {
	for msgs := range ch {
		var msg model.Massage
		err := json.Unmarshal(msgs.Body, &msg)
		if err != nil {
			log.Println(err)
			continue
		}
		err = s.rabbiMQLogic(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Service) rabbiMQLogic(msg model.Massage) error {
	switch msg.Value {
	case "dislikes":
		return s.dislike(msg)
	case "likes":
		return s.likes(msg)
	case "following":
		return s.following(msg)
	case "comments":
		return s.comments(msg)
	case "subscribers":
		return s.subscribers(msg)
	case "posts":
		return s.post(msg)
	case "pages":
		return s.post(msg)
	case "delete":
		return s.delete(msg)
	default:
		return errors.New("bad massage")
	}
	return nil
}

func (s *Service) dislike(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}

	if user.Id == "" {
		page := []model.Page{{Id: msg.Page, Dislikes: 1}}
		user = model.User{Uid: msg.Uid, Page: page}
		err := s.repo.CreateStatus(user)
		if err != nil {
			return err
		}
		return nil
	}

	pageId, err := s.sortingSlicePage(msg.Page, user.Page)
	pages := user.Page

	if err != nil {
		pages = append(pages, model.Page{Id: msg.Page, Dislikes: 1})
		user.Page = pages
		return s.repo.UpdateStatus(user)
	}

	user.Page[pageId].Dislikes += 1
	err = s.repo.UpdateStatus(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) likes(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}

	if user.Id == "" {
		page := []model.Page{{Id: msg.Page, Likes: 1}}
		user = model.User{Uid: msg.Uid, Page: page}
		err := s.repo.CreateStatus(user)
		if err != nil {
			return err
		}

		return nil
	}

	pageId, err := s.sortingSlicePage(msg.Page, user.Page)
	pages := user.Page
	if err != nil {
		pages = append(pages, model.Page{Id: msg.Page, Likes: 1})
		user.Page = pages
		return s.repo.UpdateStatus(user)
	}

	user.Page[pageId].Likes += 1

	err = s.repo.UpdateStatus(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) following(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}

	if user.Id == "" {
		user = model.User{Uid: msg.Uid, Following: 1}
		err := s.repo.CreateStatus(user)
		if err != nil {
			return err
		}
		return nil
	}

	user.Following += 1

	err = s.repo.UpdateStatus(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) comments(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}

	if user.Id == "" {

		page := []model.Page{{Id: msg.Page, Post: []model.Post{{Id: 1, Comments: 1}}}}
		user = model.User{Uid: msg.Uid, Page: page}
		err := s.repo.CreateStatus(user)
		if err != nil {
			return err
		}
		return nil
	}

	pageId, err := s.sortingSlicePage(msg.Page, user.Page)
	if err != nil {
		pages := user.Page
		pages = append(pages, model.Page{Id: msg.Page, Post: []model.Post{{Id: msg.Post, Comments: 1}}})
		user.Page = pages
		return s.repo.UpdateStatus(user)
	}

	postId, err := s.sortingSlicePost(msg.Post, user.Page[pageId].Post)
	if err != nil {
		posts := user.Page[pageId].Post
		posts = append(posts, model.Post{Id: msg.Post, Comments: 1})
		user.Page[pageId].Post = posts
		return s.repo.UpdateStatus(user)
	}

	user.Page[pageId].Post[postId].Comments += 1

	err = s.repo.UpdateStatus(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) subscribers(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}

	if user.Id == "" {
		page := []model.Page{{Id: msg.Page, Subscribers: 1}}
		user = model.User{Uid: msg.Uid, Page: page}
		err := s.repo.CreateStatus(user)
		if err != nil {
			return err
		}

		return nil
	}

	pageId, err := s.sortingSlicePage(msg.Page, user.Page)
	pages := user.Page
	if err != nil {
		pages = append(pages, model.Page{Id: msg.Page, Subscribers: 1})
		user.Page = pages
		return s.repo.UpdateStatus(user)
	}

	user.Page[pageId].Subscribers += 1

	err = s.repo.UpdateStatus(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) page(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}

	if user.Id == "" {
		page := []model.Page{{Id: msg.Page}}
		user = model.User{Uid: msg.Uid, Page: page}
		err := s.repo.CreateStatus(user)
		if err != nil {
			return err
		}

		return nil
	}

	_, err = s.sortingSlicePage(msg.Page, user.Page)
	pages := user.Page
	if err != nil {
		pages = append(pages, model.Page{Id: msg.Page})
		user.Page = pages
		return s.repo.UpdateStatus(user)
	}

	err = s.repo.UpdateStatus(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) post(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}

	if user.Id == "" {
		page := []model.Page{{Id: msg.Page, Post: []model.Post{{Id: 1}}}}
		user = model.User{Uid: msg.Uid, Page: page}
		err := s.repo.CreateStatus(user)
		if err != nil {
			return err
		}
		return nil
	}

	pageId, err := s.sortingSlicePage(msg.Page, user.Page)
	if err != nil {
		pages := user.Page
		pages = append(pages, model.Page{Id: msg.Page, Post: []model.Post{{Id: msg.Post}}})
		user.Page = pages
		return s.repo.UpdateStatus(user)
	}

	_, err = s.sortingSlicePost(msg.Page, user.Page[pageId].Post)
	if err != nil {
		posts := user.Page[pageId].Post
		posts = append(posts, model.Post{Id: msg.Post})
		user.Page[pageId].Post = posts
		return s.repo.UpdateStatus(user)
	}

	err = s.repo.UpdateStatus(user)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) delete(msg model.Massage) error {
	user, err := s.repo.GetStatusByUid(msg.Uid)
	if err != nil {
		return err
	}
	if msg.Post != 0 {
		pageId, err := s.sortingSlicePage(msg.Page, user.Page)
		if err != nil {
			return err
		}
		postId, err := s.sortingSlicePost(msg.Post, user.Page[pageId].Post)
		if err != nil {
			return err
		}
		posts := s.removeIndexPost(user.Page[pageId].Post, postId)
		user.Page[pageId].Post = posts
		err = s.repo.UpdateStatus(user)
		return err
	}
	if msg.Page != 0 {
		pageId, err := s.sortingSlicePage(msg.Page, user.Page)
		if err != nil {
			return err
		}
		pages := s.removeIndexPage(user.Page, pageId)
		user.Page = pages
		err = s.repo.UpdateStatus(user)
		return err
	}
	if msg.Uid != 0 {
		s.repo.DeleteStatus(user.Id)
		return err
	}
	return err
}
