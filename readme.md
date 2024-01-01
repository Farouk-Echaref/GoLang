# testing in Go:
- To create a unit test in go, you need to create a testfile:
    * Naming:
        - Snake Case
        - Start with the name of the file in question
        - Ends with _test.go
        - main_test.go for main.go
    * Placement::
        - Same Placement:
            feat
            |_ feat.go
            |_ feat_test.go
        - Test Package:
            feat
            |_ feat.go
            feat_test
            |_ feat_test.go
- Test Cases for individual test functions:
    * Naming:
        - Pascal Case
        - Starts with "Test" followed by the name of the function
        - TestF for function F()
        - TestT and TestT_M for: (testing type and method within type)
        ```go
        type T struct {}
        func (T) M() {}
        ```
- Usage:
    ```go
    // import package
    import "testing"

    // using testing.T struct
    func TestF(t *testing.T){
        //code
    }
    ```
