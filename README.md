# xml2json

## install
```
go get -u github.com/hidu/xml2json
```

##Useage
```
xml2json -xml b.xml
xml2json -xml b.xml -jsonschema b-schema.json  #fix json with JSONSchema
xml2json -xml b.xml -out b.json
cat b.xml|xml2json

```

###as http service
```
xml2json -port 8080
``` 