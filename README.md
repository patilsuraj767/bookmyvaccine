# Bookmyvaccine

bookmyvaccine is a simple Go program that notifies the user whenever a vaccine slot is available. 

Disclaimer - bookmyvaccine does not book the vaccine appointment. It is just used to get notified whenever a free slot is available. 


## Installation

Currently bookmyvaccine is not shipped in the form of binary, some improvement needs to be done to ship it as binary. 

1. Make sure Golang is installed on the system.
2. Fedora user should install `alsa-lib-devel` and Debian user should have installed `libasound2`
3. Clone this repository
4. Install go dependencies

```bash
# yum install alsa-lib-devel
OR
# apt install libasound2
```
```bash
# git clone https://github.com/patilsuraj767/bookmyvaccine.git
# cd bookmyvaccine
# go mod download
# go mod tidy
```

## Usage

bookmyvaccine takes below option from stdin. 

```text
# go run bookmyvaccine.go --help
bookmyvaccine - command line for continuously checking vaccine availability

Usage:
  bookmyvaccine [flags]

Flags:
      --age_group int    Age Group can have two value 18 or 45 (default 18)
      --capacity int     Number of doses you need (default 1)
      --date string      Date should be in format DD-MM-YYYY (required)
  -h, --help             help for bookmyvaccine
      --pincode string   Pin code of you area (required)

```

General use case is searching for slot in particular area. Use below command. 

```text
# go run bookmyvaccine.go --age_group 18 --capacity 1 --pincode 431001 --date 27-05-2021
```

Once the above command is run leave the terminal idle, It will make call to Cowin website every 30sec and will check for free slot. If free slot is found then the music/siren will be played.


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)