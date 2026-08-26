package miniapp

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-telegram/bot"
	"go.uber.org/zap"

	"github.com/Vlad-6894/Activity_tracking/internal/core/logger"
)

type Handler struct {
	botToken string
	log      *logger.Logger

	assets     fs.FS
	fileServer http.Handler
}

func NewHandler(botToken string, log *logger.Logger) *Handler {
	assets := os.DirFS("web/public")

	return &Handler{
		botToken:   botToken,
		log:        log,
		assets:     assets,
		fileServer: http.FileServer(http.FS(assets)),
	}
}

type meResponse struct {
	OK        bool   `json:"ok"`
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Username  string `json:"username,omitempty"`
	Error     string `json:"error,omitempty"`
}

func (h *Handler) index(w http.ResponseWriter, _ *http.Request) {
	page, err := fs.ReadFile(h.assets, "index.html")
	if err != nil {
		h.log.Error("read embedded index.html", zap.Error(err))
		http.Error(w, "mini app is unavailable", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	if _, err := w.Write(page); err != nil {
		h.log.Warn("write index.html", zap.Error(err))
	}
}

func (h *Handler) ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("pong"))
}

func (h *Handler) me(w http.ResponseWriter, r *http.Request) {
	initData := strings.TrimSpace(strings.TrimPrefix(r.Header.Get("Authorization"), "tma "))
	if initData == "" {
		h.writeMe(w, http.StatusUnauthorized, meResponse{Error: "нет initData в заголовке Authorization"})
		return
	}

	values := parseInitData(initData)
	if len(values) == 0 {
		h.writeMe(w, http.StatusBadRequest, meResponse{Error: "initData не разбирается как query string"})
		return
	}

	tgUser, ok := bot.ValidateWebappRequest(values, h.botToken)
	if !ok {
		h.log.Warn("init data hmac mismatch")
		h.writeMe(w, http.StatusUnauthorized, meResponse{Error: "подпись initData не сошлась"})
		return
	}

	h.writeMe(w, http.StatusOK, meResponse{
		OK:        true,
		ID:        tgUser.ID,
		FirstName: tgUser.FirstName,
		Username:  tgUser.Username,
	})
}

func parseInitData(raw string) url.Values {
	values := url.Values{}

	for pair := range strings.SplitSeq(raw, "&") {
		key, value, found := strings.Cut(pair, "=")
		if !found || key == "" {
			continue
		}

		values.Set(key, value)
	}

	return values
}

func (h *Handler) writeMe(w http.ResponseWriter, status int, body meResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		h.log.Warn("encode me response", zap.Error(err))
	}
}
