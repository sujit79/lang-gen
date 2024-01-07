# Language Generator
Intends to provide weighted interface (Sentiment) for text which might be large in size and gives an relative weight for the document.

# Design
Designed to parse large text and provide weight or Sentiment of Text.

# Build the Project
1. git clone https://github.com/sujit79/lang-gen.git <br/>
2. cd lang-gen <br/>
3. go build <br/>

# Usage
1. go run weighted-lang.go $(FILE_LOCATION) $(FILE_LOCATION).weight OR <br/>
2. Linux weighted-lang $(FILE_LOCATION) $(FILE_LOCATION).weight OR <br/>
3. Windows weighted-lang.exe $(FILE_LOCATION) $(FILE_LOCATION).weight OR <br/>

# Output
Shared as $(PATH_TO_FILE).weight <br/>

