## raw2curl

Converts raw HTTP based request to curl (Unix).

**Inspiration** - I was facing few un-supported HTTP verb (like `PATCH`, etc.) errors while using https://github.com/curl/h2c/ for importing over Postman.
<br>
<br>

## Install 
```
go install -v github.com/okpalindrome/raw2curl@latest
```

<br>

## Usage Example 
```
$ raw2curl
Usage: raw2curl <file-path> or pipe input via stdin

$ cat sample-get-request.txt | raw2curl
curl --http1.1 -X GET -H 'Host: www.claudeusercontent.com' -H 'Sec-Ch-Ua: \"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Microsoft Edge\";v=\"122\"' -H 'Sec-Ch-Ua-Mobile: ?0' -H 'Sec-Ch-Ua-Platform: \"macOS\"' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS \\ X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' -H 'Sec-Fetch-Site: cross-site' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Dest: iframe' -H 'Referer: https://claude.ai/' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-IN,en-GB;q=0.9,en;q=0.8,en-US;q=0.7' -H 'Connection: keep-alive' $'https://www.claudeusercontent.com/path?var=1'
```
<br>

## Note
The output flag for the curl request is not provided, as it completely depends on what/how you want to see the response (https://curl.se/docs/manpage.html).
