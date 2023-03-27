# MailGun configuration
Steps to setup Mailgun email service in order send notifications for Kata #3.
Yes, Yes, yes: even if goal of the Kata is to present a proof of concept, what if delivering stuff that works for real ?

## Create a mailGun Account

Go to https://www.mailgun.com/ and create an account

## Get your domain an API Key

1. Log to https://www.mailgun.com/ and go to 'Domains' panel
2. Select your sandbox domain in the list
3. On question 'How would you like to send your emails from sandbox.....mailgun.org?' choose 'API' answer
4. Your API Key is displayed

## Create environment variables for your domain and API Key
In your .zshrc or .bashrc (or whatever...rc) add theses two variables with value seen above

- MAILGUN_DOMAIN : your sandbox domain (without http)
- MAILGUN_PRIVATE_API_KEY : your private API key seen in step 4.
 
```
export MAILGUN_DOMAIN="your sandbox domain"
export MAILGUN_PRIVATE_API_KEY="your api Key"
```

## Add your email addresses in Authorized Recipients list
1. Log to https://www.mailgun.com/ and go to 'Domains' panel
2. Select your sandbox domain in the list 
3. On 'Authorized Recipients' block, add an email recipients that you want to use for tests

## Test your configuration

To test your Mailgun setup, you can rename and run unit test defined in [notification_sender_impl_test.go](internal/infrastructure/notification_sender_impl_test.go) :
- rename func to **TestSendEmail()**
- change 'Assignee' field with an authorized recipient

... after succeed call to Mailgun, check your spams :)
