
# Twilio Go

![alt text](http://upload.wikimedia.org/wikipedia/commons/7/73/Burrowing_owls_in_summer.jpg)
[![Build Status](https://travis-ci.org/natebrennand/twiliogo.svg?branch=master)](https://travis-ci.org/natebrennand/twiliogo)

Write Twilio applications with our Go helper library. To get started, there are some [example applications in the examples folder](examples).

All Twilio API calls are validated at runtime to insure that all necesary information is provided.

You can generate a callback parser for all Twilio callbacks that will publish the parsed callbacks into a channel.
[See the SMS example below for an example of how this works.](#sms_example)




## Installation

```
go get github.com/natebrennand/twiliogo
```

## Usage

### Linking your account(s)

You can manage your credentials either through environment variables or manually setting them.
The library will panic if you attempt to input an invalid AccountSid or token.

#### Environment Variables

```bash
export TWILIO_ACCOUNT="accountsid"
export TWILIO_TOKEN="accounttoken"
```

```go
act := twilogo.NewAccountFromEnv()
```

#### Manual Setting

```go
act := twiliogo.NewAccount("AccountSid", "AccountToken")
```




### Making a call

```go
act := twilogo.NewAccountFromEnv()

resp, err := act.Voice.Call(voice.Post{
	From: "+{ Your Twilio Number }",
	To:   "+{ Your Destination Number }",
	URL:  "http://twimlbin.com/mytwiml",
})
```




### Sending a text - with a callback handler <a name="sms_example"></a>

```go
act := twilogo.NewAccountFromEnv()

resp, err := act.Sms.Send(sms.Post{
	From:           "+{ Your Twilio Number }",
	To:             "+{ Your Destination Number }",
	Body:           "Ready to become a Go-ru?",
	StatusCallback: "{ Your IP/DNS }/sms_callback",
})

smsChan := make(chan sms.Callback)
go func(callbackChan chan sms.Callback){
	for {
		smsCallback <- callbackChan
		// process smsCallback as desired
	}
}()

// Register the sms callback handler
// All incoming SMS callbacks will be parsed and sent into "smsChan"
http.Handle("/sms_callback", sms.CallbackHandler(smsChan))

http.ListenAndServe("0.0.0.0:8000", nil)
```

