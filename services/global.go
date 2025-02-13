package services

import (
	Models "savannah/models"
	Utils "savannah/utils"
)

//Global definitions for this package

var db = Utils.ConnectDatabase()
var category Models.Categories
var categories []Models.Categories
var product Models.Products
var products []Models.Products
var customer Models.Customers
var customers []Models.Customers
var order Models.Orders
var orders []Models.Orders
var status bool
