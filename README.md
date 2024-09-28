## raw2curl

Converts raw HTTP based requests to curl (Unix).

**Inspiration** - I was facing few un-supported HTTP verb (like, `PATCH`) errors while using https://github.com/curl/h2c/.
<br>
<br>

## Usage Example 
```
$ cat sample-get-request.txt
GET /path?var=1 HTTP/1.1
Host: www.claudeusercontent.com
Sec-Ch-Ua: "Chromium";v="122", "Not(A:Brand";v="24", "Microsoft Edge";v="122"
Sec-Ch-Ua-Mobile: ?0
Sec-Ch-Ua-Platform: "macOS"
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS \ X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Sec-Fetch-Site: cross-site
Sec-Fetch-Mode: navigate
Sec-Fetch-Dest: iframe
Referer: https://claude.ai/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-IN,en-GB;q=0.9,en;q=0.8,en-US;q=0.7
Connection: keep-alive

$ raw2curl --help 
Usage of raw2curl:
  -file string
        Raw request file to parse.

$ raw2curl --file test/get.txt 
curl --http1.1 -X GET -H $'Host: www.claudeusercontent.com' -H $'Sec-Ch-Ua: \"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Microsoft Edge\";v=\"122\"' -H $'Sec-Ch-Ua-Mobile: ?0' -H $'Sec-Ch-Ua-Platform: \"macOS\"' -H $'Upgrade-Insecure-Requests: 1' -H $'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS \\ X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0' -H $'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' -H $'Sec-Fetch-Site: cross-site' -H $'Sec-Fetch-Mode: navigate' -H $'Sec-Fetch-Dest: iframe' -H $'Referer: https://claude.ai/' -H $'Accept-Encoding: gzip, deflate, br' -H $'Accept-Language: en-IN,en-GB;q=0.9,en;q=0.8,en-US;q=0.7' -H $'Connection: keep-alive' $'https://www.claudeusercontent.com/path?var=1'
```
<br>
<br>

## Keep in mind
- Sometimes, protocol errors might occur even though it’s a straightforward conversion to a curl flag (https://ec.haxx.se/http/versions/). A possible solution is to downgrade the protocol version.
- The output flag for the curl request is not provided, as it completely depends on your use case (https://curl.se/docs/manpage.html).

<br>
<br>

## Install 
```
go install -v github.com/okpalindrome/raw2curl@latest
```
