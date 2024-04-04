package errorhandling

import (
	"tracealert/pkg/models/errors"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

// 0 ------------- Success -------------

func Success(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '0'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 1 ------------- System Malfunction -------------

func SystemMalfunction(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '1'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 2 ------------- Invalid Credential -------------

func InvalidCredential(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '2'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 3 ------------- Invalid Signature -------------

func InvalidSignature(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '3'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 4 ------------- Invalid Request -------------

func InvalidRequest(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '4'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 5 ------------- Invalid Customer -------------

func InvalidCustomer(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '5'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 6 ------------- Timeout -------------

func Timeout(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '6'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 7 ------------- Invalid Customer Access Token -------------

func InvalidCustomerAccessToken(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '7'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 8 ------------- Expired Customer Access Token -------------

func ExpiredCustomerAccessToken(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '8'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 9 ------------- Unsupported Currency -------------

func UnsupportedCurrency(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '9'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 15 ------------- Invalid Amount -------------

func InvalidAmount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '15'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 16 ------------- No Activity Allowed Against The Account -------------

func NoActivityAllowedAgainstTheAccount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '16'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 20 ------------- Duplicate transaction request. Last transaction was succeed -------------

func DuplicateTransactionRequestSucceed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '20'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 21 ------------- Duplicate transaction request. Last transaction was failed -------------

func DuplicateTransactionRequestFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '21'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 30 ------------- Invalid Bill/Order ID -------------

func InvalidBillOrderID(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '30'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 31 ------------- Bill already paid -------------

func BillAlreadyPaid(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '31'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 32 ------------- Bill canceled -------------

func BillCanceled(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '32'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 33 ------------- Bill expired -------------

func BillExpired(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '33'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 34 ------------- No open bill -------------

func NoOpenBill(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '34'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 40 ------------- Invalid reversal request -------------

func InvalidReversalRequest(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '40'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 41 ------------- Duplicate reversal request -------------

func DuplicateReversalRequest(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '41'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 42 ------------- Transaction could not be reversed -------------

func TransactionReversed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '42'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 43 ------------- Account is inactive (freeze) -------------

func AccountInactive(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '43'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 50 ------------- Invalid refund request -------------

func InvalidRefundRequest(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '50'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 51 ------------- Duplicate refund request -------------

func DuplicateRefundRequest(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '51'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 52 ------------- Refund could not be processed -------------

func RefundProcessed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '52'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 53 ------------- Insufficient balance to process the refund -------------

func InsufficientBalanceRefund(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '53'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 54 ------------- Refund request not allowed -------------

func RefundRequestNotAllowed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '54'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 77 ------------- Rejected - Loan already exist for the same Product -------------

func RejectedExistProduct(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '77'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 78 ------------- TIN / SSS.ID Already Exists -------------

func TinSssIdAlreadyExists(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '78'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 79 ------------- No Specified Value of Frequency -------------

func NoSpecifiedValueFrequency(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '79'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 80 ------------- Rejected - Maximum Value Exceeded -------------

func RejectedMaximumValueExceeded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '80'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 81 ------------- Rejected - Error Validation from Core Banking -------------

func RejectedErrorValidationCoreBanking(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '81'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 97 ------------- Account Record Missing -------------

func AccountRecordMissing(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '97'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 99 ------------- Response Not Defined -------------

func ResponseNotDefined(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '99'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// A0 ------------- Data not found -------------

func DataNotFound(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'A0'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// IB-0000 ------------- Unauthorized k2C user -------------

func Unauthorizedk2CUser(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'IB-0000'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// IB-0600 ------------- k2C Server Internal Error -------------

func K2cServerInternalError(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'IB-0600'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// IB-1075 ------------- Cash Out Request already used -------------

func CashOutRequestAlreadyUsed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'IB-1075'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// IB-1076 ------------- Cash Out Request was expired -------------

func CashOutRequestExpired(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'IB-1076'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB01 ------------- Aborted Clearing Timeout -------------

func AbortedClearingTimeout(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB02 ------------- Aborted Clearing Fatal Error -------------

func AbortedClearingFatalError(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB03 ------------- Aborted Settlement Timeout -------------

func AbortedSettlementTimeout(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB04 ------------- Aborted Settlement Fatal Error -------------

func AbortedSettlementFatalError(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB05 ------------- Timeout Creditor Agent -------------

func TimeoutCreditorAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB06 ------------- Timeout Instructed Agent -------------

func TimeoutInstructedAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB07 ------------- Offline Agent -------------

func OfflineAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB08 ------------- Offline Creditor Agent -------------

func OfflineCreditorAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB08'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB09 ------------- Error Creditor Agent -------------

func ErrorCreditorAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AB10 ------------- Error Instructed Agent -------------

func ErrorInstructedAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AB10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC01 ------------- Incorrect Account Number -------------

func IncorrectAccountNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC02 ------------- Invalid Debtor Account Number-------------

func InvalidDebtorAccountNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC03 ------------- Invalid Creditor Account Number -------------

func InvalidCreditorAccountNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC04 ------------- Closed Account Number -------------

func ClosedAccountNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC05 ------------- Closed Debtor Account Number -------------

func ClosedDebtorAccountNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC06 ------------- Blocked Account -------------

func BlockedAccount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC07 ------------- Closed Creditor Account Number -------------

func ClosedCreditorAccountNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC08 ------------- Invalid Branch Code -------------

func InvalidBranchCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC08'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC09 ------------- Invalid Account Currency -------------

func InvalidAccountCurrency(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC10 ------------- Invalid Debtor Account Currency -------------

func InvalidDebtorAccountCurrency(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC11 ------------- Invalid Creditor Account Currency -------------

func InvalidCreditorAccountCurrency(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC12 ------------- Invalid Account Type -------------

func InvalidAccountType(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC12'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC13 ------------- Invalid Debtor Account Type -------------

func InvalidDebtorAccountType(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC13'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC14 ------------- Invalid Creditor Account Type -------------

func InvalidCreditorAccountType(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC14'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AC15 ------------- Account Details Changed -------------

func AccountDetailsChanged(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AC15'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG01 ------------- Transaction Forbidden -------------

func TransactionForbidden(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG02 ------------- Invalid Bank Operation Code -------------

func InvalidBankOperationCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG03 ------------- Transaction Not Supported -------------

func TransactionNotSupported(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG04 ------------- Invalid Agent Country -------------

func InvalidAgentCountry(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG05 ------------- Invalid Debtor Agent Country -------------

func InvalidDebtorAgentCountry(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG06 ------------- Invalid Creditor Agent Country -------------

func InvalidCreditorAgentCountry(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG07 ------------- Unsuccessful Direct Debit -------------

func UnsuccessfulDirectDebit(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG08 ------------- Invalid Access Rights -------------

func InvalidAccessRights(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG08'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG09 ------------- Payment Not Received -------------

func PaymentNotReceived(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG10 ------------- Agent Suspended -------------

func AgentSuspended(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AG11 ------------- Creditor Agent Suspended -------------

func CreditorAgentSuspended(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AG11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AGNT ------------- Incorrect Agent -------------

func IncorrectAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AGNT'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM01 ------------- Zero Amount -------------

func ZeroAmount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM02 ------------- Not Allowed Amount -------------

func NotAllowedAmount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM03 ------------- Not Allowed Currency -------------

func NotAllowedCurrency(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM04 ------------- Insufficient Funds -------------

func InsufficientFunds(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM05 ------------- Duplication -------------

func Duplication(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM06 ------------- Too Low Amount -------------

func TooLowAmount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM07 ------------- Blocked Amount -------------

func BlockedAmount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM09 ------------- Wrong Amount -------------

func WrongAmount(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM10 ------------- Invalid Control Sum -------------

func InvalidControlSum(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM11 ------------- Invalid Transaction Currency -------------

func InvalidTransactionCurrency(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM12 ------------- Invalid Amount -------------

func InvalidAmounts(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM12'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM13 ------------- Amount Exceeds Clearing System Limit -------------

func AmountExceedsClearingSystemLimit(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM13'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM14 ------------- Amount Exceeds Agreed Limit -------------

func AmountExceedsAgreedLimit(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM14'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM15 ------------- Amount Below Clearing System Minimum -------------

func AmountBelowClearingSystemMinimum(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM15'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM16 ------------- Invalid Group Control Sum -------------

func InvalidGroupControlSum(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM16'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM17 ------------- Invalid Payment Info Control Sum -------------

func InvalidPaymentInfoControlSum(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM17'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM18 ------------- Invalid Number Of Transactions -------------

func InvalidNumberOfTransactions(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM18'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM19 ------------- Invalid Group Number Of Transactions -------------

func InvalidGroupNumberOfTransactions(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM19'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM20 ------------- Invalid Payment Info Number Of Transactions -------------

func InvalidPaymentInfoNumberOfTransactions(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM20'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM21 ------------- Limit Exceeded -------------

func LimitExceeded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM21'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM22 ------------- Zero Amount Not Applied -------------

func ZeroAmountNotApplied(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM22'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// AM23 ------------- Amount Exceeds Settlement Limit -------------

func AmountExceedsSettlementLimit(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'AM23'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE01 ------------- Inconsistent With End Customer -------------

func InconsistentWithEndCustomer(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE04 ------------- Missing Creditor Address -------------

func MissingCreditorAddress(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE05 ------------- Unrecognised Initiating Party -------------

func UnrecognisedInitiatingParty(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE06 ------------- Unknown End Customer -------------

func UnknownEndCustomer(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE07 ------------- Missing Debtor Address -------------

func MissingDebtorAddress(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE08 ------------- Missing Debtor Name -------------

func MissingDebtorName(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE08'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE09 ------------- Invalid Country -------------

func InvalidCountry(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE10 ------------- Invalid Debtor Country -------------

func InvalidDebtorCountry(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE11 ------------- Invalid Creditor Country -------------

func InvalidCreditorCountry(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE12 ------------- Invalid Country Of Residence -------------

func InvalidCountryOfResidence(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE12'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE13 ------------- Invalid Debtor Country Of Residence -------------

func InvalidDebtorCountryOfResidence(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE13'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE14 ------------- Invalid Creditor Country Of Residence -------------

func InvalidCreditorCountryOfResidence(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE14'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE15 ------------- Invalid Identification Code -------------

func InvalidIdentificationCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE15'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE16 ------------- Invalid Debtor Identification Code -------------

func InvalidDebtorIdentificationCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE16'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE18 ------------- Invalid Contact Details -------------

func InvalidContactDetails(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE18'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE19 ------------- Invalid Charge Bearer Code -------------

func InvalidChargeBearerCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE19'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE20 ------------- Invalid Name Length -------------

func InvalidNameLength(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE20'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE21 ------------- Missing Name -------------

func MissingName(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE21'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// BE22 ------------- Missing Creditor Name -------------

func MissingCreditorName(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'BE22'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CERI ------------- Check ERI -------------

func CheckERI(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CERI'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH03 ------------- Requested Execution Date Or Requested Collection Date Too Far In Future -------------

func CheckRequestedExecutionFutureERI(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH04 ------------- Requested Execution Date Or Requested Collection Date Too Far In Past -------------

func RequestedExecutionPast(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH07 ------------- Element Is Not To Be Used At B- and C-Level -------------

func ElementLevel(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH09 ------------- Mandate Changes Not Allowed -------------

func MandateChangesNotAllowed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH10 ------------- Information On Mandate Changes Missing -------------

func InformationMandateChangesMissing(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH11 ------------- Creditor Identifier Incorrect-------------

func CreditorIdentifierIncorrect(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH12 ------------- Creditor Identifier Not Unambiguously At Transaction-Level -------------

func CreditorIdentifierLevel(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH12'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH13 ------------- Original DebtorAccountIsNotToBeUsed -------------

func OriginalDebtorAccountIsNotToBeUsed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH13'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH14 ------------- Original Debtor AgentIs Not To Be Used -------------

func OriginalDebtorUsed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH14'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH15 ------------- Element Content Includes More Than 140 Characters -------------

func ElementContentCharacters(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH15'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// CH16 ------------- Element Content Formally Incorrect -------------

func ElementContentFormallyIncorrect(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'CH16'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS17 ------------- Different Order Data In Signatures -------------

func DifferentOrderSignatures(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS17'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS18 ------------- Repeat Order -------------

func RepeatOrder(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS18'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS19 ------------- Electronic Signature Rights Insufficient -------------

func ElectronicSignatureRightsInsufficient(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS19'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS20 ------------- Signer2 Certificate Revoked -------------

func Signer2CertificateRevoked(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS20'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS21 ------------- Signer2 Certificate Not Valid -------------

func Signer2CertificateNotValid(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS21'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS22 ------------- Incorrect Signer2 Certificate -------------

func IncorrectSigner2Certificate(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS22'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS23 ------------- Signer Certification Authority Signer2 Not Valid -------------

func SignerCertificationValid(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS23'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS24 ------------- Waiting Time Expired -------------

func WaitingTimeExpired(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS24'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS25 ------------- Order File Deleted -------------

func OrderFileDeleted(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS25'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS26 ------------- User Signed Multiple Times -------------

func UserSignedMultipleTimes(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS26'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS27 ------------- User Not Yet Activated -------------

func UserNotYetActivated(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS27'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DT01 ------------- Invalid Date -------------

func InvalidDate(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DT01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DT02 ------------- Invalid Creation Date -------------

func InvalidCreationDate(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DT02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DT03 ------------- Invalid Non Processing Date -------------

func InvalidNonProcessingDate(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DT03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DT04 ------------- Future Date Not Supported -------------

func FutureDateNotSupported(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DT04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DT05 ------------- Invalid Cut Off Date -------------

func InvalidCutOffDate(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DT05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DT06 ------------- Execution Date Changed -------------

func ExecutionDateChanged(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DT06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DU01 ------------- Duplicate Message ID -------------

func DuplicateMessageID(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DU01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DU02 ------------- Duplicate Payment Information ID -------------

func DuplicatePaymentInformationID(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DU02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DU03 ------------- Duplicate Transaction -------------

func DuplicateTransaction(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DU03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DU04 ------------- Duplicate EndToEnd ID -------------

func DuplicateEndToEndID(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DU04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DU05 ------------- Duplicate Instruction ID -------------

func DuplicateInstructionID(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DU05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DUPL ------------- Duplicate Payment -------------

func DuplicatePayment(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DUPL'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// ED01 ------------- Correspondent Bank Not Possible -------------

func CorrespondentBankNotPossible(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'ED01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// ED03 ------------- Balance Info Request -------------

func BalanceInfoRequest(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'ED03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// ED05 ------------- Settlement Failed -------------

func SettlementFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'ED05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// ED06 ------------- Settlement System Not Available -------------

func SettlementSystemNotAvailable(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'ED06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// ERIN ------------- ERI Option Not Supported -------------

func ERIOptionNotSupported(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'ERIN'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF01 ------------- Invalid File Format -------------

func InvalidFileFormat(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF02 ------------- Syntax Error -------------

func SyntaxError(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF03 ------------- Invalid Payment Type Information -------------

func InvalidPaymentTypeInformation(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF04 ------------- Invalid Service Level Code -------------

func InvalidServiceLevelCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF05 ------------- Invalid Local Instrument Code -------------

func InvalidLocalInstrumentCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF06 ------------- Invalid Category Purpose Code -------------

func InvalidCategoryPurposeCode(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF07 ------------- Invalid Purpose -------------

func InvalidPurpose(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF08 ------------- Invalid EndToEnd Id -------------

func InvalidEndToEndId(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF08'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF09 ------------- Invalid Cheque Number -------------

func InvalidChequeNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF10 ------------- Bank System Processing Error -------------

func BankSystemProcessingError(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// FF11 ------------- Clearing Request Aborted -------------

func ClearingRequestAborted(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'FF11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// G000 ------------- Payment transferred and not tracked -------------

func PaymentTransferredTracked(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'G000'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// G001 ------------- Credit Debit Not Confirmed -------------

func CreditDebitNotConfirmed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'G001'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// G002 ------------- Credit Debit Not Confirmed -------------

func CreditDebitNotConfirmeds(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'G002'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// G003 ------------- Credit Pending Documents -------------

func CreditPendingDocuments(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'G003'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// G004 ------------- Credit Pending Funds -------------

func CreditPendingFunds(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'G004'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// G005 ------------- Delivered With Service Level -------------

func DeliveredWithServiceLevel(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'G005'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// G006 ------------- Delivered WIthout Service Level -------------

func DeliveredWIthoutServiceLevel(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'G006'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// ID01 ------------- Corresponding Original File Still Not Sent -------------

func CorrespondingOriginalNotSent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'ID01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// MD01 ------------- No Mandate -------------

func NoMandate(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'MD01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// MD02 ------------- Missing Mandatory Information In Mandate -------------

func MissingMandatoryInformationMandate(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'MD02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// MD05 ------------- Collection Not Due -------------

func CollectionNotDue(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'MD05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// MD06 ------------- Refund Request By End Customer -------------

func RefundRequestEndCustomer(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'MD06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// MD07 ------------- End Customer Deceased -------------

func EndCustomerDeceased(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'MD07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// MS02 ------------- Not Specified Reason Customer Generated -------------

func NotSpecifiedReasonCustomerGenerated(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'MS02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// MS03 ------------- Not Specified Reason Agent Generated -------------

func NotSpecifiedReasonAgentGenerated(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'MS03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// NARR ------------- Narrative -------------

func Narrative(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'NARR'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// NERI ------------- No ERI -------------

func NoERI(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'NERI'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC01 ------------- Bank Identifier Incorrect -------------

func BankIdentifierIncorrect(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC02 ------------- Invalid Bank Identifier -------------

func InvalidBankIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC03 ------------- Invalid Debtor Bank Identifier -------------

func InvalidDebtorBankIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC04 ------------- Invalid Creditor Bank Identifier -------------

func InvalidCreditorBankIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC05 ------------- Invalid BIC Identifier -------------

func InvalidBICIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC06 ------------- Invalid Debtor BIC Identifier -------------

func InvalidDebtorBICIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC07 ------------- Invalid Creditor BIC Identifier -------------

func InvalidCreditorBICIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC08 ------------- Invalid Clearing System Member Identifier -------------

func InvalidClearingSystemMemberIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC08'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC09 ------------- Invalid Debtor Clearing System Member Identifier -------------

func InvalidDebtorClearingSystemMemberIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC10 ------------- Invalid Creditor Clearing System Member Identifier -------------

func InvalidCreditorClearingSystemMemberIdentifier(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC11 ------------- Invalid Intermediary Agent -------------

func InvalidIntermediaryAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RC12 ------------- Missing Creditor Scheme Id -------------

func MissingCreditorSchemeId(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RC12'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RCON ------------- R-Message Conflict -------------

func R_MessageConflict(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RCON'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RF01 ------------- Not Unique Transaction Reference -------------

func NotUniqueTransactionReference(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RF01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR01 ------------- Missing Debtor Account or Identification -------------

func MissingDebtorAccountIdentification(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR02 ------------- Missing Debtor Name or Address -------------

func MissingDebtorNameAddress(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR03 ------------- Missing Creditor Name or Address -------------

func MissingCreditorNameAddress(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR04 ------------- Regulatory Reason -------------

func RegulatoryReason(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR05 ------------- Regulatory Information Invalid -------------

func RegulatoryInformationInvalid(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR05'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR06 ------------- Tax Information Invalid -------------

func TaxInformationInvalid(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR06'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR07 ------------- Remittance Information Invalid -------------

func RemittanceInformationInvalid(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR07'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR08 ------------- Remittance Information Truncated -------------

func RemittanceInformationTruncated(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR08'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR09 ------------- Invalid Structured Creditor Reference -------------

func RemittanceInformatioInvalidStructuredCreditorReferencenTruncated(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR09'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR10 ------------- Invalid Character Set -------------

func InvalidCharacterSet(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR10'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR11 ------------- Invalid Debtor Agent Service ID -------------

func InvalidDebtorAgentServiceID(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// RR12 ------------- Invalid Party ID -------------

func InvalidPartyID(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'RR12'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// S001 ------------- UETR Flagged For Cancellation -------------

func UETRFlaggedForCancellation(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'S001'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// S002 ------------- Network Stop Of UETR -------------

func NetworkStopUETR(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'S002'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// S002 ------------- Request For Canacellation Forwarded -------------

func RequestCanacellationForwarded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'S002'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// S003 ------------- Request For Canacellation Forwarded -------------

func RequestForCanacellationForwarded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'S003'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// S004 ------------- Request For Cancellation Delivery Acknowledgement -------------

func RequestCancellationDeliveryAcknowledgement(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'S004'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// SL01 ------------- Specific Service offered by Debtor Agent -------------

func SpecificServiceofferedDebtorAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'SL01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// SL02 ------------- Specific Service offered by Creditor Agent -------------

func SpecificServiceOfferedCreditorAgent(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'SL02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// SL11 ------------- Creditor not on Whitelist of Debtor -------------

func CreditorNotWhitelistDebtor(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'SL11'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// SL12 ------------- Creditor on Blacklist of Debtor -------------

func CreditorBlacklistDebtor(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'SL12'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// SL13 ------------- Maximum number of Direct Debit Transactions exceeded -------------

func MaximumNumberDirectDebitTransactionsExceeded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'SL13'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// SL14 ------------- Maximum Direct Debit Transaction Amount exceeded -------------

func MaximumDirectDebitTransactionAmountExceeded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'SL14'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// TA01 ------------- Transmisson Aborted -------------

func TransmissonAborted(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'TA01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// TD01 ------------- No Data Available -------------

func NoDataAvailable(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'TD01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// TD02 ------------- File Non Readable -------------

func FileNonReadable(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'TD02'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// TD03 ------------- Incorrect File Structure -------------

func IncorrectFileStructure(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'TD03'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// TM01 ------------- Invalid Cut Off Time -------------

func InvalidCutOffTime(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'TM01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// TS01 ------------- Transmission Successful -------------

func TransmissionSuccessful(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'TS01'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// TS04 ------------- Transfer To Sign By Hand -------------

func TransmissioTransferToSignByHandnSuccessful(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'TS04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 9912 ------------- Recipient connection is not available -------------

func RecipientConnectionNotAvailable(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '9912'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 9910 ------------- Instructed Agent is signed-off -------------

func InstructedAgentSignedOff(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '9910'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 9934 ------------- Instructing Agent is signed-off -------------

func InstructingAgentSignedOff(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '9934'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 9946 ------------- Instructing Agent is suspended -------------

func InstructingAgentSuspended(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '9946'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 9947 ------------- Instructed Agent is suspended -------------

func InstructedAgentSuspended(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '9947'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 9948 ------------- Central Switch (InstaPay IPS) service is suspended -------------

func CentralSwitchServiceSuspended(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '9948'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// NONE ------------- None -------------

func None(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'NONE'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// DS04 ------------- Order Rejected -------------

func OrderRejected(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  'DS04'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 100 ------------- Validation Failed -------------

func ValidationFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '100'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusContinue).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 101 ------------- First Login: Password Reset Required -------------

func FirstLoginPasswordResetRequired(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '101'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusSwitchingProtocols).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 102 ------------- Password Expired -------------

func PasswordExpired(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '102'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusProcessing).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 103 ------------- Invalid Password -------------

func InvalidPassword(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '103'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusEarlyHints).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 104 ------------- Invalid Token -------------

func InvalidToken(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '104'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 105 ------------- User Already Logged In -------------

func UserAlreadyLoggedIn(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '105'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 106 ------------- Password Cannot Be Reused -------------

func PasswordCannotBeReused(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '106'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 107 ------------- Invalid Date -------------

func InvalidDates(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '107'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 108 ------------- Invalid Phone number -------------

func InvalidPhoneNumber(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '108'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 109 ------------- Invalid Email Address -------------

func InvalidEmailAddress(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '109'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 110 ------------- Expired Token -------------

func ExpiredToken(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '110'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 111 ------------- Token Missing -------------

func TokenMissing(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '111'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 112 ------------- Invalid Staff Id -------------

func InvalidStaffId(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '112'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 113 ------------- Password Reuse Prohibited -------------

func PasswordReuseProhibited(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '113'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 114 ------------- Restricted: Password Reset Required -------------

func RestrictedPasswordResetRequired(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '114'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 115 ------------- Session Id Missing -------------

func SessionIdMissing(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '115'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 116 ------------- Email Reuse Prohibited -------------

func EmailReuseProhibited(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '116'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 117 ------------- Invalid OTP -------------

func InvalidOTP(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '117'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 118 ------------- Expired OTP -------------

func ExpiredOTP(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '118'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 200 ------------- Successful -------------
func OK_Response(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '200'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ErrorCode": errorLogs.Code,
		"Details":   details,
	})
}

// 201 ------------- Successfully Logged In -------------

func SuccessfullyLoggedIn(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '201'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 202 ------------- Successfully Logged Out -------------

func SuccessfullyLoggedOut(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '202'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 203 ------------- Successfully Inserted -------------

func SuccessfullyInserted(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '203'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 204 ------------- Successfully Updated -------------

func SuccessfullyUpdated(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '204'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 205 ------------- Successfully Downloaded -------------

func SuccessfullyDownloaded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '205'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusResetContent).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 300 ------------- Internal Server Error -------------

func InternalServerError(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '300'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 300 ------------- Parsing Data Failed -------------

func ParsingDataFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '300'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 301 ------------- Parsing Data Failed -------------

func ParsingDataFaileds(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '301'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMovedPermanently).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 302 ------------- Fetching Data Failed -------------

func FetchingDataFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '302'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusFound).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 303 ------------- Inserting Data Failed -------------

func InsertingDataFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '303'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusSeeOther).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 304 ------------- Updating Data Failed -------------

func UpdatingDataFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '304'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 305 ------------- Token Generation Failed -------------

func TokenGenerationFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '305'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusUseProxy).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 306 ------------- File Download Failed -------------

func FileDownloadFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '306'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusSwitchProxy).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 307 ------------- File Creation Failed -------------

func FileCreationFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '307'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusTemporaryRedirect).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 308 ------------- URL Retrieval Failed -------------

func URLRetrievalFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '307'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusPermanentRedirect).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 309 ------------- Unable to Open File -------------

func UnableOpenFile(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '309'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 310 ------------- Unmarshalling Failed -------------

func UnmarshallingFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '310'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 311 ------------- Marshalling Failed -------------

func MarshallingFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '311'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 312 ------------- Timezone Loading Failed -------------

func TimezoneLoadingFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '312'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 313 ------------- HTML Parsing Failed -------------

func HTMLParsingFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '313'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 314 ------------- Deleting Data Failed -------------

func DeletingDataFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '314'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 315 ------------- Send Emails Failed -------------

func SendEmailsFailed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '315'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMultipleChoices).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 400 ------------- BAD REQUEST -------------
func Bad_Request(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '400'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 401 ------------- UNAUTHORIZED -------------
func Unauthorized(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '401'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 402 ------------- Unauthorized Access -------------

func UnauthorizedAccess(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '402'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 403 ------------- PERMISION_DENIED -------------

func Permision_Denied(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '403'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 404  ------------- NOT FOUND -------------
func Url_Not_Found(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '404'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 405  ------------- METHOD NOT ALLOWED -------------
func Method_Not_Allowed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '405'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 409  ------------- CONFLICT -------------
func Conflict(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '409'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 415  ------------- UNSUPPORTED MEDIA TYPE -------------
func Unsupported_Media_Type(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '415'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 422  ------------- UNPROCESSABLE ENTITY -------------
func Unprocessable_Entity(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '422'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 429  ------------- RATE LIMIT EXCEEDED -------------
func Rate_Limit_Exceeded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '429'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 500  ------------- INTERNAL SERVER ERROR -------------
func Internal_Server_Error(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code = '500'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 501 ------------- Restricted IP -------------

func RestrictedIP(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '501'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 502 ------------- Restricted Username -------------

func RestrictedUsername(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '502'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 503 ------------- Service Unavailable -------------

func ServiceUnavailable(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '503'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}

// 504 ------------- Connection Exceed The Server Timeout -------------

func ConnectionExceedTheServerTimeout(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.error_logs WHERE code =  '504'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"ReasonCode":  errorLogs.Code,
					"Description": errorLogs.Description,
					"Service":     errorLogs.Service,
					"Details":     details,
				},
			},
		},
	})
}
