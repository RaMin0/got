package got

// PluginFacebookProfile type
type PluginFacebookProfile struct {
	ComponentBase
}

// Execute func
func (c *PluginFacebookProfile) Execute(ctx *BotContext) error {
	profile, err := c.client().GetProfile(ctx.userID)
	if err != nil {
		return err
	}

	ctx.Interpolations()["fb:first_name"] = profile.FirstName
	ctx.Interpolations()["fb:last_name"] = profile.LastName
	ctx.Interpolations()["fb:gender"] = profile.Gender
	ctx.Interpolations()["fb:picture_url"] = profile.PictureURL
	ctx.Interpolations()["fb:locale"] = profile.Locale
	ctx.Interpolations()["fb:timezone"] = profile.Timezone

	return nil
}
