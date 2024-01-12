package controller

import (
	"golang_template/helper"
	"golang_template/schemas"
	"golang_template/service"
	"golang_template/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Controller struct {
	Config        *util.Config
	ChatService   service.IChatService
	CardService   service.ICardService
	PromptService service.IPromptService
	Validator     *validator.Validate
}

// Chat Need header godoc
// @Summary	Chat message
// @Schemes
// @Description	Chat message
// @Tags			Chat
// @Accept			json
// @Produce			json
// @Param RequestBody body schemas.InputMessage true "Chat message"
// @Success	200	{object} string
// @Router /chat [post]
func (controller *Controller) Chat(ctx *gin.Context) {
	var request schemas.InputMessage
	// Get user input message
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusBadRequest, err, "Dữ liệu không hợp lệ")
		return
	}
	// Validate user input message
	err = controller.Validator.Struct(request)
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusBadRequest, err, "Dữ liệu không hợp lệ")
		return
	}

	// Load final prompt, user prompt, model prompt
	finalPrompt, err := controller.PromptService.GetFinalPrompt()
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lấy dữ liệu")
		return
	}
	userPrompt, err := controller.PromptService.GetUserPrompt()
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lấy dữ liệu")
		return
	}
	modelPrompt, err := controller.PromptService.GetModelPrompt()
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lấy dữ liệu")
		return
	}
	//Get cards from database
	cards, err := controller.CardService.GetCards()
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lấy dữ liệu")
		return
	}
	// Pick random cards
	randomCards, err := controller.CardService.PickRandomCards(cards, 5)
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lấy dữ liệu")
		return
	}
	response, err := controller.ChatService.GetGeneratedPrompt(request.Message, finalPrompt, userPrompt, modelPrompt, randomCards)
	if err != nil {
		// If get 400 error from Gemini Pro API, return error message
		if err.Error() == "400" {
			helper.NewHTTPError(ctx, http.StatusBadRequest, err, "Dữ liệu không hợp lệ")
			// Add user input message to database with block = true
			chatID, err := controller.ChatService.CreateChatInput(request.Message, true)
			if err != nil {
				helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lưu dữ liệu")
				return
			}
			// Add generated prompt to database with block = true
			err = controller.ChatService.CreateChatOutput(response, strings.Join(randomCards, ", "), chatID, true)
			if err != nil {
				helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lưu dữ liệu")
				return
			}
			return
		} else {
			helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lấy dữ liệu")
			return
		}
	}
	// Add user input message to database with block = false
	chatID, err := controller.ChatService.CreateChatInput(request.Message, false)
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lưu dữ liệu")
		return
	}
	// Add generated prompt to database with block = false
	err = controller.ChatService.CreateChatOutput(response, strings.Join(randomCards, ", "), chatID, false)
	if err != nil {
		helper.NewHTTPError(ctx, http.StatusInternalServerError, err, "Lỗi hệ thống khi lưu dữ liệu")
		return
	}

	// Response with status success
	log.Info().Msg("Tạo response thành công")
	helper.NewHTTPResponse(ctx, http.StatusOK, response, "Lấy kết quả thành công")
}
