package main

import (
	"djeurnie/ai/trader/data"
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	today := time.Now().UTC()
	yesterday := today.AddDate(0, 0, -1)

	for i := 0; i < 2190; i++ {
		from := yesterday.AddDate(0, 0, -i)
		to := from.AddDate(0, 0, 1)

		data.CreateData(from, to)

		// the binance api has a rate limit of 1200req/min. We do a simple sleep here to not hit that rate limit
		time.Sleep(10 * time.Millisecond)
	}
}

func doTraining() {
	// Form the training matrices.
	inputs, labels := makeInputsAndLabels("data/train.csv")

	// Define our network architecture and learning parameters.
	config := NeuralNetConfig{
		inputNeurons:  540,
		outputNeurons: 2,
		hiddenNeurons: 360,
		epochs:        5000,
		learningRate:  0.3,
	}

	// Train the neural network.
	network := NewNeuralNet(config)
	if err := network.train(inputs, labels); err != nil {
		log.Fatal(err)
	}

	network.save("data/nn.json")

	// Form the testing matrices.
	testInputs, testLabels := makeInputsAndLabels("data/test.csv")

	// Make the predictions using the trained model.
	predictions, err := network.predict(testInputs)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the accuracy of our model.
	var truePosNeg int
	numPreds, _ := predictions.Dims()
	for i := 0; i < numPreds; i++ {

		// Get the label.
		labelRow := mat.Row(nil, i, testLabels)
		var prediction int
		for idx, label := range labelRow {
			if label == 1.0 {
				prediction = idx
				break
			}
		}

		// Accumulate the true positive/negative count.
		if predictions.At(i, prediction) == floats.Max(mat.Row(nil, i, predictions)) {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy).
	accuracy := float64(truePosNeg) / float64(numPreds)

	// Output the Accuracy value to standard out.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

func makeInputsAndLabels(fileName string) (*mat.Dense, *mat.Dense) {
	// Open the dataset file.
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 7

	// Read in all the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// inputsData and labelsData will hold all the
	// float values that will eventually be
	// used to form matrices.
	inputsData := make([]float64, 4*len(rawCSVData))
	labelsData := make([]float64, 3*len(rawCSVData))

	// Will track the current index of matrix values.
	var inputsIndex int
	var labelsIndex int

	// Sequentially move the rows into a slice of floats.
	for idx, record := range rawCSVData {

		// Skip the header row.
		if idx == 0 {
			continue
		}

		// Loop over the float columns.
		for i, val := range record {

			// Convert the value to a float.
			parsedVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Add to the labelsData if relevant.
			if i == 4 || i == 5 || i == 6 {
				labelsData[labelsIndex] = parsedVal
				labelsIndex++
				continue
			}

			// Add the float value to the slice of floats.
			inputsData[inputsIndex] = parsedVal
			inputsIndex++
		}
	}
	inputs := mat.NewDense(len(rawCSVData), 4, inputsData)
	labels := mat.NewDense(len(rawCSVData), 3, labelsData)
	return inputs, labels
}
