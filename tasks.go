package cts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAllTasks - Returns all user's task
func (c *Client) GetAllTasks() (*[]TaskItem, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/tasks", c.HostURL, c.APIVersion), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequestOK(req)
	if err != nil {
		return nil, err
	}

	tasks := TaskList{}
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return nil, err
	}

	return &tasks.Tasks, nil
}

// GetTask - Returns a specifc task
func (c *Client) GetTask(taskID string) (*TaskResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/tasks/%s", c.HostURL, c.APIVersion, taskID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequestOK(req)
	if err != nil {
		return nil, err
	}
	task := &TaskResponse{}
	err = json.Unmarshal(body, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// CreateTask - Create new task
func (c *Client) CreateTask(newTask Task) (*TaskResponse, error) {

	rb, err := json.Marshal(newTask)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/tasks", c.HostURL, c.APIVersion), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequestCreated(req)
	if err != nil {
		return nil, err
	}

	task := TaskResponse{}
	err = json.Unmarshal(body, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Enable Task - Updates an task
func (c *Client) UpdateTaskEnable(taskID string, enable bool) error {
	patch := fmt.Sprintf("{\"enabled\":%v}", enable)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s/tasks/%s", c.HostURL, c.APIVersion, taskID), strings.NewReader(patch))
	if err != nil {
		return err
	}
	_, err = c.doRequestOK(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateTaskEnableWithRun(taskID string, enable bool) (*UpdateWithRunResponse, error) {
	patch := fmt.Sprintf("{\"enabled\":%v}", enable)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s/tasks/%s", c.HostURL, c.APIVersion, taskID), strings.NewReader(patch))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequestOK(req)
	if err != nil {
		return nil, err
	}

	update := &UpdateWithRunResponse{}

	err = json.Unmarshal(body, update)
	if err != nil {
		return nil, err
	}

	return update, nil
}

func (c *Client) TaskEnableInspect(taskID string, enable bool) (*bool, *string, error) {
	taskIDWithInspect := fmt.Sprintf("%s?run=inspect", taskID)
	update, err := c.UpdateTaskEnableWithRun(taskIDWithInspect, enable)
	if err != nil {
		return nil, nil, err
	}

	return &update.Inspect.ChangesPresent, &update.Inspect.Plan, nil
}

func (c *Client) TaskEnableNow(taskID string, enable bool) (*bool, *string, error) {
	taskIDWithInspect := fmt.Sprintf("%s?run=now", taskID)
	update, err := c.UpdateTaskEnableWithRun(taskIDWithInspect, enable)
	if err != nil {
		return nil, nil, err
	}

	return &update.Inspect.ChangesPresent, &update.Inspect.Plan, nil
}

// DeleteTask - Deletes an task
func (c *Client) DeleteTask(taskID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/tasks/%s", c.HostURL, c.APIVersion, taskID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequestAccepted(req)
	if err != nil {
		return err
	}

	delete := &DeleteResponse{}
	err = json.Unmarshal(body, delete)
	if err != nil {
		return err
	}

	return nil
}
