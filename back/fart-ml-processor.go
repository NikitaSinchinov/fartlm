package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"go-shared/concurrency"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	api "main/api_client"

	"github.com/gofiber/fiber/v2/log"
)

type AudioProcessor interface {
	ProcessAudio(input io.Reader) (io.Reader, error)
	Stop()
}

type FartMLAudioProcessor struct {
	client     *api.ClientWithResponses
	workerPool *concurrency.WorkerPool
}

func NewFartMLAudioProcessor(client *api.ClientWithResponses) *FartMLAudioProcessor {
	return &FartMLAudioProcessor{
		client:     client,
		workerPool: concurrency.NewDefaultWorkerPool(),
	}
}

func (f *FartMLAudioProcessor) Stop() {
	f.workerPool.Stop()
}

func (f *FartMLAudioProcessor) ProcessAudio(input io.Reader) (io.Reader, error) {
	output, err := f.makeFartAudioWithDeadline(input)
	if err != nil {
		log.Error(err)
		return f.randomFartSound()
	}
	return output, nil
}

const (
	executionTimeout = 10 * time.Second
)

func (f *FartMLAudioProcessor) makeFartAudioWithDeadline(input io.Reader) (io.Reader, error) {
	var output io.Reader
	var err error

	deadlineErr := concurrency.SyncWithDeadline(executionTimeout, func(ctx context.Context) {
		f.workerPool.SyncSubmit(func() {
			select {
			case <-ctx.Done():
			default:
				output, err = f.makeFartAudio(ctx, input)
			}
		})
	})
	if deadlineErr != nil {
		return nil, fmt.Errorf("error with deadline: %w", deadlineErr)
	}

	return output, err
}

const (
	bassBoost = 50
	deNoise   = -25
)

func (f *FartMLAudioProcessor) makeFartAudio(ctx context.Context, input io.Reader) (io.Reader, error) {
	preprocessed, err := f.makeBassBoosted(fmt.Sprintf("bass=g=%d", bassBoost), input)
	if err != nil {
		return nil, fmt.Errorf("error bass boost: %w", err)
	}

	converted, err := f.requestConvert(ctx, preprocessed)
	if err != nil {
		return nil, fmt.Errorf("error converting: %w", err)
	}

	postProcessed, err := f.makeBassBoosted(fmt.Sprintf("afftdn=nf=%d", deNoise), converted)
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}

	return postProcessed, nil
}

func makeTempFilename(prefix string) string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("%s-%d.wav", prefix, time.Now().UnixNano()))
}

func (f *FartMLAudioProcessor) makeBassBoosted(audioFilter string, input io.Reader) (io.Reader, error) {
	inputFile, err := os.Create(makeTempFilename("bass-boost-input"))
	if err != nil {
		return nil, fmt.Errorf("error creating input file: %w", err)
	}
	defer func() {
		inputFile.Close()
		os.Remove(inputFile.Name())
	}()

	_, err = io.Copy(inputFile, input)
	if err != nil {
		return nil, fmt.Errorf("error copying input: %w", err)
	}

	outputFilename := makeTempFilename("bass-boost-output")

	err = f.runBassBoost(audioFilter, inputFile.Name(), outputFilename)
	if err != nil {
		return nil, fmt.Errorf("error running bass boost: %w", err)
	}

	outputFile, err := os.Open(outputFilename)
	if err != nil {
		return nil, fmt.Errorf("error reopening output file: %w", err)
	}
	defer func() {
		outputFile.Close()
		os.Remove(outputFile.Name())
	}()

	outputBytes, err := io.ReadAll(outputFile)
	if err != nil {
		return nil, fmt.Errorf("error reading output file: %w", err)
	}
	output := bytes.NewReader(outputBytes)

	return output, nil
}

const (
	maxDurationSeconds = 10
)

func (f *FartMLAudioProcessor) runBassBoost(audioFilter string, inputFilename string, outputFilename string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputFilename,
		"-af", audioFilter,
		"-t", fmt.Sprint(maxDurationSeconds),
		outputFilename,
	)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("ffmpeg error: %v, %s", err, stderr.String())
	}
	return nil
}

func (f *FartMLAudioProcessor) requestConvert(ctx context.Context, input io.Reader) (io.Reader, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, input)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}
	base64Input := base64.StdEncoding.EncodeToString(buf.Bytes())

	resp, err := f.client.PostConvertWithResponse(ctx, api.PostConvertJSONRequestBody{AudioData: &base64Input})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("received nil response or body")
	}
	if statusCode := resp.HTTPResponse.StatusCode; statusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	return bytes.NewReader(resp.Body), nil
}

const (
	soundsDir = "/back/sounds/"
)

func (f *FartMLAudioProcessor) randomFartSound() (io.Reader, error) {
	soundFilename, err := f.randomFartSoundFilename()
	if err != nil {
		return nil, fmt.Errorf("error getting random fart sound filename: %w", err)
	}
	return os.Open(soundFilename)
}

func (f *FartMLAudioProcessor) randomFartSoundFilename() (string, error) {
	files, err := os.ReadDir(soundsDir)
	if err != nil {
		return "", fmt.Errorf("error reading directory: %w", err)
	}

	var wavFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".wav" {
			wavFiles = append(wavFiles, file.Name())
		}
	}

	if len(wavFiles) == 0 {
		return "", fmt.Errorf("no .wav files found in directory")
	}

	randomSound := wavFiles[rand.Intn(len(wavFiles))]

	return filepath.Join(soundsDir, randomSound), nil
}
