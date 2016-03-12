# xml2json

## 1.install
```
go get -u github.com/hidu/xml2json
```

## 2.Useage

```
options of xml2json:
  -jsonschema string
        json schema file path
  -out string
        json output file path
  -port uint
        run as http server,addr port eg 8080
  -xml string
        xml file path
```

### 2.1 basic
```
xml2json -xml b.xml

cat b.xml|xml2json
```

output json content to file:
```
xml2json -xml b.xml -out b.json
```
### 2.2 fix json with jsonschema
the [b.xml](test/b.xml) content:
```
<data>
    <a>va0</a>
    <b>bv0</b>
    <b>bv1</b>
</data>
```
output json content:
```
{
  "data": {
    "a": "va0",
    "b": [
      "bv0",
      "bv1"
    ]
  }
}
```

but we expect `data.a` is array not string :
```
{
  "data": {
    "a": [
      "va0"
    ],
    "b": [
      "bv0",
      "bv1"
    ]
  }
}
```
so we can use `-jsonschema` [b-schema.json](test/b-schema.json) option:
```
xml2json -xml b.xml -jsonschema b-schema.json
```

### 2.3 run as http service
```
xml2json -port 8080
``` 
and then you can visit the default ui: http:127.0.0.1:8080  

there is the API:

#### 2.3.1 POST /xml2json
params:  
`xml`: xml content ,required,not empty  
`json_schema`:  the jsons-chema to fix output josn,can empty

#### 2.3.2 POST /getxmljsonschema
get xml's json schema  
params:  
`xml`: xml content ,required,not empty  

#### 2.3.3 POST /jsonfix
fix json with json schema  
params:  
`json`: json string,required,not empty  
`json_schema`:  the jsons-chema to fix output josn,required,not emtpy

#### 2.3.4 POST /getjsonjsonschema
get json's json schema  
params:  
`json`: json content ,required,not empty  