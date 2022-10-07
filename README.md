## Scan and unzip zip files 

#### It finds zip files in the directory and subfolders you give with the source path and extracts them to the directory you specify.

#### RUN
``./unzip-linux --output="OUTPUT_PATH" --source="SOURCE_PATH"``


#### Build Linux:
``env GOOS=linux GOARCH=amd64 go build -v -o build/unzipp-linux``