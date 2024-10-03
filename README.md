# croppy

`croppy` is a command-line tool for batch cropping images based on specified parameters. It allows you to process multiple images at once, applying consistent cropping rules across your image set.

## Features

- Batch processing of images
- Customizable cropping parameters
- Support for various image formats (JPEG, PNG, etc.)
- Option to anonymize images by blurring faces

## Installation

To install `croppy`, make sure you have Go installed on your system, then clone this repository and build the project:

```bash
git clone https://github.com/raphaelluethy/croppy.git
cd croppy
go build -o croppy
```

## Usage

```bash
./croppy --top 10 --left 10 --right 100 --bottom 40 --path ./data   
```

This command will replace the pixels at the top, left, right and bottom of the image with purple pixels.

**Input**

![Input](./data/image.png)

**Output**

![Output](./output/out_image.png)
