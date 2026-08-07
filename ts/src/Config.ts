
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }


  main = {
    name: 'Luma',
  }


  feature = {
     test:     {
      "options": {
        "active": false
      }
    },

  }


  options = {
    base: 'https://public-api.luma.com',

    auth: {
      prefix: '',
    },

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      calendar: {
      },

      calendar_admin: {
      },

      calendar_coupon: {
      },

      calendar_event: {
      },

      calendar_event_approval: {
      },

      calendar_event_rejection: {
      },

      contact: {
      },

      contact_block: {
      },

      contact_restore: {
      },

      contact_tag: {
      },

      contact_tag_assignment: {
      },

      entity_lookup: {
      },

      event: {
      },

      event_cancel_request: {
      },

      event_coupon: {
      },

      event_tag: {
      },

      event_tag_assignment: {
      },

      guest: {
      },

      guest_invite: {
      },

      guest_ticket: {
      },

      host: {
      },

      image_upload: {
      },

      member: {
      },

      membership_tier: {
      },

      organization_admin: {
      },

      organization_calendar: {
      },

      organization_event: {
      },

      organization_event_transfer: {
      },

      ticket_type: {
      },

      user: {
      },

      webhook: {
      },

    }
  }


  entity = {
    "calendar": {
      "fields": [
        {
          "active": true,
          "name": "avatar_url",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "calendar_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "coordinate",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "cover_image_url",
          "req": true,
          "type": "`$ANY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "description",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "instagram_handle",
          "req": true,
          "type": "`$ANY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "is_personal",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 7
        },
        {
          "active": true,
          "name": "location",
          "req": true,
          "type": "`$ANY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "name",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "slug",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "social_image_url",
          "req": true,
          "type": "`$ANY`",
          "index$": 11
        },
        {
          "active": true,
          "name": "tint_color",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "twitter_handle",
          "req": true,
          "type": "`$ANY`",
          "index$": 13
        },
        {
          "active": true,
          "name": "url",
          "req": true,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "website",
          "req": true,
          "type": "`$ANY`",
          "index$": 15
        },
        {
          "active": true,
          "name": "youtube_handle",
          "req": true,
          "type": "`$ANY`",
          "index$": 16
        }
      ],
      "name": "calendar",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/v1/calendars/get",
              "parts": [
                "v1",
                "calendars",
                "get"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/update",
              "parts": [
                "v1",
                "calendars",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "calendar_admin": {
      "fields": [
        {
          "active": true,
          "name": "avatar_url",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "first_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "last_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 4
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        }
      ],
      "name": "calendar_admin",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/v1/calendars/admins/list",
              "parts": [
                "v1",
                "calendars",
                "admins",
                "list"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "calendar_coupon": {
      "fields": [
        {
          "active": true,
          "name": "cents_off",
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "code",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "currency",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "discount",
          "req": true,
          "type": "`$ANY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "event_ticket_type_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "percent_off",
          "req": true,
          "type": "`$ANY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "remaining_count",
          "op": {
            "create": {
              "req": false,
              "type": "`$INTEGER`"
            },
            "update": {
              "req": false,
              "type": "`$INTEGER`"
            }
          },
          "req": true,
          "type": "`$INTEGER`",
          "index$": 7
        },
        {
          "active": true,
          "name": "valid_end_at",
          "op": {
            "create": {
              "req": false,
              "type": "`$ANY`"
            },
            "update": {
              "req": false,
              "type": "`$ANY`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "valid_start_at",
          "op": {
            "create": {
              "req": false,
              "type": "`$ANY`"
            },
            "update": {
              "req": false,
              "type": "`$ANY`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 9
        }
      ],
      "name": "calendar_coupon",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/coupons/create",
              "parts": [
                "v1",
                "calendars",
                "coupons",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/calendars/coupons/list",
              "parts": [
                "v1",
                "calendars",
                "coupons",
                "list"
              ],
              "select": {
                "exist": [
                  "pagination_cursor",
                  "pagination_limit"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/coupons/update",
              "parts": [
                "v1",
                "calendars",
                "coupons",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "calendar_event": {
      "fields": [
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "status",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "submitted_by",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "tag",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 3
        }
      ],
      "name": "calendar_event",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/events/add",
              "parts": [
                "v1",
                "calendars",
                "events",
                "add"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "access",
                    "orig": "access",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "after",
                    "orig": "after",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "before",
                    "orig": "before",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "platform",
                    "orig": "platform",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_column",
                    "orig": "sort_column",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_direction",
                    "orig": "sort_direction",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/calendars/events/list",
              "parts": [
                "v1",
                "calendars",
                "events",
                "list"
              ],
              "select": {
                "exist": [
                  "access",
                  "after",
                  "before",
                  "pagination_cursor",
                  "pagination_limit",
                  "platform",
                  "sort_column",
                  "sort_direction",
                  "status"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_api_id",
                    "orig": "event_api_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_id",
                    "orig": "event_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "platform",
                    "orig": "platform",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "url",
                    "orig": "url",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/calendars/events/lookup",
              "parts": [
                "v1",
                "calendars",
                "events",
                "lookup"
              ],
              "select": {
                "exist": [
                  "event_api_id",
                  "event_id",
                  "platform",
                  "url"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.event`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "calendar_event_approval": {
      "fields": [
        {
          "active": true,
          "name": "calendar_event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        }
      ],
      "name": "calendar_event_approval",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/events/approve",
              "parts": [
                "v1",
                "calendars",
                "events",
                "approve"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "calendar_event_rejection": {
      "fields": [
        {
          "active": true,
          "name": "calendar_event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "message",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        }
      ],
      "name": "calendar_event_rejection",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/events/reject",
              "parts": [
                "v1",
                "calendars",
                "events",
                "reject"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "contact": {
      "fields": [
        {
          "active": true,
          "name": "avatar_url",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "contact",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "created_at",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "email",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "event_approved_count",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "event_checked_in_count",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 5
        },
        {
          "active": true,
          "name": "first_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "last_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "membership",
          "req": true,
          "type": "`$ANY`",
          "index$": 9
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "revenue_usd_cent",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 11
        },
        {
          "active": true,
          "name": "tag",
          "op": {
            "list": {
              "req": true,
              "type": "`$ARRAY`"
            }
          },
          "req": false,
          "type": "`$ANY`",
          "index$": 12
        },
        {
          "active": true,
          "name": "user_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 13
        }
      ],
      "name": "contact",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contacts/import",
              "parts": [
                "v1",
                "calendars",
                "contacts",
                "import"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "calendar_membership_tier_id",
                    "orig": "calendar_membership_tier_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "membership_status",
                    "orig": "membership_status",
                    "reqd": false,
                    "type": "`$ANY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "query",
                    "orig": "query",
                    "reqd": false,
                    "type": "`$ANY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_column",
                    "orig": "sort_column",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_direction",
                    "orig": "sort_direction",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ANY`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/calendars/contacts/list",
              "parts": [
                "v1",
                "calendars",
                "contacts",
                "list"
              ],
              "select": {
                "exist": [
                  "calendar_membership_tier_id",
                  "membership_status",
                  "pagination_cursor",
                  "pagination_limit",
                  "query",
                  "sort_column",
                  "sort_direction",
                  "tag"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contacts/remove",
              "parts": [
                "v1",
                "calendars",
                "contacts",
                "remove"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "contact_block": {
      "fields": [
        {
          "active": true,
          "name": "contact_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        }
      ],
      "name": "contact_block",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contacts/block",
              "parts": [
                "v1",
                "calendars",
                "contacts",
                "block"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "contact_restore": {
      "fields": [
        {
          "active": true,
          "name": "contact_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        }
      ],
      "name": "contact_restore",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contacts/restore",
              "parts": [
                "v1",
                "calendars",
                "contacts",
                "restore"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "contact_tag": {
      "fields": [
        {
          "active": true,
          "name": "color",
          "op": {
            "list": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "name",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "tag_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        }
      ],
      "name": "contact_tag",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contact-tags/create",
              "parts": [
                "v1",
                "calendars",
                "contact-tags",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/v1/calendars/contact-tags/list",
              "parts": [
                "v1",
                "calendars",
                "contact-tags",
                "list"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contact-tags/delete",
              "parts": [
                "v1",
                "calendars",
                "contact-tags",
                "delete"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contact-tags/update",
              "parts": [
                "v1",
                "calendars",
                "contact-tags",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "contact_tag_assignment": {
      "fields": [
        {
          "active": true,
          "name": "applied_count",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "skipped_count",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "tag",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "user_id",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 4
        }
      ],
      "name": "contact_tag_assignment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contact-tags/apply",
              "parts": [
                "v1",
                "calendars",
                "contact-tags",
                "apply"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/contact-tags/unapply",
              "parts": [
                "v1",
                "calendars",
                "contact-tags",
                "unapply"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "entity_lookup": {
      "fields": [],
      "name": "entity_lookup",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/entities/lookup",
              "parts": [
                "v1",
                "entities",
                "lookup"
              ],
              "select": {
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entity`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event": {
      "fields": [
        {
          "active": true,
          "name": "access",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "calendar_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "can_register_for_multiple_ticket",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 2
        },
        {
          "active": true,
          "name": "coordinate",
          "req": true,
          "type": "`$ANY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "cover_url",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            },
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "created_at",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "description",
          "req": true,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "description_md",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            },
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "display_price",
          "req": true,
          "type": "`$ANY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "duration_interval",
          "req": true,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "end_at",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            },
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "feedback_email",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 12
        },
        {
          "active": true,
          "name": "geo_address_json",
          "op": {
            "create": {
              "req": false,
              "type": "`$ANY`"
            },
            "update": {
              "req": false,
              "type": "`$ANY`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 13
        },
        {
          "active": true,
          "name": "guest_count",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 14
        },
        {
          "active": true,
          "name": "host",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 15
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "location_type",
          "req": true,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "location_visibility",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            },
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "max_capacity",
          "req": false,
          "type": "`$ANY`",
          "index$": 19
        },
        {
          "active": true,
          "name": "meeting_url",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            },
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 20
        },
        {
          "active": true,
          "name": "name",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "name_requirement",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "phone_number_requirement",
          "req": false,
          "type": "`$ANY`",
          "index$": 23
        },
        {
          "active": true,
          "name": "platform",
          "req": true,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "registration_open",
          "op": {
            "create": {
              "req": false,
              "type": "`$BOOLEAN`"
            },
            "update": {
              "req": false,
              "type": "`$BOOLEAN`"
            }
          },
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 25
        },
        {
          "active": true,
          "name": "registration_question",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 26
        },
        {
          "active": true,
          "name": "reminders_disabled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 27
        },
        {
          "active": true,
          "name": "require_approval",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 28
        },
        {
          "active": true,
          "name": "show_guest_list",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 29
        },
        {
          "active": true,
          "name": "slug",
          "req": false,
          "type": "`$STRING`",
          "index$": 30
        },
        {
          "active": true,
          "name": "spots_remaining",
          "req": true,
          "type": "`$ANY`",
          "index$": 31
        },
        {
          "active": true,
          "name": "start_at",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "suppress_notification",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 33
        },
        {
          "active": true,
          "name": "timezone",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 34
        },
        {
          "active": true,
          "name": "tint_color",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "url",
          "req": true,
          "type": "`$STRING`",
          "index$": 36
        },
        {
          "active": true,
          "name": "user_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "visibility",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            },
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 38
        },
        {
          "active": true,
          "name": "waitlist_status",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            },
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 39
        }
      ],
      "name": "event",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/create",
              "parts": [
                "v1",
                "events",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_id",
                    "orig": "event_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/events/get",
              "parts": [
                "v1",
                "events",
                "get"
              ],
              "select": {
                "exist": [
                  "event_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/cancel",
              "parts": [
                "v1",
                "events",
                "cancel"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/update",
              "parts": [
                "v1",
                "events",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event_cancel_request": {
      "fields": [
        {
          "active": true,
          "name": "cancellation_token",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "guest_count",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "is_paid",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 3
        }
      ],
      "name": "event_cancel_request",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/cancel/request",
              "parts": [
                "v1",
                "events",
                "cancel",
                "request"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event_coupon": {
      "fields": [
        {
          "active": true,
          "name": "cents_off",
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "code",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "currency",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "discount",
          "req": true,
          "type": "`$ANY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "event_ticket_type_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "percent_off",
          "req": true,
          "type": "`$ANY`",
          "index$": 7
        },
        {
          "active": true,
          "name": "remaining_count",
          "op": {
            "create": {
              "req": false,
              "type": "`$INTEGER`"
            },
            "update": {
              "req": false,
              "type": "`$INTEGER`"
            }
          },
          "req": true,
          "type": "`$INTEGER`",
          "index$": 8
        },
        {
          "active": true,
          "name": "valid_end_at",
          "op": {
            "create": {
              "req": false,
              "type": "`$ANY`"
            },
            "update": {
              "req": false,
              "type": "`$ANY`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 9
        },
        {
          "active": true,
          "name": "valid_start_at",
          "op": {
            "create": {
              "req": false,
              "type": "`$ANY`"
            },
            "update": {
              "req": false,
              "type": "`$ANY`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 10
        }
      ],
      "name": "event_coupon",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/coupons/create",
              "parts": [
                "v1",
                "events",
                "coupons",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_id",
                    "orig": "event_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/events/coupons/list",
              "parts": [
                "v1",
                "events",
                "coupons",
                "list"
              ],
              "select": {
                "exist": [
                  "event_id",
                  "pagination_cursor",
                  "pagination_limit"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/coupons/update",
              "parts": [
                "v1",
                "events",
                "coupons",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event_tag": {
      "fields": [
        {
          "active": true,
          "name": "color",
          "op": {
            "list": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "name",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "tag_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        }
      ],
      "name": "event_tag",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/event-tags/create",
              "parts": [
                "v1",
                "calendars",
                "event-tags",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/v1/calendars/event-tags/list",
              "parts": [
                "v1",
                "calendars",
                "event-tags",
                "list"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/event-tags/delete",
              "parts": [
                "v1",
                "calendars",
                "event-tags",
                "delete"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/event-tags/update",
              "parts": [
                "v1",
                "calendars",
                "event-tags",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "event_tag_assignment": {
      "fields": [
        {
          "active": true,
          "name": "applied_count",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "skipped_count",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "tag",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        }
      ],
      "name": "event_tag_assignment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/event-tags/apply",
              "parts": [
                "v1",
                "calendars",
                "event-tags",
                "apply"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/calendars/event-tags/unapply",
              "parts": [
                "v1",
                "calendars",
                "event-tags",
                "unapply"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "guest": {
      "fields": [
        {
          "active": true,
          "name": "approval_status",
          "op": {
            "create": {
              "req": false,
              "type": "`$ANY`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "check_in_qr_code",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "eth_address",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "event_ticket",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 4
        },
        {
          "active": true,
          "name": "event_ticket_order",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 5
        },
        {
          "active": true,
          "name": "guest",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "guest_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "invited_at",
          "req": true,
          "type": "`$ANY`",
          "index$": 9
        },
        {
          "active": true,
          "name": "joined_at",
          "req": true,
          "type": "`$ANY`",
          "index$": 10
        },
        {
          "active": true,
          "name": "message",
          "req": false,
          "type": "`$ANY`",
          "index$": 11
        },
        {
          "active": true,
          "name": "phone_number",
          "req": true,
          "type": "`$INTEGER`",
          "index$": 12
        },
        {
          "active": true,
          "name": "registered_at",
          "req": true,
          "type": "`$ANY`",
          "index$": 13
        },
        {
          "active": true,
          "name": "registration_answer",
          "req": true,
          "type": "`$ANY`",
          "index$": 14
        },
        {
          "active": true,
          "name": "send_email",
          "req": false,
          "type": "`$ANY`",
          "index$": 15
        },
        {
          "active": true,
          "name": "should_refund",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 16
        },
        {
          "active": true,
          "name": "solana_address",
          "req": true,
          "type": "`$ANY`",
          "index$": 17
        },
        {
          "active": true,
          "name": "status",
          "req": true,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "ticket",
          "req": false,
          "type": "`$ANY`",
          "index$": 19
        },
        {
          "active": true,
          "name": "user_email",
          "req": true,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "user_first_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 21
        },
        {
          "active": true,
          "name": "user_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "user_last_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 23
        },
        {
          "active": true,
          "name": "user_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 24
        },
        {
          "active": true,
          "name": "utm_source",
          "req": true,
          "type": "`$ANY`",
          "index$": 25
        }
      ],
      "name": "guest",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/guests/add",
              "parts": [
                "v1",
                "events",
                "guests",
                "add"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "approval_status",
                    "orig": "approval_status",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_id",
                    "orig": "event_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_column",
                    "orig": "sort_column",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_direction",
                    "orig": "sort_direction",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/events/guests/list",
              "parts": [
                "v1",
                "events",
                "guests",
                "list"
              ],
              "select": {
                "exist": [
                  "approval_status",
                  "event_id",
                  "pagination_cursor",
                  "pagination_limit",
                  "sort_column",
                  "sort_direction"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_id",
                    "orig": "event_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/events/guests/get",
              "parts": [
                "v1",
                "events",
                "guests",
                "get"
              ],
              "select": {
                "exist": [
                  "event_id",
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/guests/update-status",
              "parts": [
                "v1",
                "events",
                "guests",
                "update-status"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "guest_invite": {
      "fields": [
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "guest",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "message",
          "req": false,
          "type": "`$ANY`",
          "index$": 2
        }
      ],
      "name": "guest_invite",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/guests/send-invites",
              "parts": [
                "v1",
                "events",
                "guests",
                "send-invites"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "guest_ticket": {
      "fields": [
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "guest_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "send_email",
          "req": false,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "ticket_ids_to_remove",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "tickets_to_add",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 4
        }
      ],
      "name": "guest_ticket",
      "op": {
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/guests/update-tickets",
              "parts": [
                "v1",
                "events",
                "guests",
                "update-tickets"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "host": {
      "fields": [
        {
          "active": true,
          "name": "access_level",
          "req": false,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "is_visible",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 3
        },
        {
          "active": true,
          "name": "name",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        }
      ],
      "name": "host",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/hosts/add",
              "parts": [
                "v1",
                "events",
                "hosts",
                "add"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/hosts/remove",
              "parts": [
                "v1",
                "events",
                "hosts",
                "remove"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/hosts/update",
              "parts": [
                "v1",
                "events",
                "hosts",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "image_upload": {
      "fields": [
        {
          "active": true,
          "name": "content_type",
          "req": false,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "file_url",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "upload_url",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        }
      ],
      "name": "image_upload",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/images/create-upload-url",
              "parts": [
                "v1",
                "images",
                "create-upload-url"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "member": {
      "fields": [
        {
          "active": true,
          "name": "email",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "membership_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "membership_tier_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "registration_answer",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "skip_payment",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 4
        },
        {
          "active": true,
          "name": "status",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "user_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 6
        }
      ],
      "name": "member",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/memberships/members/add",
              "parts": [
                "v1",
                "memberships",
                "members",
                "add"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/memberships/members/update-status",
              "parts": [
                "v1",
                "memberships",
                "members",
                "update-status"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "membership_tier": {
      "fields": [
        {
          "active": true,
          "name": "access_info",
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "description",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "tint_color",
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        }
      ],
      "name": "membership_tier",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/memberships/tiers/list",
              "parts": [
                "v1",
                "memberships",
                "tiers",
                "list"
              ],
              "select": {
                "exist": [
                  "pagination_cursor",
                  "pagination_limit"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "organization_admin": {
      "fields": [
        {
          "active": true,
          "name": "api_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "avatar_url",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "email",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "first_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "last_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 5
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 6
        }
      ],
      "name": "organization_admin",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/v1/organizations/admins/list",
              "parts": [
                "v1",
                "organizations",
                "admins",
                "list"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "organization_calendar": {
      "fields": [
        {
          "active": true,
          "name": "avatar_url",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "coordinate",
          "req": true,
          "type": "`$ANY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "cover_image_url",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "description",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "instagram_handle",
          "req": true,
          "type": "`$ANY`",
          "index$": 5
        },
        {
          "active": true,
          "name": "is_personal",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 6
        },
        {
          "active": true,
          "name": "location",
          "req": true,
          "type": "`$ANY`",
          "index$": 7
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "slug",
          "op": {
            "create": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "social_image_url",
          "req": true,
          "type": "`$ANY`",
          "index$": 10
        },
        {
          "active": true,
          "name": "tint_color",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "twitter_handle",
          "req": true,
          "type": "`$ANY`",
          "index$": 12
        },
        {
          "active": true,
          "name": "url",
          "req": true,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "website",
          "req": true,
          "type": "`$ANY`",
          "index$": 14
        },
        {
          "active": true,
          "name": "youtube_handle",
          "req": true,
          "type": "`$ANY`",
          "index$": 15
        }
      ],
      "name": "organization_calendar",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v2/organizations/calendars/create",
              "parts": [
                "v2",
                "organizations",
                "calendars",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/organizations/calendars/list",
              "parts": [
                "v1",
                "organizations",
                "calendars",
                "list"
              ],
              "select": {
                "exist": [
                  "pagination_cursor",
                  "pagination_limit"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "organization_event": {
      "fields": [
        {
          "active": true,
          "name": "api_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "calendar_api_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "calendar_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "coordinate",
          "req": true,
          "type": "`$ANY`",
          "index$": 3
        },
        {
          "active": true,
          "name": "cover_url",
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "created_at",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "display_price",
          "req": true,
          "type": "`$ANY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "duration_interval",
          "req": true,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "end_at",
          "req": true,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "feedback_email",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 9
        },
        {
          "active": true,
          "name": "geo_address_json",
          "req": true,
          "type": "`$ANY`",
          "index$": 10
        },
        {
          "active": true,
          "name": "geo_latitude",
          "req": true,
          "type": "`$ANY`",
          "index$": 11
        },
        {
          "active": true,
          "name": "geo_longitude",
          "req": true,
          "type": "`$ANY`",
          "index$": 12
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "location_type",
          "req": true,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "location_visibility",
          "req": true,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "managing_calendar",
          "req": true,
          "type": "`$ARRAY`",
          "index$": 16
        },
        {
          "active": true,
          "name": "meeting_url",
          "req": true,
          "type": "`$ANY`",
          "index$": 17
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "platform",
          "req": true,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "registration_open",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 20
        },
        {
          "active": true,
          "name": "registration_question",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 21
        },
        {
          "active": true,
          "name": "require_approval",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 22
        },
        {
          "active": true,
          "name": "spots_remaining",
          "req": true,
          "type": "`$ANY`",
          "index$": 23
        },
        {
          "active": true,
          "name": "start_at",
          "req": true,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "timezone",
          "req": true,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "url",
          "req": true,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "user_api_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "user_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "visibility",
          "req": true,
          "type": "`$STRING`",
          "index$": 29
        },
        {
          "active": true,
          "name": "waitlist_status",
          "req": true,
          "type": "`$STRING`",
          "index$": 30
        },
        {
          "active": true,
          "name": "zoom_meeting_url",
          "req": true,
          "type": "`$ANY`",
          "index$": 31
        }
      ],
      "name": "organization_event",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "after",
                    "orig": "after",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "before",
                    "orig": "before",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_direction",
                    "orig": "sort_direction",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/organizations/events/list",
              "parts": [
                "v1",
                "organizations",
                "events",
                "list"
              ],
              "select": {
                "exist": [
                  "after",
                  "before",
                  "pagination_cursor",
                  "pagination_limit",
                  "sort_direction"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "organization_event_transfer": {
      "fields": [
        {
          "active": true,
          "name": "calendar_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "event_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        }
      ],
      "name": "organization_event_transfer",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/organizations/events/transfer-calendar",
              "parts": [
                "v1",
                "organizations",
                "events",
                "transfer-calendar"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "ticket_type": {
      "fields": [
        {
          "active": true,
          "name": "cent",
          "req": false,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$ANY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "is_flexible",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 4
        },
        {
          "active": true,
          "name": "is_hidden",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 5
        },
        {
          "active": true,
          "name": "max_capacity",
          "req": false,
          "type": "`$ANY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "min_cent",
          "req": false,
          "type": "`$ANY`",
          "index$": 7
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "require_approval",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 9
        },
        {
          "active": true,
          "name": "type",
          "req": true,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "valid_end_at",
          "req": false,
          "type": "`$ANY`",
          "index$": 11
        },
        {
          "active": true,
          "name": "valid_start_at",
          "req": false,
          "type": "`$ANY`",
          "index$": 12
        }
      ],
      "name": "ticket_type",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/ticket-types/create",
              "parts": [
                "v1",
                "events",
                "ticket-types",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_id",
                    "orig": "event_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "include_hidden",
                    "orig": "include_hidden",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/events/ticket-types/list",
              "parts": [
                "v1",
                "events",
                "ticket-types",
                "list"
              ],
              "select": {
                "exist": [
                  "event_id",
                  "include_hidden"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_ticket_type_id",
                    "orig": "event_ticket_type_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/events/ticket-types/get",
              "parts": [
                "v1",
                "events",
                "ticket-types",
                "get"
              ],
              "select": {
                "exist": [
                  "event_ticket_type_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/ticket-types/delete",
              "parts": [
                "v1",
                "events",
                "ticket-types",
                "delete"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/events/ticket-types/update",
              "parts": [
                "v1",
                "events",
                "ticket-types",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "user": {
      "fields": [
        {
          "active": true,
          "name": "avatar_url",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "first_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "last_name",
          "req": true,
          "type": "`$ANY`",
          "index$": 4
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        }
      ],
      "name": "user",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/v1/users/get-self",
              "parts": [
                "v1",
                "users",
                "get-self"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "webhook": {
      "fields": [
        {
          "active": true,
          "name": "created_at",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "event_type",
          "op": {
            "update": {
              "req": false,
              "type": "`$ARRAY`"
            }
          },
          "req": true,
          "type": "`$ARRAY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "id",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "secret",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "status",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "url",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        }
      ],
      "name": "webhook",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v2/webhooks/create",
              "parts": [
                "v2",
                "webhooks",
                "create"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_cursor",
                    "orig": "pagination_cursor",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "pagination_limit",
                    "orig": "pagination_limit",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v1/webhooks/list",
              "parts": [
                "v1",
                "webhooks",
                "list"
              ],
              "select": {
                "exist": [
                  "pagination_cursor",
                  "pagination_limit"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.entries`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/v2/webhooks/get",
              "parts": [
                "v2",
                "webhooks",
                "get"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v1/webhooks/delete",
              "parts": [
                "v1",
                "webhooks",
                "delete"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/v2/webhooks/update",
              "parts": [
                "v2",
                "webhooks",
                "update"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config
}

