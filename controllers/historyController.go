package controllers

import "kalvium/models"

func (c *Controller) HistoryController() ([]models.RequestLog, error) {
	history, err := c.Model.GetRequestLogs()
	return history, err
}
