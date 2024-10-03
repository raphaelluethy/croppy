# croppy

`croppy` is a command-line tool for batch cropping images based on specified parameters. It allows you to process multiple images at once, applying consistent cropping rules across your image set.

> [!WARNING]
> This program is still in development and is not yet fully functional. The images anonymizer is working, but the video anonymizer was not verified yet.

## Features

- Batch processing of images or videos
- Customizable cropping parameters
- Support for various image formats (JPEG, PNG, etc.)
- Support for MP4 video format (WIP)
- Video processing with frame extraction and reassembly (WIP)

## Flags

- `--top`: Top crop in pixels
- `--right`: Right crop in pixels
- `--bottom`: Bottom crop in pixels
- `--left`: Left crop in pixels
- `--path`: Path to the image or video file
- `--filetypes`: Filetypes to process (default: png,jpg,jpeg,mp4)
- `--video`: Flag to enable video processing mode, only works with filetypes `mp4` and will ignore all other filetypes

## Installation

To install `croppy`, make sure you have Go installed on your system, then clone this repository and build the project:

```bash
git clone https://github.com/raphaelluethy/croppy.git
cd croppy
go build -o croppy
```

You also need to have `ffmpeg` installed on your system.

## Usage

### For Images:

```bash
./croppy --top 10 --left 10 --right 100 --bottom 40 --path ./data   
```

This command will replace the pixels at the top, left, right and bottom of the image with purple pixels.

**Input**

![Input](./data/image.png)

**Output**

![Output](./output/out_image.png)

### For Videos:

```bash
./croppy --video --path ./data --output ./output --top 10 --left 10 --right 100 --bottom 40
```

This command will process the MP4 video, applying the same cropping parameters to each frame.

> [!WARNING]
> this program assumes the images are in 25fps, if the fps is different you need to change it in the code

## Video Processing Options

- `--video`: Flag to enable video processing mode
- `--input`: Path to the input video file
- `--output`: Path for the output processed video file
- `--fps`: Frames per second for processing (default: 30)
- `--start`: Start time for processing (format: HH:MM:SS)
- `--end`: End time for processing (format: HH:MM:SS)
