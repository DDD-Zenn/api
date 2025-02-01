package command

type CreateChatCommand struct {
	Prompt string `json:"prompt"`
}

func NewChatCreate(cmd CreateChatCommand) (CreateChatCommand, error) {
	// if err := c.ShouldBindJSON(&cmd); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return err
	// }

	// バリデーションなどはここで実装
	//例
	// if cmd.Prompt == "" {
	// return cmd, utils.ERR_REQUIRED_PROMPT
	// }

	return cmd, nil
}
