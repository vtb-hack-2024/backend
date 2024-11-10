package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initInfoRouter(api *gin.RouterGroup) {
	info := api.Group("/info", h.userIdentity)
	{
		info.GET("/getname", h.getName)
		info.GET("/getamount", h.getAmount)
		info.GET("/getachievements", h.getAchievements)
		info.GET("/getbaseinfo", h.getBaseInfo)
		info.GET("/getneuromean", h.getNeuroMean)
		info.GET("/getcryptodata", h.getCryptoData)
		info.GET("/getapiinfo", h.getAPIInfo)
		info.GET("/getfullapiinfo", h.getFullAPIInfo)
		info.GET("/getfines", h.getFines)
		info.GET("/getfine", h.getFineByID)
		info.GET("/getpayments", h.getPayments)
		info.GET("/getpayment", h.getPaymentByID)
		info.GET("/getstatsdata", h.getStatsData)
		info.GET("/getanalize", h.getAnalyze)
	}
}

// @Summary Get User Name
// @Description Retrieves the user's name
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getname [get]
func (h *Handler) getName(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	name, err := h.service.GetName(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": name})
}

// @Summary Get Fine Amount
// @Description Retrieves the amount of the fine
// @Tags Fine
// @Accept json
// @Produce json
// @Success 200 {object} int
// @Router /getamount [get]
func (h *Handler) getAmount(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	amount, err := h.service.GetAmount(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"amount": amount})
}

// @Summary Get User Achievements
// @Description Retrieves the user's achievements
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getachievements [get]
func (h *Handler) getAchievements(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	achievements, err := h.service.GetAchievements(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"achievements": achievements})
}

// @Summary Get Base Profile Information
// @Description Retrieves basic profile information
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getbaseinfo [get]
func (h *Handler) getBaseInfo(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	baseInfo, err := h.service.GetBaseInfo(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"baseInfo": baseInfo})
}

// @Summary Get Neuro Mean Score
// @Description Retrieves the neuro mean score by name
// @Tags Neuro
// @Accept json
// @Produce json
// @Success 200 {object} float64
// @Router /getneuromean [get]
func (h *Handler) getNeuroMean(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	neuroMean, err := h.service.GetNeuroMean(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"neuroMean": neuroMean})
}

// @Summary Get Crypto Data
// @Description Retrieves cryptocurrency data
// @Tags Crypto
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getcryptodata [get]
func (h *Handler) getCryptoData(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	cryptoData, err := h.service.GetCryptoData(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"cryptoData": cryptoData})
}

// @Summary Get API Short Info
// @Description Retrieves short information about the API
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getapiinfo [get]
func (h *Handler) getAPIInfo(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	apiInfo, err := h.service.GetAPIInfo(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"apiInfo": apiInfo})
}

// @Summary Get Full API Info
// @Description Retrieves full information about the API by ID
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getfullapiinfo [get]
func (h *Handler) getFullAPIInfo(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fullAPIInfo, err := h.service.GetFullAPIInfo(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"fullAPIInfo": fullAPIInfo})
}

// @Summary Get User Fines
// @Description Retrieves all user's fines
// @Tags Fine
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getfines [get]
func (h *Handler) getFines(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fines, err := h.service.GetFines(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"fines": fines})
}

// @Summary Get Fine by ID
// @Description Retrieves a specific fine by its ID
// @Tags Fine
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getfine [get]
func (h *Handler) getFineByID(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fine, err := h.service.GetFineByID(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"fine": fine})
}

// @Summary Get User Payments
// @Description Retrieves all user's payments
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getpayments [get]
func (h *Handler) getPayments(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	payments, err := h.service.GetPayments(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"payments": payments})
}

// @Summary Get Payment by ID
// @Description Retrieves a specific payment by its ID
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getpayment [get]
func (h *Handler) getPaymentByID(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	payment, err := h.service.GetPaymentByID(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

// @Summary Get Stats Data
// @Description Retrieves statistical data for the graph
// @Tags Stats
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /getstatsdata [get]
func (h *Handler) getStatsData(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	statsData, err := h.service.GetStatsData(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"statsData": statsData})
}

// @Summary Get User Analyze Data
// @Description Retrieves analytical data of the user
// @Tags Analyze
// @Accept json
// @Produce json
// @Success 200 {object} string
func (h *Handler) getAnalyze(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	statsData, err := h.service.GetStatsData(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"statsData": statsData})
}
