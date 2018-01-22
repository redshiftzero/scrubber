# scrubber
[![CircleCI](https://circleci.com/gh/redshiftzero/scrubber.svg?style=svg&circle-token=612ce2eb2ba545a51c7f0e73d4def1f49b431cdf)](https://circleci.com/gh/redshiftzero/scrubber)

Simple CLI to remove metadata from images

```
$ scrubber
Usage of scrubber:
  -input string
    	Image file to scrub metadata from
  -json
    	Print JSON metadata to stdout
  -output string
    	Output file of cleaned image
```

## Install

Currently you need `libexif` installed:

```
apt-get install libexif-dev
```

Or on Mac OS:

```
brew install libexif
```

Then you can:

```
go get github.com/redshiftzero/scrubber
```
