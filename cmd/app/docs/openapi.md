---
title: ClassifierAAS API v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<h1 id="classifieraas-api-schema">Schema</h1>

## get__api_schema_{id}

> Code samples

```shell
# You can also use wget
curl -X GET /api/schema/{id} \
  -H 'Accept: application/json'

```

```http
GET /api/schema/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('/api/schema/{id}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get '/api/schema/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/api/schema/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/api/schema/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/api/schema/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/api/schema/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /api/schema/{id}`

*Get schema and its actual version by ID*

<h3 id="get__api_schema_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string(uuid)|true|none|

> Example responses

> 200 Response

```json
{
  "success": true,
  "data": {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "actualVariant": {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "description": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "type": "string",
          "nextID": "de6304da-dd2b-4477-989b-4e4ffdfb5693",
          "nextErrorID": "b80602bf-bd32-44e9-a34f-0ba8fb6ff3d4",
          "data": {},
          "gridData": {}
        }
      ],
      "createdAt": "2019-08-24T14:15:22Z",
      "updatedAt": "2019-08-24T14:15:22Z"
    },
    "createdAt": "2019-08-24T14:15:22Z",
    "updatedAt": "2019-08-24T14:15:22Z"
  }
}
```

<h3 id="get__api_schema_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful response|[Schema](#schemaschema)|

<aside class="success">
This operation does not require authentication
</aside>

## post__api_schema

> Code samples

```shell
# You can also use wget
curl -X POST /api/schema \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST /api/schema HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "type": "object",
  "properties": null
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('/api/schema',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post '/api/schema',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('/api/schema', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/api/schema', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/api/schema");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/api/schema", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /api/schema`

*Create new schema*

> Body parameter

```json
{
  "type": "object",
  "properties": null
}
```

<h3 id="post__api_schema-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|false|none|

> Example responses

> 200 Response

```json
{
  "success": true,
  "data": {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "actualVariant": {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "description": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "type": "string",
          "nextID": "de6304da-dd2b-4477-989b-4e4ffdfb5693",
          "nextErrorID": "b80602bf-bd32-44e9-a34f-0ba8fb6ff3d4",
          "data": {},
          "gridData": {}
        }
      ],
      "createdAt": "2019-08-24T14:15:22Z",
      "updatedAt": "2019-08-24T14:15:22Z"
    },
    "createdAt": "2019-08-24T14:15:22Z",
    "updatedAt": "2019-08-24T14:15:22Z"
  }
}
```

<h3 id="post__api_schema-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful response|[Schema](#schemaschema)|

<aside class="success">
This operation does not require authentication
</aside>

## put__api_schema

> Code samples

```shell
# You can also use wget
curl -X PUT /api/schema \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
PUT /api/schema HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "description": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "type": "string",
      "nextID": "de6304da-dd2b-4477-989b-4e4ffdfb5693",
      "nextErrorID": "b80602bf-bd32-44e9-a34f-0ba8fb6ff3d4",
      "data": {},
      "gridData": {}
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('/api/schema',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.put '/api/schema',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.put('/api/schema', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/api/schema', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/api/schema");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/api/schema", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /api/schema`

*Update existing schema*

> Body parameter

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "description": [
    {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "type": "string",
      "nextID": "de6304da-dd2b-4477-989b-4e4ffdfb5693",
      "nextErrorID": "b80602bf-bd32-44e9-a34f-0ba8fb6ff3d4",
      "data": {},
      "gridData": {}
    }
  ]
}
```

<h3 id="put__api_schema-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|false|none|
|» id|body|string(uuid)|false|none|
|» description|body|[object]|false|none|
|»» id|body|string(uuid)|false|none|
|»» type|body|string|false|none|
|»» nextID|body|string(uuid)|false|none|
|»» nextErrorID|body|string(uuid)|false|none|
|»» data|body|object|false|none|
|»» gridData|body|object|false|none|

> Example responses

> 200 Response

```json
{
  "success": true,
  "data": {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "actualVariant": {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "description": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "type": "string",
          "nextID": "de6304da-dd2b-4477-989b-4e4ffdfb5693",
          "nextErrorID": "b80602bf-bd32-44e9-a34f-0ba8fb6ff3d4",
          "data": {},
          "gridData": {}
        }
      ],
      "createdAt": "2019-08-24T14:15:22Z",
      "updatedAt": "2019-08-24T14:15:22Z"
    },
    "createdAt": "2019-08-24T14:15:22Z",
    "updatedAt": "2019-08-24T14:15:22Z"
  }
}
```

<h3 id="put__api_schema-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Successful response|[Schema](#schemaschema)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Schema">Schema</h2>
<!-- backwards compatibility -->
<a id="schemaschema"></a>
<a id="schema_Schema"></a>
<a id="tocSschema"></a>
<a id="tocsschema"></a>

```json
{
  "success": true,
  "data": {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "actualVariant": {
      "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
      "description": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "type": "string",
          "nextID": "de6304da-dd2b-4477-989b-4e4ffdfb5693",
          "nextErrorID": "b80602bf-bd32-44e9-a34f-0ba8fb6ff3d4",
          "data": {},
          "gridData": {}
        }
      ],
      "createdAt": "2019-08-24T14:15:22Z",
      "updatedAt": "2019-08-24T14:15:22Z"
    },
    "createdAt": "2019-08-24T14:15:22Z",
    "updatedAt": "2019-08-24T14:15:22Z"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|success|boolean|false|none|none|
|data|object|false|none|none|
|» id|string(uuid)|false|none|none|
|» actualVariant|object|false|none|none|
|»» id|string(uuid)|false|none|none|
|»» description|[object]|false|none|none|
|»»» id|string(uuid)|false|none|none|
|»»» type|string|false|none|none|
|»»» nextID|string(uuid)|false|none|none|
|»»» nextErrorID|string(uuid)|false|none|none|
|»»» data|object|false|none|none|
|»»» gridData|object|false|none|none|
|»» createdAt|string(date-time)|false|none|none|
|»» updatedAt|string(date-time)|false|none|none|
|» createdAt|string(date-time)|false|none|none|
|» updatedAt|string(date-time)|false|none|none|

