package errorhandling

import (
	"tracealert/pkg/models/errors"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

func OK_Response(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '200'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
						"Details":     details,
					},
				},
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ErrorCode": errorLogs.Errorcode,
		"Details":   details,
	})
}

// 400 ------------- BAD REQUEST -------------
func Bad_Request(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '400'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
						"Details":     details,
					},
				},
			},
		})
	}

	// Process the fetched data and return a response
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 401 ------------- UNAUTHORIZED -------------
func Unauthorized(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '401'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 403 ------------- PERMISION_DENIED -------------

func Permision_Denied(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '403'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 404  ------------- NOT FOUND -------------
func Url_Not_Found(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '404'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
						"Details":     "Requested URL was not found on the server",
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     "requested URL was not found on the server",
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 405  ------------- METHOD NOT ALLOWED -------------
func Method_Not_Allowed(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '405'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 409  ------------- CONFLICT -------------
func Conflict(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '409'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 415  ------------- UNSUPPORTED MEDIA TYPE -------------
func Unsupported_Media_Type(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '415'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 422  ------------- UNPROCESSABLE ENTITY -------------
func Unprocessable_Entity(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '422'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 429  ------------- RATE LIMIT EXCEEDED -------------
func Rate_Limit_Exceeded(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '429'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}

// 500  ------------- INTERNAL SERVER ERROR -------------
func Internal_Server_Error(c *fiber.Ctx, details string) error {
	// Create an instance of errors.Error_code
	errorLogs := &errors.Errorcode{}

	err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.errorlogs WHERE errorcode = '500'`).First(errorLogs).Error
	if err != nil {
		// Handle the error when fetching data from the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Errors": fiber.Map{
				"Error": []fiber.Map{
					{
						"Source":      "Database",
						"ReasonCode":  "DatabaseError",
						"Description": "Error fetching data from the database",
						"Recoverable": false,
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
					"Source":      errorLogs.Source,
					"ReasonCode":  errorLogs.Reasoncode,
					"Description": errorLogs.Description,
					"Recoverable": errorLogs.Recoverable,
					"Details":     errorLogs.Details,
					"ErrorCode":   errorLogs.Errorcode,
				},
			},
		},
	})
}
