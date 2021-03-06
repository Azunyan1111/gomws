# gomws
Amazon mws API in Go

[![Build Status](https://travis-ci.org/svvu/gomws.svg?branch=master)](https://travis-ci.org/svvu/gomws)

# Usage
Example usage can be found in main.go

Import the packages
```go
import(
  "github.com/svvu/gomws/gmws"
  "github.com/svvu/gomws/mws/products"
)
```
Setup the configuration
```go
config := gmws.MwsConfig{
  SellerId:  "SellerId",
  AuthToken: "AuthToken",
  Region:    "US",

  // Optional if already set in env variable
  AccessKey: "AKey",
  SecretKey: "SKey",
}
```
If AccessKey and SecretKey not find in the pass in configuration, then it will try to retrieve them from env variables (**AWS_ACCESS_KEY** and **AWS_SECRET_KEY**).

Create the client
```go
productsClient, err := products.NewClient(config)
```

Call the operations, the response is a struct contains result xml string and error if operation fail
```go
fmt.Println("------GetMatchingProduct------")
response := productsClient.GetMatchingProduct([]string{"ASIN"})
// Check whether or not the response return 200.
if response.Error != nil {
	fmt.Println(response.Error.Error())
}
// result() is xml response in string
fmt.Println(response.Result())
```

Use XMLNode parser to get the data from response.

Create the xmlNode parser.
```go
/*
  xmlNode in fact is a map[string]interface{}.
  xmlNode is a wrapper over mxj map.

  All attributes will become a node with key '-attributesName'.
  Tags with attributes, their value will become a node with key '#text'.

  Ex:
    <ProductName sku="ABC">
      This will become node also.
    </ProductName>

  Will become:
    map[string]interface{
      "-sku": "ABC",
      "#text": "This will become node also.",
    }
*/
xmlNode, err = gmws.GenerateXMLNode(response.Body)
```
Check whether or not API send back error message
```go
if gmws.HasErrors(xmlNode) {
  fmt.Println(gmws.GetErrors(xmlNode))
}
```
View the response in xml format.
```go
xmlNode.PrintXML()
```
Get the data by key (xml tag name)
```go
// products is a slice of XMLNode, means individual one can be used to retrieve
// data by methods provided by XMLNode, ex: FindByKey
products := xmlNode.FindByKey("Product")
```
Many methods can be used to traverse the xml tree. For more info, refer to the [godoc](https://godoc.org/github.com/svvu/gomws/gmws)
```go
// FindByKey get the nodes in any place of the tree.
productNodes := xmlNode.FindByKey("Product")

// FindByKeys can used to retrieve nodes which are children of other nodes.
productNameNodes := xmlNode.FindByKeys("Product", "Title")

// FindByPath get the nodes with specify tree path.
// Keys in the path are separated by '.'.
// Note: the first key must be a direct child of current node, and each subsequential
//  key must be direct child of previous key.
productNameNodes := xmlNode.FindByKeys("Product.AttributeSets.ItemAttributes.Title")
```
To get the value out of node, use the corresponding type methods
```go
xmlNode.ToString()
xmlNode.ToInt()
xmlNode.ToFloat()
xmlNode.ToBool()
xmlNode.ToTime()

// Ex:
productNameNodes := xmlNode.FindByKeys("Product", "Title")
name, err := productNameNodes[0].ToString()
```
To unmarshall the data to a struct, use method
```go
xmlNode.ToStruct(structPointer)
```
The struct use json format tags.
```go
// Ex:
// To unmarshal the tag:
//  <Message>
//    <Locale>en_US</Locale>
//    <Text>Error message 1</Text>
//  </Message>
// Can use struct:
msg := struct {
		Locale string `json:"Locale"`
		Text   string `json:"Text"`
}{}
err := xmlNode.FindByKey("Message")[0].ToStruct(&msg)

// To unmarshal the attributes, use -attributeName tag.
// To unmarshal the value of tags with attributes, use #text tag.
// Ex:
// To unmarshal the tag:
//  <MessageId MarketplaceID="ATVPDKDDIKX0D" SKU="24478624">
//		173964729
//  </MessageId>
// Can use struct:
type msgID struct {
  MarketplaceID string `json:"-MarketplaceID"`
  SKU           string `json:"-SKU"`
  ID            string `json:"#text"`
}
msgid := msgID{}
err := xmlNode.FindByKey("MessageId")[0].ToStruct(&msgid)
```

Other usefull methods
```go
// Get the current tag name of the node.
xmlNode.CurrentKey()
// Get the direct children's node name.
xmlNode.Elements()
// Check whether or not the current node is the leaf, which means can't traverse deeper.
xmlNode.IsLeaf()
// Get a list of path to all the leaves.
xmlNode.LeafPaths()
// Get a list of nodes which are leaves.
xmlNode.LeafNodes()
```

# APIs

## Products
The Products API helps to get information to match your products to existing product listings on Amazon Marketplace websites.

The Products API returns product attributes, current Marketplace pricing information, and a variety of other product and listing information.

## Orders
The Orders API helps to retrieve orders information on Amazon Marketplace.

The Orders API returns orders list, items info in the order, and a variety of other orders information.

# TODO
* Add support for other APIs

# Deprecated
The build in structs to unmarshal the result are deprecated. They still can be found in branch 0.0.3.

Reasons to remove the support the build structs:
1.  Make the code eaiser to maintain and eaiser for adding support for new API.
2.  Make the code more flexiable and error tolerance.
3.  Make the access to part of the result eaiser. Ex: to get the package width of the GetMatchingProduct response
  * With old style: response.GetMatchingProductResult.Product.AttributeSets.PackageDimensions.Width
  * With new style: response.FindByKeys("PackageDimensions", "Width")
