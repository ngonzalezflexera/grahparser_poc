package main

import (
    "encoding/json"
    "fmt"
    "github.com/graphql-go/graphql"
    "github.com/graphql-go/graphql/language/ast"
    "log"
)

var TDSObjectType = graphql.NewObject(graphql.ObjectConfig{
    Name:        "TDSObjectType",
    Fields:      graphql.Fields{
        "domain_name": &graphql.Field{
            Name:              "domain_name",
            Type:              graphql.String,
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Domain name of the TDS Object",
        },
        "serial_number": &graphql.Field{
            Name:              "serial_number",
            Type:              graphql.String,
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Serial number of the TDS Object",
        },
        "host_name": &graphql.Field{
            Name:              "host_name",
            Type:              graphql.String,
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Host name of the TDS Object",
        },
        "dataset": &graphql.Field{
            Name:              "dataset",
            Type:              graphql.String,
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Host name of the TDS Object",
        },
        "source": &graphql.Field{
            Name:              "source",
            Type:              graphql.String,
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Serial number of the TDS Object",
        },
        "another_field": &graphql.Field{
            Type:              graphql.NewList(graphql.String),
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Serial number of the TDS Object",
        },
    },
    IsTypeOf:    nil,
    Description: "TDS resource object definition",
})

var CDSObjectType = graphql.NewObject(graphql.ObjectConfig{
    Name:        "CDSObjectType",
    Fields:      graphql.Fields{
        "dataset": &graphql.Field{
            Name:              "dataset",
            Type:              graphql.String,
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Host name of the TDS Object",
        },
        "source": &graphql.Field{
            Name:              "source",
            Type:              graphql.String,
            Args:              nil,
            Resolve:           nil,
            DeprecationReason: "",
            Description:       "Serial number of the TDS Object",
        },
        "annotation": &graphql.Field{
            Name:              "annotation",
            Type:              graphql.NewList(graphql.String),
            Description:       "Annotations for CDS",
        },
        //"annotation2": &graphql.Field{
        //    Name:              "annotation",
        //    Type:              graphql.NewList(graphql.String),
        //    Description:       "Annotations for CDS",
        //},
    },
    IsTypeOf:    nil,
    Description: "CDS resource object definition",
})


func main() {
    query := `
query {
  tdsresource(dataset: "computer", source: "TDS") {
  domain_name
  serial_number
  host_name
  }
  cdsresource(dataset: "super", source: "CDS"){
  annotation {
        locations
    }
  }
}
`
    schema, err := graphql.NewSchema(graphql.SchemaConfig{
        Query: graphql.NewObject(graphql.ObjectConfig{
            Name: "Query",
            Fields: graphql.Fields{
                "tdsresource": &graphql.Field{
                    Name:              "tdsresource",
                    Type:              TDSObjectType,
                    Description:       "TDS resource block. This will be used to extract data from TDS",
                    Args: graphql.FieldConfigArgument{
                        "dataset": &graphql.ArgumentConfig{
                            Type: graphql.String,
                        },
                        "source": &graphql.ArgumentConfig{
                            Type: graphql.String,
                        },
                    },
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        GetFields(p.Info.FieldASTs)
                        return nil, nil
                    },
                },
                "cdsresource": &graphql.Field{
                    Name:              "cdsresource",
                    Type:              CDSObjectType,
                    Description:       "CDS resource block. This will be used to extract data from TDS",
                    Args: graphql.FieldConfigArgument{
                        "dataset": &graphql.ArgumentConfig{
                            Type: graphql.String,
                        },
                        "source": &graphql.ArgumentConfig{
                            Type: graphql.String,
                        },
                    },
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        GetFields(p.Info.FieldASTs)
                        return nil, nil
                    },
                },
            },
        }),
    })
    if err != nil {
        log.Fatal(err)
    }
    r1 := graphql.Do(graphql.Params{
        Schema:        schema,
        RequestString: query,
    })
    if len(r1.Errors) > 0 {
        log.Fatal(r1)
    }
    b1, err := json.MarshalIndent(r1, "", "  ")
    fmt.Println(string(b1))
}

func GetFields (f []*ast.Field ) []string {
    for _, j := range f {
        if len(j.GetSelectionSet().Selections) > 0 {
            fmt.Println("There is fields to be selected")
            GetSelections(j.GetSelectionSet().Selections)
        }
    }
    return nil
}
func GetSelections ( s []ast.Selection) []string{
    var fields []string
    for _, j := range s {
        field := j.(*ast.Field)
        fields = append(fields,field.Name.Value)
    }
    return fields
}