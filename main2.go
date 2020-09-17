package main

import (
    "encoding/json"
    "fmt"
    "github.com/graphql-go/graphql"
    "github.com/graphql-go/graphql/language/ast"
    "github.com/graphql-go/graphql/language/parser"
    _ "github.com/graphql-go/graphql/language/parser"
    "github.com/graphql-go/graphql/language/source"
    _ "github.com/graphql-go/graphql/language/source"
)

type parsedQuery struct {
    operation string
    operationName string
    variablesDefinitions []map[string]string
    nestedFields  []* ast.Field
}
//type Fields struct {
//    fieldName string
//    fieldValue [] string
//    fieldArguments map[string]string
//    nestedFields [] Fields
//}

type Correlation struct {
    SerialNumber string
}
var correlationInterface * graphql.Interface
func main() {
    //correlationType := graphql.NewObject(graphql.ObjectConfig{
    //    Name:       "Correlation",
    //    Interfaces: []*graphql.Interface{correlationInterface},
    //    Fields: graphql.Fields{
    //        "serialNumber": &graphql.Field{
    //            Name:              "serialNumber",
    //            Type:              graphql.String,
    //            Args:              nil,
    //            Resolve:           nil,
    //            DeprecationReason: "",
    //            Description:       "This is a string serial Number",
    //        },
    //    },
    //    IsTypeOf:    nil,
    //    Description: "Correlation type",
    //},
    //)

    //queryType := graphql.NewObject(graphql.ObjectConfig{
    //    Name: "Query",
    //    Fields: graphql.Fields{
    //        "resource": &graphql.Field{
    //            Name: "resource",
    //            Type: graphql.String,
    //            Args: graphql.FieldConfigArgument{
    //                "dataset": &graphql.ArgumentConfig{
    //                    Type:         graphql.String,
    //                    DefaultValue: nil,
    //                    Description:  "dataset description",
    //                },
    //                "source": &graphql.ArgumentConfig{
    //                    Type:         graphql.String,
    //                    DefaultValue: nil,
    //                    Description:  "dataset source",
    //                },
    //            },
    //            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    //                //TODO maybe we can do something with this for validation
    //                return nil, nil
    //            },
    //            DeprecationReason: "",
    //            Description:       "Resource field",
    //        },
    //        //"correlation": &graphql.Field{
    //        //    Name:              "correlation",
    //        //    Type:              correlationType,
    //        //    Args:              nil,
    //        //    Resolve:           nil,
    //        //    DeprecationReason: "",
    //        //    Description:       "",
    //        //},
    //    },
    //    IsTypeOf:    nil,
    //    Description: "",
    //})
    //
    //schemaConfig := graphql.SchemaConfig{
    //    Query: queryType,
    //}
    //schema, err := graphql.NewSchema(schemaConfig)
    //if err != nil {
    //    log.Fatalf("failed to create new schema, error: %v", err)
    //}
    /*query := `query {
      applications(publisher: "Microsoft", product: "Office") {
        id
        name
        product
        version
        publisher

        # Get computers for each application
        computers {
          id
          computerName

          # Get location of computer
          location {
            id
            longName
          }
        }
      }
    }`*/
    query2 := `
query {
  tdsresource(dataset: "computer", source: "TDS") {
  domainName
  serialNumber
  host
  }

  cdsresource(dataset: "fnms", source: "CDS") {
    annotation {
        location
    }
  }

  join {
    sn: annotations{
        correlations {
            serialNumber {
                key {
                    name
                }
            }
        }
    }
            domainName {
                key {
                    name
                }
            host {
                key {
                  name
                }
            }
        }
    }
  }

  set {
    annotations {
        location {
            key {
                name 
            }
        }
    }
  }
}
`
    params := parser.ParseParams{Source: &source.Source{Body: []byte(query2), Name: "test"}}
    res, err := parser.Parse(params)
    if err != nil {
        fmt.Println(err)
    }
    b, _ := json.MarshalIndent(res, "  ", "  ")
    fmt.Println(string(b))
    // params2 := parser.ParseParams{Source: &source.Source{Body: []byte(query2), Name: "test"}}
    // res2, err := parser.Parse(params2)
    // if err != nil {
    //     fmt.Println("error ", err)
    // }

    //params := graphql.Params{Schema: schema, RequestString: query2}
    //
    //r := graphql.Do(params)
    //if len(r.Errors) > 0 {
    //    log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    //}
    //rJSON, _ := json.Marshal(r)
    //fmt.Printf("%s \n", rJSON)
}
   // v := &visitor.VisitorOptions{
      // Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
           //switch node := p.Node.(type) {
           //case *ast.Named:
           //    if node.Name != nil {
           //        fmt.Println("C: ", node.Name.Value)
           //        return visitor.ActionSkip, nil
           //    }
           //case *ast.Field:
           //    if node.Name != nil {
           //        fmt.Println("A: ", node.Name.Value)
           //        if len(node.Arguments) > 0 {
           //            for _, j := range node.Arguments {
           //                fmt.Println((*j).Name.Value, " ", (*j).Value.GetValue())
           //            }
           //        }
           //        return visitor.ActionNoChange, nil
           //    }
           //case * ast.Argument:
           //    if node.Name != nil {
           //        fmt.Println("B: ", node.Name.Value)
           //        return visitor.ActionNoChange, nil
           //    }
           //}
           //return visitor.ActionNoChange, nil
    //   },
    //}
   /*var parseQuery parsedQuery

    v1 := &visitor.VisitorOptions{
        KindFuncMap: map[string]visitor.NamedVisitFuncs{
            kinds.OperationDefinition: {
                Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
                    node := p.Node.(*ast.OperationDefinition)
                    // I think that there is only 1
                    parseQuery.operation = node.Operation
                    if node.Name != nil && node.Name.Value != "" {
                        parseQuery.operationName = node.Name.Value
                    }
                    if len(node.VariableDefinitions) > 0 {
                        for _, j := range node.VariableDefinitions {
                            variableMap := map[string]string{}
                            // Type-> Type -> Name -> Value
                            variableMap[j.Variable.Name.Value] = j.Type.String()
                        }
                    }
                    if len (node.SelectionSet.Selections) > 0 {
                        parseQuery.nestedFields = GetFields(node.SelectionSet.Selections)
                        if parseQuery.nestedFields == nil {
                            fmt.Println("Null")
                        }
                        for _, y := range parseQuery.nestedFields {
                            foo := y.GetSelectionSet().Selections
                            for _, j := range foo {
                                bar := j.(*ast.Field)
                                fmt.Println("foo ", bar.Name.Value)
                                if bar.SelectionSet != nil {
                                    zar := GetFields(bar.SelectionSet.Selections)
                                    for _, x := range zar {
                                        fmt.Println(" zar", x.Name.Value)
                                    }
                                }
                            }
                        }
                    }
                    return visitor.ActionNoChange, nil
                },
            },
            //kinds.Field: {
            //    Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
            //        fmt.Println("Entering node " , p.Node.(ast.Node).GetKind())
            //        switch node := p.Node.(type) {
            //        case *ast.Field:
            //            if node.Name != nil {
            //                fmt.Println("A1: ", node.Name.Value)
            //                if len(node.Arguments) > 0 {
            //                    for _, j := range node.Arguments {
            //                        fmt.Println((*j).Name.Value, " ", (*j).Value.GetValue())
            //                    }
            //                }
            //                return visitor.ActionNoChange, nil
            //            }
            //        case * ast.Argument:
            //            if node.Name != nil {
            //                fmt.Println("B1: ", node.Name.Value)
            //                return visitor.ActionNoChange, nil
            //            }
            //        }
            //        return visitor.ActionNoChange, nil
            //    },
            //},
        },
    }
    //            Leave: func(p visitor.VisitFuncParams) (string, interface{}) {
    //                fmt.Println("Pong")
    //                if node, ok := p.Node.(*ast.Field); ok {
    //                    fmt.Println("Ping " , node)
    //                }
    //                return visitor.ActionNoChange, nil
    //            },
    //            Kind: func(p visitor.VisitFuncParams) (string, interface{}) {
    //                fmt.Println("Pong")
    //                if node, ok := p.Node.(*ast.Field); ok {
    //                    fmt.Println("Ping " , node)
    //                }
    //                return visitor.ActionNoChange, nil
    //            },
    //        },
    //        kinds.Name: {
    //            Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
    //                fmt.Println("Pong")
    //                if node, ok := p.Node.(*ast.Name); ok {
    //                    fmt.Println("Ping " , node)
    //                }
    //                return visitor.ActionNoChange, nil
    //            },
    //        },
    //        kinds.Argument: {
    //            Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
    //                fmt.Println("Pong")
    //                if node, ok := p.Node.(*ast.Argument); ok {
    //                    fmt.Println("Ping " , node)
    //                }
    //                return visitor.ActionNoChange, nil
    //            },
    //        },
    //        kinds.OperationDefinition: {
    //            Enter: func(p visitor.VisitFuncParams) (string, interface{}) {
    //                if node, ok := p.Node.(*ast.OperationDefinition); ok {
    //                    selectionSet = node.SelectionSet
    //                    for _, j := range selectionSet.Selections {
    //                        fmt.Println("A " , j.GetSelectionSet().Kind)
    //                    }
    //                    fmt.Println(selectionSet)
    //                    return visitor.ActionUpdate, ast.NewOperationDefinition(&ast.OperationDefinition{
    //                        Loc:                 node.Loc,
    //                        Operation:           node.Operation,
    //                        Name:                node.Name,
    //                        VariableDefinitions: node.VariableDefinitions,
    //                        Directives:          node.Directives,
    //                        SelectionSet: ast.NewSelectionSet(&ast.SelectionSet{
    //                            Selections: []ast.Selection{},
    //                        }),
    //                    })
    //                }
    //                return visitor.ActionNoChange, nil
    //            },
    //            Leave: func(p visitor.VisitFuncParams) (string, interface{}) {
    //                fmt.Println("Pang")
    //                if node, ok := p.Node.(*ast.OperationDefinition); ok {
    //                    return visitor.ActionUpdate, ast.NewOperationDefinition(&ast.OperationDefinition{
    //                        Loc:                 node.Loc,
    //                        Operation:           node.Operation,
    //                        Name:                node.Name,
    //                        VariableDefinitions: node.VariableDefinitions,
    //                        Directives:          node.Directives,
    //                        SelectionSet:        selectionSet,
    //                    })
    //                }
    //                return visitor.ActionNoChange, nil
    //            },
    //        },
    //    },
    //}
*/
    //_ = visitor.Visit(res, v, nil)
   //editedAst := visitor.Visit(res2, v1, nil)
   // editedAst1 := visitor.Visit(res2, v, nil)

    //fmt.Println("Foo ", editedAst)
   //_, _ = json.MarshalIndent(editedAst2, "  ", "  ")
//    b, _ := json.MarshalIndent(res2, "  ", "  ")
//    fmt.Println(string(b))
//}

//func GetFields(p *parsedQuery, selections []ast.Selection) {
//    for _, j := range selections {
//        k := j.GetSelectionSet().Selections
//        foo := j.(*ast.Field)
//        fmt.Println("foo ", foo.Loc.Start)
//    }
//}
//func GetFields(selections []ast.Selection) []*ast.Field{
//    var ret []*ast.Field
//    for _, j := range selections {
//        if foo, ok  := j.(*ast.Field); ok {
//           ret = append(ret, foo)
//        }
//    }
//    return ret
//}