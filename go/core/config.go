package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Luma",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://public-api.luma.com",
			"auth": map[string]any{
				"prefix": "",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"calendar": map[string]any{},
				"calendar_admin": map[string]any{},
				"calendar_coupon": map[string]any{},
				"calendar_event": map[string]any{},
				"calendar_event_approval": map[string]any{},
				"calendar_event_rejection": map[string]any{},
				"contact": map[string]any{},
				"contact_block": map[string]any{},
				"contact_restore": map[string]any{},
				"contact_tag": map[string]any{},
				"contact_tag_assignment": map[string]any{},
				"entity_lookup": map[string]any{},
				"event": map[string]any{},
				"event_cancel_request": map[string]any{},
				"event_coupon": map[string]any{},
				"event_tag": map[string]any{},
				"event_tag_assignment": map[string]any{},
				"guest": map[string]any{},
				"guest_invite": map[string]any{},
				"guest_ticket": map[string]any{},
				"host": map[string]any{},
				"image_upload": map[string]any{},
				"member": map[string]any{},
				"membership_tier": map[string]any{},
				"organization_admin": map[string]any{},
				"organization_calendar": map[string]any{},
				"organization_event": map[string]any{},
				"organization_event_transfer": map[string]any{},
				"ticket_type": map[string]any{},
				"user": map[string]any{},
				"webhook": map[string]any{},
			},
		},
		"entity": map[string]any{
			"calendar": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "avatar_url",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "calendar_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "coordinate",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "cover_image_url",
						"req": true,
						"type": "`$ANY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "description",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "instagram_handle",
						"req": true,
						"type": "`$ANY`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "is_personal",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "location",
						"req": true,
						"type": "`$ANY`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "slug",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "social_image_url",
						"req": true,
						"type": "`$ANY`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "tint_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "twitter_handle",
						"req": true,
						"type": "`$ANY`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "url",
						"req": true,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "website",
						"req": true,
						"type": "`$ANY`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "youtube_handle",
						"req": true,
						"type": "`$ANY`",
						"index$": 16,
					},
				},
				"name": "calendar",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/v1/calendars/get",
								"parts": []any{
									"v1",
									"calendars",
									"get",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/update",
								"parts": []any{
									"v1",
									"calendars",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"calendar_admin": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "avatar_url",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "first_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "last_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
				},
				"name": "calendar_admin",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/v1/calendars/admins/list",
								"parts": []any{
									"v1",
									"calendars",
									"admins",
									"list",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"calendar_coupon": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "cents_off",
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "code",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "currency",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "discount",
						"req": true,
						"type": "`$ANY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "event_ticket_type_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "percent_off",
						"req": true,
						"type": "`$ANY`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "remaining_count",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$INTEGER`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$INTEGER`",
							},
						},
						"req": true,
						"type": "`$INTEGER`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "valid_end_at",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "valid_start_at",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 9,
					},
				},
				"name": "calendar_coupon",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/coupons/create",
								"parts": []any{
									"v1",
									"calendars",
									"coupons",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/calendars/coupons/list",
								"parts": []any{
									"v1",
									"calendars",
									"coupons",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"pagination_cursor",
										"pagination_limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/coupons/update",
								"parts": []any{
									"v1",
									"calendars",
									"coupons",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"calendar_event": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "submitted_by",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "tag",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 3,
					},
				},
				"name": "calendar_event",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/events/add",
								"parts": []any{
									"v1",
									"calendars",
									"events",
									"add",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "access",
											"orig": "access",
											"reqd": false,
											"type": "`$ARRAY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "after",
											"orig": "after",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "before",
											"orig": "before",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "platform",
											"orig": "platform",
											"reqd": false,
											"type": "`$ARRAY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort_column",
											"orig": "sort_column",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort_direction",
											"orig": "sort_direction",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "status",
											"orig": "status",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/calendars/events/list",
								"parts": []any{
									"v1",
									"calendars",
									"events",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"access",
										"after",
										"before",
										"pagination_cursor",
										"pagination_limit",
										"platform",
										"sort_column",
										"sort_direction",
										"status",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_api_id",
											"orig": "event_api_id",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_id",
											"orig": "event_id",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "platform",
											"orig": "platform",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "url",
											"orig": "url",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/calendars/events/lookup",
								"parts": []any{
									"v1",
									"calendars",
									"events",
									"lookup",
								},
								"select": map[string]any{
									"exist": []any{
										"event_api_id",
										"event_id",
										"platform",
										"url",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.event`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"calendar_event_approval": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "calendar_event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
				},
				"name": "calendar_event_approval",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/events/approve",
								"parts": []any{
									"v1",
									"calendars",
									"events",
									"approve",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"calendar_event_rejection": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "calendar_event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "message",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
				},
				"name": "calendar_event_rejection",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/events/reject",
								"parts": []any{
									"v1",
									"calendars",
									"events",
									"reject",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "avatar_url",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "contact",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "event_approved_count",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "event_checked_in_count",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "first_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "last_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "membership",
						"req": true,
						"type": "`$ANY`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "revenue_usd_cent",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "tag",
						"op": map[string]any{
							"list": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"req": false,
						"type": "`$ANY`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "user_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 13,
					},
				},
				"name": "contact",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contacts/import",
								"parts": []any{
									"v1",
									"calendars",
									"contacts",
									"import",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "calendar_membership_tier_id",
											"orig": "calendar_membership_tier_id",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "membership_status",
											"orig": "membership_status",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "query",
											"orig": "query",
											"reqd": false,
											"type": "`$ANY`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort_column",
											"orig": "sort_column",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort_direction",
											"orig": "sort_direction",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"reqd": false,
											"type": "`$ANY`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/calendars/contacts/list",
								"parts": []any{
									"v1",
									"calendars",
									"contacts",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"calendar_membership_tier_id",
										"membership_status",
										"pagination_cursor",
										"pagination_limit",
										"query",
										"sort_column",
										"sort_direction",
										"tag",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contacts/remove",
								"parts": []any{
									"v1",
									"calendars",
									"contacts",
									"remove",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact_block": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "contact_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
				},
				"name": "contact_block",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contacts/block",
								"parts": []any{
									"v1",
									"calendars",
									"contacts",
									"block",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact_restore": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "contact_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
				},
				"name": "contact_restore",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contacts/restore",
								"parts": []any{
									"v1",
									"calendars",
									"contacts",
									"restore",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact_tag": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "color",
						"op": map[string]any{
							"list": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"req": false,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "tag_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
				},
				"name": "contact_tag",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contact-tags/create",
								"parts": []any{
									"v1",
									"calendars",
									"contact-tags",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/v1/calendars/contact-tags/list",
								"parts": []any{
									"v1",
									"calendars",
									"contact-tags",
									"list",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contact-tags/delete",
								"parts": []any{
									"v1",
									"calendars",
									"contact-tags",
									"delete",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contact-tags/update",
								"parts": []any{
									"v1",
									"calendars",
									"contact-tags",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact_tag_assignment": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "applied_count",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "skipped_count",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "tag",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "user_id",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 4,
					},
				},
				"name": "contact_tag_assignment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contact-tags/apply",
								"parts": []any{
									"v1",
									"calendars",
									"contact-tags",
									"apply",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/contact-tags/unapply",
								"parts": []any{
									"v1",
									"calendars",
									"contact-tags",
									"unapply",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"entity_lookup": map[string]any{
				"fields": []any{},
				"name": "entity_lookup",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/entities/lookup",
								"parts": []any{
									"v1",
									"entities",
									"lookup",
								},
								"select": map[string]any{
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entity`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "access",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "calendar_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "can_register_for_multiple_ticket",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "coordinate",
						"req": true,
						"type": "`$ANY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "cover_url",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "description",
						"req": true,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "description_md",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "display_price",
						"req": true,
						"type": "`$ANY`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "duration_interval",
						"req": true,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "end_at",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "feedback_email",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "geo_address_json",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "guest_count",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "host",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "location_type",
						"req": true,
						"type": "`$STRING`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "location_visibility",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "max_capacity",
						"req": false,
						"type": "`$ANY`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "meeting_url",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 21,
					},
					map[string]any{
						"active": true,
						"name": "name_requirement",
						"req": false,
						"type": "`$STRING`",
						"index$": 22,
					},
					map[string]any{
						"active": true,
						"name": "phone_number_requirement",
						"req": false,
						"type": "`$ANY`",
						"index$": 23,
					},
					map[string]any{
						"active": true,
						"name": "platform",
						"req": true,
						"type": "`$STRING`",
						"index$": 24,
					},
					map[string]any{
						"active": true,
						"name": "registration_open",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$BOOLEAN`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$BOOLEAN`",
							},
						},
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 25,
					},
					map[string]any{
						"active": true,
						"name": "registration_question",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 26,
					},
					map[string]any{
						"active": true,
						"name": "reminders_disabled",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 27,
					},
					map[string]any{
						"active": true,
						"name": "require_approval",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 28,
					},
					map[string]any{
						"active": true,
						"name": "show_guest_list",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 29,
					},
					map[string]any{
						"active": true,
						"name": "slug",
						"req": false,
						"type": "`$STRING`",
						"index$": 30,
					},
					map[string]any{
						"active": true,
						"name": "spots_remaining",
						"req": true,
						"type": "`$ANY`",
						"index$": 31,
					},
					map[string]any{
						"active": true,
						"name": "start_at",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 32,
					},
					map[string]any{
						"active": true,
						"name": "suppress_notification",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 33,
					},
					map[string]any{
						"active": true,
						"name": "timezone",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 34,
					},
					map[string]any{
						"active": true,
						"name": "tint_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 35,
					},
					map[string]any{
						"active": true,
						"name": "url",
						"req": true,
						"type": "`$STRING`",
						"index$": 36,
					},
					map[string]any{
						"active": true,
						"name": "user_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 37,
					},
					map[string]any{
						"active": true,
						"name": "visibility",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 38,
					},
					map[string]any{
						"active": true,
						"name": "waitlist_status",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 39,
					},
				},
				"name": "event",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/create",
								"parts": []any{
									"v1",
									"events",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_id",
											"orig": "event_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/events/get",
								"parts": []any{
									"v1",
									"events",
									"get",
								},
								"select": map[string]any{
									"exist": []any{
										"event_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/cancel",
								"parts": []any{
									"v1",
									"events",
									"cancel",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/update",
								"parts": []any{
									"v1",
									"events",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_cancel_request": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "cancellation_token",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "guest_count",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "is_paid",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 3,
					},
				},
				"name": "event_cancel_request",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/cancel/request",
								"parts": []any{
									"v1",
									"events",
									"cancel",
									"request",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_coupon": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "cents_off",
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "code",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "currency",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "discount",
						"req": true,
						"type": "`$ANY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "event_ticket_type_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "percent_off",
						"req": true,
						"type": "`$ANY`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "remaining_count",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$INTEGER`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$INTEGER`",
							},
						},
						"req": true,
						"type": "`$INTEGER`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "valid_end_at",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "valid_start_at",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 10,
					},
				},
				"name": "event_coupon",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/coupons/create",
								"parts": []any{
									"v1",
									"events",
									"coupons",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_id",
											"orig": "event_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/events/coupons/list",
								"parts": []any{
									"v1",
									"events",
									"coupons",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"event_id",
										"pagination_cursor",
										"pagination_limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/coupons/update",
								"parts": []any{
									"v1",
									"events",
									"coupons",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_tag": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "color",
						"op": map[string]any{
							"list": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"req": false,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "tag_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
				},
				"name": "event_tag",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/event-tags/create",
								"parts": []any{
									"v1",
									"calendars",
									"event-tags",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/v1/calendars/event-tags/list",
								"parts": []any{
									"v1",
									"calendars",
									"event-tags",
									"list",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/event-tags/delete",
								"parts": []any{
									"v1",
									"calendars",
									"event-tags",
									"delete",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/event-tags/update",
								"parts": []any{
									"v1",
									"calendars",
									"event-tags",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_tag_assignment": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "applied_count",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "skipped_count",
						"req": true,
						"type": "`$NUMBER`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "tag",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
				},
				"name": "event_tag_assignment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/event-tags/apply",
								"parts": []any{
									"v1",
									"calendars",
									"event-tags",
									"apply",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/calendars/event-tags/unapply",
								"parts": []any{
									"v1",
									"calendars",
									"event-tags",
									"unapply",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"guest": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "approval_status",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$ANY`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "check_in_qr_code",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "eth_address",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "event_ticket",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "event_ticket_order",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "guest",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "guest_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "invited_at",
						"req": true,
						"type": "`$ANY`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "joined_at",
						"req": true,
						"type": "`$ANY`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "message",
						"req": false,
						"type": "`$ANY`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "phone_number",
						"req": true,
						"type": "`$INTEGER`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "registered_at",
						"req": true,
						"type": "`$ANY`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "registration_answer",
						"req": true,
						"type": "`$ANY`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "send_email",
						"req": false,
						"type": "`$ANY`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "should_refund",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "solana_address",
						"req": true,
						"type": "`$ANY`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": true,
						"type": "`$STRING`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "ticket",
						"req": false,
						"type": "`$ANY`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "user_email",
						"req": true,
						"type": "`$STRING`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "user_first_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 21,
					},
					map[string]any{
						"active": true,
						"name": "user_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 22,
					},
					map[string]any{
						"active": true,
						"name": "user_last_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 23,
					},
					map[string]any{
						"active": true,
						"name": "user_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 24,
					},
					map[string]any{
						"active": true,
						"name": "utm_source",
						"req": true,
						"type": "`$ANY`",
						"index$": 25,
					},
				},
				"name": "guest",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/guests/add",
								"parts": []any{
									"v1",
									"events",
									"guests",
									"add",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "approval_status",
											"orig": "approval_status",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_id",
											"orig": "event_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort_column",
											"orig": "sort_column",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort_direction",
											"orig": "sort_direction",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/events/guests/list",
								"parts": []any{
									"v1",
									"events",
									"guests",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"approval_status",
										"event_id",
										"pagination_cursor",
										"pagination_limit",
										"sort_column",
										"sort_direction",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_id",
											"orig": "event_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/events/guests/get",
								"parts": []any{
									"v1",
									"events",
									"guests",
									"get",
								},
								"select": map[string]any{
									"exist": []any{
										"event_id",
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/guests/update-status",
								"parts": []any{
									"v1",
									"events",
									"guests",
									"update-status",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"guest_invite": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "guest",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "message",
						"req": false,
						"type": "`$ANY`",
						"index$": 2,
					},
				},
				"name": "guest_invite",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/guests/send-invites",
								"parts": []any{
									"v1",
									"events",
									"guests",
									"send-invites",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"guest_ticket": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "guest_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "send_email",
						"req": false,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "ticket_ids_to_remove",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "tickets_to_add",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 4,
					},
				},
				"name": "guest_ticket",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/guests/update-tickets",
								"parts": []any{
									"v1",
									"events",
									"guests",
									"update-tickets",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"host": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "access_level",
						"req": false,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "is_visible",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
				},
				"name": "host",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/hosts/add",
								"parts": []any{
									"v1",
									"events",
									"hosts",
									"add",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/hosts/remove",
								"parts": []any{
									"v1",
									"events",
									"hosts",
									"remove",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/hosts/update",
								"parts": []any{
									"v1",
									"events",
									"hosts",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"image_upload": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "content_type",
						"req": false,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "file_url",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "upload_url",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
				},
				"name": "image_upload",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/images/create-upload-url",
								"parts": []any{
									"v1",
									"images",
									"create-upload-url",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"member": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "email",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "membership_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "membership_tier_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "registration_answer",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "skip_payment",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "user_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 6,
					},
				},
				"name": "member",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/memberships/members/add",
								"parts": []any{
									"v1",
									"memberships",
									"members",
									"add",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/memberships/members/update-status",
								"parts": []any{
									"v1",
									"memberships",
									"members",
									"update-status",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"membership_tier": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "access_info",
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "description",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "tint_color",
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
				},
				"name": "membership_tier",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/memberships/tiers/list",
								"parts": []any{
									"v1",
									"memberships",
									"tiers",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"pagination_cursor",
										"pagination_limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"organization_admin": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "api_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "avatar_url",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "first_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "last_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 6,
					},
				},
				"name": "organization_admin",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/v1/organizations/admins/list",
								"parts": []any{
									"v1",
									"organizations",
									"admins",
									"list",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"organization_calendar": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "avatar_url",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "coordinate",
						"req": true,
						"type": "`$ANY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "cover_image_url",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "description",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "instagram_handle",
						"req": true,
						"type": "`$ANY`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "is_personal",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "location",
						"req": true,
						"type": "`$ANY`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "slug",
						"op": map[string]any{
							"create": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "social_image_url",
						"req": true,
						"type": "`$ANY`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "tint_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "twitter_handle",
						"req": true,
						"type": "`$ANY`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "url",
						"req": true,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "website",
						"req": true,
						"type": "`$ANY`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "youtube_handle",
						"req": true,
						"type": "`$ANY`",
						"index$": 15,
					},
				},
				"name": "organization_calendar",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v2/organizations/calendars/create",
								"parts": []any{
									"v2",
									"organizations",
									"calendars",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/organizations/calendars/list",
								"parts": []any{
									"v1",
									"organizations",
									"calendars",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"pagination_cursor",
										"pagination_limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"organization_event": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "api_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "calendar_api_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "calendar_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "coordinate",
						"req": true,
						"type": "`$ANY`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "cover_url",
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "display_price",
						"req": true,
						"type": "`$ANY`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "duration_interval",
						"req": true,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "end_at",
						"req": true,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "feedback_email",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "geo_address_json",
						"req": true,
						"type": "`$ANY`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "geo_latitude",
						"req": true,
						"type": "`$ANY`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "geo_longitude",
						"req": true,
						"type": "`$ANY`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "location_type",
						"req": true,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "location_visibility",
						"req": true,
						"type": "`$STRING`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "managing_calendar",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "meeting_url",
						"req": true,
						"type": "`$ANY`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "platform",
						"req": true,
						"type": "`$STRING`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "registration_open",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "registration_question",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 21,
					},
					map[string]any{
						"active": true,
						"name": "require_approval",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 22,
					},
					map[string]any{
						"active": true,
						"name": "spots_remaining",
						"req": true,
						"type": "`$ANY`",
						"index$": 23,
					},
					map[string]any{
						"active": true,
						"name": "start_at",
						"req": true,
						"type": "`$STRING`",
						"index$": 24,
					},
					map[string]any{
						"active": true,
						"name": "timezone",
						"req": true,
						"type": "`$STRING`",
						"index$": 25,
					},
					map[string]any{
						"active": true,
						"name": "url",
						"req": true,
						"type": "`$STRING`",
						"index$": 26,
					},
					map[string]any{
						"active": true,
						"name": "user_api_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 27,
					},
					map[string]any{
						"active": true,
						"name": "user_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 28,
					},
					map[string]any{
						"active": true,
						"name": "visibility",
						"req": true,
						"type": "`$STRING`",
						"index$": 29,
					},
					map[string]any{
						"active": true,
						"name": "waitlist_status",
						"req": true,
						"type": "`$STRING`",
						"index$": 30,
					},
					map[string]any{
						"active": true,
						"name": "zoom_meeting_url",
						"req": true,
						"type": "`$ANY`",
						"index$": 31,
					},
				},
				"name": "organization_event",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "after",
											"orig": "after",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "before",
											"orig": "before",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort_direction",
											"orig": "sort_direction",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/organizations/events/list",
								"parts": []any{
									"v1",
									"organizations",
									"events",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"after",
										"before",
										"pagination_cursor",
										"pagination_limit",
										"sort_direction",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"organization_event_transfer": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "calendar_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "event_id",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
				},
				"name": "organization_event_transfer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/organizations/events/transfer-calendar",
								"parts": []any{
									"v1",
									"organizations",
									"events",
									"transfer-calendar",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"ticket_type": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "cent",
						"req": false,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "currency",
						"req": false,
						"type": "`$ANY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "description",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "is_flexible",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "is_hidden",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "max_capacity",
						"req": false,
						"type": "`$ANY`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "min_cent",
						"req": false,
						"type": "`$ANY`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "require_approval",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "type",
						"req": true,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "valid_end_at",
						"req": false,
						"type": "`$ANY`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "valid_start_at",
						"req": false,
						"type": "`$ANY`",
						"index$": 12,
					},
				},
				"name": "ticket_type",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/ticket-types/create",
								"parts": []any{
									"v1",
									"events",
									"ticket-types",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_id",
											"orig": "event_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "include_hidden",
											"orig": "include_hidden",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/events/ticket-types/list",
								"parts": []any{
									"v1",
									"events",
									"ticket-types",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"event_id",
										"include_hidden",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "event_ticket_type_id",
											"orig": "event_ticket_type_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/events/ticket-types/get",
								"parts": []any{
									"v1",
									"events",
									"ticket-types",
									"get",
								},
								"select": map[string]any{
									"exist": []any{
										"event_ticket_type_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/ticket-types/delete",
								"parts": []any{
									"v1",
									"events",
									"ticket-types",
									"delete",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/events/ticket-types/update",
								"parts": []any{
									"v1",
									"events",
									"ticket-types",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"user": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "avatar_url",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "first_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "last_name",
						"req": true,
						"type": "`$ANY`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
				},
				"name": "user",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/v1/users/get-self",
								"parts": []any{
									"v1",
									"users",
									"get-self",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"webhook": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "event_type",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$ARRAY`",
							},
						},
						"req": true,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "secret",
						"req": true,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "url",
						"req": true,
						"type": "`$STRING`",
						"index$": 5,
					},
				},
				"name": "webhook",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v2/webhooks/create",
								"parts": []any{
									"v2",
									"webhooks",
									"create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_cursor",
											"orig": "pagination_cursor",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "pagination_limit",
											"orig": "pagination_limit",
											"reqd": false,
											"type": "`$NUMBER`",
										},
									},
								},
								"method": "GET",
								"orig": "/v1/webhooks/list",
								"parts": []any{
									"v1",
									"webhooks",
									"list",
								},
								"select": map[string]any{
									"exist": []any{
										"pagination_cursor",
										"pagination_limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.entries`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/v2/webhooks/get",
								"parts": []any{
									"v2",
									"webhooks",
									"get",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v1/webhooks/delete",
								"parts": []any{
									"v1",
									"webhooks",
									"delete",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "POST",
								"orig": "/v2/webhooks/update",
								"parts": []any{
									"v2",
									"webhooks",
									"update",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
