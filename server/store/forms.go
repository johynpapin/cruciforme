package store

import "encoding/json"

type FormEmail struct {
	Email     string
	Confirmed bool
}

type Form struct {
	Id     string
	UserId string
	Name   string
	Emails []FormEmail
}

func (s *Store) CreateForm(form *Form) error {
	formBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	formKey := []byte("forms/" + form.Id)
	userFormIdsKey := []byte("users/" + form.UserId + "/forms")

	return s.update(func(txn txn) error {
		// Store the form
		if err := txn.set(formKey, formBytes); err != nil {
			return err
		}

		// Add the form to the list of user forms
		userFormIds, err := getFormIdsByUserId(txn, form.UserId)
		if err != nil {
			return err
		}

		userFormIds = append(userFormIds, form.Id)

		userFormIdsBytes, err := json.Marshal(userFormIds)
		if err != nil {
			return err
		}

		return txn.set(userFormIdsKey, userFormIdsBytes)
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

func getFormIdsByUserId(txn txn, userId string) ([]string, error) {
	userFormsBytes, err := txn.get([]byte("users/" + userId + "/forms"))
	if err != nil {
		return nil, err
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
