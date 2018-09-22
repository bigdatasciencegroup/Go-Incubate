# Train in Python and Infer in Golang

+ **Python** Better machine learning library support
+ **Python** Frameworks like Keras, make writing deep learning algorithms easier
+ **Python** Easier to visualize data
+ **Golang** Makes small efficient binaries
+ **Golang** Better web frameworks
+ **Golang** Concise, staticaly typed, and compiled, language
+ **Golang** Fast compile time
+ **Golang** Great concurrency features

## Steps

1. Build, train, and test the model using TensorFlow/Keras in Python
2. Name the layers in the model to enable retrieval later
3. In python code, instead of this

    ```python
    #Save model to continue later
    model.save('./cnn_lstm.hdf5')
    #Load previously saved model to continue testing
    model = load_model('./cnn_lstm.hdf5')
    ```
    do this
    ```python
    # Use TF to save the model instead of Keras, in order to load it in Golang later
    builder = tf.saved_model.builder.SavedModelBuilder("export_dir")  
    builder.add_meta_graph_and_variables(sess, ["myTag"])
    builder.save()
    ```
4. In Go, load the model and run inference

    ```go
    import (
        "fmt"
        tf "github.com/tensorflow/tensorflow/tensorflow/go"
    )

    func main() {  
        // Load the model in golang
        model, _ := tf.LoadSavedModel("export_dir", []string{"myTag"}, nil)

        defer model.Session.Close()

        // Create dummy tensor
        dummyData := [][]float64{make([]float64, 448)}
        tensor, _ := tf.NewTensor(dummyData)

        // Run the model
        result, _ := model.Session.Run(
            map[tf.Output]*tf.Tensor{
                model.Graph.Operation("nameOfInputLayer").Output(0): tensor,
            },
            []tf.Output{
                model.Graph.Operation("nameOfInferenceLayer").Output(0),
            },
            nil,
        )

        fmt.Printf("Result value: %v \n", result[0].Value())
    }

    ```

5. Tensorflow Go bindings only supported in Linux and Mac
