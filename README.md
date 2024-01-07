# Language Generator
Intends to provide weighted interface (Sentiment) for text which might be large in size and gives an relative weight for the document.

# Design
Designed to parse large text and provide weight or Sentiment of Text.

# Build the Project
1. git clone https://github.com/sujit79/lang-gen.git <br/>
2. cd lang-gen <br/>
3. go build <br/>

# Usage
1. go run weighted-lang.go $ARGUMENT_FILE_LOCATION $ARGUMENT_FILE_LOCATION.weight <br/>
                         **OR** <br/>
2. Linux <br/>
      weighted-lang $ARGUMENT_FILE_LOCATION $ARGUMENT_FILE_LOCATION.weight <br/>
                         **OR** <br/>
3. Windows <br/>
      weighted-lang.exe $ARGUMENT_FILE_LOCATION $ARGUMENT_FILE_LOCATION.weight <br/>

# Output
Shared as $ARGUMENT_FILE_LOCATION.weight <br/>

