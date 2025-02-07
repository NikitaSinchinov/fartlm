package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	api "main/api_server"
)

type Server struct {
	audioProcessor AudioProcessor
}

func NewServer(audioProcessor AudioProcessor) *Server {
	return &Server{audioProcessor: audioProcessor}
}

func (s *Server) Stop() {
	s.audioProcessor.Stop()
}

// (GET /health)
func (s *Server) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	return &api.GetHealth200Response{}, nil
}

// (POST /fart)
func (s *Server) ProcessAudio(ctx context.Context, request api.ProcessAudioRequestObject) (api.ProcessAudioResponseObject, error) {
	if request.Body == nil {
		return &api.ProcessAudio400Response{}, fmt.Errorf("no multipart form data received")
	}

	// Read the audio file from the multipart form
	audio, err := request.Body.NextPart()
	if err != nil {
		return &api.ProcessAudio400Response{}, fmt.Errorf("error reading multipart form: %w", err)
	}
	if audio.FormName() != "audio" {
		return &api.ProcessAudio400Response{}, fmt.Errorf("expected form field 'audio', got '%s'", audio.FormName())
	}

	// Process the audio
	output, err := s.audioProcessor.ProcessAudio(audio)
	if err != nil {
		return &api.ProcessAudio500Response{}, fmt.Errorf("error processing audio: %w", err)
	}

	// Read the content of output into a byte slice
	outputBytes, err := io.ReadAll(output)
	if err != nil {
		return &api.ProcessAudio500Response{}, fmt.Errorf("error reading processed audio: %w", err)
	}

	return &api.ProcessAudio200AudiowavResponse{
		Body:          bytes.NewReader(outputBytes),
		ContentLength: int64(len(outputBytes)),
	}, nil
}
