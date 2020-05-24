package usi

// FirstUSIRequest は、 `usi` で受け取るリクエストです。
type FirstUSIRequest struct {
}

// FirstUSIResponse は、 `usi` に対するレスポンスです。
type FirstUSIResponse struct {
	Name    string
	Author  string
	Options []*Option
}

// IsReadyRequest は、 `isready` で受け取るリクエストです。
type IsReadyRequest struct {
}

// IsReadyResponse は、 `isready` に対するレスポンスです。
type IsReadyResponse struct {
	IsReady bool
}

// SetOptionRequest は、 `setoption` で受け取るリクエストです。
type SetOptionRequest struct {
	ID    string
	Value string
}

// SetOptionResponse は、 `setoption` に対するレスポンスです。
type SetOptionResponse struct {
}

// NewGameRequest は、 `usinewgame` で受け取るリクエストです。
type NewGameRequest struct {
}

// NewGameResponse は、 `usinewgame` に対するレスポンスです。
type NewGameResponse struct {
}

// SetPositionRequest は、 `position` で受け取るリクエストです。
type SetPositionRequest struct {
}

// SetPositionResponse は、 `position` に対するレスポンスです。
type SetPositionResponse struct {
}

// GoRequest は、 `go` で受け取るリクエストです。
type GoRequest struct {
}

// GoResponse は、 `go` に対するレスポンスです。
type GoResponse struct {
}

// StopRequest は、 `stop` で受け取るリクエストです。
type StopRequest struct {
}

// StopResponse は、 `stop` に対するレスポンスです。
type StopResponse struct {
}

// PonderhitRequest は、 `ponderhit` で受け取るリクエストです。
type PonderhitRequest struct {
}

// PonderhitResponse は、 `ponderhit` に対するレスポンスです。
type PonderhitResponse struct {
}

// QuitRequest は、 `quit` で受け取るリクエストです。
type QuitRequest struct {
}

// QuitResponse は、 `quit` に対するレスポンスです。
type QuitResponse struct {
}

// GameOverRequest は、 `gameover` で受け取るリクエストです。
type GameOverRequest struct {
}

// GameOverResponse は、 `gameover` に対するレスポンスです。
type GameOverResponse struct {
}
