# imgcrop

`imgcrop` is a command-line tool for batch cropping images based on specified parameters. It allows you to process multiple images at once, applying consistent cropping rules across your image set.

## Features

- Batch processing of images
- Customizable cropping parameters
- Support for various image formats (JPEG, PNG, etc.)
- Option to anonymize images by blurring faces

## Installation

To install `imgcrop`, make sure you have Go installed on your system, then clone this repository and build the project:

```bash
git clone https://github.com/raphaelluethy/imgcrop.git
cd imgcrop
go build -o imgcrop
```

## Usage

```bash
./imgcrop --top 10 --left 10 --right 100 --bottom 40 --path ./data   
```

This command will replace the pixels at the top, left, right and bottom of the image with purple pixels.

**Input**

![Input](./data/image.png)

**Output**

![Output](./output/out_image.png)
