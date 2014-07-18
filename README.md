# Twilio Go
![alt text](http://upload.wikimedia.org/wikipedia/commons/7/73/Burrowing_owls_in_summer.jpg)


[![Build Status](https://travis-ci.org/natebrennand/twiliogo.svg?branch=master)](https://travis-ci.org/natebrennand/twiliogo)

## Installation
    go get github.com:natebrennand/twiliogo.git
    
## Usage

### Add your credentials
    #secrets.sh
    export TWILIO_ACCOUNT="accountsid"
    export TWILIO_TOKEN="accounttoken"
```
source secrets.sh
```
### Set up your account
    act := twiliogo.NewAccountFromEnv()

### Make a call
    resp, err := act.Voice.Call(voice.Post{
		  From: "+{ Your Twilio Number }",
		  To:   "+{ Your Destination Number }",
		  URL:  "http://twimlbin.com/mytwiml",
    })
	  
### Send a text - with a callback handler
    resp, err := act.Sms.Send(sms.Post{
  		From:           "+{ Your Twilio Number }",
  		To:             "+{ Your Destination Number }",
  		Body:           "Ready to become a Go-ru?",
  		StatusCallback: "{ Your callback endpoint }",
    })
  	
  	cp := make(chan sms.Callback)
  	
  	// Register the sms callback handler
  	http.Handle("/", sms.CallbackHandler(cp))
  	http.ListenAndServe("0.0.0.0:8000", nil)
  
