# Inject Example Server
Example Go (golang) web server with dependency injected controllers using the [Inject](http://github.com/karlkfi/inject) library.


## Dependencies

[Facebook's Grace library](http://github.com/facebookgo/grace) is used for graceful shutdown.

```
go get github.com/facebookgo/grace
```


## Usage

1. Launch the server:

    ```
    go run main.go
    ```

    (ctrl-c to quit)

1. Query Index:

    ```
    $ curl http://localhost:8080/; echo
    Welcome to the Inject Example Server!
    ```

1. Query User List:

    ```
    $ curl http://localhost:8080/users; echo
    {"users":[{"id":"0","name":"Alice"},{"id":"1","name":"Bob"},{"id":"2","name":"Charlie"}]}
    ```

1. Query User:

    ```
    $ curl http://localhost:8080/user/0; echo
    {"id":"0","name":"Alice"}
    $ curl http://localhost:8080/user/1; echo
    {"id":"1","name":"Bob"}
    $ curl http://localhost:8080/user/2; echo
    {"id":"2","name":"Charlie"}
    ```


## License

   Copyright 2015 Karl Isenberg

   Licensed under the [Apache License Version 2.0](LICENSE) (the "License");
   you may not use this project except in compliance with the License.

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
