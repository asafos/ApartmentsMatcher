package web

// func Index(session *session.Session, db *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		auth := IsAuthenticated(session, ctx)

// 		// Bind data to template
// 		bind := fiber.Map{
// 			"name": "Fiber",
// 			"auth": auth,
// 		}

// 		if auth {
// 			store := session.Get(ctx)
// 			// Get User ID from session store
// 			userID, _ := store.Get("userid").(int64)
// 			user, err := FindUserByID(db, userID)
// 			if err != nil {
// 				return utils.SendError(ctx, "Error when finding user by ID", fiber.StatusInternalServerError)
// 			}
// 			bind["username"] = user.Name
// 		}

// 		// Render template
// 		err := ctx.Render("index", bind)
// 		if err != nil {
// 			err2 := ctx.Status(500).SendString(err.Error())
// 			if err2 != nil {
// 				return utils.SendError(ctx, err2.Error(), fiber.StatusInternalServerError)
// 			}
// 		}
// 		return err
// 	}
// }
