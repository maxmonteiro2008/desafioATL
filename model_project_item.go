/*
 * API Inventory
 *
 * This is the Challenge API
 *
 * API version: 1.0.0
 * Contact: maxmonte.vix@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
//package swagger
package main

import (
	"time"
)

type ProjectItem struct {
	Name string `bson:"name"`

	DisplayName string `bson:"display_name,omitempty"`

	Description string `bson:"description,omitempty"`

	CreateTime time.Time `bson:"create_time,omitempty"`

	UpdateTime time.Time `bson:"update_time,omitempty"`

	Apis []ApiItem `bson:"apis,omitempty"`
}

type Projects []ProjectItem