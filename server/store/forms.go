package store

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/dchest/uniuri"
)

const (
	formIdBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345789"
	formIdLength = 8
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Form struct {
	Id       string
	UserId   string
	Name     string
	Email    string
	Verified bool
}

func (s *Store) CreateForm(form *Form) error {
	return s.update(func(txn txn) error {
		if form.Id == "" {
			formId, err := generateFormId(txn)
			if err != nil {
				return fmt.Errorf("could not generate the form id: %w", err)
			}

			form.Id = formId
		}

		formBytes, err := json.Marshal(form)
		if err != nil {
			return fmt.Errorf("could not marshal the form: %w", err)
		}

		// Store the form
		if err := txn.set([]byte("forms/"+form.Id), formBytes); err != nil {
			return fmt.Errorf("could not store the form: %w", err)
		}

		// Add the form to the list of user forms
		userFormIds, err := getFormIdsByUserId(txn, form.UserId)
		if err != nil {
			return fmt.Errorf("could not get the form ids of the user: %w", err)
		}

		userFormIds = append(userFormIds, form.Id)

		userFormIdsBytes, err := json.Marshal(userFormIds)
		if err != nil {
			return fmt.Errorf("could not marshal the form ids of the user: %w", err)
		}

		return txn.set([]byte("users/"+form.UserId+"/forms"), userFormIdsBytes)
	})
}

func (s *Store) GetFormById(formId string) (*Form, error) {
	formBytes, err := s.get([]byte("forms/" + formId))
	if err != nil {
		return nil, err
	}

	form := &Form{}
	if err = json.Unmarshal(formBytes, form); err != nil {
		return nil, err
	}

	return form, nil
}

func (s *Store) GetFormsByUserId(userId string) ([]*Form, error) {
	var userForms []*Form

	if err := s.view(func(txn txn) error {
		userFormIds, err := getFormIdsByUserId(txn, userId)
		if err != nil {
			return err
		}

		for _, formId := range userFormIds {
			form, err := getFormByFormId(txn, formId)
			if err != nil {
				return err
			}

			userForms = append(userForms, form)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return userForms, nil
}

func generateFormId(txn txn) (string, error) {
	for {
		formId := uniuri.NewLen(formIdLength)

		form, err := txn.get([]byte("forms/" + formId))
		if err != nil {
			return "", err
		}

		if form == nil {
			return formId, nil
		}
	}
}

func getFormIdsByUserId(txn txn, userId string) ([]string, error) {
	userFormsBytes, err := txn.get([]byte("users/" + userId + "/forms"))
	if err != nil {
		return nil, err
	}
	if userFormsBytes == nil {
		return nil, nil
	}

	var userForms []string
	if err = json.Unmarshal(userFormsBytes, &userForms); err != nil {
		return nil, err
	}

	return userForms, nil
}

func getFormByFormId(txn txn, formId string) (*Form, error) {
	formBytes, err := txn.get([]byte("forms/" + formId))
	if err != nil {
		return nil, err
	}

	form := &Form{}
	if err = json.Unmarshal(formBytes, form); err != nil {
		return nil, err
	}

	return form, nil
}
